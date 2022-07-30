package main

import (
	"Simple-Emu-Downloader/downloader"
	"Simple-Emu-Downloader/get"
	"fmt"
	"github.com/gookit/color"
	"strconv"
)

var emu, input string
var breaker int

func main() {
	c := color.HEX("#7fecad")
	for {
		fmt.Println()
		color.HEX("#eedeb0").Println("	* 欢迎使用Simple Emu Downloader by lkidszzzz *")
		fmt.Println()
		color.Set(color.FgLightCyan)
		fmt.Println("	输入 1 获取 Nintendo Switch 模拟器（推荐）：Yuzu")
		fmt.Println("	输入 2 获取 Nintendo Switch 模拟器：Ryujinx")
		fmt.Println("	输入 3 获取 Nintendo 3DS 模拟器：Citra3DS")
		fmt.Println("	输入 4 获取 Sony PSP 模拟器：PPSSPP")
		fmt.Println("	输入 5 退出下载器。")
		fmt.Print("	请输入对应的数字以选择您要下载的模拟器：")
		_, err := fmt.Scan(&emu)
		if err != nil {
			fmt.Println("Unknown err.")
			return
		}
		e, err := strconv.Atoi(emu)
		switch {
		case e == 1:
			for {
				fmt.Println()
				color.HEX("#eedeb0").Println("	Yuzu-下载：")
				fmt.Println()
				color.Set(color.FgLightCyan)
				fmt.Println("	输入 1 选择Github普通下载通道（国内部分地区可能体验不佳甚至无法访问,若自己有代理&加速服务则优先选择）；")
				fmt.Println("	输入 2 选择Github公共CDN下载通道1（推荐）（使用人数较多时可能速度较慢，请避开高峰期使用）。")
				fmt.Println("	输入 3 选择Github公共CDN下载通道2（使用人数较多时可能速度较慢，请避开高峰期使用）。")
				fmt.Println("	输入 4 返回。")
				fmt.Printf("	请输入对应的数字以选择下载通道：")
				_, err := fmt.Scan(&input)
				if err != nil {
					fmt.Println("Unknown err.")
					return
				}
				e, err := strconv.Atoi(input)
				switch {
				case e == 1:
					link := get.YuzuGetter()
					fmt.Println()
					fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
					c.Println(link)
					fmt.Println()
					d := downloader.NewDownloader("./")
					d.AppendResource(downloader.Namer(link), link)
					d.Start()
					breaker = 1
				case e == 2:
					link := get.YuzuCDN1()
					fmt.Println()
					fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
					c.Println(link)
					fmt.Println()
					d := downloader.NewDownloader("./")
					d.AppendResource(downloader.Namer(link), link)
					d.Start()
					breaker = 1
				case e == 3:
					link := get.YuzuCDN2()
					fmt.Println()
					fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
					c.Println(link)
					fmt.Println()
					d := downloader.NewDownloader("./")
					d.AppendResource(downloader.Namer(link), link)
					d.Start()
					breaker = 1
				case e == 4:
					breaker = 1
				case err != nil || e < 1 || e > 3:
					color.Error.Prompt("	请按要求正确输入数字。")
				}
				if breaker == 1 {
					breaker = 0
					break
				}
			}
		case e == 2:
			for {
				fmt.Println()
				color.HEX("#eedeb0").Println("	Ryujinx-下载：")
				fmt.Println()
				color.Set(color.FgLightCyan)
				fmt.Println("	输入 1 选择Github普通下载通道（国内部分地区可能体验不佳甚至无法访问,若自己有代理&加速服务则优先选择）；")
				fmt.Println("	输入 2 选择Github公共CDN下载通道1（推荐）（使用人数较多时可能速度较慢，请避开高峰期使用）。")
				fmt.Println("	输入 3 选择Github公共CDN下载通道2（使用人数较多时可能速度较慢，请避开高峰期使用）。")
				fmt.Println("	输入 4 返回。")
				fmt.Printf("	请输入对应的数字以选择下载通道：")
				_, err := fmt.Scan(&input)
				if err != nil {
					fmt.Println("Unknown err.")
					return
				}
				e, err := strconv.Atoi(input)
				switch {
				case e == 1:
					link := get.RyuGetter()
					fmt.Println()
					fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
					c.Println(link)
					fmt.Println()
					d := downloader.NewDownloader("./")
					d.AppendResource(downloader.Namer(link), link)
					d.Start()
					breaker = 1
				case e == 2:
					link := get.RyuCDN1()
					fmt.Println()
					fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
					c.Println(link)
					fmt.Println()
					d := downloader.NewDownloader("./")
					d.AppendResource(downloader.Namer(link), link)
					d.Start()
					breaker = 1
				case e == 3:
					link := get.RyuCDN2()
					fmt.Println()
					fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
					c.Println(link)
					fmt.Println()
					d := downloader.NewDownloader("./")
					d.AppendResource(downloader.Namer(link), link)
					d.Start()
					breaker = 1
				case e == 4:
					breaker = 1
				case err != nil || e < 1 || e > 4:
					color.Error.Prompt("	Warn:请按要求正确输入数字。")
				}
				if breaker == 1 {
					breaker = 0
					break
				}
			}
		case e == 3:
			for {
				fmt.Println()
				color.HEX("#eedeb0").Println("	Citra3DS-下载：")
				fmt.Println()
				color.Set(color.FgLightCyan)
				fmt.Println("	输入 1 选择Citra-Nightly（更新较慢，相当于稳定版）；")
				fmt.Println("	输入 2 选择Citra-Canary（每日更新，相当于测试版）。")
				fmt.Println("	输入 3 返回。")
				fmt.Printf("	请输入对应的数字以选择您想要下载的版本：")
				_, err := fmt.Scan(&input)
				if err != nil {
					fmt.Println("Unknown err.")
					return
				}
				e, err := strconv.Atoi(input)
				switch {
				case e == 1:
					for {
						fmt.Println()
						color.HEX("#eedeb0").Println("	Citra3DS-Nightly-下载：")
						fmt.Println()
						color.Set(color.FgLightCyan)
						fmt.Println("	输入 1 选择Github普通下载通道（国内部分地区可能体验不佳甚至无法访问,若自己有代理&加速服务则优先选择）；")
						fmt.Println("	输入 2 选择Github公共CDN下载通道1（推荐）（使用人数较多时可能速度较慢，请避开高峰期使用）。")
						fmt.Println("	输入 3 选择Github公共CDN下载通道2（使用人数较多时可能速度较慢，请避开高峰期使用）。")
						fmt.Println("	输入 4 返回。")
						fmt.Printf("	请输入对应的数字以选择下载通道：")
						_, err := fmt.Scan(&input)
						if err != nil {
							fmt.Println("Unknown err.")
							return
						}
						e, err := strconv.Atoi(input)
						switch {
						case e == 1:
							link := get.CitraNightlyGetter(get.CitraNightlyVerGetter())
							fmt.Println()
							fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
							c.Println(link)
							fmt.Println()
							d := downloader.NewDownloader("./")
							d.AppendResource(downloader.Namer(link), link)
							d.Start()
							breaker = 1
						case e == 2:
							link := get.CitraNightlyCDN1()
							fmt.Println()
							fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
							c.Println(link)
							fmt.Println()
							d := downloader.NewDownloader("./")
							d.AppendResource(downloader.Namer(link), link)
							d.Start()
							breaker = 1
						case e == 3:
							link := get.CitraNightlyCDN2()
							fmt.Println()
							fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
							c.Println(link)
							fmt.Println()
							d := downloader.NewDownloader("./")
							d.AppendResource(downloader.Namer(link), link)
							d.Start()
							breaker = 1
						case e == 4:
							breaker = 1
						case err != nil || e < 1 || e > 4:
							color.Error.Prompt("	Warn:请按要求正确输入数字。")
						}
						if breaker == 1 {
							breaker = 0
							break
						}
						breaker = 1
					}
				case e == 2:
					for {
						fmt.Println()
						color.HEX("#eedeb0").Println("	Citra3DS-Canary-下载：")
						fmt.Println()
						color.Set(color.FgLightCyan)
						fmt.Println("	输入 1 选择Github普通下载通道（国内部分地区可能体验不佳甚至无法访问,若自己有代理&加速服务则优先选择）；")
						fmt.Println("	输入 2 选择Github公共CDN下载通道1（推荐）（使用人数较多时可能速度较慢，请避开高峰期使用）。")
						fmt.Println("	输入 3 选择Github公共CDN下载通道2（使用人数较多时可能速度较慢，请避开高峰期使用）。")
						fmt.Println("	输入 4 返回。")
						fmt.Printf("	请输入对应的数字以选择下载通道：")
						_, err := fmt.Scan(&input)
						if err != nil {
							fmt.Println("Unknown err.")
							return
						}
						e, err := strconv.Atoi(input)
						switch {
						case e == 1:
							link := get.CitraCanaryGetter(get.CitraCanaryVerGetter())
							fmt.Println()
							fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
							c.Println(link)
							fmt.Println()
							d := downloader.NewDownloader("./")
							d.AppendResource(downloader.Namer(link), link)
							d.Start()
							breaker = 1
						case e == 2:
							link := get.CitraCanaryCDN1()
							fmt.Println()
							fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
							c.Println(link)
							fmt.Println()
							d := downloader.NewDownloader("./")
							d.AppendResource(downloader.Namer(link), link)
							d.Start()
							breaker = 1
						case e == 3:
							link := get.CitraCanaryCDN2()
							fmt.Println()
							fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
							c.Println(link)
							fmt.Println()
							d := downloader.NewDownloader("./")
							d.AppendResource(downloader.Namer(link), link)
							d.Start()
							breaker = 1
						case e == 4:
							breaker = 1
						case err != nil || e < 1 || e > 4:
							color.Error.Prompt("	Warn:请按要求正确输入数字。")
						}
						if breaker == 1 {
							breaker = 0
							break
						}
						breaker = 1
					}
				case e == 3:
					breaker = 1
				case err != nil || e < 1 || e > 3:
					color.Error.Prompt("	Warn:请按要求正确输入数字。")
				}
				if breaker == 1 {
					breaker = 0
					break
				}
			}
		case e == 4:
			for {
				fmt.Println()
				color.HEX("#eedeb0").Println("	PPSSPP-下载：")
				fmt.Println()
				color.Set(color.FgLightCyan)
				fmt.Println("	输入 1 下载；")
				fmt.Println("	输入 2 返回。")
				fmt.Printf("	请输入对应的数字以选择下载或返回：")
				_, err := fmt.Scan(&input)
				if err != nil {
					fmt.Println("Unknown err.")
					return
				}
				e, err := strconv.Atoi(input)
				switch {
				case e == 1:
					link := get.PpssppGetter()
					fmt.Println()
					fmt.Println("	下载链接（您也可以复制到您喜欢的下载器去进行下载）：")
					c.Println(link)
					fmt.Println()
					d := downloader.NewDownloader("./")
					d.AppendResource(downloader.Namer(link), link)
					d.Start()
					breaker = 1
				case e == 2:
					breaker = 1
				case err != nil || e < 1 || e > 2:
					color.Error.Prompt("	Warn:请按要求正确输入数字。")
				}
				if breaker == 1 {
					breaker = 0
					break
				}
			}
		case e == 5:
			breaker = 1
		case err != nil || e < 1 || e > 5:
			color.Error.Prompt("	Warn:请按要求正确输入数字。")
		}
		if breaker == 1 {
			breaker = 0
			break
		}
	}
}
