package main

import (
	"fmt"
	"net/http"
)

func logic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Execute the logic")
	w.Write([]byte("OK!!!"))
}
func main() {
	handlerLogic := http.HandlerFunc(logic)
	http.Handle("/", middleware(handlerLogic))
	http.ListenAndServe(":9000", nil)
}
func middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware execution before request")
		w.Write([]byte("Response in middleware "))
		handler.ServeHTTP(w, r)
		fmt.Println("Middleware execution after response")
	})
}
