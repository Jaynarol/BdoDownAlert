package source

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/jaynarol/BdoDownAlert/source/alarm"
	"github.com/jaynarol/BdoDownAlert/source/command"
	"github.com/jaynarol/BdoDownAlert/source/val"
	"gopkg.in/ini.v1"
	"jaynarol.com/utility/console"
	"log"
	"os"
	"strings"
	"time"
)

var (
	lastStatus = val.LastStatus{}
)

func Main() {
	welcome()
	if loadSettings() && checkSound() {
		loopPing()
	}
	exit()
}

func welcome() {
	yellow := color.New(color.FgYellow).SprintFunc()

	color.Yellow(val.TextWelcome)
	color.Cyan(val.TextEnjoy)
	fmt.Println()
	fmt.Fprintf(color.Output, val.TextCredit, yellow(val.Developer), yellow(val.BdoTHFamily))
	color.Yellow(strings.Repeat("=", 85))
	fmt.Println()
}

func loopPing() {
	for {
		client, err := command.IsAlive()
		if err != nil {
			break
		}

		lastStatus = alarm.ShouldAlert(lastStatus, client)
		updateConsole()
	}
}

func loadSettings() bool {
	cfg, err := ini.Load(val.FileSetting)
	if err != nil {
		log.Printf(val.TextFailReadSetting, err)
		return false
	}
	val.Setting = cfg
	return true
}

func checkSound() bool {
	if _, err := os.Stat(val.FileSound); os.IsNotExist(err) {
		log.Printf(val.TextFailReadSetting, err)
		return false
	}
	return true
}

func updateConsole() {
	intervalPing := val.Setting.Section("interval").Key("ping").RangeInt(10, 3, 86400)
	for second := intervalPing; second > 0; second-- {
		title := fmt.Sprintf(val.TextTitle, val.Status[lastStatus.Alive], second)
		console.SetTitle(title)
		time.Sleep(time.Second)
	}
}

func exit() {
	fmt.Print(val.TextExit)
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
