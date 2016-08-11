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
type TicketController struct  {
	interceptors.JWTAuthorization
	*revel.Controller
}

func (c TicketController) Create() revel.Result  {
	var ticket = encoders.EncodeTicket(c.Request.Body)
	c.Params.Bind(&ticket.PlaneID, "id")
	c.Params.Bind(&ticket.PasId, "id")
	ticket.UserID, _ = strconv.ParseInt(c.Session["id"],10,0)
	if err := app.Db.Create(&ticket).Error; err != nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Creation failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(ticket))
}

func ( c TicketController)Update() revel.Result  {
	var ticket = encoders.EncodeTicket(c.Request.Body)
	var id int
	var tickets models.Tickets
	c.Params.Bind(&id , "id")
	if rowsCount := app.Db.First(&tickets,id).RowsAffected; rowsCount <1 {
		return c.RenderJson(util.ResponseError("ticket informetion not founded"));
	}
	if err := app.Db.Model(&tickets).Updates(&ticket).Error;err !=nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("ticket Update failed"));
	}
	return c.RenderJson(util.ResponseSuccess(ticket))
}

func (c TicketController)Delete() revel.Result  {
	var (
		id int
		ticket models.Tickets
	)
	c.Params.Bind(&id ,"id")
	if rowsCount := app.Db.First(&ticket,id).RowsAffected; rowsCount <1 {
		return c.RenderJson(util.ResponseError("ticket informetion not founded to delete"));
	}
	if err := app.Db.Delete(&ticket).Error; err != nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("ticket delete Failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(ticket))
}

func (c TicketController) Get() revel.Result {
	var tickets [] models.Tickets
	var limitQuery = c.Request.URL.Query().Get("limit")
	if limitQuery == ""{
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset")
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&tickets).RowsAffected; founded<1{
		return c.RenderJson(util.ResponseError("No Founded tickets"))
	}
	return c.RenderJson(tickets)
}

