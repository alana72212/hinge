package helper

import (
	"fmt"
	"hinge/structs"
	"strings"
	"time"
)

type SMS struct {
	Key string
	NID string
}

func (s *SMS) GetNumber() structs.NumberAPI {
	body := Request(fmt.Sprintf("https://smshub.org/stubs/handler_api.php?api_key=%v&action=getNumber&service=vz", s.Key))
	data := string(body)
	var numberData structs.NumberAPI
	numberData.ID = strings.Split(strings.Split(data, "ACCESS_NUMBER:")[1], ":")[0]
	s.NID = numberData.ID
	numberData.Number = "+" + strings.Split(strings.Split(data, "ACCESS_NUMBER:")[1], ":")[1]
	return numberData
}

func (s *SMS) GetCode() string {
	for {
		resp := string(Request(fmt.Sprintf("https://smshub.org/stubs/handler_api.php?api_key=%v&action=getStatus&id=%v", s.Key, s.NID)))
		if resp == "STATUS_WAIT_CODE" {
			time.Sleep(500 * time.Millisecond)
		} else if strings.Contains(resp, "STATUS_OK") {
			code := strings.Split(resp, "STATUS_OK:")[1]
			return code
		}
	}
}
