package initapp

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

var myApp fyne.App

func InitApp() {
	myApp = app.New()
}

func GetMyApp() fyne.App {
	return myApp
}
