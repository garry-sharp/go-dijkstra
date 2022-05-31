package main

type Node struct {
	Name    string
	Visited bool
}

type Path struct {
	Node   *Node
	Weight uint
}

type Graph struct {
	NodeNameLookup map[string]*Node //An O(1) hashmap to quickly get a node from the string name of the node
	Nodes          []*Node
	Edges          map[*Node][]*Path
	Tracker        map[*Node]*Tracker
}

type Tracker struct {
	Distance uint
	Previous *Node
}

func (g *Graph) Init() {
	g.Edges = make(map[*Node][]*Path)
	g.NodeNameLookup = make(map[string]*Node)
	g.Tracker = make(map[*Node]*Tracker)
}

//This implementation is bidirectional
func (g *Graph) AddEdge(node1 *Node, node2 *Node, weight uint, addUniDirectional bool) {
	g.Edges[node1] = append(g.Edges[node1], &Path{Node: node2, Weight: weight})
	if addUniDirectional {
		g.Edges[node2] = append(g.Edges[node2], &Path{Node: node1, Weight: weight})
	}
}

func (g *Graph) AddNode(name string) *Node {
	if nodeLookup := g.NodeNameLookup[name]; nodeLookup == nil {
		node := &Node{Name: name, Visited: false}
		g.Nodes = append(g.Nodes, node)
		g.NodeNameLookup[name] = node
		g.Tracker[node] = &Tracker{Distance: ^uint(0), Previous: nil}
		return g.Nodes[len(g.Nodes)-1]
	} else {
		return nodeLookup
	}
}

func (g *Graph) GetUnvisitedNeighbours(node *Node) []*Path {
	var result []*Path
	for _, el := range g.Edges[node] {
		if el.Node.Visited == false {
			result = append(result, el)
		}
	}
	return result
}
