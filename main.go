package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dvsekhvalnov/jose2go"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
)

type scriptoria struct {
	secret string
}

func (s *scriptoria) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, _ := ioutil.ReadAll(r.Body)
	unsafe := blackfriday.MarkdownBasic(b)
	html := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
	token, _ := jose.Sign(string(html), jose.HS256, []byte(s.secret))

	w.Header().Set("Token", token)
	w.Write(html)
}

func main() {

	app := &scriptoria{
		secret: "TESTING123",
	}

	s := &http.Server{
		Addr:           ":8080",
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
