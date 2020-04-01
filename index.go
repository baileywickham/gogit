package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type indexEntry struct {
	mode     string
	filepath string
	hash     string
}

func addOrUpdateIndex(path []string) {

}

func writeIndex(files []indexEntry) {
	index := readIndex()
	for _, file := range files {
		index = append(index, file)
	}
	indexFile, err := os.OpenFile(gogitRelPath(".gogit/index"), os.O_RDWR, 0755)
	if err != nil {
		log.Fatal("Error opening index file")
		return
	}
	defer indexFile.Close()

	// TODO sort index entry list
	for _, entry := range index {
		writeIndexEntry(indexFile, entry)
	}
}

func writeIndexEntry(file *os.File, e indexEntry) {
	_, err := file.WriteString(fmt.Sprintf("%s %s %s\n", e.mode, e.filepath, e.hash))
	if err != nil {
		panic(err)
	}
}

func inIndexByName(filename string) bool {
	index := readIndex()
	for _, entry := range index {
		if entry.filepath == filename {
			return true
		}
	}
	return false
}

func inIndexByHash(hash string) bool {
	index := readIndex()
	for _, entry := range index {
		if entry.hash == hash {
			return true
		}
	}
	return false
}

func readIndex() []indexEntry {
	index := make([]indexEntry, 0)
	indexpath := gogitRelPath(".gogit/index")
	file, err := os.Open(indexpath)
	defer file.Close()

	if err != nil {
		log.Fatal("No index file")
	}

	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")
		index = append(index, indexEntry{mode: line[0], filepath: line[1], hash: line[2]})
	}
	return index
}
