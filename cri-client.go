package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	// TODO: proper implementation -> https://github.com/containerd/containerd
	crictl "github.com/jolorunyomi/golang-crictl/crictl"
)

const X_VER = "1.0.0"

func main() {
	fmt.Print("CRI-O compatible CLI client X\n\n")

	socket := flag.String("cri-sock", "/var/run/docker.sock", "Socket CRI-O compatible daemon listens on.")
	flag.Parse()

	cl := crictl.Client(*socket)
	ping := &crictl.Ping{}
	err := ping.New(*cl)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("client-x version: %s\n", X_VER)
	fmt.Printf("server-x version: %s\n\n\n", ping.ServerVersion)

	if len(flag.Args()) < 1 {
		os.Exit(0)
	}

	handleCommand(cl, flag.Arg(0), flag.Args())
}

func handleCommand(cl *http.Client, command string, extraData []string) {
	switch command {
	case "ps":
		ps := &crictl.Ps{}
		err := ps.New(*cl)

		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("CONTAINER ID\tIMAGE\tCOMMAND\t\tCREATED\tSTATUS\tPORTS \tNAMES")
		for _, container := range *ps {
			fmt.Printf(
				"%s\t%s\t%s\t\t%d\t%s\t%d\t%s\n",
				container.ID,
				container.Image,
				container.Command,
				container.Created,
				container.Status,
				container.Ports[0].PrivatePort,
				container.Names,
			)
		}
		// TODO: possible enhancements
		// * clean up ps output
		// * add more commands
		// * better error handling
		// * cleaner interfaces
		// * handle extra data
		// * rootless support

	default:
		fmt.Printf("%s is not a supported command\n", command)
		os.Exit(1)
	}
}
