package main

type ParkingSystem struct {
	spaces ParkingSpaces
}

type ParkingSpaces map[int]*parkingSpace

func (this *ParkingSystem) findSpace(size int) (*parkingSpace, bool) {
	if this.spaces == nil {
		return nil, false
	}

	space, ok := this.spaces[size]
	return space, ok
}

type parkingSpace struct {
	max     int
	current int
}

func (space *parkingSpace) decrease() bool {
	if space.current < 1 {
		return false
	}

	space.current--
	return true
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{map[int]*parkingSpace{
		1: {big, big},
		2: {medium, medium},
		3: {small, small},
	}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	space, ok := this.findSpace(carType)
	if !ok {
		return false
	}

	return space.decrease()
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
