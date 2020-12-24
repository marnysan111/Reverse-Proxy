package write

import (
	"io/ioutil"
	"os"
)

func Write(ip string, port string) {
	var lines = []string{"HOST_IP=" + ip + "\nHOST_PORT=" + port}
	b := []byte{}
	for _, line := range lines {
		ll := []byte(line)
		for _, l := range ll {
			b = append(b, l)
		}
	}

	err := ioutil.WriteFile(".env", b, 0666)
	if err != nil {
		os.Exit(1)
	}
}
