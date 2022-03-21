package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("file_may_not_exist")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(os.Stdout, f)
}
