package status

import (
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/modules/storage"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/transfer"
	"go.uber.org/zap"
	"os"
	"path"
	"strconv"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	logging.Log().Info("init db...")
	dbName := "status_testing_database"
	testPort := 6481

	err := database.OpenDB(dbName)
	if err != nil {
		logging.Log().Error("db err", zap.Error(err))
		panic(err)
	}
	go services.StartServer(testPort)
	time.Sleep(time.Second * 3)

	storage.TestPeerAddr = "https://localhost:" + strconv.Itoa(testPort)
	storage.TestPeerInfo.Port = testPort
	storage.TestPeerPieceInfo = storage.PeerPieceInfo{
		storage.TestPeerAddr: storage.TestPieceInfo,
	}
	storage.TestPPInfoList = map[string]storage.PeerPieceInfo{
		storage.TestFileInfoHash: storage.TestPeerPieceInfo,
	}

	SaveTestData()
	transfer.SaveTestData()

	m.Run()

	database.CloseDBEngine()
	dbPath := path.Join(utils.AppDataDir(), dbName)
	utils.RemoveFilePath(dbPath)
	os.Exit(0)
}
