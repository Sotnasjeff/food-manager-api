# Project Title: Food-Manager-Api

# Description:
 - This is a project demo created to make some presentation about GRPC in study group called RNA, and this project persists data in a database using unary and binary communication.

 # Prerequisites to run locally:
 - Make the following steps to do it:
    - Install Go in your machine: https://go.dev/doc/install 
    - Install Protobuf Compiler: https://grpc.io/docs/protoc-installation/
    - Install Protoc-gen-go and protoc-gen-go-grpc GO plugins: https://grpc.io/docs/languages/go/quickstart/
    - Install Docker: https://docs.docker.com/get-docker/
    - Install Evans Client (GRPC's client to make request to server): https://github.com/ktr0731/evans
    - Run the commands created on Makefile:
        - First, type in your terminal: make proto, to remove the protofiles and generate a new protofile
        - In your terminal type: Go Mod Tidy to install the dependencies not downloaded.
        - After updated the project, type in terminal: Make Server
        - After running the server, open a new tab in terminal and type: Make Evans
        - In Evans, you can make the call for the methods to test it and send request.