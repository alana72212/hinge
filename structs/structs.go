package structs

type SendCode struct {
	Number   string `json:"phoneNumber"`
	DeviceID string `json:"deviceId"`
}

type OTP struct {
	DeviceID  string `json:"deviceId"`
	InstallID string `json:"installId"`
	Number    string `json:"phoneNumber"`
	OTP       string `json:"otp"`
}

type NumberAPI struct {
	ID     string
	Number string
}

type RegisterResp struct {
	IdentityID string `json:"identityId"`
	Token      string `json:"token"`
	Exp        string `json:"expires"`
	IsMigrated bool   `json:"migratedUser"`
}

type EmailAPI struct {
	Status int    `json:"status"`
	Mail   string `json:"mail"`
	MailId int    `json:"mailId"`
	Error  string `json:"error"`
	Code   string `json:"code"`
}
