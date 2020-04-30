// Simply presents an html web page. 
package main


import (
    "fmt"
    "log"
    "net/http"
)

// Next Up is creating a listener to handel the number of incoming connections
// Learning more about the html template features.
// Maybe start developing some basic unit tests and learning about go's testing


func main() {

	srv := &http.Server{
		Addr:           ":8080",
	}

	http.HandleFunc("/", servePage)
	log.Fatal(srv.ListenAndServe())
	
}

func servePage (w http.ResponseWriter, r *http.Request) {

	location := "./test_page/"
	pageName := r.URL.Path[len("/"):]
	
	fmt.Println(pageName)
	
	http.ServeFile(w, r, location + pageName + ".html")
}
