package storage

import (
	"github.com/p2pNG/core/internal/logging"
	"github.com/p2pNG/core/services/discovery"
	"net"
	"time"
)

// path
var testDataPath = "F://GoModProjects/p2pNG/core/test/testdata"
var TestFilePath = testDataPath + "/TestFile.txt"
var TestSeedPath string = testDataPath + "/Download/TestSeed"
var TestSeedFilePath = "/TestFile/TestFile_Downloaded.txt"

// attribute
var TestPeerAddr string = "https://localhost:6480"
var TestPieceLength int64 = 2048
var TestFileSize int64 = 6332
var TestModifyTime time.Time = getTestModifyTime()
var TestPiece []byte = getTestPiece()
var TestPieceIndex int64 = 0
var TestSeedWellKnown = []string{
	TestPeerAddr,
}
var TestFileWellKnown = []string{
	TestPeerAddr,
}
var TestPieceInfo = []byte{
	1, 1, 1, 1,
}

// hash
var TestPieceHash = TestFileInfo.PieceHash[TestPieceIndex]
var TestFileHash = "uNFYuGhEqWXuRalepnlm0ZNopsNDVI403aYZARGoNucZ5JWXmlmKaGw1hVCT--qLzskkPiNEkQvsuU5P1ADftA=="
var TestFileInfoHash string = "zfS2uU_kPwnt-U0NABEKRg5ICJbTOf7o25tGVSmlqYqMdZTWANf734X3mMiE8KOR3LXbvWLbRYQCzbKhRK-Ijg=="
var TestPieceHashList []string = []string{
	"uM6epHato4tAUIfcw7jfJJpE-bbhkfuIQJ_n0NOWqRQ=",
	"b52XLEZ2ycisrNPtrs0i9yq8d-PhgvFx-vv7aJZjEGk=",
	"Fur0Wg5ENkjW2QXxAPMLIgwdO-1KjQ48DO6SkzK5lPQ=",
	"Bq6-sMIDeGXQ3T2kLDMLGXklJTUrEBxaoWB2yf75cug=",
}
var TestSeedInfoHash = "Ln3rQCfCE5P1MN-DtFtJOjhk52oyIrcTBEtem4pahouG_zFSah1HwPc3h8Rjp80PCY5TdegXHu5TEZE1wfngFg=="

// struct
var TestFileInfo = FileInfo{
	Size:        TestFileSize,
	Hash:        TestFileHash,
	PieceLength: TestPieceLength,
	PieceHash:   TestPieceHashList,
	WellKnown:   TestFileWellKnown,
}
var TestLocalFileInfo = LocalFileInfo{
	FileInfo:   TestFileInfo,
	Path:       TestFilePath,
	LastModify: TestModifyTime,
}
var TestSeedInfo = SeedInfo{
	Title: "TestSeedInfoTitle",
	Files: []SeedFileItem{
		{
			Path:            TestSeedFilePath,
			Size:            TestFileSize,
			Hash:            TestFileHash,
			RecFileInfoHash: TestFileInfoHash,
			RecPieceLength:  TestPieceLength,
		},
	},
	ExtraInfo: nil,
	WellKnown: TestSeedWellKnown,
}
var TestPeerInfo = discovery.PeerInfo{
	Address:  net.ParseIP("127.0.0.1"),
	Port:     6060,
	DNS:      []string{"dns"},
	LastSeen: time.Date(2020, 11, 19, 21, 30, 0, 0, time.Local),
}

var TestPeerPieceInfo = PeerPieceInfo{
	TestPeerAddr: TestPieceInfo,
}

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
