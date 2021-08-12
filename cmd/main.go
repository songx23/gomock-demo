package cmd

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"gomock-showcase/internal/client"
	"gomock-showcase/internal/server"
)

func main() {
	c := client.NewClient(http.DefaultClient, url.URL{})
	s := server.NewServer(c)
	f, err := s.OnFire("What is this?")
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(f)
}
