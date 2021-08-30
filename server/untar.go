package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func untar(srcFile, dstDir string) error {

	f, err := os.Open(srcFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	gzf, err := gzip.NewReader(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tarReader := tar.NewReader(gzf)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		path := filepath.Join(dstDir, header.Name)

		if err != nil {
			log.Fatalf("ExtractTarGz: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(path, 0755); err != nil {
				fmt.Println(err.Error())
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
				fmt.Println(err.Error())
				log.Fatalf("ExtractTarGz: Mkdir() failed: %s", err.Error())
			}

			outFile, err := os.Create(path)
			if err != nil {
				fmt.Println(err.Error())
				log.Fatalf("ExtractTarGz: Create() failed: %s", err.Error())
			}
			defer outFile.Close()
			if _, err := io.Copy(outFile, tarReader); err != nil {
				fmt.Println(err.Error())
				log.Fatalf("ExtractTarGz: Copy() failed: %s", err.Error())
			}
		default:
			log.Fatalf(
				"ExtractTarGz: uknown type: %s in %s",
				header.Typeflag,
				header.Name)
		}
	}
	return nil
}
