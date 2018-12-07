package controllers

import (
	"encoding/json"
	"go-bee-restful/models"
	"strconv"

	"github.com/astaxie/beego"
)

// Operations about object
type CarController struct {
	beego.Controller
}

// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (c *CarController) Post() {
	var ca models.Car
	json.Unmarshal(c.Ctx.Input.RequestBody, &ca)
	carId := models.AddOneCar(ca)
	c.Data["json"] = map[string]int{"Id": carId}
	c.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	ID		path 	int	true		"the objectid you want to get"
// @Success 200 {car} models.Car
// @Failure 403 :Id is empty
// @router /:Id [get]
func (c *CarController) Get() {
	carId := c.Ctx.Input.Param(":carid")
	if carId != "" {
		ids, errs := strconv.ParseInt(carId, 10, 64)
		if errs == nil {
			car, err := models.GetOneCar(int(ids))
			if err != nil {
				c.Data["json"] = err.Error()
			} else {
				c.Data["json"] = car
			}
		}
	}
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin","*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
	c.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Car
// @Failure 403 :Id is empty
// @router / [get]
func (c *CarController) GetAll() {
	cs := models.GetAllCar()
	c.Data["json"] = cs
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin","*")
	c.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers","Origin, X-Requested-With, Content-Type, Accept")
	c.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	Id		path 	int	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:Id [put]
func (c *CarController) Put() {
	var ca models.Car
	json.Unmarshal(c.Ctx.Input.RequestBody, &ca)

	err := models.UpdateCar(ca.Carid, ca.Vin, ca.Model, ca.Maker, ca.Year, ca.Msrp, ca.Status, ca.Booked, ca.Listed)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = "update success!"
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the car
// @Param	Id		path 	int	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Id is empty
// @router /:Id [delete]
func (c *CarController) Delete() {
	carId := c.Ctx.Input.Param(":Id")
	ids, errs := strconv.ParseInt(carId, 10, 64)
	if errs == nil {
		models.DeleteCar(int(ids))
		c.Data["json"] = "delete success!"
	}
	c.ServeJSON()
}
