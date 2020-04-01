package main

import (
	"log"
	"os"

	r "github.com/baileywickham/runner"
	"gopkg.in/yaml.v2"
)

type userConfig struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email"`
}

func main() {
	shell := r.NewShell()
	shell.Add_command(r.Command{
		Cmd:      "init",
		Callback: gitInit,
		Helptext: "init .git directory",
	}, r.Command{
		Cmd: "hash-file",
		Callback: func(filename string) {
			println(hashObject)
		},
		Helptext: "hash a file to sha256",
	}, r.Command{
		Cmd:      "cat-file",
		Callback: catObjectFile,
		Helptext: "cat hashed file",
	}, r.Command{
		Cmd:      "add",
		Callback: add,
		Helptext: "add object",
	})
	shell.Start()

}

func gitInit() {
	if wdInRepo() {
		log.Fatal("already in gogit repo")
		return
	}
	err := os.MkdirAll(".gogit/objects", 0755)
	if err != nil {
		panic(err)
	}
	_, err = os.Create(".gogit/index")
	if err != nil {
		panic(err)
	}
	// can ignore errors because we created the dir
	file, _ := os.Create(".gogit/user.yaml")
	data, err := yaml.Marshal(userconfig)
	if err != nil {
		panic(err)
	}
	file.Write(data)
}
