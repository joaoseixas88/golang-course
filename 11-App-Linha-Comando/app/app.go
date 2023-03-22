package app

import (
	"fmt"
	"log"
	"net"
	"strconv"

	"github.com/urfave/cli"
)

func Gerar() *cli.App {
	app := cli.NewApp()

	app.Name = "CommandLine App"
	app.Usage = "Find Server Names and Ips on Internet"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Find ip address on internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: findIps,
		},
		{
			Name:  "calculate",
			Usage: "Calculate 2 variables",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "operation",
					Value: "+",
				},
				cli.StringFlag{
					Name:  "first",
					Value: "0",
				},
				cli.StringFlag{
					Name:  "second",
					Value: "",
				},
			}, Action: calculate,
		},
	}

	return app
}

func findIps(c *cli.Context) {
	host := c.String("host")
	ip, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	for _, value := range ip {
		fmt.Println(value)
	}
}

func calculate(c *cli.Context) {
	operation := c.String("operation")
	first, _ := strconv.Atoi(c.String("first"))
	second, _ := strconv.Atoi(c.String("second"))
	switch operation {
	case "-":
		fmt.Println("Resultado: ", first-second)
	case "x":
		fmt.Println("Resultado: ", first*second)
	case "/":
		fmt.Println("Resultado: ", first/second)
	default:
		fmt.Println("Resultado: ", first+second)
	}
}
