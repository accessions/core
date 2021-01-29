package swagger

import (
	"bytes"
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"testing/fstest"
)

//go:generate go run generate.go

//go:embed embed
var swagfs embed.FS

// SpecType represents one of the two types of Swagger spec
type SpecType string


const (
	SpecTypeJSON = `json`
	SpecTypeYAML = `yaml`
)

type swagWrapFS struct {
	static  fs.FS
	overlay fstest.MapFS
}

func (s *swagWrapFS) Open(name string) (fs.File, error) {
	if _, err := s.overlay.Stat(name); err == nil {
		return s.overlay.Open(name)
	}
	return s.static.Open(name)
}

// Handler
func Handler(spec []byte, spectype SpecType) http.Handler {

	static, _ := fs.Sub(swagfs, "embed")
	tmpl, _ := template.ParseFS(static, "index.html")
	idxbuf := new(bytes.Buffer)
	_ = tmpl.ExecuteTemplate(idxbuf, "index.html", struct{ SpecType string }{string(spectype)})
	m := fstest.MapFS{
		"index.html":                  {Data: idxbuf.Bytes()},
		"swagger." + string(spectype): {Data: spec},
	}
	return http.FileServer(http.FS(&swagWrapFS{
		static:  static,
		overlay: m,
	}))
}
