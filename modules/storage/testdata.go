package storage

import (
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/services/discovery"
	"net"
	"os"
	"path/filepath"
	"time"
)

var testDataPath = getTestDataPath()
var testSeedFilePath = "/TestFile/TestFile_Downloaded.txt"

// TestFilePath path to provide file resource
var TestFilePath = testDataPath + "/TestFile.txt"

// TestDownloadFilePath path to test file download
var TestDownloadFilePath = testDataPath + "/Download/DownloadedFile.txt"

// TestSeedPath path to download seed file
var TestSeedPath = testDataPath + "/Download/TestSeed"

// TestPeerAddr peer address
var TestPeerAddr = "https://localhost:6480"

// TestPieceLength piece length for TestFileInfo
var TestPieceLength int64 = 2048

// TestFileSize file size for TestFileInfo
var TestFileSize int64 = 6332

// TestModifyTime modify time for TestLocalFileInfo
var TestModifyTime = getTestModifyTime()

// TestPiece piece data for test
var TestPiece = getTestPiece()

// TestPieceIndex piece index to test download
var TestPieceIndex int64 = 0

// TestSeedWellKnown well known for TestSeedInfo
var TestSeedWellKnown = []string{
	TestPeerAddr,
}

// TestFileWellKnown well known for TestFileInfo
var TestFileWellKnown = []string{
	TestPeerAddr,
}

// TestPieceInfo piece info for TestPeerPieceInfo
var TestPieceInfo = []byte{
	1, 1, 1, 1,
}

// TestPieceHash piece hash at TestFileInfo.PieceHash
var TestPieceHash = TestFileInfo.PieceHash[TestPieceIndex]

// TestFileHash file hash of TestFilePath
var TestFileHash = "uNFYuGhEqWXuRalepnlm0ZNopsNDVI403aYZARGoNucZ5JWXmlmKaGw1hVCT--qLzskkPiNEkQvsuU5P1ADftA=="

// TestFileInfoHash file info hash of TestFileInfo
var TestFileInfoHash = "zfS2uU_kPwnt-U0NABEKRg5ICJbTOf7o25tGVSmlqYqMdZTWANf734X3mMiE8KOR3LXbvWLbRYQCzbKhRK-Ijg=="

// TestPieceHashList piece hash list of TestFileInfo
var TestPieceHashList = []string{
	"uM6epHato4tAUIfcw7jfJJpE-bbhkfuIQJ_n0NOWqRQ=",
	"b52XLEZ2ycisrNPtrs0i9yq8d-PhgvFx-vv7aJZjEGk=",
	"Fur0Wg5ENkjW2QXxAPMLIgwdO-1KjQ48DO6SkzK5lPQ=",
	"Bq6-sMIDeGXQ3T2kLDMLGXklJTUrEBxaoWB2yf75cug=",
}

// TestSeedInfoHash seed info hash of SeedInfo
var TestSeedInfoHash = "Ln3rQCfCE5P1MN-DtFtJOjhk52oyIrcTBEtem4pahouG_zFSah1HwPc3h8Rjp80PCY5TdegXHu5TEZE1wfngFg=="

// TestFileInfo for test
var TestFileInfo = FileInfo{
	Size:        TestFileSize,
	Hash:        TestFileHash,
	PieceLength: TestPieceLength,
	PieceHash:   TestPieceHashList,
	WellKnown:   TestFileWellKnown,
}

// TestLocalFileInfo for test
var TestLocalFileInfo = LocalFileInfo{
	FileInfo:   TestFileInfo,
	Path:       TestFilePath,
	LastModify: TestModifyTime,
}

// TestSeedInfo for test
var TestSeedInfo = SeedInfo{
	Title: "TestSeedInfoTitle",
	Files: []SeedFileItem{
		{
			Path:            testSeedFilePath,
			Size:            TestFileSize,
			Hash:            TestFileHash,
			RecFileInfoHash: TestFileInfoHash,
			RecPieceLength:  TestPieceLength,
		},
	},
	ExtraInfo: nil,
	WellKnown: TestSeedWellKnown,
}

// TestPeerInfo for test
var TestPeerInfo = discovery.PeerInfo{
	Address:  net.ParseIP("127.0.0.1"),
	Port:     6480,
	DNS:      []string{"dns"},
	LastSeen: time.Date(2020, 11, 19, 21, 30, 0, 0, time.Local),
}

// TestPeerPieceInfo for test
var TestPeerPieceInfo = PeerPieceInfo{
	TestPeerAddr: TestPieceInfo,
}

// TestPPInfoList for test
var TestPPInfoList = map[string]PeerPieceInfo{
	TestFileInfoHash: TestPeerPieceInfo,
}

func getTestPiece() []byte {
	piece, err := ReadFilePiece(TestLocalFileInfo, TestPieceIndex)
	if err != nil {
		logging.Log().Error(err.Error())
	}
	return piece
}

func getTestModifyTime() time.Time {
	t, err := time.Parse(TimeLayout, "2020-11-22 21:45:42")
	if err != nil {
		logging.Log().Error(err.Error())
	}
	return t
}

func getTestDataPath() (path string) {
	err := filepath.Walk("../..", func(p string, info os.FileInfo, err error) error {
		if info.IsDir() && filepath.Base(p) == "testdata" {
			path, err = filepath.Abs(p)
			if err != nil {
				panic(err)
			}
			logging.Log().Info("set testdata path to : " + path)
			return nil
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	return path
}
