package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	targetApkFileName, err := filepath.Abs(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal("error occurred during parsing path: ", err)
	}
	fInfo, err := os.Stat(targetApkFileName)
	if os.IsNotExist(err) {
		log.Fatal("no such file or directory: ", err)
	}
	if fInfo.IsDir() {
		log.Fatal("target path is not file: ")
	}

}
