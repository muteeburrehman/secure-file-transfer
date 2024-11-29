package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"secure-file-transfer/crypto"
)

var aesKey = []byte("32-byte-long-key-for-aes-encrypt")

func Upload(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File upload failed", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Encrypt file
	encryptedFile, err := crypto.Encrypt(fileBytes, aesKey)
	if err != nil {
		http.Error(w, "File encryption failed", http.StatusInternalServerError)
		return
	}

	// Sign the encrypted file
	signature, err := crypto.Sign([]byte(encryptedFile))
	if err != nil {
		http.Error(w, "Signature generation failed", http.StatusInternalServerError)
		return
	}

	// Save the encrypted file
	err = os.WriteFile("./static/uploads/encrypted_file.txt", []byte(encryptedFile), 0644)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "File uploaded successfully",
		"signature": string(signature),
	})
}
