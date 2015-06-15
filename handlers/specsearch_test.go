package handlers

import (
	"github.com/stretchr/testify/assert"
	//	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	searchTermAllResults = ""
	searchTermNoResults  = "ujsdfusdfipsfipsfjöksfösfipsf"
)

func TestSearchMsgSpecsNoResults(t *testing.T) {
	results := searchMsgSpecs(searchTermNoResults)
	assert.Equal(t, 0, len(results))
}

func TestSearchMsgSpecsAllResults(t *testing.T) {
	results := searchMsgSpecs(searchTermAllResults)
	assert.Equal(t, 194, len(results))
}

func TestSearchSegSpecsNoResults(t *testing.T) {
	results := searchSegSpecs(searchTermNoResults)
	assert.Equal(t, 0, len(results))
}

func TestSearchSegSpecsAllResults(t *testing.T) {
	results := searchSegSpecs(searchTermAllResults)
	assert.Equal(t, 156, len(results))
}

func TestSearchCompositeDataElemSpecsNoResults(t *testing.T) {
	results := searchCompositeDataElemSpecs(searchTermNoResults)
	assert.Equal(t, 0, len(results))
}

func TestSearchCompositeDataElemSpecsAllResults(t *testing.T) {
	results := searchCompositeDataElemSpecs(searchTermAllResults)
	assert.Equal(t, 198, len(results))
}

func TestSearchSimpleDataElemSpecsNoResults(t *testing.T) {
	results := searchSimpleDataElemSpecs(searchTermNoResults)
	assert.Equal(t, 0, len(results))
}

func TestSearchSimpleDataElemSpecsAllResults(t *testing.T) {
	results := searchSimpleDataElemSpecs(searchTermAllResults)
	assert.Equal(t, 649, len(results))
}

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
	req, err := http.NewRequest("GET", "?searchterm="+searchTermNoResults, nil)
	assert.Nil(b, err)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(SpecSearch)
		handler.ServeHTTP(w, req)
		assert.Equal(b, http.StatusOK, w.Code)
	}
}
