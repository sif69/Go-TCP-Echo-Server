# 🚀 Go TCP Echo Server

A simple, concurrent TCP echo server written in Go. This server listens on a user-specified port, accepts multiple client connections simultaneously, reads data sent by clients, and echoes the same data back. The project also includes learning notes on TCP server internals, Go’s concurrency model, and practical usage examples.

---

## ✨ Features

- **Custom Port Listening:** Specify the port via command-line arguments.
- **Concurrent Client Handling:** Uses goroutines to serve multiple clients at once.
- **Raw TCP Stream Processing:** Reads and writes directly to TCP connections.
- **Echo Functionality:** Sends received data back to the client.
- **Resource Cleanup:** Uses `defer` to ensure sockets are closed properly.
- **Server and Connection Logging:** Logs server status and client interactions.

---

## 🛠️ How It Works

### 1. Listening for Connections

```go
listener, err := net.Listen("tcp", port)
```
Binds to all interfaces on the specified port and starts listening for incoming TCP connections.

### 2. Accepting Clients

```go
conn, err := listener.Accept()
```
Waits for a client to connect (completes the TCP handshake). Returns a `net.Conn` object for communication with the client.

### 3. Handling Clients Concurrently

```go
go handleConnection(conn)
```
Each client connection is handled in its own goroutine, allowing the server to serve many clients at the same time without blocking.

### 4. Reading and Echoing Data

```go
bytes, err := reader.ReadBytes('\n')
conn.Write([]byte("Echo: " + string(bytes)))
```
Reads data from the client until a newline. Sends the same data back, prefixed with "Echo: ".

### 5. Resource Management

```go
defer listener.Close()
defer conn.Close()
```
Ensures that sockets are closed when no longer needed, preventing resource leaks.

---

## 🧩 Function Breakdown

- **main:**  
  Parses the port from command-line arguments, sets up the TCP listener, enters an infinite loop to accept new connections, and launches a new goroutine for each client.

- **handleConnection:**  
  Reads data from the client, echoes the data back, handles errors, and closes the connection when done.

---

## ⚡ Goroutines and Concurrency

- **Goroutines** are lightweight threads managed by Go.
- Each client connection is handled in a separate goroutine, allowing the server to process many clients in parallel.
- This model makes Go ideal for building scalable network servers.

---

## 🌟 Why Use Go for Servers?

- **Simplicity:** Minimal boilerplate for networking code.
- **Concurrency:** Goroutines and channels make concurrent programming easy and efficient.
- **Performance:** Go’s runtime efficiently manages thousands of concurrent connections.
- **Cross-platform:** Go binaries run on all major operating systems.

---

## 🎯 Purpose and Use Cases for an Echo Server

- **Testing:**  
  - Verify network connectivity and client implementations.
  - Benchmark network performance and latency.
- **Learning:**  
  - Understand TCP/IP, sockets, and concurrency.
- **Debugging:**  
  - Ensure clients can send and receive data correctly.
- **Foundation:**  
  - Serve as a starting point for more complex protocols (e.g., chat servers).

---

## 🚦 How to Run

### Start the Server

```sh
go run main.go 8080
```
You’ll see:  
`Listening on port :8080...`

### Test the Server

#### Windows

```sh
telnet localhost 8080
```
Type messages and see them echoed back.

#### Linux

```sh
echo "Hello World" | nc localhost 8080
```

#### PowerShell Example

```powershell
$client = New-Object System.Net.Sockets.TcpClient
$client.Connect("localhost", 8080)
$stream = $client.GetStream()
$writer = New-Object System.IO.StreamWriter($stream)
$reader = New-Object System.IO.StreamReader($stream)
$writer.WriteLine("Hello from PowerShell")
$writer.Flush()
$response = $reader.ReadLine()
Write-Host "Received: $response"
$client.Close()
```

---

## 📚 TCP Theory (Learning Notes)

- **TCP Server Lifecycle:**
  1. **Bind & Listen:** OS reserves the port for the server.
  2. **Accept:** Waits for a client to connect (TCP handshake).
  3. **Data Transfer:** Reliable, ordered byte stream between client and server.
  4. **Close:** Graceful connection teardown.

- **TCP Connection Stages:**
  1. **Handshake:** SYN → SYN+ACK → ACK
  2. **Data Transfer:** Client/server exchange data.
  3. **Teardown:** FIN → ACK → FIN → ACK

---

## 🧪 Example Session

**Terminal 1 (Server):**
```sh
go run main.go 8080
Listening on port :8080...
```

**Terminal 2 (Client using netcat):**
```sh
nc localhost 8080
Hello Server!
Echo: Hello Server!
How are you?
Echo: How are you?
```

---

## 📝 Client Testing Commands

```go
// To write response from client side:
//   i) In Windows: telnet localhost <port>
//   ii) In Linux: echo Hello world | nc localhost 9090
```

---

## 🚀 Future Improvements

- Develop a concurrent TCP server in Go handling multiple clients using goroutines and channels.
- Extend into a chat application with logging, monitoring, and custom protocol support.
- Optimize the server to handle 10,000+ concurrent connections with TLS security.

---

## 🏃 Agile Step-by-Step Improvement Plan

1. **Backlog Creation:**  
   - List features: authentication, chat rooms, logging, monitoring, TLS, etc.

2. **Sprint 1:**  
   - Refactor codebase for modularity.
   - Add structured logging.

3. **Sprint 2:**  
   - Implement basic chat functionality (broadcast messages to all clients).
   - Add user identification.

4. **Sprint 3:**  
   - Integrate monitoring (metrics, connection counts).
   - Add error handling and recovery.

5. **Sprint 4:**  
   - Implement TLS for secure connections.
   - Optimize for high concurrency (benchmark and tune).

6. **Sprint 5:**  
   - Add custom protocol support (commands, message types).
   - Write comprehensive tests and documentation.

7. **Review & Retrospective:**  
   - Gather feedback, prioritize next improvements, and iterate.

---

**This project is a solid foundation for learning Go networking, building scalable servers, and experimenting with real-world TCP applications.**
