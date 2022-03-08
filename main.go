package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Getenv("TOKEN"))
	resp, err := http.Get("https://en.wikipedia.org/wiki/Main_Page")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
