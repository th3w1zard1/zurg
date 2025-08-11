package version

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GetFile serves version information as a file response
func GetFile(w http.ResponseWriter, r *http.Request) {
	versionInfo := map[string]string{
		"version":    Version,
		"git_commit": GitCommit,
		"built_at":   BuiltAt,
	}

	w.Header().Set("Content-Type", "application/json")
	
	if err := json.NewEncoder(w).Encode(versionInfo); err != nil {
		http.Error(w, fmt.Sprintf("Error encoding version info: %v", err), http.StatusInternalServerError)
		return
	}
}