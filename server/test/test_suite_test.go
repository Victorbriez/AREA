package test_test

import (
	"context"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"net/http/httptest"
	"server/src/config"
	v1 "server/src/controllers/v1"
	"server/src/middleware"
	"server/src/router"
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestMainApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Test Suite")
}

var _ = Describe("Test for the whole app", func() {
	config.InitDB()
	config.InitRedis()

	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middleware.CORSMiddleware())
	r.GET("/about.json", v1.About)
	v1Router := r.Group("/v1")
	router.SetupV1Router(v1Router)

	Context("Run the whole app.", func() {
		It("Basic test to check that about still works", func() {
			req, _ := http.NewRequest(http.MethodGet, "/about.json", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("Basic test to check that url are correctly generated (register)", func() {
			req, _ := http.NewRequest(http.MethodGet, "/v1/oauth/google/url?type=register", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("Basic test to check that url are correctly generated (login)", func() {
			req, _ := http.NewRequest(http.MethodGet, "/v1/oauth/google/url?type=login", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("Basic test to check that url are block if not log (link)", func() {
			req, _ := http.NewRequest(http.MethodGet, "/v1/oauth/google/url?type=link", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})

		It("Basic test to check that the state is not valid (callback)", func() {
			req, _ := http.NewRequest(http.MethodGet, "/v1/oauth/google/callback?code=token&state=state", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})

		It("Basic test to check that the code is not valid (callback)", func() {
			config.Redis.Set(context.Background(), "fake_state", "testing", time.Minute*1)

			req, _ := http.NewRequest(http.MethodGet, "/v1/oauth/google/callback?code=token&state=fake_state", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusInternalServerError))
		})
	})
})
