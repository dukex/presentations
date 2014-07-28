package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/dukex/moraraqui/crawlers"
	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	// START OMIT
	m.Get("/api/imoveis/:state/:city/:neighborhood", func(res http.ResponseWriter, req *http.Request, params martini.Params) {
		res.Header().Set("Content-Type", "text/event-stream")

		item, timeout := crawlers.Get(params["state"], params["city"], params["neighborhood"])

		for {
			select {
			case property, ok := <-item:
				if !ok {
					item = nil
				}
				b, _ := json.Marshal(property)
				b = append(b, []byte("\n")...)
				res.Write(b)
				res.(http.Flusher).Flush()
			case <-timeout:
				item = nil
				return
			}
		}
	})
	// END OMIT
}
