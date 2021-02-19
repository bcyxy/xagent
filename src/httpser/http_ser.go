package httpser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"xagent/src/common"
	"xagent/src/shelltool"
)

type handlerFuncT func(string, uint32) (interface{}, error)

//请求处理函数映射表
var handlerMap map[string]handlerFuncT = make(map[string]handlerFuncT)

func init() {
	handlerMap["shell"] = shelltool.Do
}

//ReqT 请求结构
type ReqT struct {
	ReqType string          `json:"req_type"`
	Timeout uint32          `json:"timeout"`
	Params  json.RawMessage `json:"params"`
}

//MyHandler xxx
func MyHandler(rspWriter http.ResponseWriter, req *http.Request) {
	common.LogDebug("Receive http request. client=%s", req.RemoteAddr)

	// 暂时只支持POST请求
	if strings.ToUpper(req.Method) != "POST" {
		eMsg := fmt.Sprintf("Unsupported http method '%s'", req.Method)
		rspWriter.WriteHeader(500)
		fmt.Fprintln(rspWriter, eMsg)
		common.LogInfo(eMsg)
		return
	}

	bodyByte, err := ioutil.ReadAll(req.Body)
	if err != nil {
		rspWriter.WriteHeader(500)
		fmt.Fprintf(rspWriter, "Read body failed. err=%v", err)
		return
	}

	reqBody := ReqT{}
	err = json.Unmarshal(bodyByte, &reqBody)
	if err != nil {
		rspWriter.WriteHeader(500)
		fmt.Fprintf(rspWriter, "Loads json failed. err=%v", err)
		return
	}

	//TODO 处理超时

	hdFunc, ok := handlerMap[reqBody.ReqType]
	if !ok {
		rspWriter.WriteHeader(500)
		fmt.Fprintln(rspWriter, "Find func failed.")
		return
	}

	rst, err := hdFunc(string(reqBody.Params), reqBody.Timeout)
	if err != nil {
		rspWriter.WriteHeader(500)
		fmt.Fprintln(rspWriter, "Exec func failed.")
		return
	}
	rspWriter.Header().Set("Content-Type", "application/json")
	jsonStr, _ := json.Marshal(rst)
	rspWriter.Write(jsonStr)
}

//Start 启动HTTP服务
func Start() error {
	http.HandleFunc("/", MyHandler)

	chanErr := make(chan error)
	go func() {
		time.Sleep(time.Duration(1) * time.Second)
		chanErr <- nil
	}()

	go func() {
		portStr := common.GetConfStr("glb", "http_port", "6661")
		chanErr <- http.ListenAndServe("0.0.0.0:"+portStr, nil)
	}()

	return <-chanErr
}
