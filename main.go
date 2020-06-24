package main

import (
	"fyneGui/initapp"
	"fyneGui/service"
)

func main() {
	initapp.InitApp()
	service.LoginScreen()
	service.Run()
	initapp.GetMyApp().Run()
}
