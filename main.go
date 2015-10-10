package main

import (
	"flag"
	"fmt"
	"image/png"
	"log"
	. "net/http"
	"strings"
)

var (
	port = flag.String("port", "8080", "port to listent to")

	// map of distribution names
//	dists = map[string]
)

func main() {
	flag.Parse()
	fmt.Println("start web server to generate and display plots")

	// register handlers
	HandleFunc("/", dstr)

	// start server
	log.Fatal(ListenAndServe(":"+*port, nil))

}

// dstr handler generates distributions
func dstr(w ResponseWriter, r *Request) {
	// get distribution name from url path
	dist := strings.Replace(r.URL.Path, "/", "", -1)
	// define image params
	img := NewDumper(2, 2, 400, 300)
	histChart(dist, "test histo", true, true, false, img)
	png.Encode(w, img.I)
}
