package main

import (
	"os"
	"path"
	"strings"
)

func gogitPathFromWD() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return gogitPath(wd)

}

func wdInRepo() bool {
	if gogitPathFromWD() == "" {
		return false
	}
	return true
}

// assumes we are in a gogit repo
func gogitRelPath(fdpath string) string {
	return path.Join(gogitPathFromWD(), fdpath)
}

// assumes we are in a gogit repo
func gogitPath(dirpath string) string {
	// this may cause bugs I'm not expecting?
	if dirpath == "/" || dirpath == "" {
		return ""
	}
	if _, err := os.Stat(".gogit"); os.IsNotExist(err) {
		return gogitPath(dirpath[:strings.LastIndex(dirpath, "/")])
	}
	return dirpath
}
