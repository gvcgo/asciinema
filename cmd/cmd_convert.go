package cmd

import (
	"os"
	"strings"

	"github.com/gvcgo/goutils/pkgs/gtea/gprint"
	"github.com/gvcgo/goutils/pkgs/gutils"
)

func isAggInstalled() bool {
	_, err := gutils.ExecuteSysCommand(true, "", "agg", "--help")
	return err == nil
}

func (r *Runner) ConvertToGif(fPath, outFilePath string) (err error) {
	if !isAggInstalled() {
		gprint.PrintError("agg<https://github.com/asciinema/agg> is not installed.")
		gprint.PrintInfo("Please use vm<https://github.com/gvcgo/version-manager> to install agg.")
		return
	}
	if !strings.HasSuffix(outFilePath, ".gif") {
		outFilePath += ".gif"
	}
	homeDir, _ := os.UserHomeDir()
	_, err = gutils.ExecuteSysCommand(false, homeDir,
		"agg", fPath, outFilePath,
	)
	return
}
