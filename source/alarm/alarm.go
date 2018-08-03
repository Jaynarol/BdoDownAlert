package alarm

import (
	"fmt"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/fatih/color"
	"github.com/jaynarol/BdoDownAlert/source/shutdown"
	"github.com/jaynarol/BdoDownAlert/source/val"
	"jaynarol.com/utility/console"
	"jaynarol.com/utility/messagebox"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const MB_TOPMOST = 0x00040000

var (
	showMessageBox    = false
	unsendLineMessage = true
	situations        = map[bool]map[bool]string{
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

	message := val.TextSituation[section]["message"]
	console.SetTitle(fmt.Sprintf(val.TextTitle, message))
	fmt.Printf("\r")
	log.Printf("%s\r\n", message)

	enableLine := val.Setting.Section(section).Key("line_message").MustBool(false)
	inputToken := len(val.Setting.Section("system").Key("line_token").MustString("")) > 30
	enableSound := val.Setting.Section(section).Key("sound").MustBool(false)
	intervalAlert := val.Setting.Section("interval").Key("alert").RangeInt(10, 3, 86400)
	shutdownSetting := shutdown.Setting(section)
	showMessageBox = true
	unsendLineMessage = true

	go func() {
		for second := 0; showMessageBox == true; second++ {
			if enableSound && second%intervalAlert == 0 {
				playSound()
			}
			if enableLine && inputToken && unsendLineMessage && second%10 == 0 {
				lineNotify(shutdownSetting.Active, message)
			}
			if shutdownSetting.Active {
				color.Set(color.BgRed, color.Bold)
				fmt.Printf(val.TextShutingDown, shutdownSetting.Method, shutdownSetting.Delay-second-1)
				color.Unset()
				if (second+1)%shutdownSetting.Delay == 0 {
					shutdown.Run(shutdownSetting.Method)
				}
			}
			time.Sleep(time.Second)
		}
		fmt.Printf("\r                                                                          ")
	}()

	messagebox.Show(val.AppName, val.TextSituation[section]["popup"]+shutdownSetting.Message, messagebox.MB_ICONEXCLAMATION|MB_TOPMOST|messagebox.MB_OK)
	showMessageBox = false
}

func playSound() {
	f, _ := os.Open(val.FileSound)
	s, format, _ := mp3.Decode(f)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	playing := make(chan struct{})
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		close(playing)
	})))
}

func lineNotify(countdownShutdown bool, text string) {
	if countdownShutdown {
		fmt.Printf("\r")
	}

	client := &http.Client{}
	params := url.Values{}
	params.Set("message", text)
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(params.Encode()))
	if err != nil {
		log.Printf("LINE MESSAGE ERROR: %s\r\n", err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", val.Setting.Section("system").Key("line_token").MustString("")))
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != 200 {
		log.Printf("LINE MESSAGE ERROR: %s - %s\r\n", resp.Status, err)
		return
	}
	log.Print("send Line Message successful\r\n")
	unsendLineMessage = false
}
