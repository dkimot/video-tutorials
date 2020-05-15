package main

import (
	"fmt"

	vidtuts "github.com/dkimot/video-tutorials"
)

func main() {
	typedMsg := (*vidtuts.VideoViewed)(nil)

	fmt.Println(typedMsg.Type())
}
