package services

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/p2pNG/core"
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/listener"
	"github.com/p2pNG/core/modules/request"
	"github.com/p2pNG/core/services/discovery"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

type contextType int

// StatusContext request context for status service
// TransferContext request context for transfer service
// P2PNGCode a custom header to describe error code
// P2PNGMsg a custom header to describe error msg
// NoPermissions lack of permissions
const (
	// request context
	StatusContext   contextType = iota
	TransferContext contextType = iota
	// custom header
	P2PNGCode string = "p2pNG-code"
	P2PNGMsg  string = "p2pNG-msg"
	// error code
	NoPermissions string = "3001"
)

// SeedHashToSeedDB 		Key=SeedHash,Value=SeedInfo
// FileInfoHashToFileDB 	Key=FileInfoHash,Value=FileInfo
// FileHashToFileDB 		Key=FileHash,Value=FileInfo
// FileInfoHashToLocalFileDB 	Key=FileInfoHash,Value=LocalFileInfo
// SeedHashToPeerDB 	Key=SeedHash,Value=PeerInfo
// FileInfoHashToPeerDB Key=FileInfoHash,Value=PeerInfo
// FileHashToPeerDB 	Key=FileHash,Value=PeerInfo
// FileInfoHashToPeerPieceDB 	Key=FileInfoHash,Value=PeerPieceInfo
const (
	SeedHashToSeedDB          = "SeedInfoHash-SeedInfo"
	FileInfoHashToFileDB      = "FileInfoHash-FileInfo"
	FileHashToFileDB          = "FileHash-FileInfo"
	FileInfoHashToLocalFileDB = "FileInfoHash-LocalFileInfo"
	SeedHashToPeerDB          = "SeedInfoHash-PeerInfo"
	FileInfoHashToPeerDB      = "FileInfoHash-PeerInfo"
	FileHashToPeerDB          = "FileHash-PeerInfo"
	FileInfoHashToPeerPieceDB = "FileInfoHash-PeerPieceInfo"
)

// DataBaseBuckets database buckets to be init
var DataBaseBuckets = []string{
	//seed
	SeedHashToSeedDB,
	//file
	FileInfoHashToFileDB,
	FileHashToFileDB,
	FileInfoHashToLocalFileDB,
	//peer
	SeedHashToPeerDB,
	FileInfoHashToPeerDB,
	FileHashToPeerDB,
	//peer piece
	FileInfoHashToPeerPieceDB}

// GetHTTPClient returns a http client
func GetHTTPClient() (client *http.Client, err error) {
	client, _, err = request.GetDefaultHTTPClient()
	return
}

// WriteRespDataAsJSON convert data into json format and response to client
func WriteRespDataAsJSON(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(&data)
	if err != nil {
		WriteErrorToResp(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		logging.Log().Error("fail response", zap.Error(err))
	}
}

// ReadJSONBody read responseBody and use json to unmarshal
func ReadJSONBody(resp *http.Response, data interface{}) error {
	//length, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	//logging.Log().Warn("head "+resp.Header.Get("Content-Length"))
	//if err != nil {
	//	return err
	//}
	//body := make([]byte, length)
	//_, err = resp.Body.Read(body)
	//if err != nil {
	//	return err
	//}
	// todo: change to resp.Body.Read(body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, data)
}

// WriteRespDataAsOctetStream write response data with Content-Type = application/octet-stream header
func WriteRespDataAsOctetStream(w http.ResponseWriter, data []byte) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(data)
	if err != nil {
		logging.Log().Error("fail response", zap.Error(err))
	}
}

// WriteErrorToResp response error msg to client
func WriteErrorToResp(w http.ResponseWriter, err error, statusCode int) {
	logging.Log().Warn("fail request:", zap.Error(err))
	w.Header().Set(P2PNGMsg, err.Error())
	w.WriteHeader(statusCode)
}

// GetAllKeyFromBucket get all keys from specified bucket
func GetAllKeyFromBucket(bucket string) (keys []string, err error) {
	db, err := database.GetDBEngine()
	if err != nil {
		return nil, err
	}
	err = db.View(func(tx *bolt.Tx) error {
		buk := tx.Bucket([]byte(bucket))
		if buk == nil {
			return errors.New("database error : bucket [" + bucket + "] does not exist")
		}
		err = buk.ForEach(func(k, v []byte) error {
			keys = append(keys, string(k))
			return nil
		})
		return err
	})
	return
}

// PeerInfoToStringAddr returns http addr format from discovery.PeerInfo
func PeerInfoToStringAddr(info discovery.PeerInfo) string {
	return "https://" + info.Address.String() + ":" + strconv.Itoa(info.Port)
}

// StartServer start server at port
func StartServer(port int) {
	logging.Log().Info("start server at " + strconv.Itoa(port) + " ...")
	db, err := database.GetDBEngine()
	defer database.CloseDBEngine()
	if err != nil {
		logging.Log().Error("db err", zap.Error(err))
		panic(err)
	}
	err = database.InitBuckets(db, DataBaseBuckets)
	if err != nil {
		logging.Log().Error("db err", zap.Error(err))
		panic(err)
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	plugins := core.GetRouterPluginRegistry()
	//todo: Replace with real config data
	x := "{\"BuildName\":\"Hello World\",\"LocalDiscoveryPort\":6553}"
	for _, plugin := range plugins {
		info := plugin.PluginInfo()
		// todo: delete
		logging.Log().Info("loading router plugin",
			zap.String("plugin", info.Name), zap.String("version", info.Version))
		//todo: Use Real config
		err := json.Unmarshal([]byte(x), plugin.Config())
		if err != nil {
			logging.Log().Fatal("load config for plugin failed", zap.Error(err), zap.String("plugin", info.Name))
		}

		err = plugin.Init()
		if err != nil {
			logging.Log().Fatal("init for plugin failed", zap.Error(err), zap.String("plugin", info.Name))
		}

		router.Mount(info.Prefix, plugin.GetRouter())

		if err = database.InitBuckets(db, info.Buckets); err != nil {
			logging.Log().Fatal("init buckets in database failed", zap.Error(err), zap.String("plugin", info.Name))
		}
	}
	go func() {
		err = listener.ListenBoth(router, ":"+strconv.Itoa(port))
		if err != nil {
			logging.Log().Fatal("start http service failed", zap.Error(err))
		}
	}()
	{
		osSignals := make(chan os.Signal, 1)
		signal.Notify(osSignals, os.Interrupt, syscall.SIGTERM)
		<-osSignals
	}
}
