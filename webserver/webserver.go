package webserver

import "net/http"

func MyWebserver() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":3000", nil)
}

func home(writer http.ResponseWriter, reader *http.Request) {
	http.ServeFile(writer, reader, "./webserver/index.html")
}
