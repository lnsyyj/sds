package tests

import (
	"fmt"

	"github.com/go-resty/resty"
	. "github.com/lnsyyj/sds/config"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ bool = Describe("Sds", func() {

	resp, err := resty.R().
		SetHeaders(
			map[string]string{
				//"Content-Type": "application/json",
				"LOG_USER":     "admin",
				"X-Auth-Token": Token,
			}).
		Get(SDSPreURL + "softwareversion")
	if err != nil {
	}
	fmt.Println(resp)

	It("should populate the fields correctly", func() {
		Expect(resp.StatusCode()).To(Equal(200))
	})
})
