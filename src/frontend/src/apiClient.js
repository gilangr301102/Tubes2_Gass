// import axios from "axios";

// export const axiosClient = axios.create({
//   baseURL: `https://localhost:8080`,
//   headers: {
//     Accept: "application/json",
//     "Content-Type": "application/x-www-form-urlencoded",
//   },
// });

// export const postIDSAlgo = async (parameter) => {
//   return await axiosClient
//     .post("/wikiraceIDS", parameter)
// };

// export const postBFSAlgo = async (parameter) => {
//   return await axiosClient
//   .post("/wikiraceBFS", parameter);
// };

import axios from "axios";

export const axiosClient = axios.create({
  baseURL: `http://localhost:8080`,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/x-www-form-urlencoded",
  },
});

export const postIDSAlgo = async (parameter) => {
  try {
    const response = await axiosClient.post("/wikiraceIDS", parameter);
    return response.data;
  } catch (error) {
    console.error("Error in postIDSAlgo:", error);
    throw error;
  }
};

export const postBFSAlgo = async (parameter) => {
  try {
    const response = await axiosClient.post("/wikiraceBFS", parameter);
    return response.data;
  } catch (error) {
    console.error("Error in postBFSAlgo:", error);
    throw error;
  }
};
