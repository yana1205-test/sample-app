package hello

import (
	"fmt"
	"os"
)

func Run() error {
	fmt.Fprintln(os.Stdout, Message)
	return nil
}
