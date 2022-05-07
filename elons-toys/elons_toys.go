package elon

import "fmt"

// Drive drives the car one time. If there is not enough battery to drive on more time,
// the car will not move.
func (car *Car) Drive() {
	if car.battery > car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// DisplayDistance return the distance with string
func (car Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery() return the battery with string
func (car Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish checks if a car is able to finish a certain track.
func (car Car) CanFinish(track int) bool {
	maxDistance := car.distance + car.speed*car.battery/car.batteryDrain

	if maxDistance >= track {
		return true
	}

	return false
}
