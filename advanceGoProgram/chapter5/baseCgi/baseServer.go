package baseCgi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
 Go中使用`net/http`包实现简单的http server
*/
func echo(wr http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		wr.Write([]byte("echo error"))
		return
	}

	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
	fmt.Println("request body data length: ", writeLen)
}

func BaseServer() {
	http.HandleFunc("/", echo)
	fmt.Println("Server run on socket 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
