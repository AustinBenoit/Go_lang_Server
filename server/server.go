// Simply presents an html web page. 
package main


import (
	//"fmt"
    "log"
    "net/http"
)

func main() {
	handler := http.FileServer(http.Dir("./server/test_page"))
	http.Handle("/", handler)
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}
