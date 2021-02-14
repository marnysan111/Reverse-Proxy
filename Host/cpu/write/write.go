package write

import (
	"fmt"
	"io/ioutil"
)

/* Host/proxyのenvに書き込み */
func WriteEnv(ip string, port string) error {
	var lines = []string{"HOST_IP=" + ip + "\nHOST_PORT=" + port}
	b := []byte{}
	for _, line := range lines {
		ll := []byte(line)
		for _, l := range ll {
			b = append(b, l)
		}
	}

	err := ioutil.WriteFile("../proxy/.env", b, 0666)
	if err != nil {
		fmt.Println("err:", err)
		return err
	}
}
