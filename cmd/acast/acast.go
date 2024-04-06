package main

import (
	"fmt"
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
		GroupID: GroupID,
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

	// Play.
	play := &cobra.Command{
		Use:     "play",
		Aliases: []string{"p"},
		GroupID: GroupID,
		Short:   "Plays a record.",
		Long:    "Example: acast play <xxx.cast>",
		Run: func(cc *cobra.Command, args []string) {
			if len(args) == 0 {
				cc.Help()
				return
			}
			c.cmd.Title, c.cmd.FilePath = handleFilePath(args[0])
			c.cmd.Play()
		},
	}
	c.rootCmd.AddCommand(play)

	// Upload.
	upload := &cobra.Command{
		Use:     "upload",
		Aliases: []string{"u"},
		GroupID: GroupID,
		Short:   "Uploads a record file to asciinema.org.",
		Long:    "Example: acast upload <xxx.cast>",
		Run: func(cc *cobra.Command, args []string) {
			if len(args) == 0 {
				cc.Help()
				return
			}
			c.cmd.Title, c.cmd.FilePath = handleFilePath(args[0])
			if respStr, err := c.cmd.Upload(); err == nil {
				gprint.PrintInfo(respStr)
			} else {
				gprint.PrintError("upload failed: %+v", err)
			}
		},
	}
	c.rootCmd.AddCommand(upload)

	// Convert to gif.
	convert := &cobra.Command{
		Use:     "convert-to-gif",
		Aliases: []string{"cg"},
		GroupID: GroupID,
		Short:   "Converts an asciinema cast to gif.",
		Long:    "Example: acast cg <input.cast> <output.gif>",
		Run: func(cc *cobra.Command, args []string) {
			if len(args) < 2 {
				cc.Help()
				return
			}
			if err := c.cmd.ConvertToGif(args[0], args[1]); err != nil {
				gprint.PrintError("convert failed: %+v", err)
			}
		},
	}
	c.rootCmd.AddCommand(convert)

	// Cut.
	cut := &cobra.Command{
		Use:     "cut",
		Aliases: []string{"c"},
		GroupID: GroupID,
		Short:   "Removes a certain range of time frames.",
		Long:    "Example: acast cut --start=1.0 --end=5.0 <in.cast> <out.cast>",
		Run: func(cc *cobra.Command, args []string) {
			start, _ := cc.Flags().GetFloat64("start")
			end, _ := cc.Flags().GetFloat64("end")
			if len(args) < 2 || end <= start {
				cc.Help()
				return
			}
			c.cmd.Cut(args[0], args[1], start, end)
		},
	}
	cut.Flags().Float64P("start", "s", 0, "start time")
	cut.Flags().Float64P("end", "e", 0, "end time")
	c.rootCmd.AddCommand(cut)

	// Speed.
	speed := &cobra.Command{
		Use:     "speed",
		Aliases: []string{"s"},
		GroupID: GroupID,
		Short:   "Updates the cast speed by a certain factor.",
		Long:    "Example: acast speed --factor=0.7 --start=1.0 --end=5.0 <in.cast> <out.cast>",
		Run: func(cc *cobra.Command, args []string) {
			factor, _ := cc.Flags().GetFloat64("factor")
			start, _ := cc.Flags().GetFloat64("start")
			end, _ := cc.Flags().GetFloat64("end")
			if len(args) < 2 || end <= start || factor <= 0 {
				cc.Help()
				return
			}
			c.cmd.Speed(args[0], args[1], factor, start, end)
		},
	}
	speed.Flags().Float64P("factor", "f", 0.7, "speed factor")
	speed.Flags().Float64P("start", "s", 0, "start time")
	speed.Flags().Float64P("end", "e", 0, "end time")
	c.rootCmd.AddCommand(speed)

	// Quantize.
	quantize := &cobra.Command{
		Use:     "quantize",
		Aliases: []string{"q"},
		GroupID: GroupID,
		Short:   "Updates the cast delays following quantization ranges.",
		Long:    "Example: acast quantize --ranges=1.0,5.0 <in.cast> <out.cast>",
		Run: func(cc *cobra.Command, args []string) {
			ranges, _ := cc.Flags().GetStringArray("ranges")
			if len(ranges) == 0 || len(args) < 2 {
				cc.Help()
				return
			}
			c.cmd.Quantize(args[0], args[1], ranges)
		},
	}
	quantize.Flags().StringArrayP("ranges", "r", []string{}, "quantization ranges")
	c.rootCmd.AddCommand(quantize)

	version := &cobra.Command{
		Use:     "version",
		Aliases: []string{"v"},
		GroupID: GroupID,
		Short:   "Shows version info of acast.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(GitHash) > 7 {
				GitHash = GitHash[:7]
			}
			fmt.Println(gprint.CyanStr("%s(%s)", GitTag, GitHash))
		},
	}
	c.rootCmd.AddCommand(version)
}

func (c *Cli) Run() {
	if c.rootCmd == nil {
		return
	}
	if err := c.rootCmd.Execute(); err != nil {
		gprint.PrintError("%+v", err)
	}
}

func main() {
	cli := NewCli()
	cli.Run()
}
