package val

import (
	"gopkg.in/ini.v1"
)

const (
	AppName     = "BdoDownAlert"
	Version     = "v1.1.0"
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
	TextEnjoy  = "  It's time to relax and assign this work to me... Don't worry I will notify if something wrong!\n"

	TextTitle           = AppName + ": BDO [ %s ]"
	TextChecking        = "\rCheck again %d seconds..."
	TextExit            = "Press 'Enter' to continue..."
	TextShutingDown     = "\rComputer will [%s] in [%d] seconds "
	TextRunning         = "Running"
	TextDead            = "Dead"
	TextFailReadSetting = "Fail to read file: %v"
	TextShutdown        = "============================\n\nคอมพิวเตอร์จะ %s ในอีก %d วินาที\n\nหากต้องการยกเลิกให้กดปุ่ม OK หรือ ปล่อยไว้เผื่อให้คอมปิดอัตโนมัติ\n\n"

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
	ShutdownDoing = map[string]string{
		"shutdown":  "shutting down",
		"hibernate": "hibernating",
	}
	ShutdownCmd = map[string][]string{
		"shutdown":  {"cmd", "/C", "shutdown", "-s -f -t 0"},
		"hibernate": {"cmd", "/C", "shutdown", "-h -f -t 0"},
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

type ShutdownSetting struct {
	Active  bool
	Method  string
	Delay   int
	Message string
}
