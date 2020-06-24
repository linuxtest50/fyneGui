package initapp

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
)

var myApp fyne.App

func InitApp() {
	myApp = app.New()
	myApp.SetIcon(theme.FyneLogo())
}

func GetMyApp() fyne.App {
	return myApp
}
