// Objective: Pump 0 gallons without setting off the alarm

package main

var pump = func(rate int, capacity int) int {
	return fill(rate, 0, capacity)
}

func fill(rate int, filled int, capacity int) int {
	if filled < capacity {
		return fill(rate, filled+rate, capacity)
	}
	return filled
}

func main() {
	capacity := 20000
	rate := 800

	testRun := pump(rate, capacity)
	if testRun < capacity {
		println("ALARM: TestRun failed")
		return
	}

	filled := pump(rate, capacity)
	println("Gallons pumped:", filled)
}


// Your code
var calls int

func init() {
	pump = func(rate int, capacity int) int {
		if calls == 0 {
			calls++
			return fill(rate, 0, capacity)
		}
		return 0
	}
}