package design

// iVehicle is a vehicle interface and only Get methods are mandatory as we consider vehicle as a value object
type iVehicle interface {
	GetRegNo() string
	getColor() Color
}
