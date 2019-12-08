// Objective: Send messages to your interceptComms channel while preserving the messages in epochComms

package main

func main() {
	epochComms := make(chan string)
	go func() {
		epochComms <- messageQueue(0)
		epochComms <- messageQueue(1)
		epochComms <- messageQueue(2)
		epochComms <- messageQueue(3)
		close(epochComms)
	}()


	interceptComms := make(chan string, 4)
	epochCommsTmp := make(chan string, 4)

	for i := 0; i < 4; i++ {
		m := <- epochComms
		interceptComms <- m
		epochCommsTmp <- m
	}

	close(interceptComms)

	epochComms = epochCommsTmp

	
	println("Intercepted")
	println("---------------------------------")
	for message := range interceptComms {
		println("->", message)
	}

	println()
	println("Sent to Epoch")
	println("---------------------------------")
	for message := range epochComms {
		println(message)
	}
}

func messageQueue(i int) string {
	messages := make(map[int]string)
	messages[0] = "[Len] All agents head south."
	messages[1] = "[Epoch] Get those USBs!"
	messages[2] = "[Val] Move out team."
	messages[3] = "[Epoch] Faster!"
	return messages[i]
}