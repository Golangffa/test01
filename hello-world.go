package main

import (
	"fmt"
	"net/http"
	// "os"
	// "io"
	// "log"
	// // /Users/masimli/go/download.jpeg"image"
  // // "image/jpeg"
)

func hello(w http.ResponseWriter, r *http.Request) {
	var name string

	found := r.URL.Query().Get("name")
	if found != "" {
		name = found
	} else {
		name = "world"
	}

	fmt.Fprintf(w, "Hello, %s!", name)

	//var Path = "/Users/masimli/go/download.jpeg"
  //ResizeImage(Path, 500)
//
// img, err := os.Open("/Users/masimli/go/download.jpeg")
// if err != nil {
// 		log.Fatal(err) // perhaps handle this nicer
// }
// defer img.Close()
// w.Header().Set("Content-Type", "image/jpeg") // <-- set the content-type header
// io.Copy(w, img)

}

func main() {
	fmt.Print("Go to http://localhost:7078/?name=Fuad\n")

	http.HandleFunc("/", hello)
	http.ListenAndServe(":7078", nil)
}
