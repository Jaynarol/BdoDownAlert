package line

import (
	"fmt"
	"github.com/jaynarol/BdoDownAlert/source/val"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func Notify(isCoundown bool, text string, sendLineMessage *bool) {
	if isCoundown {
		fmt.Printf("\r")
	}

	resp, err := makeRequest(text)
	defer resp.Body.Close()
	if err != nil || resp.StatusCode != 200 {
		log.Printf("LINE MESSAGE ERROR: %s - %s\r\n", resp.Status, err)
		return
	}
	log.Print("send Line Message successful\r\n")
	*sendLineMessage = true
}

func makeRequest(text string) (*http.Response, error) {

	client := &http.Client{}
	params := url.Values{}
	params.Set("message", text)
	req, err := http.NewRequest("POST", "https://notify-api.line.me/api/notify", strings.NewReader(params.Encode()))
	if err != nil {
		log.Printf("LINE MESSAGE ERROR: %s\r\n", err)
		return nil, err
	}

	tokenHeader := fmt.Sprintf("Bearer %s", val.Setting.Section("system").Key("line_token").MustString(""))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", tokenHeader)

	return client.Do(req)
}
