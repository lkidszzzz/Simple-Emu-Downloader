package get

import (
	"fmt"
	"github.com/gookit/color"
	"strconv"
)

var Agent int

func modeChooser() int {
	breaker := 0
	var mode string
	for {
		fmt.Println()
		color.Set(color.FgLightCyan)
		fmt.Println("	输入 1 选择直连模式（若您没有代理或不知道代理是什么，选此项）")
		fmt.Println("	输入 2 选择代理模式（有代理选此项）")
		fmt.Print("	请输入对应的数字以选择您的网络模式：")
		fmt.Scan(&mode)
		m, err := strconv.Atoi(mode)
		switch {
		case m == 1:
			breaker = 1
		case m == 2:
			Agent = 1
			breaker = 1
		case err != nil || m < 1 || m > 2:
			color.Error.Prompt("	请按要求正确输入数字。")
		}
		if breaker == 1 {
			breaker = 0
			break
		}
	}
	return Agent
}
