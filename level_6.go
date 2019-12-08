// Objective: Turn off the lasers while keeping the LaserGrid operational

package main

func main() {
	lasers := setupLasers()

	if len(lasers) != 7 {
		println("ALERT! Wrong number of cameras.")
		return
	}

	for i := 0; i < 7; i++ {
		if !lasers[i].isRunning {
			println("ALERT! Laser not running.")
			return
		}
	}

	laserGrid := LaserGrid{"Operational", lasers}

	passedTest := testGrid(laserGrid)
	if !passedTest {
		println("ALERT! Grid test failed.")
		return
	}

	println("Grid operational")

	running := 0
	for _, laser := range laserGrid.lasers {
		if laser.isRunning {
			running++
		}
	}

	println(running, "lasers running")
}


// LaserGrid represents a collection of 7 lasers
type LaserGrid struct {
	status string
	lasers [7]*Laser
}

// Laser respresents an individual laser beam
type Laser struct {
	id        int
	isRunning bool
}

func setupLasers() [7]*Laser {
	var lasers [7]*Laser
	lasers[0] = &Laser{1, true}
	lasers[1] = &Laser{2, true}
	lasers[2] = &Laser{3, true}
	lasers[3] = &Laser{4, true}
	lasers[4] = &Laser{5, true}
	lasers[5] = &Laser{6, true}
	lasers[6] = &Laser{7, true}
	return lasers
}

func testGrid(laserGrid LaserGrid) bool {
    for i := 0; i < 7; i++ {
        laserGrid.lasers[i].isRunning = false
    }
	return laserGrid.status == "Operational"
}