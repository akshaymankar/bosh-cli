package main

import (
	"fmt"
	"os"

	"github.com/akshaymankar/int-yaml/cmd"
	boshui "github.com/akshaymankar/int-yaml/ui"
	boshlogger "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"
	goflags "github.com/jessevdk/go-flags"
)

func main() {
	logger := boshlogger.NewLogger(boshlogger.LevelInfo)

	opts := cmd.InterpolateOpts{
		Args: cmd.InterpolateArgs{
			Manifest: cmd.FileBytesArg{
				FS: boshsys.NewOsFileSystemWithStrictTempRoot(logger),
			},
		},
	}

	_, err := goflags.ParseArgs(&opts, os.Args[1:])
	if err != nil {
		fmt.Errorf(err.Error())
		os.Exit(1)
	}

	command := cmd.NewInterpolateCmd(boshui.NewConsoleUI(logger))
	err = command.Run(opts)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
