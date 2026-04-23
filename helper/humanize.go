package helper

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Noooste/azuretls-client"
)

type Humanizer struct {
	Auth    string
	Session *azuretls.Session
	DID     string
	IID     string
	SID     string
}

func (c *Humanizer) AddBday(bday string) int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"birthday": bday,
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddGender() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"genderId": 1,
			"genderIdentityId": map[string]interface{}{
				"visible": true,
				"value":   54,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddOrient() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"sexualOrientations": map[string]interface{}{
				"value":   []int{1},
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddLookingFor() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"sexualOrientations": map[string]interface{}{
				"value":   []int{1},
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddRelationship() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"datingIntention": map[string]interface{}{
				"visible": true,
				"value":   2,
			},
			"datingIntentionText": "",
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddRelType() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"relationshipTypeIds": map[string]interface{}{
				"value":   []int{1},
				"visible": true,
			},
			"relationshipTypesText": "",
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddHeight() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"height": 166,
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddEthnicity() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"ethnicitiesText": "",
			"ethnicities": map[string]interface{}{
				"value": []int{
					8,
				},
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddHaveChild() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"children": map[string]interface{}{
				"value":   1,
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddWantChild() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"children": map[string]interface{}{
				"value":   1,
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddFamily() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"familyPlans": map[string]interface{}{
				"visible": true,
				"value":   4,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddHometown() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"hometown": map[string]interface{}{
				"value":   "New York City",
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddWorkplace() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"works": map[string]interface{}{
				"value":   "Panera",
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddJobTitle() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"jobTitle": map[string]interface{}{
				"visible": true,
				"value":   "cashier",
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddSchool() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"educations": map[string]interface{}{
				"visible": true,
				"value": []string{
					"Berkely University",
				},
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddGrad() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"educationAttained": 3,
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddReligion() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"religions": map[string]interface{}{
				"value": []int{
					0,
				},
				"visible": false,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddPolitical() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"politics": map[string]interface{}{
				"value":   0,
				"visible": false,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddDrinking() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"drinking": map[string]interface{}{
				"visible": true,
				"value":   3,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddSmoking() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"smoking": map[string]interface{}{
				"value":   3,
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddWeed() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"marijuana": map[string]interface{}{
				"value":   3,
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddDrugs() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"drugs": map[string]interface{}{
				"visible": true,
				"value":   3,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddLocation() int {
	p := map[string]interface{}{
		"lastActiveOptIn": true,
		"profile": map[string]interface{}{
			"location": map[string]interface{}{
				"latitude":        39.613725399999964,
				"countryShort":    "US",
				"longitude":       -105.0161769,
				"adminArea1Short": "CO",
				"adminArea2":      "Arapahoe County",
				"locality":        "Littleton",
				"postalSuffix":    "1952",
				"postalCode":      "80120",
				"source":          "google",
				"adminArea1Long":  "Colorado",
				"countryLong":     "United States",
				"name":            "Littleton",
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddPronouns() int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"pronouns": map[string]interface{}{
				"value": []int{
					1,
					2,
					3,
				},
				"visible": true,
			},
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) GetUID() string {

	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Get("https://prod-api.hingeaws.net/user/v3")
	if err != nil {
		log.Fatal(err)
	}
	return string(resp.Header["X-User-Id"][0])
}

func (c *Humanizer) AddName(first string, last string) int {
	p := map[string]interface{}{
		"profile": map[string]interface{}{
			"firstName": first,
			"lastName":  last,
		},
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Patch("https://prod-api.hingeaws.net/user/v3", payload)
	if err != nil {
		log.Fatal(err)
	}
	return int(resp.StatusCode)
}

func (c *Humanizer) AddMail(mailAPI Email) (int, string) {
	mail := mailAPI.GetMail()
	p := map[string]interface{}{
		"email": mail,
	}
	payload, _ := json.Marshal(p)
	c.Session.OrderedHeaders = azuretls.OrderedHeaders{
		{"Accept", "*/*"},
		{"Accept-Encoding", "gzip, deflate, br"},
		{"Accept-Language", "en"},
		{"Authorization", "Bearer " + c.Auth},
		{"Connection", "keep-alive"},
		{"Content-Length", fmt.Sprintf("%d", len([]byte(payload)))},
		{"Content-Type", "application/json"},
		{"Host", "prod-api.hingeaws.net"},
		{"User-Agent", "Hinge/11637 CFNetwork/1404.0.5 Darwin/22.3.0"},
		{"X-App-Version", "9.89.1"},
		{"X-Build-Number", "11637"},
		{"X-Device-Id", c.DID},
		{"X-Device-Model", "unknown"},
		{"X-Device-Model-Code", "iPhone14,5"},
		{"X-Device-Platform", "iOS"},
		{"X-Device-Region", "US"},
		{"X-Install-Id", c.IID},
		{"X-OS-Version", "16.3.1"},
		{"X-Session-Id", c.SID},
	}
	resp, err := c.Session.Put("https://prod-api.hingeaws.net/user/v2/profile/email", payload)
	if err != nil {
		log.Fatal(err)
	}
	if int(resp.StatusCode) == 200 {
		p := map[string]interface{}{
			"code": mailAPI.GetCode(),
		}
		payload, _ := json.Marshal(p)
		c.Session.OrderedHeaders[2][1] = fmt.Sprintf("%d", len([]byte(payload)))
		resp_, err := c.Session.Post("https://prod-api.hingeaws.net/codeval/email/validate", payload)
		if err != nil {
			log.Fatal(err)
		}
		return int(resp_.StatusCode), mail
	}
	return 0, ""
}
