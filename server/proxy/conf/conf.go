package conf

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Host struct {
	IP   string
	Port string
}

func HostCheck() Host {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	var host Host = Host{
		IP:   os.Getenv("HOST_IP"),
		Port: os.Getenv("HOST_PORT"),
	}

	fmt.Println("IP:"+host.IP, "Port:"+host.Port)
	return host
}
