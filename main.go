package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bcgov/efv-api/config"
	"github.com/bcgov/efv-api/internal/handlers"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to EFV API!!")
}

func main() {
	cfg := config.Load()

	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	// register mock API endpoints from internal handlers
	handlers.RegisterRoutes(mux)
	// also serve the CRA-named OpenAPI file for compatibility
	mux.HandleFunc("/openapi-efv.yaml", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "docs/openapi-efv.yaml")
	})
	// serve the swagger-ui index and assets (we use a small index.html + CDN for assets)
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs/swagger-ui"))))

	srv := &http.Server{
		Addr:         ":" + cfg.Port,
		Handler:      cors(mux, cfg.AllowedOrigins, cfg.AllowCredentials),
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	log.Printf("starting server on %s", srv.Addr)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server failed: %v", err)
	}
}

func cors(next http.Handler, allowedOrigins []string, allowCredentials bool) http.Handler {
	allowAll := false
	origins := map[string]struct{}{}
	for _, o := range allowedOrigins {
		if o == "*" {
			allowAll = true
			break
		}
		origins[o] = struct{}{}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			if allowAll {
				w.Header().Set("Access-Control-Allow-Origin", "*")
			} else if _, ok := origins[origin]; ok {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				w.Header().Set("Vary", "Origin")
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			w.Header().Set("Access-Control-Expose-Headers", "Content-Length")
			if allowCredentials {
				w.Header().Set("Access-Control-Allow-Credentials", "true")
			}
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
