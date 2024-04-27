import React from 'react';
import Graph from 'react-graph-vis';

const GraphVisualization = () => {
  // JSON data representing nodes and edges
  const data = {
    nodes: [
      { id: 1, label: 'Sweden' },
      { id: 2, label: 'Republic of Ireland' },
      { id: 3, label: 'Dracula' },
      { id: 4, label: 'Gothic (architecture)' },
    ],
    edges: [
      { from: 1, to: 2 },
      { from: 2, to: 3 },
      { from: 1, to: 4 },
      { from: 4, to: 3 },
    ],
  };

  // Options for the graph visualization
  const options = {
    layout: {
      hierarchical: false,
    },
    edges: {
      color: '#000000',
    },
    height: '500px',
  };

  // Highlighted paths between start node (Sweden) and goal node (Dracula)
  const highlightedEdges = [
    { from: 1, to: 2, color: 'red' },
    { from: 2, to: 3, color: 'red' },
    { from: 1, to: 4, color: 'blue' },
    { from: 4, to: 3, color: 'blue' },
  ];

  // Combine original edges and highlighted edges
  const allEdges = data.edges.concat(highlightedEdges);

  // Data with all edges for visualization
  const graphData = { nodes: data.nodes, edges: allEdges };

  return <Graph graph={graphData} options={options} />;
};

const App = () => {
  const path = [
    ["Sweden", "Republic of Ireland", "Dracula"],
    ["Sweden", "Gothic (architecture)", "Dracula"]
  ];

  return (
    <div className="App">
      <h1>Graph Visualization</h1>
      <GraphVisualization path={path} />
    </div>
  );
};

export default App;