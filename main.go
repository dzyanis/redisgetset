package main

import (
	"github.com/go-redis/redis"
	"net/http"
	"fmt"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		err := client.Set(
			r.URL.Query().Get("name"),
			r.URL.Query().Get("value"),
			0).Err()

		if err != nil {
			fmt.Fprintf(w, "Error %s", err.Error())
		}

		fmt.Fprintf(w, "OK")
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		val, err := client.Get(r.URL.Query().Get("name")).Result()
		if err != nil {
			fmt.Fprintf(w, "Error %s", err.Error())
		}
		fmt.Fprintf(w, val)
	})

	http.ListenAndServe(":8080", nil)
}
