package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func inIndexByName(filename string) bool {
	index := readIndex()
	_, ok := index[filename]
	return ok
}

func inIndexByHash(hash string) bool {
	index := readIndex()
	for _, val := range index {
		if val == hash {
			return true
		}
	}
	return false

}

func readIndex() map[string]string {
	index := make(map[string]string)
	indexpath := gogitPath("index")
	file, err := os.Open(indexpath)
	defer file.Close()

	if err != nil {
		log.Fatal("No index file")
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")
		index[line[0]] = line[1]
	}
	return index
}

func addToIndex(filename, hash string) {
	indexpath := gogitPath("index")
	file, err := os.Open(indexpath)
	defer file.Close()

	if err != nil {
		panic(err)
	}

}
