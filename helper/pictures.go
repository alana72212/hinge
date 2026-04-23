package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/Noooste/azuretls-client"
)

type Pictures struct {
	Session *azuretls.Session
	Token   string
	Stamp   string
	Status  string
	Key     string
	Auth    string
	DID     string
	IID     string
	SID     string
	UID     string
}

func (p *Pictures) GetToken() {
	p_ := map[string]interface{}{
		"params": map[string]interface{}{
			"tags":  p.UID,
			"phash": "true",
		},
	}
	payload, _ := json.Marshal(p_)
	p.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + p.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", p.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", p.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", p.SID},
	}
	resp, err := p.Session.Post("https://prod-api.hingeaws.net/cdn/token", payload)
	if err != nil {
		log.Fatal(err)
	}
	var r map[string]string
	err = json.Unmarshal(resp.Body, &r)
	if err != nil {
		log.Fatal(err)
	}
	p.Token = r["token"]
	p.Stamp = r["timestamp"]
	p.Key = r["key"]
	p.Status = resp.Header.Get("X-Auth-State-Id")
}

func (p *Pictures) Upload(photo string) (int, string) {
	file, err := os.Open(photo)
	if err != nil {
		return 0, ""
	}
	defer file.Close()
	boundary := "jkL9234nSKDJfLSKM8u12l31238SDIUVd8834nSKDjNKS"
	var b bytes.Buffer
	writeField := func(name, value string) {
		b.WriteString("--" + boundary + "\r\n")
		b.WriteString(`Content-Disposition: form-data; name="` + name + `"` + "\r\n\r\n")
		b.WriteString(value + "\r\n")
	}
	writeField("signature", p.Token)
	writeField("api_key", p.Key)
	writeField("timestamp", p.Stamp)
	b.WriteString("--" + boundary + "\r\n")
	b.WriteString(`Content-Disposition: form-data; name="file"; filename="myUpload"` + "\r\n")
	b.WriteString("Content-Type: image/jpeg\r\n\r\n")
	io.Copy(&b, file)
	b.WriteString("\r\n")
	writeField("tags", p.UID)
	writeField("phash", "true")
	b.WriteString("--" + boundary + "--\r\n")
	p.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en-US,en;q=0.9"},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", b.Len())},
		{"Content-Type", "multipart/form-data; boundary=" + boundary},
		{"Host", "upload.hingenexus.com"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
	}
	resp, err := p.Session.Post("https://upload.hingenexus.com/v1_1/hinge/image/upload", b.Bytes())
	if err != nil {
		return 0, ""
	}
	var r map[string]interface{}
	err = json.Unmarshal(resp.Body, &r)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode), r["secure_url"].(string)
}
