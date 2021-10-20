package main

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		outputDir string
	)

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
	if !strings.HasSuffix(fInfo.Name(), ".apk") {
		log.Fatal("not apk file")
	}
	decompiledTo := decompileApkToSmali(targetApkFileName)
	clientsLocationPath := filepath.Join(decompiledTo, "jp", "naver", "line", "android", "thrift", "client")

	err = filepath.Walk(clientsLocationPath, func(path string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if strings.HasSuffix(info.Name(), "Client.smali") {
			service, err := parseThriftClientInterfaceFile(path)
			if err != nil {
				return err
			}
			outputThriftDefinitionFile(service, outputDir+"/"+info.Name()[:len(info.Name())-len("Client.smali")]+".thrift")
		}
		return nil
	})
	if err != nil {
		log.Fatal("failed to walk func: ", err)
	}

}
