package main

import (
	"compress/zlib"
	"crypto/sha256"
	"io"
	"io/ioutil"
	"os"
)

func hash(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	sum := sha256.Sum256(data)
	println(sum[:])
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
