package main

import (
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

func ProcessFile(pathname string) error {
	f, err := os.Open(pathname)
	if err != nil {
		return errors.Wrap(err, "unable to open file")
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
	return nil
}

func main() {
	pathname := "file_may_not_exist"
	err := ProcessFile(pathname)
	if err != nil {
		log.Printf("ProcessFile(%s) failed. \nerror: %+v\n", pathname, err)
	}
}
