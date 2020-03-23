package main

import (
	"os"

	r "github.com/baileywickham/runner"
)

func main() {
	shell := r.NewShell()
	shell.Add_command(r.Command{
		Cmd:      "init",
		Callback: gitInit,
		Helptext: "init .git directory",
	})
	shell.Start()

}

func gitInit() {
	if _, err := os.Stat(".gogit/"); !os.IsNotExist(err) {
		println("gogit directory already exists")
		return
	}
	err := os.Mkdir(".gogit", 0755)
	if err != nil {
		panic(err)
	}
	// shouldn't error because we created the dir
	os.Mkdir(".gogit/objects", 0755)
}
