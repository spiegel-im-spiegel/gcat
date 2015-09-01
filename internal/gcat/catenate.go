package gcat

import (
	"github.com/spiegel-im-spiegel/gutil"
)

// Context of catenate
type Context struct {
	//Context for command-line.
	Cli *gutil.CliUi
}

//Catenation (raw data)
func (cxt *Context) Catenate() error {
	reader, err := cxt.Cli.NewReader()
	if err != nil {
		return err
	}
	_, err = reader.WriteTo(cxt.Cli.Writer)
	return err
}
