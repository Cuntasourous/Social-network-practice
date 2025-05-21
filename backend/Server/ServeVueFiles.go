package Server

import (
	"net/http"
	"os"
	"path/filepath"
)

func ServeDistFiles() {
	// Serve static files from the dist directory
	fs := http.FileServer(http.Dir("frontend/Social-network/dist"))
	http.Handle("/assets/", fs)
}

func SpaHandler(w http.ResponseWriter, r *http.Request) {
	// get path to dist directoty to serve the index.html inside it
	// dist = frontend\Social-network\dist
	distDir := filepath.Join("frontend", "Social-network", "dist")
	indexFile := filepath.Join(distDir, "index.html")

	requestedFile := filepath.Join(distDir, r.URL.Path)

	if _, err := os.Stat(requestedFile); os.IsNotExist(err) {
		//if it doesnt exist then serve index.html
		http.ServeFile(w, r, indexFile)
	} else {
		// If the file exists, serve it
		http.FileServer(http.Dir(distDir)).ServeHTTP(w, r)
	}
}
