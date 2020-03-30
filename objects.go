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
)

func add(objpath string) {
	if !wdInRepo() {
		log.Fatal("not in gogit repo")
		return
	}

}

func hashObject(filename string, filedata []byte) string {
	sum := sha512.Sum512(filedata)
	//NOTE, this should probably be returned as a [64]byte, but that is hard to deal with
	return hex.EncodeToString(sum[:])
}

func writeObject(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	header := fmt.Sprint("blob ", len(data), '\000') //append header for git

	sum := hashObject(filename, data)

	dir := fmt.Sprint(".gogit/objects/", string(sum[:2]), "/")
	os.MkdirAll(dir, 0755)

	hashfile, err := os.Create(dir + string(sum[2:42]))
	if err != nil {
		panic(err)
	}
	defer hashfile.Close()

	w := zlib.NewWriter(hashfile)
	defer w.Close()
	w.Write([]byte(header))
	w.Write(data)

}

func catFile(filename string) {
	//data, err := ioutil.ReadFile(filename)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	r, err := zlib.NewReader(file)
	defer r.Close()

	io.Copy(os.Stdout, r)
}
