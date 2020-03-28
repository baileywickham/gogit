package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var userconfig = userConfig{"bailey", "b@g"}

func parseConfig() userConfig {
	var user userConfig
	data, err := ioutil.ReadFile(".gogit/user.yaml")
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(data, &user)
	if err != nil {
		panic(err)
	}
	return user
}

//const (
//	NAME  = "bailey"
//	EMAIL = "b@g"
//)
