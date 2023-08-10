package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

const defaultHost = "rafaelemery.github.io"

// Generate function will return CLI application to be executed
func Generate() *cli.App {
	app := cli.NewApp()

	defineAppParams(app)
	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: defaultHost,
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "search IP address",
			Flags:  flags,
			Action: searchIps,
		},
		{
			Name:   "server",
			Usage:  "seach servers",
			Flags:  flags,
			Action: searchServers,
		},
	}

	return app
}

func defineAppParams(app *cli.App) {
	app.Name = "CLI test application"
	app.Usage = "Search for IPs and server names"
}

func searchIps(c *cli.Context) {
	host := getHost(c)

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func searchServers(c *cli.Context) {
	host := getHost(c)

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}

func getHost(c *cli.Context) string {
	host := c.String("host")
	if host == defaultHost {
		fmt.Println("no host defined. using default host - ", defaultHost)
	}

	return host
}
