package timelinehandlers

import (
	"fin4-core/server/datatype"
	"fin4-core/server/routermiddleware"
	"github.com/gin-gonic/gin"
)

//InjectHandlers injects asset handlers into the application
func InjectHandlers(sc datatype.ServiceContainer, rg *gin.RouterGroup) {
	authenticator := routermiddleware.SessionMustAuth()
	rg.GET("/v2/timeline", authenticator, FindEntries(sc, datatype.HOMETIMELINE))
	rg.GET("/v2/timeline/asset/:assetID", authenticator, FindEntries(sc, datatype.ASSETTIMELINE))
	rg.GET("/v2/timeline/user/:userID", authenticator, FindEntries(sc, datatype.USERTIMELINE))
}
