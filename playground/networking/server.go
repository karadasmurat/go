package networking

import (
	"io"
	"log"
	"net/http"
)

func RunServer() {

	handleRoot := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello world!\n")
	}

	handleTodos := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Todo Handler!\n")
	}

	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/todos", handleTodos)

	// ListenAndServe listens on the TCP network address addr
	// The handler is typically nil, in which case the DefaultServeMux is used.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
