package router

import (
	"github.com/gin-gonic/gin"
	"server/src/controllers/v1/action"
	"server/src/controllers/v1/action/actionId"
	"server/src/controllers/v1/flow"
	"server/src/controllers/v1/flow/flowId"
	"server/src/controllers/v1/flowStep"
	"server/src/controllers/v1/oauth/providerSlug"
	"server/src/controllers/v1/providers"
	providerSlugController "server/src/controllers/v1/providers/providerSlug"
	"server/src/controllers/v1/providers/providerSlug/scopes"
	"server/src/controllers/v1/providers/providerSlug/scopes/scopeId"
	"server/src/controllers/v1/users"
	"server/src/controllers/v1/users/userId"
	providerSlug2 "server/src/controllers/v1/users/userId/providerSlug"
	"server/src/middleware"
)

func SetupV1Router(r *gin.RouterGroup) {

	oauthProviderGroup := r.Group("oauth/:provider")
	{
		oauthProviderGroup.GET("/url", providerSlug.URL)
		oauthProviderGroup.GET("/callback", providerSlug.Callback)
	}

	providerGroup := r.Group("providers")
	{
		providerGroup.GET("/", providers.GetAllProviders)
		providerGroup.POST("/", providers.AddProvider)
		providerGroup.DELETE("/:provider", providerSlugController.DeleteProvider)
		providerGroup.GET("/:provider", providerSlugController.GetProviderDetails)
		providerGroup.GET("/:provider/users", providerSlugController.GetUsers)

		scopesGroup := providerGroup.Group("/:provider/scopes")
		{
			scopesGroup.GET("/", scopes.GetAllScopes)
			scopesGroup.POST("/", scopes.AddScope)
			scopesGroup.DELETE("/:scopeId", scopeId.DeleteScope)
			scopesGroup.GET("/:scopeId", scopeId.GetScopeDetails)
			scopesGroup.GET("/:scopeId/users", scopeId.GetUsers)
		}
	}

	actionGroup := r.Group("action")
	{
		actionGroup.GET("/", action.GetAllActions)
		actionGroup.POST("/", action.AddAction)
		actionGroup.DELETE("/:actionId", actionId.DeleteAction)
		fieldsGroup := actionGroup.Group("/:actionId/fields")
		{
			fieldsGroup.GET("/", actionId.GetActionFields)
			fieldsGroup.POST("/", actionId.AddActionField)
		}
	}

	flowGroup := r.Group("flow")
	{
		flowGroup.GET("/", flow.GetAllFlow)
		flowGroup.Use(middleware.UUIDAuthMiddleware())
		flowGroup.POST("/", flow.AddFlow)
		flowGroup.PUT("/:flowId", flowId.UpdateFlow)
	}
	flowStepGroup := r.Group("flowsteps")
	{
		flowStepGroup.GET("/", flowStep.GetAllFlowStep)
		flowStepGroup.POST("/", flowStep.AddFlowStep)
	}

	r.POST("/users/register", users.Register)
	r.POST("/users/login", users.Login)
	userGroup := r.Group("users")
	userGroup.Use(middleware.UUIDAuthMiddleware())
	{
		userGroup.GET("/:userId", userId.GetUser)
		userGroup.PUT("/:userId", userId.UpdateUser)
		userGroup.GET("/:userId/providers", userId.GetProviders)
		userGroup.GET("/:userId/providers/:provider/scopes", providerSlug2.GetScopes)
		userGroup.GET("/:userId/flows", userId.GetFlows)
		userGroup.POST("/logout", users.Logout)
	}
}
