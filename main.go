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
	img := NewImg(2, 2, 800, 400)
	// gen points from defined distribution
	points := genDist(dist, 100, 2, 1)
	// create scatter chart
	scatterChart(genFloats(100), points, img)
	// encode RGBA to png
	png.Encode(w, img.I)
}
