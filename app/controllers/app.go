package controllers

import (
	"github.com/revel/revel"
	"encoding/json"
	"io/ioutil"

	//"strconv"
	"log"
	"github.com/gureosman/AdvancedGo/app/models"
	"github.com/gureosman/AdvancedGo/app"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type App struct {
	*revel.Controller
}

//get
func (c App) Index() revel.Result {
	var user models.User
	var id int
	c.Params.Bind(&id ,"id")
       if err := app.Db.First(&user ,id); err!=nil{
	       return c.RenderJson(err)
       }
	return c.RenderJson(user)
}
// Post
func (c App)Create() revel.Result{
	var user models.User
	var data,_ =ioutil.ReadAll(c.Request.Body)
	if err:= json.Unmarshal(data, &user);err != nil{
            return c.RenderText("error not found")
	}
	if err := app.Db.Create(&user).Error; err !=nil{
		log.Println(err);
		return c.RenderJson(err)
	}
	return c.RenderText("Welcome Mr " + user.Name)
}


