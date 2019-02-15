package bysonhandlers

import (
	"fin4-core/server/datatype"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// ApproveClaim approves a mining claim
func GetDevFund(sc datatype.ServiceContainer) gin.HandlerFunc {
	return func(c *gin.Context) {

		devFund , err := sc.Ethereum.FundManager.GetDevFund(&bind.CallOpts{})  //&bind.CallOpts{}
		if err == nil {
			panic(err)
		}
		log.Print(devFund)
		//user := auth.MustGetUser(c)

		c.String(http.StatusOK, (*devFund).String())


		/*body := struct {
			ClaimID int `json:"claimId"`
		}{}
		c.BindJSON(&body)
		claim, err := sc.TokenService.FindClaim(datatype.ID(body.ClaimID))
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		token, err := sc.TokenService.FindTokenForUser(claim.TokenID, user.ID)
		if token == nil || token.CreatorID != user.ID || err != nil {
			c.String(http.StatusBadRequest, "Token not found or not authorized")
			return
		}
		// doer, err := sc.UserService.FindByID(claim.UserID)
		// if err != nil {
		// 	c.String(http.StatusBadRequest, err.Error())
		// 	return
		// }
		err = sc.TokenService.ApproveClaim(datatype.ID(body.ClaimID), user.ID)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		c.String(http.StatusOK, "")**/
	}
}
