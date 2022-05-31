# Dijistra's shortest path algorithm

This project is a (relatively) lightweight implementation of Dijistra's shortest path algorithm. It seeks to both code it in an effective but also educational way. There are many logs to the console to show exactly how the algorithm is working.

## Overview

The code is layed out in 3 files.

- `main.go` goes through the algorithm step by step and calls lower level functions to keep it as clean as possible.
- `types.go` stores the various structs and their methods.
- `helper.go` stores other code non-essential to the functioning of the algorithm but which would be too clunky to keep in the other two files.

## File Types

The code will run on 2 file types. The first being a "traditional" `.graph` file in the form of a whitespace and new line separated matrix. The number or type of whitespace characters doesn't matter. The following is an example of an acceptable matrix

```
0 4  0 0  0  0  0 8   0
4 0  8 0  0  0  0 11  0
0 8  0 7  0  4  0 0   2
0 0  7 0  9  14 0 0   0
0 0  0 9  0  10 0 0   0
0 0  4 14 10 0  2 0   0
0 0  0 0  0  2  0 1   6
8 11 0 0  0  0  1 0   7
0 0  2 0  0  0  6 7   0
```

The second type of file to import is a custom one for this project of type `.graphSimple` The format of this is multiple lines of 3 whitespace separated values representing 2 node points and a weight. These points are read as bidirectional. For a single directional use the `.graph` format. This is an example of the `.graphSimple` format 

```
A B 6
B C 5
A D 1
B E 2
D B 2
C E 5
E D 1
```

For the avoidance of any doubt the above example is the same as the following `.graph` example.

```
0 6 0 1 0
6 0 5 2 2
0 5 0 0 5
1 2 0 0 1
0 2 5 1 0
```

## Data Types

The implementation has 4 key data types/structs.

- `Node` represents a single node. It has a name and visited property. The names assigned when importing graphs are set to be upper case letters. The visited property is used by the algorithm.
- `Path` essentially a subtype of the `Graph` type. An initial implementation preferred a linked list approach but this proved to be more effective. A path contains both a weight and a target
- `Graph` the main overarching struct used. Its properties are the following:
  - `NodeNameLookup` - An O(1) hashmap to quickly get a node from the string name of the node. It is only used in the initial parse of the data in "simple" format and is irrelevant for the functioning of the algorithm
  - `Nodes` - A list is all Nodes used in the Graph
  - `Edges` - A map of type `map[*Node][]*Path`. Passing a Node to the map basically yields a slice or Paths which are basically other nodes plus their weight
  - `Tracker` - A map of type `map[*Node]*Tracker` which is used exclusively by the algorithm.
- `Tracker` - Used by the algorithm to literally keep track of total distances as well as the "previous" node (the last node from which the distance was calculated)

## Other comments and notes

- The program can be run with a `go build . && ./go-dijkstra` command.
- The length of an edge cannot be negative nor can it be larger than 2<sup>64</sup>-1 on 64 bit machines and 2<sup>32</sup>-1 on 32 bit machines.
- Not all functions and methods are 100% safe, especially the helper functions. Files are expected to be in the right format otherwise the program will panic and not fail gracefully.
- The following terms in comments and variable names are used interchangably
    * Vertex and Node
    * Path and Edge
    * Weight and Distance
    * Neighbour and Adjacent
    
    While I appreciate this may be confusing, I do hope it is obvious enough.
- I want to stress this project was done as part of a personal learning experience for me and also to share with people who want to learn and visualise an algorithm as well as view the code.
