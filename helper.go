package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ParseFileGraphFormat(fn string) Graph {
	var graph Graph
	graph.Init()
	if body, err := ioutil.ReadFile(fn); err != nil {
		log.Panic(err)
	} else {
		bodyString := string(body)
		bodyString = regexp.MustCompile(`[ \t]+`).ReplaceAllString(bodyString, " ")
		rows := strings.Split(bodyString, "\n")
		for i, row := range rows {
			vals := strings.Split(row, " ")
			for j, val := range vals {
				node1Name := string(65 + i)
				node2Name := string(65 + j)
				if weight, err := strconv.ParseUint(val, 0, 0); err != nil {
					log.Panicf("Unable to parse %s as uint in location %d %d\n", val, i, j)
				} else {
					node1 := graph.AddNode(node1Name)
					node2 := graph.AddNode(node2Name)
					if weight != 0 {
						graph.AddEdge(node1, node2, uint(weight), false)
					}
				}
			}
		}
	}
	return graph
}

func ParseFileSimpleFormat(fn string) Graph {
	var graph Graph
	graph.Init()
	if body, err := ioutil.ReadFile(fn); err != nil {
		log.Panic(err)
	} else {
		exp, _ := regexp.Compile(`([A-Za-z]+)\s+([A-Za-z]+)\s+([0-9]+)`)
		for _, line := range exp.FindAllStringSubmatch(string(body), -1) {
			var node1name string
			var node2name string
			for i, match := range line {
				switch i {
				case 1:
					node1name = match
				case 2:
					node2name = match
				case 3:
					if weight, err := strconv.ParseUint(match, 10, 0); err != nil {
						log.Panicf("Unable to parse file character %s as int", match)
					} else {
						if weight != 0 {
							node1 := graph.AddNode(node1name)
							node2 := graph.AddNode(node2name)
							graph.AddEdge(node1, node2, uint(weight), true)
						}
					}
				}
			}
		}
	}
	return graph
}

//Reads user input and removes some bloat from main.go
func LoadFile() Graph {
	var filename string
	fmt.Printf("Please enter the filename to parse (enter for example_data/example1.graph): ")
	if filename == "" {
		filename = "example_data/example1.graph"
	}
	fmt.Scanln(&filename)
	var graph Graph
	if v := strings.Split(filename, "."); v[len(v)-1] == "graph" {
		graph = ParseFileGraphFormat(filename)
	} else if v[len(v)-1] == "graphSimple" {
		graph = ParseFileSimpleFormat(filename)
	} else {
		log.Panicf("Cannot parse filetype %s\n", v[len(v)-1])
	}
	return graph
}

func SelectNode(graph *Graph) *Node {
	var nodeName string
	fmt.Printf("Please enter which node you would like to calculate from (enter for 'A'): ")
	fmt.Scanln(&nodeName)
	if nodeName == "" {
		nodeName = "A"
	}
	currentNode := graph.NodeNameLookup[nodeName]
	if currentNode == nil {
		log.Panicf("Cannot find any node called %s", nodeName)
	}
	return currentNode
}

//This funcion is not safe, it is only used for my personal tests. Doesn't check different row lengths and out of bounds indexes for example
func PrettyPrintTracker(g *Graph) string {
	fmt.Println()
	var dat [][]string
	var output string
	maxLengths := []int{0, 0, 0}
	addRow := func(row []string) {
		for i, entry := range row {
			if maxLengths[i] < len(entry) {
				maxLengths[i] = len(entry)
			}
		}
		dat = append(dat, row)
	}
	addRow([]string{"Node", "Distance", "Previous"})
	for _, node := range (*g).Nodes {
		tracker := g.Tracker[node]
		name := node.Name
		if node.Visited == true {
			name = "!" + name
		}
		shortestDistance := fmt.Sprint(tracker.Distance)
		if tracker.Distance == ^uint(0) {
			shortestDistance = "inf"
		}
		var previous string
		if tracker.Previous != nil {
			previous = tracker.Previous.Name
		} else {
			previous = "None"
		}
		row := []string{name, shortestDistance, previous}
		addRow(row)
	}
	for _, row := range dat {
		for j, el := range row {
			paddingLength := maxLengths[j]
			output += "| " + fmt.Sprintf("%*s", paddingLength, el) + " "
			if j == len(row)-1 {
				output += "|\n"
			}
		}
	}
	return output
}

func PrintShortestPath(startingNode *Node, node *Node, graph *Graph) string {
	output := node.Name
	for node != startingNode {
		node = graph.Tracker[node].Previous
		output = node.Name + " -> " + output
	}
	return output
}
