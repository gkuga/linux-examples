package main

import (
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/creack/pty"
)

func main() {
	c := exec.Command("/bin/sh")

	ptmx, err := pty.Start(c)
	if err != nil {
		log.Fatal(err)
	}
	defer ptmx.Close()

	go func() {
		_, _ = io.Copy(ptmx, os.Stdin)
	}()

	_, _ = io.Copy(os.Stdout, ptmx)
}
