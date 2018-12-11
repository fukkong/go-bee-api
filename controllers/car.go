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
// @Param	body		body 	models.Car	true		"The car content"
// @Success 200 {string} models.Car.Carid
// @Failure 403 body is empty
// @router / [post]
func (c *CarController) Post() {
	var ca models.Car
	json.Unmarshal(c.Ctx.Input.RequestBody, &ca)
	carId := models.AddOneCar(ca)
	c.Data["json"] = map[string]int{"Carid": carId}
	c.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	Carid		path 	int	true		"the Carid you want to get"
// @Success 200 {Car} models.Car
// @Failure 403 :Carid is empty
// @router /:Carid [get]
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
	c.ServeJSON()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {Car} models.Car
// @Failure 403 :Carid is empty
// @router / [get]
func (c *CarController) GetAll() {
	cs := models.GetAllCar()
	c.Data["json"] = cs
	c.ServeJSON()
}

// @Title Update
// @Description update the object
// @Param	Carid		path 	int	true		"The carId you want to update"
// @Param	body		body 	models.Car	true		"The body"
// @Success 200 {car} models.Car
// @Failure 403 :Carid is empty
// @router /:Carid [put]
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
// @Param	Carid		path 	int	true		"The carId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Carid is empty
// @router /:Carid [delete]
func (c *CarController) Delete() {
	carId := c.Ctx.Input.Param(":Carid")
	ids, errs := strconv.ParseInt(carId, 10, 64)
	if errs == nil {
		models.DeleteCar(int(ids))
		c.Data["json"] = "delete success!"
	}
	c.ServeJSON()
}
