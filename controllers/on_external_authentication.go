package controllers

import "fmt"

func OnExternalAuthentication(b []byte) {
	fmt.Println("user was authenticated")
}
