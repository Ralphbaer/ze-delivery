package main

import (
	"log"

	"github.com/Ralphbaer/ze-delivery/common"
	"github.com/Ralphbaer/ze-delivery/partner/gen"
)

func main() {
	common.InitLocalEnvConfig()
	gen.InitializeApp().Run()
	log.Print("partner service terminatedd")
}

