package transfer

import (
	"github.com/go-chi/chi"
	"net/http"
)

// getFilePiece to transfer the specified file piece
func getFilePiece(w http.ResponseWriter, r *http.Request) {

}

// getFileInfoByFileHash returns FileInfo that matches the FileHash
func getFileInfoByFileHash(w http.ResponseWriter, r *http.Request) {
	chi.URLParam(r, "")
}

// getFileInfoByFileInfoHash returns FileInfo that matches the FileInfoHash
func getFileInfoByFileInfoHash(w http.ResponseWriter, r *http.Request) {

}

// getSeedInfo returns SeedInfo that matches the SeedInfoHash
func getSeedInfo(w http.ResponseWriter, r *http.Request) {

}
