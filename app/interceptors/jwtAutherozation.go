package interceptors

import (
	"github.com/revel/revel"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"time"
	"github.com/gureosman/AdvancedGo/app/util"
	"strconv"
)

type JWTAuthorization struct  {
	*revel.Controller
}
func ( c JWTAuthorization)checkUser() revel.Result {
	var tokenString = c.Request.Header.Get("token")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected Signing Method:%v", token.Header["alg"])
		}
		appSecret, _ := revel.Config.String("app.secret");
		return []byte(appSecret), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			var expDate = time.Unix(int64(claims["exp"].(float64)),0)
			if expDate.Before(time.Now()){
				return c.RenderJson(util.ResponseError("Expired Token"))
			}
			c.Session["email"] = claims["email"].(string)
			c.Session["id"] = strconv.Itoa(int(claims["id"].(float64)));
			return nil
		}
	} else {
		return c.RenderJson(util.ResponseError("Invalid Token Key"))

	}
	return c.RenderJson(util.ResponseSuccess(tokenString))
}

func init()  {
	revel.InterceptMethod(JWTAuthorization.checkUser, revel.BEFORE)

}