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