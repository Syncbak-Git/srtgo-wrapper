package main

import (
	"bufio"
	"fmt"
	"os"

	srtgo "github.com/Syncbak-Git/srtgo-wrapper"
)

func main() {
	options := make(map[string]string)
	options["blocking"] = "0"
	options["transtype"] = "file"
	hostname := "0.0.0.0"
	port := 8090

	fmt.Printf("(srt://%s:%d) Listening", hostname, port)
	a := srtgo.NewSrtSocket(hostname, uint16(port), options)
	err := a.Listen(2)
	defer a.Close()
	if err != nil {
		panic("Error on Listen")
	}

	for {
		s, _, err := a.Accept()
		if err != nil {
			panic("Error on Accept")
		}

		buff := make([]byte, 2048)
		fo, err := os.Create("sample.ts")
		w := bufio.NewWriter(fo)
		for {
			n, err := s.Read(buff)

			if err != nil {
				fmt.Println(err)
				break
			}

			if n == 0 {
				break
			}

			w.Write(buff[:n])
		}
		w.Flush()
		s.Close()
		fo.Close()
	}
}
