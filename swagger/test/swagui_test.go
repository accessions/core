package test

import (
	_ "embed"
	"github.com/accessions/core/swagger"
	"net/http"
	"testing"
)
//go:embed swagger.json
var spec []byte
func TestSwaggerUi (t *testing.T) {
	http.Handle("/swagger/", http.StripPrefix("/swagger", swagger.Handler(spec, swagger.SpecTypeJSON)))
	_ = http.ListenAndServe(":8080", nil)
}
