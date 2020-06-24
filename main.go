package main

import (
	"fyneGui/client"
	"fyneGui/handle"
	"fyneGui/initapp"
	"fyneGui/log"
)

func main() {
	log.InitLog()
	initapp.InitApp()
	handle.LoginScreen()
	handle.Run()
	go client.ClientRun()
	initapp.GetMyApp().Run()
}
