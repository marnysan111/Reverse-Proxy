package form

import (
	"log"
	"net/http"
	"net/url"
)

func PostForm() {
	values := url.Values{}
	values.Add("id", "123")
	values.Add("name", "ポメラニアン")
	values.Encode()
	_, err := http.PostForm("http://192.168.1.10:50000", values)
	if err != nil {
		log.Fatal(err)
	}
	/*
		// header
		fmt.Printf("[status] %d\n", res.StatusCode)
		for k, v := range res.Header {
			fmt.Print("[header] " + k)
			fmt.Println(": " + strings.Join(v, ","))
		}

		// body
		defer res.Body.Close()
		body, error := ioutil.ReadAll(res.Body)
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println("[body] " + string(body))
	*/
}
