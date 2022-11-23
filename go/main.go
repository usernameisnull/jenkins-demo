package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	n := generateRandomNum(1, 30000)
	now := time.Now().Format("2006-01-02 15:04:05")
	res := fmt.Sprintf("time: %s\nrandom number: %d\n", now, n)
	fmt.Fprintf(w, res)
}

func main() {
	var port int
	flag.IntVar(&port, "p", 8080, "端口号,默认为8080")
	flag.Parse()
	fmt.Println("监控在端口", port)
	http.HandleFunc("/", indexHandler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("错误: ", err)
	}
}

func generateRandomNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
