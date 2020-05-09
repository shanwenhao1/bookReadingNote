package baseCgi

import (
	"log"
	"net/http"
	"time"
)

func hello(wr http.ResponseWriter, r *http.Request) {
	wr.Write([]byte("hello"))
}

// time middleware
func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		// next handler
		next.ServeHTTP(wr, r)

		timeElapsed := time.Since(timeStart)
		log.Println(timeElapsed)
	})
}

func BaseMiddleware() {
	http.Handle("/", timeMiddleware(http.HandlerFunc(hello)))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
