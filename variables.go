package types

import (
	"os"
	"path/filepath"
)

// add
var (
	DBName  = "fsd.bolt"
	LogName = "log.bolt"
	Port    = "32104"
)

var (
	HOME                    = os.Getenv("HOME")
	FSD_FOLDER              = ".fsd"
	FSD_FULL_FOLDER_NAME    = filepath.Join(HOME, FSD_FOLDER)
	DBFullName, LogFullName = filepath.Join(FSD_FULL_FOLDER_NAME, DBName), filepath.Join(FSD_FULL_FOLDER_NAME, LogName)
	Output                  = STRING
)
