package main

import (
	"th3y3m/e-commerce-platform/API"
)

func main() {
	router := API.Controller()
	router.Run("localhost:8080")
}
