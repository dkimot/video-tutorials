package main

import (
	"log"
	"net/http"

	recordviewings "github.com/dkimot/video-tutorials/apps/record_viewings"
	"github.com/dkimot/video-tutorials/messagedb"
)

func main() {
	msgDB := messagedb.New(env.MessageDBOpts)

	app := recordviewings.New(msgDB)

	log.Fatal(http.ListenAndServe(env.ListenAddr, app))
}
