package function

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		input = body
	}

	fmt.Printf("request body: %s\n", string(input))

	response := struct {
		Payload    string              `json:"payload"`
		Headers    map[string][]string `json:"headers"`
		Enviroment []string            `json:"enviroment"`
	}{
		Payload:    string(input),
		Headers:    r.Header,
		Enviroment: os.Environ(),
	}

	resBody, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resBody)
}
