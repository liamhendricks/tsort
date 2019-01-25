package main

import (
    "github.com/liamhendricks/golib/dag"
    "fmt"
)

var visited map[string]*dag.Node
var ordering []*dag.Node
var g *dag.Dag
var name string
var nodes map[string]*dag.Node

func TopologicalSort(g *dag.Dag) []*dag.Node {
    // loop through nodes
    for _, v := range nodes {
        if _, ok := visited[v.Name()]; !ok {
            // perform dfs on each node
            depthFirstSearch(v)
        }
    }

    return ordering
}

func depthFirstSearch(n *dag.Node) {
    // mark node as visited
    visited[n.Name()] = n

    // get child nodes
    edges := nodes[n.Name()].Children()

    // loop through child nodes
    for _, e := range edges {
        if _, ok := visited[e.Name()]; !ok {
            // perform dfs on each child
            depthFirstSearch(e)
        }
    }
    ordering = append(ordering, n)
}

func main() {
    visited = make(map[string]*dag.Node)
    g := dag.New()
    nodes = g.Nodes()
    g.AddVertex("p1", "Open Peanut Butter Jar")
    g.AddVertex("j2", "Spread Jelly on Bread")
    g.AddVertex("j1", "Open Jelly Jar")
    g.AddVertex("e1", "Eat the Dang Sandwich")
    g.AddVertex("n1", "Slice Banana")
    g.AddVertex("p2", "Spread PB on Bread")
    g.AddVertex("b1", "Slice Bread")
    g.AddVertex("n2", "Put Bananas on Bread")
    g.AddVertex("b2", "Put bread together")
    g.AddEdge("b1", "n1")
    g.AddEdge("j1", "j2")
    g.AddEdge("n1", "j2")
    g.AddEdge("n1", "p2")
    g.AddEdge("n1", "n2")
    g.AddEdge("b1", "j1")
    g.AddEdge("b1", "p1")
    g.AddEdge("p1", "p2")
    g.AddEdge("p1", "j2")
    g.AddEdge("j1", "p2")
    g.AddEdge("j2", "n2")
    g.AddEdge("p2", "n2")
    g.AddEdge("n2", "b2")
    g.AddEdge("b2", "e1")

    TopologicalSort(g)

    for _, o := range ordering {
        fmt.Printf("%s: %s\n", o.Val(), o.Name())
    }
}
