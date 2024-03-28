package util

import (
	"log"

	"github.com/fatih/color"
)

func ServerLog(msg string) {
	log.Println(color.BlueString("[grpc-server] => " + msg))
}

func ClientLog(msg string) {
	log.Println(color.YellowString("[grpc-client] => " + msg))
}

func ErrLog(msg string) {
	log.Println(color.RedString("error: " + msg))
}
