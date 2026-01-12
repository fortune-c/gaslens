package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/analyze", handleAnalyze)
	fmt.Println("GasLens Web Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html><head><title>GasLens</title></head><body>
<h1>GasLens - Bytecode Analyzer</h1>
<form action="/analyze" method="post">
<textarea name="bytecode" placeholder="Enter bytecode..." rows="10" cols="80"></textarea><br>
<button type="submit">Analyze</button>
</form></body></html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleAnalyze(w http.ResponseWriter, r *http.Request) {
	bytecodeStr := r.FormValue("bytecode")
	bytecodeStr = strings.TrimSpace(strings.ReplaceAll(bytecodeStr, "0x", ""))
	
	code, err := hex.DecodeString(bytecodeStr)
	if err != nil {
		http.Error(w, "Invalid bytecode", 400)
		return
	}

	// Basic response - would need analyzer integration
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "analyzed", 
		"bytecode_length": len(code),
	})
}
