/* Alta3 Research | Tara
   Testing - file on which to run a simple test */

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpDisplayImageHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	httpDisplayImage(res, req)

	if res.Code == http.StatusOK {
		t.Log("got status http.StatusOK")
	}
}
