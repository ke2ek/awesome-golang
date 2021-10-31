# Concept

- `Shortest Path Problem` is to find the shortest path between given two vertices in a weighted-graph.
    - If there is no weighted edge, BFS can solve it.
        - Single-Source Shortest Path: BFS can find all the shortest path from the start to each vertex.
    - `Shortest` means the closest distance, the smallest cost or the lowest weight.
- These kinds of problems here assume that there is `no edge with a negative weight`.
    - Basically, given a graph containing a cycle with negative edges, the total cost of the cycle will diverge to the infinity of negative.
- These kinds of problems here assume that only `directed` graph will be handled.
    - To solve problems about undirected graphs using these algorithms, we must divide a given undirected graph into two directed graphs.
        - Each directed graph has only edges with one-direction.
- The following solutions provide not the list of vertices in the shortest path but the shortest cost of path.
    - e.g., Dijkstra, Bellman-Ford, Floyd-Warshall
- To know about the specific information like the list of vertices, it will need some extra data for storing other information.
