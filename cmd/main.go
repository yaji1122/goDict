package main

import "log"

func main() {
	//setup router
	r := setupRouter()

	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		log.Panic(err)
	}
}