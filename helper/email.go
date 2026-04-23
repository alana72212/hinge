package helper

import (
	"encoding/json"
	"fmt"
	"hinge/structs"
	"time"
)

type Email struct {
	Mail   string
	Key    string
	MailId int
}

func (e *Email) GetMail() string {
	req := Request(fmt.Sprintf("https://smsbower.online/api/mail/getActivation?api_key=%v&service=vz&domain=gmail.com", e.Key))
	var resp structs.EmailAPI
	json.Unmarshal(req, &resp)
	if resp.Status == 0 {
		return resp.Error
	}
	e.Mail = resp.Mail
	e.MailId = resp.MailId
	return resp.Mail
}

func (e *Email) GetCode() string {
	for {
		req := Request(fmt.Sprintf("https://smsbower.org/api/mail/getCode?api_key=%v&mailId=%v", e.Key, e.MailId))
		var resp structs.EmailAPI
		json.Unmarshal(req, &resp)
		if resp.Status == 0 {
			if resp.Error == "Code has not been received yet, please try again later" {
				time.Sleep(500 * time.Millisecond)
			} else {
				return resp.Error
			}
		} else {
			return resp.Code
		}
	}
}
