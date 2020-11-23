package main

import "TTT/mod/helpers"

func main() {
	r := helpers.SetupRouter()
	err := r.Run()
	if err != nil {
		helpers.Log.Error(err)
		return
	}
}
