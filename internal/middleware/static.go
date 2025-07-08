package middleware

import (
	"log"
	"net/http"
	"path/filepath"
)

//ServeStatic – prepares static directory and serves it to http.Handler
func ServeStatic(path string) http.Handler{
  // Abstract static folder path
	absStaticDir, absErr := filepath.Abs(path)
  if absErr != nil {
    log.Fatalf("invalid static directory path %q: %v", path, absErr)
  }
    
  // Serve /static/* from that dir
  fs := http.FileServer(http.Dir(absStaticDir))
  staticDir := http.StripPrefix("/staticFiles/", fs)

  
  return staticDir
} 

/*
func ServeStatic(path string) func(http.Handler) http.Handler{
  // Abstract static folder path
	absStaticDir, absErr := filepath.Abs(path)
  if absErr != nil {
    log.Fatalf("invalid static directory path %q: %v", path, absErr)
  }
    
  // Serve /static/* from that dir
  fs := http.FileServer(http.Dir(absStaticDir))
  staticDir := http.StripPrefix("/staticFiles/", fs)


  return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if strings.HasPrefix(r.URL.Path, "/staticFiles/") {
                staticDir.ServeHTTP(w, r)
                return
            }
            next.ServeHTTP(w, r) 
        })
    }
} 
*/