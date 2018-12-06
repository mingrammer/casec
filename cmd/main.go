package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/mingrammer/cfmt"
	"github.com/sergi/go-diff/diffmatchpatch"
	"github.com/urfave/cli"
)

// CLI information
const (
	AppName   = "casec"
	Author    = "mingrammer"
	Email     = "mingrammer@gmail.com"
	Version   = "0.0.1"
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
			Usage: "Source `CASE` in {upper,lower,title,camel,pascal,snake,kebab,lisp}",
		},
		cli.StringFlag{
			Name:  "to, t",
			Usage: "Target `CASE` in {upper,lower,title,camel,pascal,snake,kebab,lisp}",
		},
		cli.StringSliceFlag{
			Name:  "ignore, i",
			Usage: "Ignore the text that matches the `PATTERNS`",
		},
		cli.BoolFlag{
			Name:  "text, x",
			Usage: "Read the input as raw text, not a file path",
		},
		cli.BoolFlag{
			Name:  "dry-run, d",
			Usage: "Show diff without actual converting",
		},
	}
	app.Action = func(ctx *cli.Context) error {
		var original string
		var converted string
		var err error
		var writer io.WriteCloser

		if len(ctx.Args()) == 0 {
			return cli.NewExitError(cfmt.Serror("There are no arguments"), errNoArguments)
		}
		if ctx.String("to") == "" {
			return cli.NewExitError(cfmt.Serror("You must specify target case"), errInvalidOptions)
		}
		input := ctx.Args().Get(0)

		from := ctx.String("from")
		to := ctx.String("to")
		ignore := ctx.StringSlice("ignore")

		if ctx.Bool("text") {
			original, converted, err = convertFromText(input, from, to, ignore)
			writer = os.Stdout
		} else {
			original, converted, err = convertFromFile(input, from, to, ignore)
			writer, _ = os.OpenFile(input, os.O_RDWR, 0644)
		}
		if err != nil {
			return err
		}

		if ctx.Bool("dry-run") {
			dmp := diffmatchpatch.New()
			diffs := dmp.DiffMain(original, converted, false)
			fmt.Println(dmp.DiffPrettyText(diffs))
		} else {
			writer.Write([]byte(converted))
			writer.Close()
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
