package main

import (
	"encoding/json"
	"hinge/helper"
	"hinge/structs"
	"log"

	"github.com/Noooste/azuretls-client"
)

type Hinge struct {
	session      *azuretls.Session
	session_id   string
	device_id    string
	install_id   string
	phone_number string
	NID          string
	UID          string
	logger       *log.Logger
}

func (h *Hinge) sendSMS() int {
	h.session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Connection", "keep-alive"},
		{"Content-Type", "application/json"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", h.device_id},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", h.session_id},
	}
	p := structs.SendCode{
		Number:   h.phone_number,
		DeviceID: h.device_id,
	}
	payload, _ := json.Marshal(p)
	resp, err := h.session.Post("https://prod-api.hingeaws.net/auth/sms/v2/initiate", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (h *Hinge) register(otp string) structs.RegisterResp {
	h.session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Connection", "keep-alive"},
		{"Content-Length", "146"},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", h.device_id},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", h.install_id},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", h.session_id},
	}
	p := structs.OTP{
		DeviceID:  h.device_id,
		InstallID: h.install_id,
		Number:    h.phone_number,
		OTP:       otp,
	}
	payload, _ := json.Marshal(p)
	resp, err := h.session.Post("https://prod-api.hingeaws.net/auth/sms/v2", payload)
	if err != nil {
		log.Fatal(err)
	}
	var accData structs.RegisterResp
	_err := json.Unmarshal(resp.Body, &accData)
	if _err != nil {
		log.Fatal(_err)
	}
	return accData
}

func (h *Hinge) install() int {
	h.session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Connection", "keep-alive"},
		{"Content-Length", "52"},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", h.device_id},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", h.session_id},
	}
	p := map[string]string{
		"installId": h.install_id,
	}
	payload, _ := json.Marshal(p)
	resp, err := h.session.Post("https://prod-api.hingeaws.net/identity/install", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (h *Hinge) humanize(user structs.RegisterResp) int {
	c := helper.Humanizer{
		Auth:    user.Token,
		Session: h.session,
		DID:     h.device_id,
		IID:     h.install_id,
		SID:     h.session_id,
	}

	// The logging below is for debugging reasons while testing if specific functions worked.
	// You may comment out, delete, or whatever else cus they aren't needed
	// Also lots of this is static, you should not keep it static

	first, last := "Jacob", "Staples"
	if name := c.AddName(first, last); name == 200 {
		h.logger.Printf("Successfuly Added Name -> %v, %v", first, last)
	}
	h.UID = c.GetUID()
	m := helper.Email{
		Key: "API_Key", // Obtain key from smsbower.online
	}
	if status, mail := c.AddMail(m); status == 200 {
		h.logger.Printf("Successfuly Verified Mail -> %v", mail)
	} else {
		h.logger.Printf("Error Adding Mail -> %v", status)
	}
	bday1 := "2000-03-23T00:00:00Z"
	if bday := c.AddBday(bday1); bday == 200 {
		h.logger.Printf("Successfuly Added Birthday -> %v", bday1)
	}
	if addy := c.AddLocation(); addy == 200 {
		h.logger.Printf("Successfuly Added Location -> %v, %v", 39.613725399999964, -105.0161769)
	}
	if pronouns := c.AddPronouns(); pronouns == 200 {
		h.logger.Printf("Successfuly Added Pronouns -> %v, %v, %v", 1, 2, 3)
	}
	if Gender := c.AddGender(); Gender == 200 {
		h.logger.Printf("Successfuly Added Gender -> %v", 1)
	}
	if sOrient := c.AddOrient(); sOrient == 200 {
		h.logger.Printf("Successfuly Added Orientation -> %v", 1)
	}
	if lookingFor := c.AddLookingFor(); lookingFor == 200 {
		h.logger.Printf("Successfuly Added Looking For -> %v", 1)
	}
	if relationship := c.AddRelationship(); relationship == 200 {
		h.logger.Printf("Successfuly Added Relationship -> %v", 2)
	}
	if reltype := c.AddRelType(); reltype == 200 {
		h.logger.Printf("Successfuly Added Relationship Type -> %v", 1)
	}
	if height := c.AddHeight(); height == 200 {
		h.logger.Printf("Successfuly Added Height -> %vcm", 166)
	}
	if eth := c.AddEthnicity(); eth == 200 {
		h.logger.Printf("Successfuly Added Ethnicity -> %v", 8)
	}
	if hc := c.AddHaveChild(); hc == 200 {
		h.logger.Printf("Successfuly Added Has Child -> %v", 1)
	}
	if wc := c.AddWantChild(); wc == 200 {
		h.logger.Printf("Successfuly Added Want Child -> %v", 1)
	}
	if fam := c.AddWantChild(); fam == 200 {
		h.logger.Printf("Successfuly Added Family -> %v", 4)
	}
	if ht := c.AddHometown(); ht == 200 {
		h.logger.Printf("Successfuly Added Hometown -> %v", "NYC")
	}
	if wp := c.AddWorkplace(); wp == 200 {
		h.logger.Printf("Successfuly Added Workplace -> %v", "Panera")
	}
	if wt := c.AddJobTitle(); wt == 200 {
		h.logger.Printf("Successfuly Added Job Title -> %v", "Cashier")
	}
	if sc := c.AddSchool(); sc == 200 {
		h.logger.Printf("Successfuly Added School -> %v", "UC Berkely")
	}
	if grad := c.AddGrad(); grad == 200 {
		h.logger.Printf("Successfuly Added Grad -> %v", "Postgrad")
	}
	if rel := c.AddReligion(); rel == 200 {
		h.logger.Printf("Successfuly Added Religion -> %v", "Rather not say")
	}
	if pol := c.AddPolitical(); pol == 200 {
		h.logger.Printf("Successfuly Added Political -> %v", "Rather not say")
	}
	if drink := c.AddDrinking(); drink == 200 {
		h.logger.Printf("Successfuly Added Drinking -> %v", 3)
	}
	if smoke := c.AddSmoking(); smoke == 200 {
		h.logger.Printf("Successfuly Added Smoking -> %v", 3)
	}
	if weed := c.AddWeed(); weed == 200 {
		h.logger.Printf("Successfuly Added Weed -> %v", 3)
	}
	if drugs := c.AddDrugs(); drugs == 200 {
		h.logger.Printf("Successfuly Added Drugs -> %v", 3)
	}
	return 0
}

func gen(proxy string) {
	sms := helper.SMS{
		Key: "API_KEY", // Obtain key from smshub.org
	}
	h := Hinge{
		session_id: helper.UUID(),
		device_id:  helper.UUID(),
		install_id: helper.UUID(),
		logger:     log.Default(),
	}
	session := azuretls.NewSession()
	defer session.Close()
	session.Browser = azuretls.Ios
	if proxy != "" {
		err := session.SetProxy(proxy)
		if err != nil {
			log.Fatal(err)
		}
	}
	number := sms.GetNumber()
	h.phone_number = number.Number
	h.logger.Printf("Got Number -> %v", h.phone_number)
	h.session = session
	sentSMS := h.sendSMS()
	if sentSMS == 200 {
		h.logger.Printf("Sent OTP -> %v", h.phone_number)
		if h.install() == 201 {
			otp := sms.GetCode()
			h.logger.Printf("Got OTP -> %v", otp)
			reg := h.register(otp)
			h.logger.Printf("Successfuly Registered -> %v", reg.Token)
			h.humanize(reg)
			p := helper.Pictures{
				Session: h.session,
				Auth:    reg.Token,
				DID:     h.device_id,
				IID:     h.install_id,
				SID:     h.session_id,
				UID:     h.UID,
			}
			p.GetToken()
			if pfp, pUrl := p.Upload("./1.jpg"); pfp == 200 {
				h.logger.Printf("Successfuly Added Picture -> %v", pUrl)
			}
		} else {
			h.logger.Println("Failed to install")
		}
	} else {
		h.logger.Println("Failed to send OTP")
	}
}

func main() {
	gen("") // proxyless
	//gen("http://username:password@ip:port") // with proxy
}
