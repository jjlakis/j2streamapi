package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
)

type srsClientResponse struct {
	code    int               `json:"code"`
	server  int               `json:"server"`
	clients []srsSingleClient `json:"$clients"`
}

type srsSingleClient struct {
	id         string  `json:"id"`
	vhost      string  `json:"vhost"`
	ip         string  `json: "stream"`
	pageUrl    string  `json: "ip"`
	swfUrl     string  `json: "pageUrl"`
	tcUrl      string  `json: "swfUrl"`
	url        string  `json: "tcUrl"`
	streamType string  `json: "type"`
	publish    bool    `json: "publish"`
	alive      float32 `json: "alive"`
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
			w.Write([]byte("Unable to unmarshal response"))
			return
		}

		fmt.Println(kaczka)

		w.Write((streamsBody))
		// w.Write([]byte("welcome"))
	})
	http.ListenAndServe(":8080", r)
}
