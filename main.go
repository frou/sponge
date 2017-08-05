package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/frou/stdext"
)

var flagQuiet = flag.Bool(
	"q", false, "Don't print non-fatal warnings")

func main() {
	flag.Parse()
	stdext.Exit(run())
}

func run() error {
	if flag.NArg() != 1 {
		return fmt.Errorf("Usage: %s path/to/wring/sponge/into", argv0())
	}
	var outPath = flag.Arg(0)
	_, err := os.Stat(outPath)
	outFileDidntExist := os.IsNotExist(err)

	buf, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(outPath, buf, stdext.OwnerWritableReg)
	if err == nil && outFileDidntExist && !*flagQuiet {
		fmt.Fprintf(
			os.Stderr,
			"Warning: Output file (%s) didn't exist before writing it - "+
				"unnecessary use of %s?\n", outPath, argv0())
	}
	return err
}

func argv0() string {
	return filepath.Base(os.Args[0])
}
