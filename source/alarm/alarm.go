package alarm

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/go-vgo/robotgo"
	"github.com/jaynarol/BdoDownAlert/source/line"
	"github.com/jaynarol/BdoDownAlert/source/shutdown"
	"github.com/jaynarol/BdoDownAlert/source/sound"
	"github.com/jaynarol/BdoDownAlert/source/val"
	"jaynarol.com/utility/console"
	"jaynarol.com/utility/messagebox"
	"log"
	"strings"
	"time"
)

const MbTopmost = 0x00040000

var (
	showMessageBox  = false
	sendLineMessage = false
	situations      = map[bool]map[bool]string{
		true: {
			true:  val.SituationStillRuning,
			false: val.SituationDying,
		},
		false: {
			true:  val.SituationStarting,
			false: val.SituationKeepDead,
		},
	}
)

func ShouldAlert(lastStatus val.LastStatus, client val.Client) val.LastStatus {
	situation := situations[lastStatus.Alive][client.Found]

	switch situation {
	case val.SituationStarting:
		return val.LastStatus{Alive: true, Port: client.Port}
	case val.SituationStillRuning:
		return checkReconnect(lastStatus, client)
	case val.SituationDying:
		alert("disconnect_alert")
	case val.SituationKeepDead:
	}

	return val.LastStatus{}
}

func checkReconnect(lastStatus val.LastStatus, client val.Client) val.LastStatus {
	if lastStatus.Port != client.Port {
		alert("reconnect_alert")
		return val.LastStatus{Alive: true, Port: client.Port}
	}
	return lastStatus
}

func alert(section string) {

	shortMessage := val.TextSituation[section]["shortMessage"]
	if haveActivity(shortMessage) {
		return
	}

	alertSetting := val.AlertSetting{
		Message:         val.TextSituation[section]["message"],
		EnableLine:      val.Setting.Section(section).Key("line_message").MustBool(false),
		ValidToken:      len(val.Setting.Section("system").Key("line_token").MustString("")) > 30,
		EnableSound:     val.Setting.Section(section).Key("sound").MustBool(false),
		IntervalAlert:   val.Setting.Section("interval").Key("alert").RangeInt(10, 3, 86400),
		ShutdownSetting: shutdown.Setting(section),
	}

	fmt.Printf("\r")
	log.Printf("%s\r\n", alertSetting.Message)
	console.SetTitle(fmt.Sprintf(val.TextTitle, shortMessage))

	go loopAlert(alertSetting)
	showMessagebox(section, alertSetting.ShutdownSetting)
}

func loopAlert(alert val.AlertSetting) {
	shutdownSetting := alert.ShutdownSetting
	sendLineMessage = false

	for second := 0; showMessageBox == true; second++ {
		if alert.EnableSound && second%alert.IntervalAlert == 0 {
			sound.PlaySound()
		}
		if alert.EnableLine && alert.ValidToken && !sendLineMessage && second%10 == 0 {
			line.Notify(shutdownSetting.Active, alert.Message, &sendLineMessage)
		}
		if shutdownSetting.Active {
			color.Set(color.BgRed, color.Bold)
			fmt.Printf(val.TextShutdownCounting, shutdownSetting.Method, shutdownSetting.Delay-second-1)
			color.Unset()
			if (second+1)%shutdownSetting.Delay == 0 {
				shutdown.Run(shutdownSetting.Method)
			}
		}
		time.Sleep(time.Second)
	}
	clearLineConsole()
}

func showMessagebox(section string, shutdownSetting val.ShutdownSetting) {
	showMessageBox = true
	messagebox.Show(val.AppName, val.TextSituation[section]["popup"]+shutdownSetting.Message, messagebox.MB_ICONEXCLAMATION|MbTopmost|messagebox.MB_OK)
	showMessageBox = false
}

func haveActivity(shortMessage string) bool {
	maxIdle := val.Setting.Section("interval").Key("idle").RangeInt(15, 0, 3600)
	x, y := robotgo.GetMousePos()
	foundActivity := false
	for second := maxIdle; second > 0; second-- {
		x2, y2 := robotgo.GetMousePos()
		if x != x2 || y != y2 {
			foundActivity = true
			break
		}
		color.Set(color.FgCyan)
		fmt.Printf(val.TextIdleCounting, shortMessage, second)
		color.Unset()
		time.Sleep(time.Second)
	}
	clearLineConsole()
	return foundActivity
}

func clearLineConsole() {
	fmt.Printf("\r%s", strings.Repeat(" ", 90))
}
