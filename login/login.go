package login

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"fyneGui/client_handler"
	"fyneGui/initapp"
	"fyneGui/ntf"
	"fyneGui/obj"
)

func Login() {
	myApp := initapp.GetMyApp()
	myWin := myApp.NewWindow("Entry")

	myWin.Resize(fyne.NewSize(400, 400))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("input name")
	nameEntry.OnChanged = func(content string) {
		//fmt.Println("name:", nameEntry.Text, "entered")
	}

	passEntry := widget.NewPasswordEntry()
	passEntry.SetPlaceHolder("input password")

	nameBox := widget.NewHBox(widget.NewLabel("Name"), layout.NewSpacer(), nameEntry)
	passwordBox := widget.NewHBox(widget.NewLabel("Password"), layout.NewSpacer(), passEntry)

	loginBtn := widget.NewButton("Login", func() {
		fmt.Println("name:", nameEntry.Text, "password:", passEntry.Text, "login in")
		//myWin.Close()
		client_handler.Login(nameEntry.Text)
	})

	signBtn := widget.NewButton("Sign", func() {
		myWin.Close()
		ntf.Tell(&obj.NtfSign{})
	})
	multiEntry := widget.NewEntry()
	multiEntry.SetPlaceHolder("please enter\nyour description")
	multiEntry.MultiLine = true

	content := widget.NewVBox(nameBox, passwordBox, loginBtn, signBtn, multiEntry)
	myWin.SetContent(content)
	myWin.Show()
}
