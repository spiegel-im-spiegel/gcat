package facade

import (
	"flag"
	"os"

	"github.com/spiegel-im-spiegel/gutil"
	"github.com/spiegel-im-spiegel/gcat/internal/gcat"
)

// Exit codes are int values that represent an exit code for a particular error.
const (
	ExitCodeOK    int = 0
	ExitCodeError int = 1 + iota
)

// Context is the context of facade
type Context struct {
	//Context for command-line.
	Cli *gutil.CliContext

	//Information of command
	CommandName, Version string
}

// Run invokes the Context with the given arguments.
func (cxt *Context) Run(args []string) int {
	var (
		out  string
		verFlag bool
	)

	// Define option flag parse
	flags := flag.NewFlagSet(cxt.CommandName, flag.ContinueOnError)
	flags.SetOutput(cxt.Cli.ErrorWriter)

	flags.StringVar(&out, "out", "", "output file")
	flags.BoolVar(&verFlag, "version", false, "Print version information and quit.")

	// Parse commandline flag
	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeError
	}

	// Show version
	if verFlag {
		cxt.Cli.OutputErrln(cxt.CommandName, "version", cxt.Version)
		return ExitCodeOK
	}

	// Parse commandline sub-arguments
	infiles := flags.Args()

	//OutputFile
	if len(out) > 0 {
		file, err := os.Create(out) //maybe file path
		if err != nil {
			cxt.Cli.OutputErrln(err.Error())
			return ExitCodeError
		}
		defer file.Close()
		cxt.Cli.Writer = file
	}

	//Create Context
	gcatCxt := gcat.Context {Cli: cxt.Cli}

	//Input File
	if len(infiles) == 0 {
		if err := gcatCxt.Catenate(); err != nil {
			cxt.Cli.OutputErrln(err.Error())
			return ExitCodeError
		}
	} else {
		for _, infile := range infiles {
			file, err := os.Open(infile) //maybe file path
			if err != nil {
				cxt.Cli.OutputErrln(err.Error())
				return ExitCodeError
			}
			defer file.Close()
			cxt.Cli.Reader = file
			if err := gcatCxt.Catenate(); err != nil {
				cxt.Cli.OutputErrln(err.Error())
				return ExitCodeError
			}
			file.Close()
		}
	}

	return ExitCodeOK
}
