package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/frou/stdext"
)

var flagQuiet = flag.Bool(
	"q", false, "don't print non-fatal warnings")

func main() {
	stdext.Exit(run())
}

func run() error {
	stdext.SetPreFlagsUsageMessage(
		fmt.Sprintf("%s soaks up the entirety of standard input and then "+
			"wrings it into an output file.", stdext.ExecutableBasename()),
		false,
		"path/to/wring.into")

	if err := stdext.ParseFlagsExpectingNArgs(1); err != nil {
		return err
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
			"warning: output file '%s' didn't exist before writing it - "+
				"unnecessary use of %s?\n",
			outPath, stdext.ExecutableBasename())
	}
	return err
}
