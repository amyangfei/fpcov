package hello

import (
	"os"
)

var (
	// OsExit is function placeholder for os.Exit
	OsExit func(int)
)

func init() {
	OsExit = os.Exit
}
