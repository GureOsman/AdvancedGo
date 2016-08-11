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
type ScheduleController struct  {
	interceptors.JWTAuthorization
	*revel.Controller
}

func (c ScheduleController) Create() revel.Result  {
	var schedule = encoders.EncodeSchedules(c.Request.Body)
	schedule.UserID, _ = strconv.ParseInt(c.Session["id"],10,0)
	if err := app.Db.Create(&schedule).Error; err != nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Creation failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(schedule))
}

func ( c ScheduleController)Update() revel.Result  {
	var schedule = encoders.EncodeSchedules(c.Request.Body)
	var id int
	var schedules models.Schedule
	c.Params.Bind(&id , "id")
	if rowsCount := app.Db.First(&schedules,id).RowsAffected; rowsCount <1 {
		return c.RenderJson(util.ResponseError("schedules informetion not founded"));
	}
	if err := app.Db.Model(&schedules).Updates(&schedule).Error;err !=nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("schedules Update failed"));
	}
	return c.RenderJson(util.ResponseSuccess(schedule))
}

func (c ScheduleController)Delete() revel.Result  {
	var (
		id int
		schedules models.Schedule
	)
	c.Params.Bind(&id ,"id")
	if rowsCount := app.Db.First(&schedules,id).RowsAffected; rowsCount <1 {
		return c.RenderJson(util.ResponseError("schedules informetion not founded to delete"));
	}
	if err := app.Db.Delete(&schedules).Error; err != nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("schedules delete Failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(schedules))
}

func (c ScheduleController) Get() revel.Result {
	var schedules [] models.Schedule
	var limitQuery = c.Request.URL.Query().Get("limit")
	if limitQuery == ""{
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset")
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&schedules).RowsAffected; founded<1{
		return c.RenderJson(util.ResponseError("No Founded schedules"))
	}
	return c.RenderJson(schedules)
}

