package design

// Slot struct, has boolean to mark if it is occupied and if Occupied then stores Car
type Slot struct {
	Occupied bool
	Car      Car
}

// Marks Occupied boolean as true and parks the car
func (s *Slot) carToSlot(c Car) {
	s.Occupied = true
	s.Car = c
}

// Marks slot as unoccupied when car leaves
func (s *Slot) freeSlot() {
	s.Occupied = false
	// s.Car = nil
}

// Returns regNo and Color of car in the Slot
func (s Slot) vehicleDetails() (regNo string, color Color) {
	return s.Car.GetRegNo(), s.Car.getColor()
}
