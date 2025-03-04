package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	http.HandleFunc("/", serveHTML)
	http.HandleFunc("/upload", uploadFile)
	http.HandleFunc("/run-python", runPython)
	http.HandleFunc("/converted/", serveConvertedFiles)

	fmt.Println("Server started at http://<IP-Address>:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

// serveHTML serves the index.html file and clears the uploads folder
func serveHTML(w http.ResponseWriter, r *http.Request) {
	// Clear the uploads and converted folders
	clearFolder("uploads")
	clearFolder("converted")

	http.ServeFile(w, r, "index.html")
}

// uploadFile handles file uploads
func uploadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// 1GB max upload size
	if err := r.ParseMultipartForm(1 << 30); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	dst, err := os.Create(filepath.Join("uploads", "up_file"))
	if err != nil {
		http.Error(w, "Failed to create file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully!")
}

// runPython executes the Python script ca.py and returns the content of the generated file
func runPython(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Execute the Python script
	cmd := exec.Command("python3", "ca.py")
	_, err := cmd.CombinedOutput()
	if err != nil {
		http.Error(w, "Failed to execute Python script", http.StatusInternalServerError)
		return
	}

	// Read the content of the generated file (cv.srt)
	fileContent, err := os.ReadFile("converted/cv.srt")
	if err != nil {
		http.Error(w, "Failed to read generated file", http.StatusInternalServerError)
		return
	}

	// Return the result as JSON
	response := map[string]interface{}{
		"success":     true,
		"fileContent": string(fileContent),
	}
	json.NewEncoder(w).Encode(response)
}

// serveConvertedFiles serves files from the converted folder
func serveConvertedFiles(w http.ResponseWriter, r *http.Request) {
	filePath := filepath.Join("converted", filepath.Base(r.URL.Path))
	http.ServeFile(w, r, filePath)
}

// clearFolder deletes all files in the specified folder
func clearFolder(folder string) {
	files, err := os.ReadDir(folder)
	if err != nil {
		fmt.Println("Error reading folder:", err)
		return
	}

	for _, file := range files {
		filePath := filepath.Join(folder, file.Name())
		if err := os.Remove(filePath); err != nil {
			fmt.Println("Error deleting file:", err)
		}
	}
}
