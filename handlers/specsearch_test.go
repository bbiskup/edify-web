package handlers

import (
	"github.com/stretchr/testify/assert"
	//	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

/*func TestSearchMsgSpecsNoResults(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SpecSearch)
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}*/

func TestSpecSearchAllResults(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	assert.Nil(t, err)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(SpecSearch)
	handler.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	//log.Printf("Body: %s", w.Body)
}

func BenchmarkSpecSearchAllResults(b *testing.B) {
	req, err := http.NewRequest("GET", "", nil)
	assert.Nil(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(SpecSearch)
		handler.ServeHTTP(w, req)
		assert.Equal(b, http.StatusOK, w.Code)
	}
}

func BenchmarkSpecSearchNoResults(b *testing.B) {
	req, err := http.NewRequest("GET", "?searchterm=sdfusfuspdfsdfjsfjsldfjopsfoa", nil)
	assert.Nil(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(SpecSearch)
		handler.ServeHTTP(w, req)
		assert.Equal(b, http.StatusOK, w.Code)
	}
}
