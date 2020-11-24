package transfer

import (
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/internal/utils"
	"github.com/p2pNG/core/modules/database"
	"github.com/p2pNG/core/services"
	"github.com/p2pNG/core/services/status"
	"go.uber.org/zap"
	"os"
	"path"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	logging.Log().Info("init db...")
	dbName := "testing_database"
	err := database.OpenDB(dbName)
	if err != nil {
		logging.Log().Error("db err", zap.Error(err))
		panic(err)
	}
	go services.StartServer(6480)
	time.Sleep(time.Second * 3)
	status.SaveTestData()
	SaveTestData()
	m.Run()
	database.CloseDBEngine()
	dbPath := path.Join(utils.AppDataDir(), dbName)
	utils.RemoveFilePath(dbPath)
	os.Exit(0)
}
