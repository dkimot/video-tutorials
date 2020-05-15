package main

import (
	"os"

	"github.com/dkimot/video-tutorials/messagedb"
)

var env environ

type environ struct {
	ListenAddr    string
	MessageDBOpts messagedb.Options
}

func init() {
	opts := messagedb.Options{
		Addr: os.Getenv("MESSAGE_DB_ADDR"),
		Pass: os.Getenv("MESSAGE_DB_PASS"),
		Port: os.Getenv("MESSAGE_DB_PORT"),
		User: os.Getenv("MESSAGE_DB_USER"),
	}
	listenAddr := os.Getenv("LISTEN_ADDR")

	env = environ{
		ListenAddr:    listenAddr,
		MessageDBOpts: opts,
	}
}
