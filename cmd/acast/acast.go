package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gvcgo/asciinema/cmd"
	"github.com/gvcgo/asciinema/util"
	"github.com/gvcgo/goutils/pkgs/gtea/gprint"
	"github.com/gvcgo/goutils/pkgs/gutils"
	"github.com/spf13/cobra"
)

var (
	GitTag  string
	GitHash string
	GroupID string = "asciinema"
)

func SetWorkDir() {
	homeDir, _ := os.UserHomeDir()
	gvcDir := filepath.Join(homeDir, ".gvc")
	var workdir string
	if ok, _ := gutils.PathIsExist(gvcDir); ok {
		workdir = filepath.Join(gvcDir, "asciinema")
	} else {
		workdir = filepath.Join(homeDir, ".config", "asciinema")
	}
	os.MkdirAll(workdir, os.ModePerm)
	os.Setenv(util.DefaultHomeEnv, workdir)
}

func getName(base string) string {
	if base == "" {
		return base
	}
	return strings.Split(base, ".")[0]
}

func handleFilePath(fpath string) (title, result string) {
	cwd, _ := os.Getwd()
	if fpath == "" {
		return "default_cast", filepath.Join(cwd, "default.cast")
	}
	base := filepath.Base(fpath)
	if base == fpath {
		return getName(base), filepath.Join(cwd, base)
	}
	return getName(base), fpath
}

type Cli struct {
	rootCmd *cobra.Command
	cmd     *cmd.Runner
}

func NewCli() *Cli {
	SetWorkDir()
	c := &Cli{
		rootCmd: &cobra.Command{
			Short: "asciinema terminal recorder.",
			Long:  "acast <Command> <SubCommand> --flags args...",
		},
		cmd: cmd.New(),
	}
	c.rootCmd.AddGroup(&cobra.Group{ID: GroupID, Title: "Command list: "})
	c.initiate()
	return c
}

func (c *Cli) initiate() {
	if c.rootCmd == nil || c.cmd == nil {
		return
	}

	// Auth to aciinema.org
	auth := &cobra.Command{
		Use:     "auth",
		Aliases: []string{"a"},
		GroupID: GroupID,
		Short:   "Authrization to asciinema.org.",
		Run: func(cc *cobra.Command, args []string) {
			authUrl, info := c.cmd.Auth()
			gprint.PrintInfo(info)
			var cmd *exec.Cmd
			if runtime.GOOS == gutils.Darwin {
				cmd = exec.Command("open", authUrl)
			} else if runtime.GOOS == gutils.Linux {
				cmd = exec.Command("x-www-browser", authUrl)
			} else if runtime.GOOS == gutils.Windows {
				cmd = exec.Command("cmd", "/c", "start", authUrl)
			} else {
				gprint.PrintError("unsupported os")
			}

			if err := cmd.Run(); err != nil {
				gprint.PrintError("auth failed: %+v", err)
			}
		},
	}
	c.rootCmd.AddCommand(auth)

	// Record.
	record := &cobra.Command{
		Use:     "record",
		Aliases: []string{"r"},
		Short:   "Creates a record.",
		Long:    "Example: acast record <xxx.cast>",
		Run: func(cc *cobra.Command, args []string) {
			if len(args) == 0 {
				cc.Help()
				return
			}
			c.cmd.Title, c.cmd.FilePath = handleFilePath(args[0])
			err := c.cmd.Rec()
			if err != nil {
				gprint.PrintError("record failed: %+v", err)
			}
		},
	}
	c.rootCmd.AddCommand(record)
}
