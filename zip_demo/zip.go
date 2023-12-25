package main

import (
	"archive/zip"
	"io"
	"os"
)

const ZipName = "example.zip"

// create zip file
func createZip() error {
	file, err := os.Create(ZipName)
	if err != nil {
		return err
	}
	defer file.Close()
	zipWriter := zip.NewWriter(file)
	defer zipWriter.Close()
	return nil
}

// add zip file
func addFileToZip(zipWriter *zip.Writer, filePath string) error {
	fileToZip, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer fileToZip.Close()
	fileInfo, err := fileToZip.Stat()
	if err != nil {
		return err
	}
	header, err := zip.FileInfoHeader(fileInfo)
	if err != nil {
		return err
	}
	writer, err := zipWriter.CreateHeader(header)
	if err != nil {
		return err
	}
	_, err = io.Copy(writer, fileToZip)
	return err

}

// param control compress level
func createZipFileWithOptions() error {
	zipFile, err := os.Create("example.zip")
	if err != nil {
		return err
	}
	defer zipFile.Close()
	//0-9 0表示不压缩  9表示最高压缩
	zipWrite, err := zip.NewWriterLevel(zipFile, 9)
	if err != nil {
		return err
	}
	defer zipWrite.Close()
	// 这里添加文件到zip文件中
	return nil

}
func main() {

}
