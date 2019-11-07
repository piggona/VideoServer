package main

import (
	"log"
	"net/http"
)

var httpClient *http.Client

func init() {
	httpClient = &http.Client{}
}

func request(b *ApiBody, w http.ResponseWriter, r *http.Request) {
	// var resp *http.Response
	var err error

	switch b.Method {
	case http.MethodGet:
		req, _ := http.NewRequest("Get", b.Url, nil)
		req.Header = r.Header
		// resp, err = httpClient.Do(req)
		if err != nil {
			log.Printf("%s", err)
			return
		}
	}
}

func normalResponse(w http.ResponseWriter, r *http.Response) {
	// res, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// re, _ := json.Unmarshal(res,)
	// }
}
