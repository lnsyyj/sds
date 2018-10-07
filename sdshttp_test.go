package sdshttp_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Sdshttp", func() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {

	}

	Describe("Get baidu", func() {
		Context("when the JSON parses succesfully", func() {
			It("should populate the fields correctly", func() {
				Expect(resp.Status).To(Equal("200 OK"))
			})

		})
	})
})