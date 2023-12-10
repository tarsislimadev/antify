package main

import "fmt"

// import "flag"

func main() {
	data := "/data" // *flag.String("data", "/", "Path to save all data.")

	port := "555" // *flag.String("port", "555", "Port to expose HTTP project files.")

	fmt.Printf("Antify; Data: %s; Port: %s; \n", data, port)
}
