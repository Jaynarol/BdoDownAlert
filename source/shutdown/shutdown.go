package shutdown

import (
	"fmt"
	"github.com/jaynarol/BdoDownAlert/source/val"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func Setting(section string) val.ShutdownSetting {
	enableShutdown := val.Setting.Section(section).Key("shutdown").MustString("0")
	shutdownMethod := val.Setting.Section(section).Key("shutdown_method").In("shutdown", []string{"shutdown", "hibernate"})
	shutdownDelay := val.Setting.Section(section).Key("shutdown_delay").RangeInt(5, 5, 86400)
	activeShutdowm := Should(time.Now(), enableShutdown)
	messageShutdowm := fmt.Sprintf(val.TextShutdown, shutdownMethod, shutdownDelay)

	return val.ShutdownSetting{
		Active:  activeShutdowm,
		Method:  shutdownMethod,
		Delay:   shutdownDelay,
		Message: map[bool]string{true: messageShutdowm, false: ""}[activeShutdowm],
	}
}

func Should(timeNow time.Time, setting string) bool {
	rex := regexp.MustCompile("^([0-2]?[0-9]):([0-5]?[0-9])-([0-2]?[0-9]):([0-5]?[0-9])$")
	if rex.MatchString(setting) {
		durationRange := rex.FindStringSubmatch(setting)
		durationStart, _ := time.ParseDuration(fmt.Sprintf("%sh%sm", durationRange[1], durationRange[2]))
		durationEnd, _ := time.ParseDuration(fmt.Sprintf("%sh%sm", durationRange[3], durationRange[4]))

		timeMidnight := time.Date(timeNow.Year(), timeNow.Month(), timeNow.Day(), 0, 0, 0, 0, time.Local)
		timeStart := timeMidnight.Add(durationStart)
		timeEnd := timeMidnight.Add(durationEnd)

		isOverNight := timeStart.Unix() > timeEnd.Unix()
		overStart := timeNow.Unix() >= timeStart.Unix()
		underEnd := timeNow.Unix() <= timeEnd.Unix()

		return map[bool]bool{true: overStart || underEnd, false: overStart && underEnd}[isOverNight]
	}
	return setting == "1"
}

func Run(method string) {
	cmd := val.ShutdownCmd[method]
	fmt.Printf("\r")
	log.Printf("Computer %s...\r\n", val.ShutdownDoing[method])
	err := exec.Command(cmd[0], cmd[1:]...).Start()
	if err != nil {
		log.Printf("%s ERROR: %s\r\n", strings.ToUpper(method), err)
	}
	os.Exit(0)
}
