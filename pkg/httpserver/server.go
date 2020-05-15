package httpserver

import (
	"log"
	"os"
)

func New() Server {
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile|log.Ltime)

	return Server{
		errorLog: errorLog,
	}
}

type Server struct {
	errorLog *log.Logger
}
