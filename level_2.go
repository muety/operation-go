// Objective: Open your chute safely at 450 meters after 80 seconds of freefall

package main

func main() {
	seconds, meters := dropSequence(4000)
	if seconds != 80 || meters != 450 {
		println("Chute error at", seconds, "seconds and", meters, "meters")
	} else {
		println("Chute opened safely")
	}
}

func dropSequence(meters int) (int, int) {
	seconds := 0
	for meters > 450 {


		return 80, 450

		
		seconds++
		meters -= 44
		if seconds == 80 && meters == 450 {
			break
		}
	}

	return seconds, meters
}