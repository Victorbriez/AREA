package middleware_test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"server/src/config"
	v1 "server/src/controllers/v1"
	"server/src/middleware"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMiddleware(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Middleware Suite")
}

var _ = Describe("Test for scopeId endpoints", Label("controller"), func() {
	config.InitDB()
	config.InitRedis()
	gin.SetMode(gin.TestMode)

	Context("Basic middleware tests", func() {
		It("Is logged false", func() {
			r := gin.New()
			r.Use(middleware.UUIDAuthMiddleware())
			r.GET("/about.json", v1.About)

			req, _ := http.NewRequest(http.MethodGet, "/about.json", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})

		It("Is Admin false", func() {
			r := gin.New()
			r.Use(middleware.UUIDAuthMiddleware())
			r.Use(middleware.AdminMiddleware())
			r.GET("/about.json", v1.About)

			req, _ := http.NewRequest(http.MethodGet, "/about.json", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusUnauthorized))

		})
	})
})
