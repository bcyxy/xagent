package httpser

import (
	"fmt"
	"net/http"
	"time"
	"xagent/common"
)

//MyHandler xxx
func MyHandler(w http.ResponseWriter, r *http.Request) {
	common.LogDebug("Receive http request.")

	fmt.Fprintln(w, "hello world")
}

//Start xxx
func Start() error {
	http.HandleFunc("/", MyHandler)

	chanErr := make(chan error)
	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		chanErr <- nil
	}()

	go func() {
		chanErr <- http.ListenAndServe("0.0.0.0:6661", nil)
	}()

	return <-chanErr
}
