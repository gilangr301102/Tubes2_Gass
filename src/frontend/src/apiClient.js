import axios from "axios";
import React from "react";
  
export const axiosClient = axios.create({
  baseURL: `http://localhost:8080`,
  headers: {
    Accept: "application/json",
    "Content-Type": "application/json",
  },
});

export const postIDSAlgo = async (parameter) => {
  return await axiosClient
    .post("/wikiraceIDS", parameter)
};

export const postBFSAlgo = async (parameter) => {
  return await axiosClient.post("/wikiraceBFS", parameter);
};
