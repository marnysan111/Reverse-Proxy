package manage

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func Top(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./manage/template/index.html")
	if err != nil {
		fmt.Println(err)
		log.Fatalf("template ERROR")
	}
	t.Execute(w, nil)
}
