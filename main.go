package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func ProcessFile(pathname string) error {
	_, err := os.Open(pathname)
	if err != nil {
		return errors.Wrap(err, "unable to open file")
	}
	return nil
}

func main() {
	err := ProcessFile("file_may_not_exist")
	if err != nil {
		switch errors.Cause(err).(type) {
		case *os.PathError:
			fmt.Printf("path error: %+v\n", err)
			return
		default:
			fmt.Printf("a different error: %+v\n", err)
			return
		}
	}
}
