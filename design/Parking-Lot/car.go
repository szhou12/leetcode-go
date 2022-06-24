package design

// A Car struct containing registration no. and its color
type Car struct {
	iVehicle // 这里为什么？？？
	regNo    string
	color    Color
}

// NewCar creates new Car and returns its reference.
func NewCar(regNo string, color Color) *Car {
	return &Car{
		regNo: regNo,
		color: color,
	}
}

// MakeCar creates new Car and returns it.
func MakeCar(regNo string, color Color) Car {
	return Car{
		regNo: regNo,
		color: color,
	}
}

// 注意: setters 都是 指针接收者； getters 都是 值接收者
// SetRegNo for setting Reg No. of Car
func (v *Car) SetRegNo(regNo string) {
	v.regNo = regNo
}

// GetRegNo for getting Reg No. of Car
func (v Car) GetRegNo() string {
	return v.regNo
}

// SetColor for setting Color of Car
func (v *Car) SetColor(color Color) {
	v.color = color
}

// GetColor for getting Color of Car
func (v Car) GetColor() Color {
	return v.color
}
