package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli/v2"
)

// Gerar retorna a aplicação de linha de comando para ser executada
func Gerar() *cli.App {
	flags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "devbook.com.br",
		},
	}

	app := &cli.App{
		Name:  "Aplicação de Linha de Comando",
		Usage: "Busca de IPs Nomes de servidor na internet",
		Commands: []*cli.Command{
			{
				Name:   "ip",
				Usage:  "Busca IPs de endereço na internet",
				Flags:  flags,
				Action: buscarIPs,
			},
			{
				Name:   "servidores",
				Usage:  "Busca o nome do servidor na internet",
				Flags:  flags,
				Action: buscarServidores,
			},
		},
	}

	return app
}

func buscarServidores(c *cli.Context) error {
	host := c.String("host")

	ser, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range ser {
		fmt.Println(s.Host)
	}
	return nil
}

func buscarIPs(c *cli.Context) error {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
	return nil
}
