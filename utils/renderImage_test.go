/* Alta3 Research | Tara
   Testing - file on which to run a simple test */

package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRenderImageHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()
	url := "https://apod.nasa.gov/apod/image/2112/IridescenzaLunaPleiadi1024.jpg"
	fileName := "IridescenzaLunaPleiadi1024.jpg"
	renderImage(res, req, url, fileName)

	if res.Code != http.StatusOK {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusOK)
	}
}
