package main

import "fmt"

import "flag"

func main() {
	data := *flag.String("data-path", "/data", "Path to save all data.")

	port := *flag.String("port", "555", "Port expose HTTP project files.")

	fmt.Printf("Antify; Data: %s; Port: %s; \n", data, port)
}
