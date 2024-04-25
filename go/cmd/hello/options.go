package hello

import (
	"github.com/spf13/pflag"
)

var Message string

func AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&Message, "message", "m", "hello world!", "output message")
}

func Complete() error {
	return nil
}

func Validate() error {
	return nil
}
