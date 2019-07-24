package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":3000", "http service address")

var templ = template.Must(template.New("dummy").Parse(templateStr))

func main() {

	flag.Parse()
	http.Handle("/", http.HandlerFunc(Home))

	log.Println("ListenAndServe will listen on %s", *addr)
	err := http.ListenAndServe(*addr, nil)

	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func Home(w http.ResponseWriter, req *http.Request) {

	log.Println("home handler.")
	templ.Execute(w, nil)
}

const templateStr = "hello world"