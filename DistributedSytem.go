package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Node struct {
	Id int
}

type Message struct {
	From int
	To   int
	Data string
}

func (n *Node) SendMessage(to int, data string) {
	fmt.Printf("Node %d sends message to Node %d: %s\n", n.Id, to, data)
}

func (n *Node) ReceiveMessage(m Message) {
	fmt.Printf("Node %d receives message from Node %d: %s\n", n.Id, m.From, m.Data)
}

func main() {
	// Create a set of nodes
	nodes := []Node{{Id: 1}, {Id: 2}, {Id: 3}}

	// Initialize the nodes with random timeouts
	for i := range nodes {
		timeout := time.Duration(rand.Intn(500)+500) * time.Millisecond
		go func(node Node, timeout time.Duration) {
			for {
				time.Sleep(timeout)
				to := rand.Intn(len(nodes))
				if to != node.Id {
					data := fmt.Sprintf("Message from Node %d to Node %d", node.Id, to+1)
					node.SendMessage(to+1, data)
				}
			}
		}(nodes[i], timeout)
	}

	// Simulate receiving messages
	for {
		from := rand.Intn(len(nodes))
		to := rand.Intn(len(nodes))
		if from != to {
			data := fmt.Sprintf("Message from Node %d to Node %d", from+1, to+1)
			message := Message{From: from + 1, To: to + 1, Data: data}
			nodes[to].ReceiveMessage(message)
		}
		time.Sleep(1 * time.Second)
	}
}
