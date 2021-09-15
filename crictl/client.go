package golangcrictl

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
)

func Client(socketAddr string) *http.Client {
	_, err := os.Stat(socketAddr)
	if os.IsNotExist(err) {
		fmt.Printf("no daemon listening on %s: %v\n", socketAddr, err)
		os.Exit(1)
	}

	cl := http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", socketAddr)
			},
		},
	}
	return &cl
}
