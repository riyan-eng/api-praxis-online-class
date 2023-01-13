package main

import (
	"fmt"

	"github.com/riyan-eng/api-praxis-online-class/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDatabase()
}

func main() {
	fmt.Println("hello riyan")
}
