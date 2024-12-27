package common

import (
	"io"
	"net/http"
)

func ForwardResponse(url string, w http.ResponseWriter) {
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	// Read the entire response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the content type from the response and set it in the response writer
	contentType := resp.Header.Get("Content-Type")
	w.Header().Set("Content-Type", contentType)

	// Write the body to the response writer
	w.Write(body)
}
