package main

import (
	"fmt"
	"io"
	"os"

	"github.com/pkg/errors"
)

func processFile(path string) error {
	if path == "" {
		return errors.New("path is empty")
	}
	f, err := os.Open(path)
	if err != nil {
		return errors.Wrap(err, "unable to open file")
	}
	if f.Name() != "expectedname" {
		return errors.Errorf("unexpected file name: %s", f.Name())
	}
	return nil
}

func properErrorMessage(path string) (io.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Errorf("properMessage('%s'), os.Open() call failed with error '%s'", path, err)
	}
	return f, nil
}

func main() {
	err := processFile("")
	if err != nil {
		switch errors.Cause(err).(type) {
		case *os.PathError:
			fmt.Printf("file not found: %+v\n", err)
		default:
			fmt.Printf("unknown error: %+v", err)
		}
		return
	}
	fmt.Println("OK")

	_, err2 := properErrorMessage("ghost")
	if err2 != nil {
		fmt.Println(err2)
	}
}
