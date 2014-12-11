package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"io"
)

func helloHandler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
	<head>
		<title>Hello Gopher</title>
	</head>
	<body>
		Hello Gopher!
	</body>
</html>`,
	)
}
func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Go web apps!")
}
func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(getPort(),nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}
// Get the Port from the environment
func getPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the hosting environment
	if port == "" {
		port = "8080"
	}
	fmt.Println("listening at port: " + port)
	return ":" + port
}
