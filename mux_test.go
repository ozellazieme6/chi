package chi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNestedRoutingWithPathRewrite(t *testing.T) {
	r := NewRouter()
	
	r.Route("/api/{version}", func(r Router) {
		r.Use(func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				req.URL.Path = "/api/v1/target"
				next.ServeHTTP(w, req)
			})
		})

		subRouter := NewRouter()
		subRouter.Get("/target", func(w http.ResponseWriter, req *http.Request) {
			version := URLParam(req, "version")
			if version != "v1" {
				t.Errorf("expected version parameter to be 'v1', got '%s'", version)
			}
			w.Write([]byte("ok"))
		})

		r.Mount("/", subRouter)
	})

	ts := httptest.NewServer(r)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/api/v1/target")
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status 200, got %d", res.StatusCode)
	}
}