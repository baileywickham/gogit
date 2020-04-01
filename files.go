package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

func gogitAbsPathFromWD() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return gogitPath(wd)

}

func wdInRepo() bool {
	if gogitAbsPathFromWD() == "" {
		return false
	}
	return true
}

// gogitRelPath returns the path to a file relitive to the top level gogit repo
func gogitRelPath(fdpath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// err should never occur
	p, _ := filepath.Rel(gogitAbsPathFromWD(), path.Join(cwd, fdpath))

	return p
}

// Returns the path of the gogit directory, starting from dirpath.
// For internal use only
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
