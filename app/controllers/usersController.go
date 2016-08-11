package controllers
import (
	"github.com/revel/revel"
	//"os/user"
	"time"
	//"fmt"
	"github.com/gureosman/AdvancedGo/app/models"
	"github.com/gureosman/AdvancedGo/app/encoders"
	"log"
	"github.com/gureosman/AdvancedGo/app/util"
	"github.com/dgrijalva/jwt-go"

	"github.com/gureosman/AdvancedGo/app"
	//"fmt"
	//"fmt"
	//"strconv"
	"strconv"
)
type UsersController struct {
	*revel.Controller
}
func (c UsersController)Create() revel.Result  {
	var user = encoders.EncodeSingleUsers(c.Request.Body);
	if user.Email == "" || user.Password == ""{
		return c.RenderJson(util.ResponseError("User Information Is Empty"));
	}

	if err := app.Db.Create(&user).Error; err != nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Creation failed"));
	}
	return c.RenderJson(util.ResponseSuccess(user));
}
func (c UsersController) Login() revel.Result  {
	var user = encoders.EncodeSingleUsers(c.Request.Body)
	if user.Email == "" || user.Password == ""{
		return c.RenderJson(util.ResponseError("User Login failed"));
	}
	if founded:= app.Db.Where(&user).First(&user).RowsAffected; founded <1 {
		return c.RenderJson(util.ResponseError("User Not Founded"));
	}
        token := jwt.NewWithClaims(jwt.SigningMethodHS256 , jwt.MapClaims{
		"id": user.ID,
		"email":user.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	appSecret, _:= revel.Config.String("app.secret");
	tokenString, err := token.SignedString([]byte(appSecret));
	if err != nil {
		log.Println(err)
		return c.RenderJson(util.ResponseError("Key Generation Failed"))

	}
	var tokenmodel models.Token
	tokenmodel.Email = user.Email
	tokenmodel.Name = user.Name
	tokenmodel.Token = tokenString
	return c.RenderJson(util.ResponseSuccess(tokenmodel))
}

func ( c UsersController)Update() revel.Result  {
	var user = encoders.EncodeSingleUsers(c.Request.Body)
	user.ID, _ = strconv.ParseInt(c.Session["id"],10,0)
	var id int
	var users models.User
	c.Params.Bind(&id , "id")
	if rowsCount := app.Db.First(&users, id).RowsAffected; rowsCount <1 {
		return c.RenderJson(util.ResponseError("Passengers informetion not founded"));
	}
	if err := app.Db.Model(&users).Updates(&user).Error;err !=nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Post Update failed"));
	}
	return c.RenderJson(util.ResponseSuccess(user))
}

