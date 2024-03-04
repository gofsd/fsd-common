package types

import (
	"os"
	"path/filepath"
)

// add
var (
	DbName  = "fsd.bolt"
	LogName = "log.bolt"
	Port    = "32104"
)

var (
	HOME       = os.Getenv("HOME")
	FSD_FOLDER = ".fsd"
	DBFullPath = filepath.Join(HOME, FSD_FOLDER)
)
