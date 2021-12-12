/* Consuming RESTful APIs | Tara
API - NASA's APOD API lookup, this will download the image & displays the image in browser
when we open http://localhost:8089/  in a browser */

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model "github.com/tjmulaka/alta3research-go-cert/model"
	utils "github.com/tjmulaka/alta3research-go-cert/utils"
)

func main() {

	// The HandleFunc registers the handler function for the given URL pattern
	http.HandleFunc("/", httpDisplayImage)
	fmt.Println("Server started at port 8089")
	log.Fatal(http.ListenAndServe(":8089", nil)) // listens on TCP network address for incoming HTTP requests

}

func httpDisplayImage(w http.ResponseWriter, r *http.Request) {

	// define URL (NASA's APOD API) as a string
	url := "https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY"

	// Build the request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("APOD API NewRequest: ", err)
		return
	}

	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client - you can load it up with
	// parameters if you wanted (timeouts, tls, retries, etc.)
	// You only need ONE of these (supported by goroutines)
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("APOD API Do: ", err)
		return
	}

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var pictureData model.APOD

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&pictureData); err != nil {
		log.Fatal("APOD API json Decode error : ", err)
	}

	// pictureData is the data - display API response values
	fmt.Println("Capsule Record     =\n", pictureData)
	fmt.Println("Copyright     =", pictureData.Copyright)
	fmt.Println("Title     =", pictureData.Title)
	fmt.Println("URL     =", pictureData.URL)

	// image rendering starts
	utils.RenderImage(w, r, pictureData.URL, pictureData.Title)
	fmt.Println("successfully displayed the image     =", pictureData.URL)

}
