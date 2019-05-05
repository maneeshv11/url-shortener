package service

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
	"url-shortener/cache"
	"url-shortener/model"
)

func HandleOriginalUrlRequest(w http.ResponseWriter, r *http.Request) {

	if params := mux.Vars(r); params["shortCode"] != "" {
		cacheKey := shortUrlCacheKey(params["shortCode"])

		if orgUrl, err := cache.Get(cacheKey); err == nil {
			http.Redirect(w, r, orgUrl, http.StatusFound)
			return
		} else {
			log.Printf("%s", err)
		}
	}
	w.WriteHeader(http.StatusNotFound)
	_, _ = fmt.Fprintf(w, "<h1>404 Not Found</h1>")

}

func ShortenUrl(w http.ResponseWriter, r *http.Request) {

	var request model.AddShortUrlRequest
	_ = json.NewDecoder(r.Body).Decode(&request)

	if shortCode, err := nextHashCode(); err == nil && request.OrgUrl != "" {

		if err := cache.Set(shortUrlCacheKey(shortCode), request.OrgUrl); err == nil {

			err := json.NewEncoder(w).Encode(model.AddShortUrlResponse{
				OrgUrl:   request.OrgUrl,
				ShortUrl: shortCode,
			})

			if err == nil {
				return
			}
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusBadRequest)

}

func HandlePopulateQueue(w http.ResponseWriter, _ *http.Request) {

	if shortCode, err := cache.RPop("tiny::shortKey"); err == nil || err == redis.Nil {

		if shortCode == "" {
			shortCode = "1"
		}

		if start, err := strconv.ParseInt(shortCode, 10, 64); err == nil {

			items := populateHashQueue(start, 100)
			_, _ = fmt.Fprintf(w, "successfully added %d items in queue", items)
			return

		} else {

			log.Printf("error in parsing integer from last pushed short code. error: %s", err)

		}
	} else {

		log.Printf("error in get last pushed element error: %s", err)

	}

	w.WriteHeader(http.StatusInternalServerError)

}

func shortUrlCacheKey(shortCode string) string {

	return "tiny::" + strings.TrimSpace(shortCode)

}

func nextHashCode() (string, error) {

	if shortCode, err := cache.LPop("tiny::shortKey"); err == nil {

		return shortCode, nil

	} else {

		log.Printf("Not abel to find the new short code from queue. error: %s", err)
		return "", err

	}

}

func populateHashQueue(start int64, size int64) int64 {

	var shortCodes []string

	end := start + size

	for i := start; i < end; i++ {
		shortCodes = append(shortCodes, strconv.FormatInt(i, 10))
	}

	if itemsAdded, err := cache.RPush("tiny::shortKey", shortCodes...); err == nil {
		log.Printf("Added %d items in tiny::shortKey successfully", itemsAdded)
		return itemsAdded
	}

	return 0

}
