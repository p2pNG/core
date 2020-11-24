package transfer

import (
	"github.com/go-chi/chi"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	"net/http"
	"strconv"
)

// serverGetFilePiece to transfer the specified localFileInfo piece
func serverGetFilePiece(w http.ResponseWriter, r *http.Request) {
	fileInfoHash := chi.URLParam(r, "fileInfoHash")
	pieceIndex, err := strconv.Atoi(chi.URLParam(r, "pieceIndex"))
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}

	localFileInfo, err := getLocalFileInfoByFileInfoHash(fileInfoHash)
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

// serverGetFileInfoByFileHash returns FileInfo that matches the FileHash
func serverGetFileInfoByFileHash(w http.ResponseWriter, r *http.Request) {
	fileHash := chi.URLParam(r, "fileHash")
	fileInfoList, err := getFileInfoByFileHash(fileHash)
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}
	services.WriteRespDataAsJSON(w, fileInfoList)
}

// serverGetFileInfoByFileInfoHash returns FileInfo that matches the FileInfoHash
func serverGetFileInfoByFileInfoHash(w http.ResponseWriter, r *http.Request) {
	fileInfoHash := chi.URLParam(r, "fileInfoHash")
	fileInfo, err := getFileInfoByFileInfoHash(fileInfoHash)
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}
	services.WriteRespDataAsJSON(w, fileInfo)
}

// serverGetSeedInfo returns SeedInfo that matches the SeedInfoHash
func serverGetSeedInfo(w http.ResponseWriter, r *http.Request) {
	seedInfoHash := chi.URLParam(r, "seedInfoHash")
	seedInfoList, err := getSeedInfoBySeedInfoHash(seedInfoHash)
	if err != nil {
		services.WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}
	services.WriteRespDataAsJSON(w, seedInfoList)
}
