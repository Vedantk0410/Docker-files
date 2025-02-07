
# Toofan backend Project Setup 

Project Overview

1)Setting Up the Golang Backend Service

2)Setting Up the Flask Bot Service


3)Creating the Docker Compose Configuration

4)Building and Running the Application






## Setting Up the Golang Backend Service

## Step 1: Create the Dockerfile

1) Base Image: Use the official Golang image from Docker Hub.

2) Specify the base image as golang:1.21-alpine.
Set WorkingDirectory: Define the working directory inside the container.

3) Set the working directory to /app.Copy Dependency Files: Copy go.mod and go.sum files to the container.Copy these files to the working directory.Download Dependencies: Use go mod download to install dependencies.

4) Run the command to download the necessary Go modules.Copy Application Code: Copy the application code from the app directory.

5) Copy the rest of the application code to the container.Build the Application: Build the Go application and name the output binary main.

6) Use the go build command to compile the application.Expose Port: Expose port 8003 to allow external access.

7) Specify the port that the application will use.Run Command: Define the command to run the application.Set the command to run the compiled binary.

## Step2: Build the Docker Image
docker build -t toofan-core .

## Setting Up the Flask Bot Service


## Step 1: Create the Dockerfile for flask
1) Base Image: Use the official Python image from Docker Hub.

2) Specify the base image as python:3.13-alpine.Set Working Directory: Define the working directory inside the container.Set the working directory to /app.Copy Application Code: Copy the entire contents of the current directory to the container.

3) Copy all files to the working directory.Install Dependencies: Install Python dependencies from requirements.txt.Use pip to 

4) Install the required packages.Expose Port: Expose port 5003 to allow external access.

5) Specify the port that the application will use.Environment Variable: Set the FLASK_APP environment variable.Define the entry point for the Flask application.

6) Run Command: Define the command to run the Flask application.Set the command to start the Flask server.
## Step 2: Build the Docker Image
docker build -t toofan-bot-service .


## Creating the Docker Compose Configuration
## Step 1: Define Services

1) Version: Specify the version of the Docker Compose file format Use version 3.8.

2) Services: Define the toofan-core and toofan-bot-service services.
1)Specify the build context and Dockerfile for each service.

2)Name the containers toofan-core and toofan-bot-service.

3)Map ports 8003 and 5003 on the host to the same ports in the containers.

3) Networks: Connect both services to a custom bridge network called toofan-network.Define a network to enable communication between the services.
## Building and Running the Application

To build and run the application using Docker Compose, navigate to the directory containing the docker-compose.yml file and execute:

docker-compose up --build
## Accessing the Services
1) The Golang backend service will be accessible at http://localhost:8003.
2) The Flask bot service will be accessible at http://localhost:5003.
