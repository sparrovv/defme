package hydra

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/sparrovv/defme/configuration"
)

var config configuration.Config

func Serve(port string, wordinkConfig configuration.Config) {
	log.Println("start server on port " + port)
	config = wordinkConfig

	http.HandleFunc("/", translationHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func translationHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	word := r.FormValue("word")
	translateTo := r.FormValue("to")
	returnJSON := true

	// TODO: implement:
	// err - timeout - should not be more than 5s
	// err - network
	formattedResponse := BuildResponse(word, config, translateTo, returnJSON)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, formattedResponse)

	log.Println(fmt.Sprintf("GET %v (took: %v)", "/word="+word+"&to="+translateTo, time.Since(now)))
}
