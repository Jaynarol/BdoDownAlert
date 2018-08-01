package val

import (
	"gopkg.in/ini.v1"
)

const (
	AppName     = "BdoDownAlert"
	Version     = "v1.0.0"
	FileSetting = "setting.ini"
	FileSound   = "assets/alarm.mp3"
	Developer   = "Jaynarol"
	BdoTHFamily = "Noxia"

	TextWelcome = "\n" +
		"   ____       _         _____                                    _              _\n" +
		"  |  _ \\     | |       |  __ \\                            /\\    | |            | |\n" +
		"  | |_) |  __| |  ___  | |  | |  ___ __      __ _ __     /  \\   | |  ___  _ __ | |_\n" +
		"  |  _ <  / _` | / _ \\ | |  | | / _ \\\\ \\ /\\ / /| '_ \\   / /\\ \\  | | / _ \\| '__|| __|\n" +
		"  | |_) || (_| || (_) || |__| || (_) |\\ V  V / | | | | / ____ \\ | ||  __/| |   | |_\n" +
		"  |____/  \\__,_| \\___/ |_____/  \\___/  \\_/\\_/  |_| |_|/_/    \\_\\|_| \\___||_|    \\__|\n "

	TextCredit = "  " + AppName + " " + Version + " Powered by [%s] in BdoTH family is [%s]\n\n"
	TextEnjoy = "  It's time to relax and assign this work to me... Don't worry I will notify if something wrong!"

	TextTitle           = AppName + ": BDO [ %s ] Check again %d seconds..."
	TextTitle2          = AppName + ": BDO [ %s ]"
	TextExit            = "Press 'Enter' to continue..."
	TextRunning         = "Running"
	TextDead            = "Dead"
	TextFailReadSetting = "Fail to read file: %v"

	SituationStillRuning = "StillRuning"
	SituationDying       = "Dying"
	SituationStarting    = "Starting"
	SituationKeepDead    = "KeepDead"
)

var (
	PingCmd  = [4]string{"cmd", "/C", "netstat", "-ano"}
	BdoPorts = [2]string{":8888 ", ":8889 "}
	Setting  = ini.Empty()
	Status   = map[bool]string{
		true:  TextRunning,
		false: TextDead,
	}
	TextSituation = map[string]map[string]string{
		"reconnect_alert": {
			"message": "BDO Reconnect!",
			"popup":   "                Black Desert Online                \n\n                [ - Reconnect - ]                \n\n",
		},
		"disconnect_alert": {
			"message": "BDO Disconnect!",
			"popup":   "                Black Desert Online                \n\n                [ - Disconnect - ]                \n\n",
		},
	}
)

type LastStatus struct {
	Alive bool
	Port  string
}

type Client struct {
	Found bool
	Port  string
}
