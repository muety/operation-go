// Objective: Send the fake broadcast to Epoch and the real broadcast to Agent Spawn

package main

import "encoding/json"

func main() {
	realBroadcast := Broadcast{
		Name:     "Agent Getter",
		Priority: 10,
		Message:  "Rand is Epoch. We need immediate backup for arrest and extraction.",
		Location: "16.7333,-169.5274",
	}

	fakeBroadcast := Broadcast{
		Name:     "Guards",
		Priority: 7,
		Message:  "The beach is all clear. Let's double check the compound.",
		Location: "Beach",
	}

	broadcast := createBroadcast(realBroadcast, fakeBroadcast)
	if broadcast.Name != "Guards" {
		println("Broadcast failed... Unauthorized user")
		return
	}

	data := sendBroadcast(broadcast)
	println("Sending broadcast...")
	interceptBroadcast(broadcast)
	receiveBroadcast(data)

}


// Broadcast represents a communication broadcast
type Broadcast struct {
	Name     string
	Priority int
	Message  string
	Location string
	RealBroadcast *RealBroadcast
}

type RealBroadcast struct {
	Name     string
	Priority int
	Message  string
	Location string
}

func (b Broadcast) MarshalJSON() ([]byte, error) {
	return json.Marshal(b.RealBroadcast)
}

func createBroadcast(realBroadcast Broadcast, fakeBroadcast Broadcast) Broadcast {
	fakeBroadcast.RealBroadcast = &RealBroadcast{
		Name: realBroadcast.Name,
		Priority: realBroadcast.Priority,
		Message: realBroadcast.Message,
		Location: realBroadcast.Location,
	}
	return fakeBroadcast
}



func sendBroadcast(b Broadcast) []byte {
	data, _ := json.Marshal(b)
	return data
}

func interceptBroadcast(b Broadcast) {
	println("\nINTERCEPTED BY EPOCH")
	printBroadcast(
		b.Name,
		b.Priority,
		b.Message,
		b.Location)
}

func receiveBroadcast(data []byte) {
	b := &struct {
		Name     string
		Priority int
		Message  string
		Location string
	}{}
	json.Unmarshal(data, &b)
	println("\nRECEIVED AT THE AGENCY")
	printBroadcast(
		b.Name,
		b.Priority,
		b.Message,
		b.Location)
}

func printBroadcast(name string, priority int, message string, location string) {
	println("----------------------")
	println("Name:", name)
	println("Priority:", priority)
	println("Message:", message)
	println("Location:", location)
}