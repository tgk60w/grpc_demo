/*
@Time : 2020/6/19 上午10:16
@Author : gongyikun
@File : config.go
@Software: GoLand
*/
package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type Grpc struct {
	Client  string `json:"client"`
	Service string `json:"service"`
	WebAddr string `json:"webAddr"`
	ChanNum int    `json:"chanNum"`
}

func New() *Grpc {

	fpath := "./config/env.toml"

	var conf Grpc

	if _, err := toml.DecodeFile(fpath, &conf); err != nil {
		fmt.Println("decode config file err")
	}
	//fmt.Println(conf.Client)
	return &conf
}
