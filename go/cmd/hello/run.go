package hello

import (
	"fmt"
	"os"
)

func Run() error {
	fmt.Fprintln(os.Stdout, "hi2", Message)
	return nil
}
