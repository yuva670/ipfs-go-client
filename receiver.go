package main

import (
	"os"
	"os/exec"
	//"fmt"
	"io"
	"bufio"
)

func main() {

	hash := os.Args[1]
	filename := os.Args[2]
	DownloadFromIPFS(filename, hash)
}

func DownloadFromIPFS(filename string, hash string) {

	cmd := exec.Command("ipfs", "cat", hash)

	outfile, err := os.Create(filename)

	if err != nil {
        panic(err)
    }
    defer outfile.Close()

    stdoutPipe, err := cmd.StdoutPipe()
    if err != nil {
        panic(err)
    }

    writer := bufio.NewWriter(outfile)
	defer writer.Flush()

    err = cmd.Start()
    if err != nil {
        panic(err)
    }

    go io.Copy(writer, stdoutPipe)
    cmd.Wait()
}
