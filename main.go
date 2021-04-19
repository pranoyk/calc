package main

import "net/http"

func main () {
	 router := http.NewServeMux()

	 router.Handle("/", http.FileServer(http.Dir("public")))
	 http.ListenAndServe("0.0.0.0:8080", router)
}