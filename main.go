package main

import (
	"fmt"
	"os"
)

// cloud_cost_optimizer - Optimize cloud spending
func cloud_cost_optimizer(path string) {
	fmt.Println("========================================")
	fmt.Println("  Cloud-Cost-Optimizer")
	fmt.Println("  Optimize cloud spending")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	cloud_cost_optimizer(path)
}
