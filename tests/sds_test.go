package tests

import (
	"fmt"

	"github.com/go-resty/resty"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ bool = Describe("Sds", func() {

	BeforeEach(func() {
		token = GetToken()
		fmt.Println(token)
	})

	resp, err := resty.R().
		SetHeaders(
			map[string]string{
				//"Content-Type": "application/json",
				"LOG_USER":     "admin",
				"X-Auth-Token": token,
			}).
		Get(SDSPreURL + "softwareversion")
	if err != nil {
	}
	fmt.Println(resp)

	It("should populate the fields correctly", func() {
		Expect(resp.StatusCode()).To(Equal(200))
	})
})
