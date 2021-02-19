package shelltool

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	expect "github.com/google/goexpect"
)

type interactItemT struct {
	Expect  string `json:"expect"`
	Send    string `json:"send"`
	Timeout uint32 `json:"timeout"`
}

type shellParamsT struct {
	Spawn    string          `json:"spawn"`
	Interact []interactItemT `json:"interact"`
}

//Do xxx
func Do(paramsStr string, timeout uint32) (interface{}, error) {
	rstData := []string{}

	// 解析参数
	params := shellParamsT{}
	err := json.Unmarshal([]byte(paramsStr), &params)
	if err != nil {
		return rstData, err
	}

	// 命令交互
	exp, _, err := expect.Spawn(fmt.Sprintf(params.Spawn), -1)
	if err != nil {
		return rstData, err
	}
	defer exp.Close()

	for _, interItem := range params.Interact {
		expRE := regexp.MustCompile(interItem.Expect)
		output, _, err := exp.Expect(expRE, time.Duration(interItem.Timeout)*time.Second)
		if err != nil { //如果有错误，怎么退出？
			break
		}
		rstData = append(rstData, output)
		exp.Send(interItem.Send + "\n")
	}

	return rstData, nil
}
