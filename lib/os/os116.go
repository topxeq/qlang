package os

import (
	"os"
)

func init() {
	Exports["ReadDir"] = os.ReadDir
	Exports["ReadFile"] = os.ReadFile
	Exports["MkdirTemp"] = os.MkdirTemp
	Exports["CreateTemp"] = os.CreateTemp
	Exports["WriteFile"] = os.WriteFile

}
