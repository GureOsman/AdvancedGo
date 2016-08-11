package controllers

import (
	"github.com/revel/revel"
	"github.com/gureosman/AdvancedGo/app/encoders"
	"github.com/gureosman/AdvancedGo/app/util"
	"strconv"
	"log"
	"github.com/gureosman/AdvancedGo/app/interceptors"
	"github.com/gureosman/AdvancedGo/app"
	"github.com/gureosman/AdvancedGo/app/models"
)
type PlaneController struct  {
	interceptors.JWTAuthorization
	*revel.Controller
}

func (c PlaneController) Create() revel.Result  {
	var plane = encoders.EncodePlane(c.Request.Body)
	plane.UserID, _ = strconv.ParseInt(c.Session["id"],10,0)
	if err := app.Db.Create(&plane).Error; err != nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Creation failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(plane))
}

func ( c PlaneController)Update() revel.Result  {
	var plane = encoders.EncodePlane(c.Request.Body)
	var id int
	var passengers models.Planes
	c.Params.Bind(&id , "id")
	if rowsCount := app.Db.First(&passengers,id).RowsAffected; rowsCount <1 {
		return c.RenderJson(util.ResponseError("Planes informetion not founded"));
	}
	if err := app.Db.Model(&passengers).Updates(&plane).Error;err !=nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Post Update failed"));
	}
	return c.RenderJson(util.ResponseSuccess(plane))
}

func (c PlaneController)Delete() revel.Result  {
	var (
		id int
		planes models.Planes
	)
	c.Params.Bind(&id ,"id")
	if rowsCount := app.Db.First(&planes,id).RowsAffected; rowsCount <1 {
		return c.RenderJson(util.ResponseError("Plane informetion not founded to delete"));
	}
	if err := app.Db.Delete(&planes).Error; err != nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Plane delete Failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(planes))
}

func (c PlaneController) Get() revel.Result {
	var planes [] models.Planes
	var limitQuery = c.Request.URL.Query().Get("limit")
	if limitQuery == ""{
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset")
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&planes).RowsAffected; founded<1{
		return c.RenderJson(util.ResponseError("No Founded Planes"))
	}
	return c.RenderJson(planes)
}

