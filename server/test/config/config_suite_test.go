package config_test

import (
	"server/src/config"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Config Suite")
}

var _ = Describe("Connection to dockers service", func() {
	Context("Connect to Db", func() {
		It("should return the correct result", func() {
			config.InitDB()
		})
	})

	Context("Connect to Redis", func() {
		It("should return the correct result", func() {
			config.InitRedis()
		})
	})
})
