package sdshttp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSdshttp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sdshttp Suite")
}
