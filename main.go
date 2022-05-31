package main

import (
	"fmt"
)

func main() {
	//Read in data
	graph := LoadFile()

	//Set current node to selected first node
	currentNode := SelectNode(&graph)
	startingNode := currentNode

	//For the algorithm to work the initial distance to itself needs to be set to 0
	graph.Tracker[currentNode].Distance = 0

	//Set up the sptSet, to avoid multiple paths to the same node we make it a map
	sptSetPaths := make(map[*Node]*Path)

	//Start of algorithm proper
	for {
		fmt.Println(PrettyPrintTracker(&graph))

		//Get an list of all adjacent (neighbour) nodes that have not been assessed
		neighbours := graph.GetUnvisitedNeighbours(currentNode)

		// If no neighbours are found either all nodes have been visited or we need to a previous node to continue traversal. Handle these cases

		fmt.Printf("  Working on Node %s\n", currentNode.Name)
		fmt.Printf("    There are %d direct neighbours\n", len(neighbours))

		// Set un initial dummy path with the highest weight possible.
		// Since we handle no neighbours in the above case this is safe as the default state will never be returned and stops us having to write extra conditions
		nextNode := &Path{Weight: ^uint(0)}
		for _, neighbour := range neighbours {
			if sptSetPaths[neighbour.Node] == nil || neighbour.Weight < sptSetPaths[neighbour.Node].Weight {
				fmt.Printf("    Adding Node %s to sptSet\n", neighbour.Node.Name)
				sptSetPaths[neighbour.Node] = neighbour
			}

			// These value represent the total current distance from the initial node to the new neighbour and the previous distance for readability
			distance := graph.Tracker[currentNode].Distance + neighbour.Weight
			lastKnownDistance := graph.Tracker[neighbour.Node].Distance
			fmt.Printf("    Node %s has a distance of %d\n", neighbour.Node.Name, neighbour.Weight)

			// If a new shortest distance is found is found update the Tracker accordingly
			if distance < lastKnownDistance {
				fmt.Printf("      New shortest distance!\n")
				graph.Tracker[neighbour.Node].Distance = distance
				graph.Tracker[neighbour.Node].Previous = currentNode
			}
		}

		// The next node to look at is the one with the lowest weight

		currentNode.Visited = true

		// Get the next node
		var candidateToRemove *Node
		for key, value := range sptSetPaths {
			candidate := value
			//remove visited nodes from sptSet on the fly
			if candidate.Node.Visited {
				fmt.Printf("  Removing Node %s from sptSet\n", candidate.Node.Name)
				delete(sptSetPaths, key)
				continue
			}

			if candidate.Weight < nextNode.Weight {
				nextNode = candidate
				candidateToRemove = candidate.Node
			}
		}

		fmt.Printf("    Next node set at Node %s. Removing from sptSet\n", nextNode.Node.Name)
		currentNode = nextNode.Node
		delete(sptSetPaths, candidateToRemove)
		fmt.Printf("    %d Candidates remaining in sptSet\n", len(sptSetPaths))

		if len(sptSetPaths) == 0 {
			break
		}

		//time.Sleep(3 * time.Second)
	}

	fmt.Printf("Algoritm Terminated\n\n")
	for _, node := range graph.Nodes {
		if node.Name == startingNode.Name {
			continue
		}
		fmt.Printf("Shortest route to %s  - ", node.Name)
		fmt.Printf(PrintShortestPath(startingNode, node, &graph))
		fmt.Printf(" (%d)\n", graph.Tracker[node].Distance)
	}
}
