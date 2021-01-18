package main

import (
	"client/form"
	"time"
)

func main() {
	for {
		time.Sleep(time.Second * 2)
		form.PostForm()
	}
}
