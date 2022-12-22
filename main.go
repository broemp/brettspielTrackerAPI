package main

import (
	"github.com/broemp/brettspielTrackerAPI/initializers"
	"github.com/broemp/brettspielTrackerAPI/internal"
)

func init() {
	initializers.Setup()
	initializers.ConnectToDB()
}

func main() {
	internal.StartServer()
}
