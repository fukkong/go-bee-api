package models

import (
	"errors"
)

var (
	Cars map[int]*Car
)

type Car struct {
	Carid  int
	Vin    string
	Model  string
	Maker  string
	Year   uint
	Msrp   uint64
	Status string
	Booked bool
	Listed bool
}

func init() {
	Cars = make(map[int]*Car)
	Cars[1] = &Car{1, "MNBUMF050FW496402", "320i", "BMW", 2013,
		10000, "Ordered", true, true}
	Cars[2] = &Car{2, "4JDBLMF080FW468775", "Carmry", "Toyota", 2015,
		12000, "In stock", true, false}
	Cars[3] = &Car{3, "TFBAXXMAWAFS71274", "Focus", "Ford", 2016,
		13000, "Ordered", false, true}
	Cars[4] = &Car{4, "G3SBUMF080FW470449", "Civic", "Honda", 2016,
		14000, "Sold", false, false}
}

func AddOneCar(car Car) (Id int) {
	car.Carid = len(Cars) + 1
	Cars[car.Carid] = &car
	return car.Carid
}

func GetOneCar(Id int) (Car *Car, err error) {
	if v, ok := Cars[Id]; ok {
		return v, nil
	}
	return nil, errors.New("Id Not Exist")
}

func GetAllCar() map[int]*Car {
	return Cars
}

func UpdateCar(Id int, Vin string, Model string, Maker string, Year uint, MSRP uint64, Status string, Booked bool, Listed bool) (err error) {
	if v, ok := Cars[Id]; ok {
		v.Vin = Vin
		v.Model = Model
		v.Maker = Maker
		v.Year = Year
		v.Msrp = MSRP
		v.Status = Status
		v.Booked = Booked
		v.Listed = Listed
		return nil
	}
	return errors.New("Id Do Not Exist")
}

func DeleteCar(Id int) {
	delete(Cars, Id)
}
