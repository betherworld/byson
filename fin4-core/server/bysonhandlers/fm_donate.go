package bysonhandlers

import (
	"fin4-core/server/auth"
	"fin4-core/server/datatype"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// GetTokens fetch tokens from db
func Donate(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {

		devFund , err := sc.Ethereum.FundManager.GetDevFund(&bind.CallOpts{})

		log.Print(devFund)
		log.Print("hallo")

		u := auth.MustGetUser(c)
		tokens, err := sc.TokenService.FindAll(u.ID)
		if err != nil {
			c.String(http.StatusServiceUnavailable, err.Error())
			return
		}
		c.JSON(http.StatusOK, struct {
			Count   int
			Limit   int
			Page    int
			Entries []datatype.Token
		}{
			Count:   len(tokens),
			Limit:   10,
			Page:    0,
			Entries: tokens,
		})
	}
}
