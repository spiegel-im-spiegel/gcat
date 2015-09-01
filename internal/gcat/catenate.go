package gcat

import (
	"io"

	"github.com/spiegel-im-spiegel/gutil"
)

// Context of catenate
type Context struct {
	//Context for command-line.
	Cli *gutil.CliUi
}

//Catenation (raw data)
func (cxt *Context) Catenate() error {
	_, err := io.Copy(cxt.Cli.Writer, cxt.Cli.Reader)
	return err
}
