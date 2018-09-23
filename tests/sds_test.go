package tests

import (
	"fmt"

	"github.com/go-resty/resty"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ bool = Describe("Sds", func() {
	var response = &resty.Response{}
	BeforeEach(func() {
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
		response = resp
	})

	Describe("Categorizing book length", func() {
		Context("The first book", func() {
			It("should populate the fields correctly", func() {
				Expect(response.StatusCode()).To(Equal(200))
				fmt.Println(response.String())
				fmt.Println(response.Body())
			})
		})
	})
})
