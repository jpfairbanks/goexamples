package main

import "fmt"
import "log"
import "net/http"
import "html"
import "flag"

//import "encoding/json"

func f(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %q", html.EscapeString(r.URL.Path))
}

type config struct {
	port string
	path string
}

var cnf config = config{port: ":8080", path: "/test/"}

func init() {
	flag.StringVar(&cnf.port, "port", ":8080", "port on which to listen")
	flag.StringVar(&cnf.path, "path", "/", "path on which to serve")
}

func main() {
	flag.Parse()
	http.HandleFunc(cnf.path, f)
	err := http.ListenAndServe(cnf.port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
