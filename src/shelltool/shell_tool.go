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
func Do(paramsStr string, timeout uint32) ([]string, error) {
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
		fmt.Printf("yxytest:%v|%v\n", expRE, interItem.Send)
		output, mcList, err := exp.Expect(expRE, time.Duration(interItem.Timeout)*time.Second)
		if err != nil { //如果有错误，怎么退出？
			fmt.Printf("%v|%v\n", err, output)
			break
		}
		fmt.Printf("%v\n", mcList)
		rstData = append(rstData, output)
		exp.Send(interItem.Send + "\n")
	}

	fmt.Printf("%v\n", rstData)

	return []string{}, nil
}

//TestSellTool xxx
func TestSellTool() {
	testReq := `
	{
		"spawn": "ssh localhost",
		"interact": [
			{
				"expect": "[Pp]assword:",
				"send": "666666",
				"timeout": 5
			},
			{
				"expect": "[\\$\\]#>]",
				"send": "exit",
				"timeout": 4
			},
			{
				"expect": "[\\$\\]#>]",
				"send": "quit",
				"timeout": 1
			}
		]
	}`

	_, err := Do(testReq, 30)
	if err != nil {
		//fmt.Printf("%v\n", err)
	}
}

/*
var (
	userRE   = regexp.MustCompile("username:")
	passRE   = regexp.MustCompile("password:")
	promptRE = regexp.MustCompile(`[\$#\]>]`)
)

exp, _, err := expect.Spawn(fmt.Sprintf("ssh localhost"), -1)
if err != nil {
	log.Fatal(err)
}
defer exp.Close()

//exp.Expect(userRE, timeout)
//exp.Send("yxy\n")

exp.Expect(passRE, timeout)
exp.Send("666666\n")

exp.Expect(promptRE, timeout)
exp.Send("ls\n")
result, v1, err := exp.Expect(promptRE, timeout)
if err != nil {
	return
}
fmt.Printf("yxytest:%v|%v\n", result, v1)

exp.Send("pwd\n")
result, _, _ = exp.Expect(promptRE, timeout)
fmt.Printf("yxytest:%v\n", result)

exp.Send("exit\n")
*/
