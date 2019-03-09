package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"time"
// )

// func main() {
// 	fmt.Println("Please visit http://127.0.0.1:12345/")
// 	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
// 		s := fmt.Sprintf("你好, 世界! -- Time: %s", time.Now().String())
// 		fmt.Fprintf(w, "%v\n", s)
// 		log.Printf("%v\n", s)
// 	})
// 	if err := http.ListenAndServe(":12345", nil); err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

import (
	"io/ioutil"
	"log"
	"net/http"
)

func echo(wr http.ResponseWriter, r *http.Request) {
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		if _, e := wr.Write([]byte("echo error")); e != nil {

		}

		return
	}

	writeLen, err := wr.Write(msg)
	if err != nil || writeLen != len(msg) {
		log.Println(err, "write len:", writeLen)
	}
}

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello everyone!"))
	})
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
