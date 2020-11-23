package transfer

import (
	"github.com/go-chi/chi"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	"net/http"
	"strconv"
)

// getFilePiece to transfer the specified localFileInfo piece
func getFilePiece(w http.ResponseWriter, r *http.Request) {
	fileInfoHash := chi.URLParam(r, "fileInfoHash")
	pieceIndex, err := strconv.Atoi(chi.URLParam(r, "pieceIndex"))
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}

	localFileInfo, err := GetLocalFileInfoByFileInfoHash(fileInfoHash)
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}

	piece, err := storage.ReadFilePiece(localFileInfo, int64(pieceIndex))
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}

	services.WriteRespDataAsOctetStream(w, piece)
}

// getFileInfoByFileHash returns FileInfo that matches the FileHash
func getFileInfoByFileHash(w http.ResponseWriter, r *http.Request) {
	fileHash := chi.URLParam(r, "fileHash")
	fileInfoList, err := GetFileInfoByFileHash(fileHash)
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}
	services.WriteRespDataAsJSON(w, fileInfoList)
}

// getFileInfoByFileInfoHash returns FileInfo that matches the FileInfoHash
func getFileInfoByFileInfoHash(w http.ResponseWriter, r *http.Request) {
	fileInfoHash := chi.URLParam(r, "fileInfoHash")
	fileInfo, err := GetFileInfoByFileInfoHash(fileInfoHash)
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}
	services.WriteRespDataAsJSON(w, fileInfo)
}

// getSeedInfo returns SeedInfo that matches the SeedInfoHash
func getSeedInfo(w http.ResponseWriter, r *http.Request) {
	seedInfoHash := chi.URLParam(r, "seedInfoHash")
	seedInfoList, err := GetSeedInfo(seedInfoHash)
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}
	services.WriteRespDataAsJSON(w, seedInfoList)
}
