import React, { useState } from 'react';
import Graph from 'react-graph-vis';
import './App.css'; // Import CSS file for styling

const GraphVisualization = ({ paths, nodes }) => {
  // JSON data representing nodes and edges
  const data = {
    nodes: nodes.map((label, index) => ({ id: index + 1, label })),
    edges: [],
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

  // Generate highlighted edges based on input paths
  const highlightedEdges = paths.map((p, index) => {
    return p.reduce((edges, node, i) => {
      if (i < p.length - 1) {
        edges.push({
          from: nodes.findIndex(n => n === node) + 1,
          to: nodes.findIndex(n => n === p[i + 1]) + 1,
          color: index === 0 ? 'red' : 'blue', // Change color based on index
        });
      }
      return edges;
    }, []);
  }).flat();

  // Combine original edges and highlighted edges
  const allEdges = data.edges.concat(highlightedEdges);

  // Data with all edges for visualization
  const graphData = { nodes: data.nodes, edges: allEdges };

  return <Graph graph={graphData} options={options} />;
};

const App = () => {
  const [algorithm, setAlgorithm] = useState('IDS');
  const [sourceTitle, setsourceTitle] = useState('');
  const [goalTitle, setgoalTitle] = useState('');
  const [maxDepth, setMaxDepth] = useState('');
  const [isFindAll, setIsFindAll] = useState('');
  const [nodes, setNodes] = useState([]);
  const [paths, setPaths] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [responseMessage, setResponseMessage] = useState('');
  const [elapsedTime, setElapsedTime] = useState('');
  const [numOfArticleChecked, setNumOfArticleChecked] = useState('');
  const [numOfNodeArticleVisited, setNumOfNodeArticleVisited] = useState('');
  const [numberOfPath, setNumberOfPath] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    setIsLoading(true);
    const formData = new FormData();
    formData.append('sourceTitle', sourceTitle);
    formData.append('goalTitle', goalTitle);
    if (algorithm === 'IDS') {
      formData.append('maxDepth', maxDepth);
      formData.append('isFindAll', isFindAll);
    } else {
      formData.append('isFindAll', isFindAll);
    }

    const apiEndpoint = algorithm === 'IDS' ? 'http://localhost:8080/wikiraceIDS' : 'http://localhost:8080/wikiraceBFS';

    try {
      const response = await fetch(apiEndpoint, {
        method: 'POST',
        body: formData,
      });
      const data = await response.json();
      setResponseMessage(data.message);
      setElapsedTime(data.elapsed_time);
      setNumOfArticleChecked(data.num_of_article_checked);
      setNumOfNodeArticleVisited(data.num_of_node_article_visited);
      setNumberOfPath(data.number_of_path);
      const parsedPaths = data.path || [];
      const allNodes = parsedPaths.reduce((acc, path) => {
        path.forEach(node => {
          if (!acc.includes(node.trim())) {
            acc.push(node.trim());
          }
        });
        return acc;
      }, []);
      setNodes(allNodes);
      setPaths(parsedPaths);
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="App">
      <h1>Graph Visualization</h1>
      <div className="algorithm-menu">
        <label>
          Choose Algorithm:
          <select value={algorithm} onChange={(e) => setAlgorithm(e.target.value)}>
            <option value="IDS">IDS</option>
            <option value="BFS">BFS</option>
          </select>
        </label>
      </div>
      <div className="form-container">
        <form onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="sourceTitle">Source Title:</label>
            <input
              type="text"
              id="sourceTitle"
              value={sourceTitle}
              onChange={(e) => setsourceTitle(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="goalTitle">Goal Title:</label>
            <input
              type="text"
              id="goalTitle"
              value={goalTitle}
              onChange={(e) => setgoalTitle(e.target.value)}
              required
            />
          </div>
          {algorithm === 'IDS' && (
            <div className="form-group">
              <label htmlFor="maxDepth">Max Depth:</label>
              <input
                type="number"
                id="maxDepth"
                value={maxDepth}
                onChange={(e) => setMaxDepth(e.target.value)}
                required
              />
            </div>
          )}
          <div className="form-group">
            <label>
              Find All:
              <input
                type="checkbox"
                checked={isFindAll}
                onChange={(e) => setIsFindAll(e.target.checked)}
              />
            </label>
          </div>
          <button type="submit" disabled={isLoading}>
            {isLoading ? 'Loading...' : 'Generate Graph'}
          </button>
        </form>
      </div>
      {/* Display response data */}
      {responseMessage && (
        <div className="response-container">
          <h2>Response Data</h2>
          <p>Response Message: {responseMessage}</p>
          <p>Elapsed Time: {elapsedTime}</p>
          <p>Number of Articles Checked: {numOfArticleChecked}</p>
          <p>Number of Node Articles Visited: {numOfNodeArticleVisited}</p>
          <p>Number of Paths: {numberOfPath}</p>
        </div>
      )}
      {/* Display graph visualization */}
      {nodes.length > 0 && (
        <div className="graph-container">
          <h2>Graph Visualization</h2>
          <GraphVisualization paths={paths} nodes={nodes} />
        </div>
      )}
    </div>
  );
};

export default App;
