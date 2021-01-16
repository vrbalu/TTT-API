package main

import (
	"TTT/mod/controllers"
	"TTT/mod/helpers"
)

func main() {
	r := controllers.SetupRouter()
	err := r.Run()
	if err != nil {
		helpers.Log.Error(err)
		return
	}
}
