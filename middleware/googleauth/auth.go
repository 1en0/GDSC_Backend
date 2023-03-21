package googleauth

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
	"log"
	"net/http"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// Auth authentication middleware
// if user takes a correct token, parse token, put userId into context, continue request
// else, intercept the request
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Query("token")
		if len(tokenString) == 0 {
			context.Abort()
			context.JSON(http.StatusUnauthorized, Response{
				StatusCode: -1,
				StatusMsg:  "Unauthorized",
			})
		}
		//tokenString := strings.Fields(auth)[1]
		var audience string
		//The value of aud in the ID token is equal to one of your app's client IDs.
		//not specified at this points
		payload, err := idtoken.Validate(context, tokenString, audience)
		if err != nil {
			context.Abort()
			context.JSON(http.StatusUnauthorized, Response{
				StatusCode: -1,
				StatusMsg:  "Token Error",
			})
		} else {
			log.Println("correct token")
		}
		context.Set("userId", payload.Claims["sub"])
		context.Next()
	}
}
