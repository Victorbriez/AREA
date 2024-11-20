package dto_test

import (
	"server/src/models/dto"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestDto(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dto Suite")
}

var _ = Describe("Test for Data type object", func() {
	Context("Error", func() {
		It("should return the correct result", func() {
			requestError := dto.Error("An error. Woupsi")
			Expect(requestError.Error).To(Equal("An error. Woupsi"))
		})
	})
})
