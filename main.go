package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func webtenso(w http.ResponseWriter, r *http.Request) {

	url := r.FormValue("tenso")
	if len(url) != 0 {
		fmt.Println(url)
		res, err := http.Get(url)
		checkError(err)
		data, err := ioutil.ReadAll(res.Body)
		checkError(err)
		fmt.Fprintf(w, string(data))

	} else {
		fmt.Fprintf(w, "Ooops!\nFormat Example: http://xxx.com:8080/?tenso=http://yyy.com/api")
	}

}

func main() {

	http.HandleFunc("/", webtenso)
	err := http.ListenAndServe(":8080", nil)
	checkError(err)

}

func checkError(err error) {
	if err != nil {
		log.Fatal("Ooops! : ", err)
	}
}
