function countComponents(n: number, edges: number[][]): number {
  const visited: boolean[] = new Array(n).fill(false);
  const graph = buildGraph(n, edges);

  let count = 0;

  for (let i = 0; i < n; i++) {
    if (!visited[i]) {
      count++;
      dfs(i, graph, visited);
    }
  }

  return count;
}

const buildGraph = (n: number, edges: number[][]): number[][] => {
  const graph = Array.from({ length: n }, () => []);

  for (let edge of edges) {
    let [src, dist] = edge;
    graph[src].push(dist);
    graph[dist].push(src);
  }

  return graph;
};

const dfs = (idx: number, graph: number[][], visited: boolean[]): void => {
  visited[idx] = true;

  const nodes = graph[idx];
  for (let node of nodes) {
    if (!visited[node]) dfs(node, graph, visited);
  }
};
