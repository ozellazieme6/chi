package chi

import (
	"net/http"
)

func (mx *Mux) routeHTTP(w http.ResponseWriter, r *http.Request) bool {
	ctx := RouteContext(r.Context())
	if ctx == nil {
		ctx = NewRouteContext()
		r = r.WithContext(context.WithValue(r.Context(), RouteCtxKey, ctx))
	}

	// ... existing routing logic ...
	// Ensure we update the existing ctx instead of replacing it
	return mx.router.Match(ctx, r.Method, r.URL.Path)
}