package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type srsClientResponse struct {
	Code    int64             `json:"code"`
	Server  int64             `json:"server"`
	Clients []srsSingleClient `json:"clients"`
}

type srsSingleClient struct {
	ID         int64   `json:"id"`
	Vhost      int64   `json:"vhost"`
	Stream     int64   `json:"stream"`
	IP         string  `json:"ip"`
	PageURL    string  `json:"pageUrl"`
	SwfURL     string  `json:"swfUrl"`
	TcURL      string  `json:"tcUrl"`
	URL        string  `json:"url"`
	StreamType string  `json:"type"`
	Publish    bool    `json:"publish"`
	Alive      float64 `json:"alive"`
}

const apiurl = "http://lakis.eu:1985"

func main() {
	r := chi.NewRouter()
	r.Get("/streams", func(w http.ResponseWriter, r *http.Request) {

		streams, err := http.Get(apiurl + "/api/v1/clients/")
		if err != nil {
			w.Write([]byte("Unable to get stream info from SRS API"))
			return
		}
		streamsBody, err := ioutil.ReadAll(streams.Body)
		if err != nil {
			w.Write([]byte("Unable to read stream info retrieven from SRS API"))
			return
		}

		kaczka := srsClientResponse{}
		err = json.Unmarshal(streamsBody, &kaczka)
		if err != nil {
			fmt.Println(err)
			w.Write([]byte("Unable to unmarshal response"))
			return
		}

		fmt.Println(kaczka.Clients)

		w.Write((streamsBody))
		// w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":8080", r)
}
