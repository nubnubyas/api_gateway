# Orbital X ByteDance Api Gateway

This project is an API Gateway built with the Cloudwego Hertz server, Kitex client and RPC servers, external databases for RPC servers and Nacos registry for service registration and discovery. It allows clients to send JSON-encoded HTTP requests, which are then translated to Thrift binary format and forwarded to the appropriate RPC backend server using Kitex's Generic-Call mechanism. 

## Table of Contents

- [Getting Started](#getting-started)
- [Dependencies](#dependencies)
- [Usage](#usage)
- [Acknowledgements](#acknowledgements)

## Getting Started

To get started with this project, follow these steps:

1. Clone the repository to your local machine. (nacos not included in the repo as of now, please install separately)
2. Install the necessary dependencies, such as Go, Hertz, Kitex, and Nacos (https://nacos.io/en-us/docs/quick-start.html).
3. Start the Nacos registry on your machine.
4. Run the API Gateway server by running the main.go file under hertz_gateway folder.
5. Send a JSON-encoded HTTP request to the API Gateway (host :8080) using a tool like cURL, Postman or Thunderclient .
6. Check that the response from the API Gateway server is in JSON-encoded HTTP format.

## Dependencies

This project depends on the following packages:

Cloudwego Hertz  
Cloudwego Kitex  
Cloudwego/Kitex-contrib/Nacos

Run ``` go mod tidy ``` in the root directory if required 

## Usage
1. Start the Nacos Server and access the Nacos Console
Run ```bin\startup.cmd -m standalone``` in the Nacos Directory (you should have installed)
Access Console on a web browser ```http://localhost:8848/nacos```

3. Start the Hertz Gateway
Run ```go run .``` in the Hertz gateway directory in a terminal

4. Start the RPC Servers
In the respective Kitex_servers directory, run ```go run .``` in a separate terminal.
Hearbeat/Periodic Status sent to the Nacos Server can be observed in the terminal.

6. Testing by sending HTTP Requests

Here's the details of the Example API Service in the repo provided for testing.  
Format: (path to make request, request method, request body/params expected)  
  a. Student_Api  
    - student_api/queryStudent, GET, [int num]  
    - student_api/insertStudent, POST, [int num, string name, string gender]  
  b. calculator  
    - calculator/get, POST, [string name, int num1, int num2, string operation]  

  ```curl http://localhost:8080/```  
This command can be used to manually updates the gateway after any IDL files are being updates, removed or added.

Here's an example of how to send a JSON-encoded HTTP request to the API Gateway using cURL:  

```curl -X POST -H "Content-Type: application/json" -d '{"Num":"2","Name":"Ben","Gender":"Male"}' -u monday:mahjong http://localhost:8080/student_api/insert```  
This command sends a POST request with a JSON-encoded body containing a num, name and gender field to the API Gateway server running on localhost at port 8080. 

Reference our [Demo video](https://drive.google.com/file/d/1fzpVKpczA3NTpi2iYsOFELMmpU3IMR0v/view?usp=sharing) if needed 

## Acknowledgements

Cloudwego for developing Hertz, Kitex, and Nacos
The open source community for their contributions to the Go programming language and related packages.
