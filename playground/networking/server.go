package networking

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"playground/model"
)

func RunServer() {

	handleRoot := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello world!\n")
	}

	handleWizards := func(w http.ResponseWriter, _ *http.Request) {
		potter := model.Wizard{Name: "Potter", House: "Gryffindoor"}
		// Marshaling: Converting a Go struct to JSON
		potterJSON, err := json.Marshal(potter)

		// simulate error in marshalling
		// err = errors.New("dummy error")

		if err != nil {
			// fmt.Println("Error:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		io.WriteString(w, string(potterJSON))
	}

	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/wizards", handleWizards)

	// ListenAndServe listens on the TCP network address addr
	// The handler is typically nil, in which case the DefaultServeMux is used.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
