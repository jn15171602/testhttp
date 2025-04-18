package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// This function handles incoming HTTP requests.
// It writes a simple HTML page back to the browser.
func handler(writer http.ResponseWriter, request *http.Request) {
	// request.URL.Path[1:] gets the part of the URL after the first slash.
	fmt.Fprintf(writer, "<html><body><h1>h1, very cool</h1><p>%s</p></body></html>", request.URL.Path[1:])
}

// This function is a middleware — it wraps around your real handler.
// It logs info about every request, like what the user asked for and how long it took.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now() // Record the current time

		// Log the HTTP method (GET, POST, etc.), the path, and the IP address of the client
		log.Printf("Request: %s %s from %s", r.Method, r.RequestURI, r.RemoteAddr)

		// Call the actual handler (the one that generates the response)
		next.ServeHTTP(w, r)

		// Calculate how long the handler took to run, and log it
		duration := time.Since(start)
		log.Printf("Handled in %v", duration)
	})
}

func main() {
	// Create a new "ServeMux", which is like a router that maps paths to functions
	mux := http.NewServeMux()

	// Tell the router that when someone visits "/", it should call our handler function
	mux.HandleFunc("/", handler)

	// Wrap our router in the logging middleware so every request gets logged
	loggedMux := LoggingMiddleware(mux)

	// Start the web server on port 8080 using our wrapped router
	log.Println("Starting web server on http://localhost:8080")
	http.ListenAndServe(":8080", loggedMux)
}
