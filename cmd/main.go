package main

import (
	"log"
	"x14nfile"
)

func main() {
	tree, err := x14nfile.BuildDirectoryTree("./download/clash")
	if err != nil {
		log.Fatal(err)
	}
	x14nfile.FileDir = tree
	x14nfile.Serverstart()
}
