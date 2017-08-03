package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/frou/stdext"
)

func main() {
	stdext.Exit(run())
}

func run() error {
	if len(os.Args) != 2 {
		return fmt.Errorf(
			"usage: %s path/to/wring/sponge/into",
			filepath.Base(os.Args[0]))
	}
	var outPath = os.Args[1]

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outPath, buf, stdext.OwnerWritableReg)
}
