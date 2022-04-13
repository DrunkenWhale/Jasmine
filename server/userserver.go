package server

import "net/http"



func StartUserServer(caches []string,port int) {
	http.HandleFunc("/api", func(writer http.ResponseWriter, request *http.Request) {

	})
}
