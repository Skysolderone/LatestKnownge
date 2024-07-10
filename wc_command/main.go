package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// wc command build
func main() {
	fileName := flag.String("c", "", "Specify the file")
	if *fileName == "" {
		panic("Must specify a file name")
	}
	currenDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	pathTh := []string{}
	err = filepath.Walk(currenDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err.Error())
		} else {
			pathTh = append(pathTh, path)
		}
		return err
	})
	if err != nil {
		log.Fatal(err)
	}
	if !contains(pathTh, *fileName) {
		log.Fatal("file is not found")
	}
	fileInfo, err := os.Stat(*fileName)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(fileInfo.Size(), *&fileName)
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
		return false
	}
	return false
}
