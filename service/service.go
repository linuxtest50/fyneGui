package service

import (
	"fyneGui/login"
	"fyneGui/ntf"
	"fyneGui/obj"
	"fyneGui/sign"
)

func Run() {
	go func() {
		for msg := range ntf.FyneGuiMsg {
			switch msg.(type) {
			case *obj.NtfLogin:
				//login.Login()
			case *obj.NtfSign:
				sign.SignIn()
			case *obj.NtfSignSuccess:
				login.Login()
			}
		}
	}()
}

func LoginScreen() {
	login.Login()
}
