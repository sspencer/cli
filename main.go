package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/sspencer/cli/internal/sumslib"
)

// Handler functions
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, `<!DOCTYPE html>
<html>
<head>
    <title>Grok's Go Server</title>
    <style>
        body { font-family: Arial, sans-serif; text-align: center; padding: 50px; background: #0d1117; color: #c9d1d9; }
        h1 { color: #58a6ff; }
        .info { margin-top: 30px; font-size: 1.1em; }
    </style>
</head>
<body>
    <h1>🚀 Hello from Pure Go!</h1>
    <p>A minimal web server using only the Go standard library.</p>
    
    <div class="info">
        <p><strong>Path:</strong> %s</p>
        <p><strong>Method:</strong> %s</p>
        <p><strong>Time:</strong> %s</p>
        <p><strong>Remote Addr:</strong> %s</p>
    </div>
</body>
</html>`, r.URL.Path, r.Method, time.Now().Format(time.RFC1123), r.RemoteAddr)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status":"healthy","service":"go-stdlib-server","time":"`+time.Now().Format(time.RFC3339)+`"}`)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"message":   "Welcome to the API endpoint!",
		"version":   "1.0.0",
		"framework": "net/http (stdlib)",
	}

	fmt.Fprintf(w, `{
  "message": "%s",
  "version": "%s",
  "framework": "%s",
  "timestamp": "%s"
}`, response["message"], response["version"], response["framework"], time.Now().Format(time.RFC3339))
}

func sumsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	if _, ok := r.URL.Query()["help"]; ok {
		fmt.Fprintln(w, "Usage: /sums/<number>?[x=1,4]&[i=5,7]")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Find all 3-number combinations from 1-9 that sum to <number>.")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Arguments:")
		fmt.Fprintln(w, "  <number>   Target sum, must be between 6 and 24 (inclusive)")
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "Query parameters:")
		fmt.Fprintln(w, "  x=<nums>   Comma-separated numbers to exclude (e.g., x=1,4)")
		fmt.Fprintln(w, "  i=<nums>   Comma-separated numbers, one of which must be included (e.g., i=5,7)")
		fmt.Fprintln(w, "  help       Show this help message")
		return
	}

	targetStr := r.PathValue("target")
	includeStr := r.URL.Query().Get("i")
	excludeStr := r.URL.Query().Get("x")
	err := sumslib.FindCombinationsConv(w, targetStr, includeStr, excludeStr)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
		return
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/sums/{target}", sumsHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback for local development
	}

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	fmt.Printf("🌟 Starting server on http://localhost:%s\n", port)
	log.Fatal(server.ListenAndServe())
}
