package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc(
		"/",
		func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(
				w,
				`<h1> Under construction </h1>
                 <img src="https://media.tenor.com/bnG9NZdIQe4AAAAi/pudgy-pudgypenguin.gif" />`,
			)
		},
	)

	s := &http.Server{
		Handler:      r,
		Addr:         fmt.Sprintf(":%d", 80),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := s.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
