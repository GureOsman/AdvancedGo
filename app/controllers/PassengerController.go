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
type PassengerController struct  {
	interceptors.JWTAuthorization
	*revel.Controller
}
func (c PassengerController) Create() revel.Result  {
	var passenger = encoders.EncodePassengers(c.Request.Body)
	passenger.UserID, _ = strconv.ParseInt(c.Session["id"],10,0)
	if err := app.Db.Create(&passenger).Error; err != nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Creation failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(passenger))
}
func ( c PassengerController)Update() revel.Result  {
	var passenger = encoders.EncodePassengers(c.Request.Body)
	passenger.UserID, _ = strconv.ParseInt(c.Session["id"],10,0)
	passenger.ID, _ = strconv.ParseInt(c.Session["id"],10,0)
	var id int
	var passengers models.Passengers
	c.Params.Bind(&id , "id")
	if rowsCount := app.Db.First(&passengers, id).RowsAffected; rowsCount <1 {
		return c.RenderJson(util.ResponseError("Passengers informetion not founded"));
	}
	if err := app.Db.Model(&passengers).Updates(&passenger).Error;err !=nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Post Update failed"));
	}
	return c.RenderJson(util.ResponseSuccess(passenger))
}
func (c PassengerController)Delete() revel.Result  {
	var (
		id int
		passengers models.Passengers
	)
	c.Params.Bind(&id ,"id")
	if rowsCount := app.Db.First(&passengers,id).RowsAffected; rowsCount <1 {
		return c.RenderJson(util.ResponseError("Passsenger informetion not founded to delete"));
	}
	if err := app.Db.Delete(&passengers).Error; err != nil{
		log.Println(err)
		return c.RenderJson(util.ResponseError("Passenger delete Failed"));
	}
	return  c.RenderJson(util.ResponseSuccess(passengers))
}
func (c PassengerController) List() revel.Result {
	var passengers [] models.Passengers
	var limitQuery = c.Request.URL.Query().Get("limit")
	if limitQuery == ""{
		limitQuery = "0"
	}
	var offsetQuery = c.Request.URL.Query().Get("offset")
	if founded := app.Db.Limit(limitQuery).Offset(offsetQuery).Find(&passengers).RowsAffected; founded<1{
		return c.RenderJson(util.ResponseError("No Founded Passengers"))
	}
	for i ,passenger := range passengers {
		app.Db.First(&passengers[i].Tickets ,passenger.UserID)
	}
	return c.RenderJson(passengers)
}
func (c PassengerController) Get() revel.Result {
	var passenger models.Passengers
	var id int
	c.Params.Bind(&id, "id")
	if err := app.Db.First(&passenger, id).Error; err != nil {
		return c.RenderJson(util.ResponseError("passaenger not founded"))
	}
	return c.RenderJson(util.ResponseSuccess(passenger))
}


