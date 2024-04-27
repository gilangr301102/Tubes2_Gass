import React, { useState } from 'react';
import Graph from 'react-graph-vis';
import axios from 'axios'; // Import Axios
import './App.css'; // Import CSS file for styling
import { postBFSAlgo, postIDSAlgo } from './apiClient';
import { useEffect } from 'react';

const App = () => {
  const [algorithm, setAlgorithm] = useState('IDS');
  const [sourceTitle, setSourceTitle] = useState('');
  const [goalTitle, setGoalTitle] = useState('');
  const [maxDepth, setMaxDepth] = useState('');
  const [isFindAll, setIsFindAll] = useState(false); // Changed to false initially
  const [nodes, setNodes] = useState([]);
  const [paths, setPaths] = useState([]);
  const [paths2, setPaths2] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [responseMessage, setResponseMessage] = useState('');
  const [elapsedTime, setElapsedTime] = useState('');
  const [numOfArticleChecked, setNumOfArticleChecked] = useState('');
  const [numOfNodeArticleVisited, setNumOfNodeArticleVisited] = useState('');
  const [numberOfPath, setNumberOfPath] = useState('');

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

  const handleSubmit = (e) => {
    e.preventDefault();
    setNodes([]);
    setPaths([]);
    setResponseMessage('');
    setElapsedTime('');
    setNumOfArticleChecked('');
    setNumOfNodeArticleVisited('');
    setNumberOfPath('');
    getResult();
  };

  const getResult = async () => {
    try {
      setIsLoading(true);
      const request = {
        sourceTitle: sourceTitle,
        goalTitle: goalTitle,
        isFindAll: isFindAll ? "1" : "0",
      };
      if (algorithm === 'IDS') {
        request.maxDepth = maxDepth;
      }
      let response;
      if (algorithm === 'IDS') {
        response = await postIDSAlgo(request);
      } else {
        response = await postBFSAlgo(request);
      }
      setResponseMessage(response.message);
      setElapsedTime(response.elapsed_time);
      setNumOfArticleChecked(response.num_of_article_checked);
      setNumOfNodeArticleVisited(response.num_of_node_article_visited);

      if (isFindAll) {
        setNumberOfPath(response.number_of_path);
        setPaths2(response.path);
      }else{
        setPaths2([response.path]);
      }
    }
    catch (error) {
      console.error('Error in getResult:', error);
      setResponseMessage('Error occurred. Please check console for more details.');
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
      const parsedPaths = paths2 || [];
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
  }, [paths2]);

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
            <label htmlFor="sourceTitle">Start Title:</label>
            <input
              type="text"
              id="sourceTitle"
              value={sourceTitle}
              onChange={(e) => setSourceTitle(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label htmlFor="goalTitle">Goal Title:</label>
            <input
              type="text"
              id="goalTitle"
              value={goalTitle}
              onChange={(e) => setGoalTitle(e.target.value)}
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
          <button type="submit" disabled={isLoading} onClick={getResult}>
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
}

export default App;
