package main

import "fmt"

import "io"

import "log"

import "net"

var AntifyData = "/data"
var AntifyHost = ""
var AntifyPort = "555"

func onError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func isLastChars(str1 string, lenght int, bytes []byte) string {
	if (len(str1) < len(bytes)) { lenght = len(str1) }

	return []byte(str1[len(str1) - lenght:]) == bytes
}

func ReadUntil(r io.Reader) (str string, err error) {
	result := ""

	buf := make([]byte, 1)

	for isLastChars(result, 6, []byte{13,10,13,10,13,10}) {
		_, err := r.Read(buf)

		if (err != nil) { return result, err }

		result += string(buf)
	}

	return result, nil
}

func handleConn(c net.Conn) {
	defer c.Close()

	for {
		b, err := ReadUntil(c)
		onError(err)
		fmt.Print(b)
	}
}

func main() {
	fmt.Printf("Antify; Data: %s; Port: %s; \n", AntifyData, AntifyPort)

	l, err := net.Listen("tcp", AntifyHost + ":" + AntifyPort)
	defer l.Close()
	onError(err)

	for {
		conn, err := l.Accept()
		onError(err)
		go handleConn(conn)
	}
}
