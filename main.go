import (
	"fmt"
	"io"
	"net/http"
	"os/exec"
)
func handleUpload(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form
	err := r.ParseMultipartForm(32 << 20) // Limit file size to 32MB
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	// Get the uploaded file
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error getting file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Print the file content to the default printer
	cmd := exec.Command("cmd", "/c", "type", handler.Filename)
	cmd.Stdin = file
	err = cmd.Run()
	if err != nil {
		http.Error(w, "Error printing file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded and printed successfully!")
}

func main() {
	for {
		http.HandleFunc("/upload", handleUpload)
		fmt.Println("Server listening on port 8080")
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Println("Error starting server:", err)
			// Add a delay before restarting
			time.Sleep(5 * time.Second)
		}
	}
}
