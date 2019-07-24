package main

import (
	"flag"
	"log"
	"html/template"
	"net/http"
	"github.com/gorilla/mux"
)

var (

	addr    = flag.String("addr", ":3000", "http service address")
	templ   = template.Must(template.New("dummy").Parse(templateStr))
)

func main() {

	flag.Parse()

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)

	log.Println("ListenAndServe will listen on %s", *addr)
	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe", err)
	}
}

func HomeHandler(w http.ResponseWriter, req *http.Request) {

	log.Print("home handler.")
	templ.Execute(w, nil)
}

const templateStr = "hello world"