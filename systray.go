package main

import (
	"fmt"
	"time"

	"github.com/getlantern/systray"
)

var status = 1

func main() {
	onExit := func() {
		status = 2
	}
	loadConf()
	systray.Run(onReady, onExit)
}

func onReady() {
	//systray.SetTemplateIcon(icon.Data, icon.Data)
	systray.SetTemplateIcon(picok, picok)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")

	go func() {
		total := len(ips)
		if total == 0 {
			return
		}

		for {
			time.Sleep(1 * time.Second)
			t := time.Now()
			sec := time.Now().Second()
			str := fmt.Sprintf(t.Format("15:04:05"))
			if status == 2 {
				return
			} else if status == 0 {
				systray.SetTooltip(str + "\nDisabled")
			} else {
				if sec%20 == 0 {
					go func() {
						for i := 0; i < total; i++ {
							ips[i].pingIP()
						}
					}()

				}
				down := 0
				up := 0
				badip := ""
				for i := 0; i < total; i++ {
					if ips[i].avgRtt == -1 {
						down += 1
						if down < 3 { //只能显示3个down掉的IP
							badip += ips[i].ipv4.To4().String() + " "
						} else if down < 4 {
							badip += ips[i].ipv4.To4().String()
						}
					} else {
						up += 1
					}
				}
				str += fmt.Sprintf("\nUp:%d Down:%d", up, down)

				if up == total {
					systray.SetIcon(picok)
				} else {
					systray.SetIcon(picnok)
					str += "\n" + badip
				}
				systray.SetTooltip(str)
			}
		}
	}()

	// We can manipulate the systray in other goroutines
	go func() {
		systray.SetTemplateIcon(picok, picok)
		systray.SetTitle("Awesome App")
		systray.SetTooltip("Working...")

		mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		mDisabled := systray.AddMenuItem("Disabled", "Disabled")
		systray.AddMenuItem("Version:2024-02-22", "Ignored")
		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

		for {
			select {
			case <-mEnabled.ClickedCh:
				status = 1
			case <-mDisabled.ClickedCh:
				status = 0
			case <-mQuit.ClickedCh:
				systray.Quit()
				status = 2
				return
			}
		}
	}()
}
