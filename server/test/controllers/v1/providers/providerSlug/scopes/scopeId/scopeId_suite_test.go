package scopeId_test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"server/src/config"
	"server/src/controllers/v1/providers/providerSlug/scopes"
	"server/src/controllers/v1/providers/providerSlug/scopes/scopeId"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestScopeId(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ScopeId Suite")
}

var _ = Describe("Test for scopeId endpoints", Label("controller"), func() {
	config.InitDB()
	config.InitRedis()
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.GET("/v1/providers/:provider/scopes", scopes.GetAllScopes)
	r.GET("/v1/providers/:provider/scopes/:scopeId", scopeId.GetScopeDetails)

	Context("Basic scopes id", func() {
		It("List ids", func() {
			req, _ := http.NewRequest(http.MethodGet, "/v1/providers/google/scopes", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("Get details of scopes id", func() {
			req, _ := http.NewRequest(http.MethodGet, "/v1/providers/google/scopes", nil)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
			req, _ = http.NewRequest(http.MethodGet, "/v1/providers/google/scopes/1", nil)
			w = httptest.NewRecorder()
			r.ServeHTTP(w, req)
			Expect(w.Code).To(Equal(http.StatusOK))
		})
	})
})
