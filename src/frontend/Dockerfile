# Use an official Node runtime as a parent image for frontend
FROM node:14 AS frontend

# Set the Current Working Directory inside the container
WORKDIR /

# Copy package.json and package-lock.json to the PWD(Present Working Directory) inside the container
COPY package*.json ./

# Clear npm cache and install dependencies
RUN npm cache clean --force && npm install --verbose

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY / .

# Expose port 3000 to the outside world
EXPOSE 3000

# Command to run the frontend
CMD ["npm", "run", "start"]