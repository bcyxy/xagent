package httpser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"xagent/src/common"
)

//HdlFunc xxx
type HdlFunc func(string, uint32) ([]string, error)

//HandlersMap xxx
var HandlersMap map[string]HdlFunc

//ReqT xxx
type ReqT struct {
	ReqType string          `json:"req_type"`
	Timeout uint32          `json:"timeout"`
	Params  json.RawMessage `json:"params"`
}

//MyHandler xxx
func MyHandler(rspWriter http.ResponseWriter, req *http.Request) {
	common.LogDebug("Receive http request.")

	bodyByte, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rspWriter.WriteHeader(500)
		fmt.Fprintln(rspWriter, "Read body failed.")
		return
	}

	reqBody := ReqT{}
	err = json.Unmarshal(bodyByte, &reqBody)
	if err != nil {
		rspWriter.WriteHeader(500)
		fmt.Fprintln(rspWriter, "Read body failed.")
		return
	}

	//处理超时

	hdFunc, ok := HandlersMap[strings.ToUpper(reqBody.ReqType)]
	if !ok {
		rspWriter.WriteHeader(500)
		fmt.Fprintln(rspWriter, "Read body failed.")
		return
	}

	rst, err := hdFunc(string(reqBody.Params), reqBody.Timeout)
	if err != nil {
		rspWriter.WriteHeader(500)
		fmt.Fprintln(rspWriter, "Read body failed.")
		return
	}
	rspWriter.Header().Set("Content-Type", "application/json")
	jsonStr, _ := json.Marshal(rst)
	rspWriter.Write(jsonStr)
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
