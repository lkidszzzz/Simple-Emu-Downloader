package main

import (
	sed "Simple-Emu-Downloader"
	"fmt"
	"strconv"
)

var emu, yuzu string
var b int

func main() {
	for {
		fmt.Println("输入 1 获取 Nitendo Switch 模拟器（推荐）：YUZU")
		fmt.Println("输入 2 获取 Nitendo Switch 模拟器：Ryujinx")
		fmt.Println("输入 3 获取 Nitendo 3DS 模拟器：Citra3DS")
		fmt.Println("输入 4 获取 Sony PSP 模拟器：PPSSPP")
		fmt.Println("输入 5 退出下载器。")
		fmt.Println("请输入对应的数字以选择您要下载的模拟器：")
		_, err := fmt.Scan(&emu)
		if err != nil {
			fmt.Println("Unknown error.")
			return
		}
		e, err := strconv.Atoi(emu)
		switch {
		case e == 1:
			for {
				fmt.Println("输入 1 选择Github普通下载通道（国内部分地区可能体验不佳甚至无法访问）；")
				fmt.Println("输入 2 选择Github公共CDN下载通道1（推荐）（使用人数较多时可能速度较慢，请避开高峰期使用）。")
				fmt.Println("输入 3 选择Github公共CDN下载通道2（使用人数较多时可能速度较慢，请避开高峰期使用）。")
				fmt.Println("输入 4 返回。")
				fmt.Println("请输入对应的数字以选择下载通道：")
				_, err := fmt.Scan(&yuzu)
				if err != nil {
					fmt.Println("Unknown error.")
					return
				}
				e, err := strconv.Atoi(yuzu)
				switch {
				case e == 1:
					fmt.Println("下载链接（您也可以复制到您喜欢的下载器去进行下载）：", sed.YuzuGetter())
					b = 1
				case e == 2:
					fmt.Println("下载链接（您也可以复制到您喜欢的下载器去进行下载）：", sed.YuzuCDN1())
					b = 1
				case e == 3:
					fmt.Println("下载链接（您也可以复制到您喜欢的下载器去进行下载）：", sed.YuzuCDN2())
					b = 1
				case e == 4:
					b = 1
				case err != nil || e < 1 || e > 3:
					fmt.Println("Warn:请按要求正确输入数字。")
				}
				if b == 1 {
					b = 0
					break
				}
			}
		case e == 2:
		case e == 3:
		case e == 4:
		case e == 5:
			b = 1
			fmt.Println("over!!!")
		case err != nil || e < 1 || e > 5:
			fmt.Println("Warn:请按要求正确输入数字。")
		}
		if b == 1 {
			b = 0
			break
		}
	}
}
