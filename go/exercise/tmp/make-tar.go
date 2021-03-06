package main

import (
	"os"
	"archive/tar"
	"log"
	"io"
	"compress/gzip"
)

func addFile(tw * tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball	
		header, err := tar.FileInfoHeader(stat, "")
		if err != nil {
			panic(err)
		}
		/*
		header := new(tar.Header)
		header.Name = path
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()
		*/
		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// copy the file data to the tarball 
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	// set up the output file
	file, err := os.Create("output.tar.gz")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	// set up the gzip writer	
	gw := gzip.NewWriter(file)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()
	// grab the paths that need to be added in
	paths := []string{
		"./t.go",
		"./test.go",
	}
	// add each file as needed into the current tar archive
	for i := range paths {
		if err := addFile(tw, paths[i]); err != nil {
			log.Fatalln(err)
		}
	}
}

