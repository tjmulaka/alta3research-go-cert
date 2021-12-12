/* Consuming RESTful APIs | Tara
API - NASA's APOD API lookup, this will download the image & displays the image in browser
when we open http://localhost:8089/  in a browser */

package utils

import (
	"io"
	"log"
	"net/http"
)

func renderImage(w http.ResponseWriter, r *http.Request, url string, fileName string) {

	// create a GET request to the image
	resp, err := http.Get(url)

	// check for errors
	if err != nil {
		log.Fatal(" error for get url in renderImage ", err)
	}

	// the client must close the response body when finished
	defer resp.Body.Close()

	var imageName = fileName + ".jpg"

	//stream the body to the client without fully loading it into memory. Set response headers.

	//w.Header().Set("Content-Disposition", "attachment; filename="+imageName) //this will download the image.
	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", resp.Header.Get("Content-Length"))

	/* The io.Copy() function copies from source to destination until
	   either EOF is reached on source or an error occurs. Returns the number of bytes
	   copied and the first error encountered while copying, if any. */
	_, err = io.Copy(w, resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	http.ServeFile(w, r, imageName)
	w.WriteHeader(http.StatusOK)
}
