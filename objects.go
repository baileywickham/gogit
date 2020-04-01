package main

import (
	"compress/zlib"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func add(fdpath string) {
	if !wdInRepo() {
		log.Fatal("not in gogit repo")
		return
	}

	files := make([]indexEntry, 0)

	err := filepath.Walk(fdpath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("error walking %s", path)
			return nil
		}
		if info.IsDir() || info.Name() == ".gogit" {
			return nil
		}

		data, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalf("error reading %s", path)
			return err
		}

		hash := hashObject(path, data)

		writeObject(path, hash, data)

		files = append(files,
			indexEntry{mode: info.Mode().String(),
				filepath: gogitRelPath(path),
				hash:     hash})
		return nil
	})
	if err != nil {
		panic(err)
	}
	// after walking directory, write index
	writeIndex(files)

}

func hashObject(filename string, filedata []byte) string {
	sum := sha512.Sum512(filedata)
	//NOTE, this should probably be returned as a [64]byte, but that is hard to deal with
	return hex.EncodeToString(sum[:])
}

func writeObject(filename, hash string, filedata []byte) {
	header := fmt.Sprint("blob ", len(filedata), '\000') //append header for git

	dir := fmt.Sprint(".gogit/objects/", string(hash[:2]), "/")
	os.MkdirAll(dir, 0755)

	compfile, err := os.Create(dir + string(hash[2:42]))
	if err != nil {
		panic(err)
	}
	defer compfile.Close()

	w := zlib.NewWriter(compfile)
	defer w.Close()
	w.Write([]byte(header))
	w.Write(filedata)
}

func catObjectFile(filename string) {
	//data, err := ioutil.ReadFile(filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	r, err := zlib.NewReader(file)
	defer r.Close()

	io.Copy(os.Stdout, r)
}
