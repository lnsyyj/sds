package tests

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty"
	. "github.com/lnsyyj/sds/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tidwall/gjson"
)

func TestSds(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sds Suite")
}

var _ = BeforeSuite(func() {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{
			"auth": {
				"tenantName": "admin",
					"passwordCredentials": {
						"username": "admin",
						"password": "Admin_123456"
					}
			}
		}`).
		Post(TokenPreURL + "tokens")

	if err != nil {
	}

	fmt.Println(resp)
	Token = gjson.Get(resp.String(), "access.token.id").String()
	fmt.Printf("get sds token: %s", Token)
	//value := gjson.Get(resp.String(), "access.token.id")
	//fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++")
	//fmt.Println(value)
	//fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++")
})
