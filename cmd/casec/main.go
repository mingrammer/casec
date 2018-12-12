package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/mingrammer/cfmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/urfave/cli"
)

// CLI information
const (
	AppName   = "casec"
	Author    = "mingrammer"
	Email     = "mingrammer@gmail.com"
	Version   = "0.1.2"
	Usage     = "A text case converter"
	UsageText = "casec [OPTIONS/FLAGS] [-x <text>] [<path>]"
	ArgsUsage = "[path and text]"
	// ArgsUsage = "<path> The file path that you want to convert case of contents"
)

// Custom exit codes
const (
	errNoArguments = iota + 10
	errInvalidOptions
)

func main() {
	app := cli.NewApp()
	app.Name = AppName
	app.Author = Author
	app.Email = Email
	app.Version = Version
	app.Usage = Usage
	app.UsageText = UsageText
	app.ArgsUsage = ArgsUsage
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "from, f",
			Usage: "source `CASE` in {upper,lower,title,camel,pascal,snake,kebab,lisp}",
		},
		cli.StringFlag{
			Name:  "to, t",
			Usage: "target `CASE` in {upper,lower,title,camel,pascal,snake,kebab,lisp}",
		},
		cli.StringFlag{
			Name:  "lines, l",
			Usage: "set lines `M:N` to be converted. empty means first or last line",
		},
		cli.StringSliceFlag{
			Name:  "ignore, i",
			Usage: "ignore the text that matches the `PATTERNS`",
		},
		cli.BoolFlag{
			Name:  "text, x",
			Usage: "read the input as raw text, not a file path",
		},
		cli.BoolFlag{
			Name:  "dry-run, n",
			Usage: "show how would have been converted",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		var err error
		var orig string
		var conv string
		var writer io.WriteCloser

		var src string
		var tgt string
		var start int
		var end int
		var ignoreRe *regexp.Regexp
		if len(ctx.Args()) == 0 {
			return errors.New(cfmt.Serror("There are no arguments"))
		}
		if src, err = parseSource(ctx.String("from")); err != nil {
			return err
		}
		if tgt, err = parseTarget(ctx.String("to")); err != nil {
			return err
		}
		if start, end, err = parseLines(ctx.String("lines")); err != nil {
			return err
		}
		if ignoreRe, err = parseIgnore(ctx.StringSlice("ignore")); err != nil {
			return err
		}
		input := ctx.Args().Get(0)

		if ctx.Bool("text") {
			orig, conv = convertFromText(input, src, tgt, start, end, ignoreRe)
			writer = os.Stdout
		} else {
			orig, conv, err = convertFromFile(input, src, tgt, start, end, ignoreRe)
			if err != nil {
				return err
			}
			writer, _ = os.OpenFile(input, os.O_RDWR, 0644)
		}

		if ctx.Bool("dry-run") {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(orig, conv, false)
			fmt.Println(dmp.DiffPrettyText(diffs))
		} else {
			writer.Write([]byte(conv))
			writer.Close()
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err.Error())
	}
}
