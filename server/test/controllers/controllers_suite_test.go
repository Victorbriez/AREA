package controllers_test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	v1 "server/src/controllers/v1"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}

var _ = Describe("controller", Label("controller"), func() {
	var router *gin.Engine

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		router = gin.New()
		router.GET("/about.json", v1.About)
	})

	Context("simple GET on /v1/about.json", func() {
		It("should return the correct result", func() {
			req, _ := http.NewRequest(http.MethodGet, "/about.json", nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusOK))
			Expect(w.Header().Get("Content-Type")).To(Equal("application/json; charset=utf-8"))
		})
	})
})
