package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"time"
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
			str := fmt.Sprintf(t.Format("time 15:04:05"))
			if status == 2 {
				return
			} else if status == 0 {
				systray.SetTooltip(str + "\nDisabled")
			} else {
				flag := false
				if sec%10 == 0 {
					flag = true
				}
				stat := 0
				for _, elem := range ips {
					if flag {
						elem.pingIP()
						if elem.avgRtt != -1 {
							stat++
						}
					}
					str += fmt.Sprintf("\n%s: %dms", elem.ipv4.String(), elem.avgRtt)
				}
				if flag {
					if stat == total {
						systray.SetIcon(picok)
					} else {
						systray.SetIcon(picnok)
					}
					systray.SetTooltip(str)
				}
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
		systray.AddMenuItem("Version:2024-02-19", "Ignored")
		mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

		systray.AddSeparator()

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
