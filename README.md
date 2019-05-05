# url-shortener
This Url shortener service recieves long url and creates a mapping with short hash.
The short hash is a sequential number which is pushed in redis cache against long url. For maintaining counter we are pushing integer sequence in a redis queue.

