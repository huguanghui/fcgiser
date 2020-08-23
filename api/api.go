package api

import "fmt"

var (
	CmdRuner = NewRunner()
)

type Result struct {
	Reusult     int    `json:"result"`
	Description string `json:"description"`
}

func StartPrint() {
	fmt.Println("hello")
}
