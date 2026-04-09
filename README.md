# RPC Made Simple (Go Server + Multi-language Clients)

This repository demonstrates how to build a simple RPC system in Go and interact with it using different protocols and client languages (Go, JavaScript, and PHP).

It covers:

- Native Go RPC over TCP
- JSON-RPC over TCP sockets
- HTTP-based RPC
- HTTP + JSON-RPC
- Cross-language clients (JS & PHP)

📖 **Blog Post:**  
Read the full explanation here:  
👉 https://your-blog-link-here.com

---

## 📁 Project Structure

```
.
├── http
│   ├── client
│   └── server
├── jsonrpc
│   ├── client
│   └── server
├── tcp
│   ├── client
│   └── server
├── http-jsonrpc
│   ├── client
│   └── server
├── utils
│   └── checkError.go
├── socket-js-rpc.js
├── socket-php-rpc.php
├── http-js-rpc.js
└── http-php-rpc.php
```

---

## ⚙️ Features

- Simple RPC service (`Arith`) with:
  - `Multiply`
  - `Divide`
- Transport layers:
  - TCP
  - HTTP
- Encoding formats:
  - Go RPC (gob)
  - JSON-RPC
- Cross-language interoperability:
  - JavaScript (Node.js)
  - PHP

---

## 🚀 Getting Started

### Clone the repository

```
git clone https://github.com/omept/rpc.git
cd rpc
```

---

## 🖥️ Running the Servers

Each implementation runs on port `1234`.

### TCP RPC Server
```
cd tcp/server
go run main.go
```

### JSON-RPC (TCP) Server
```
cd jsonrpc/server
go run main.go
```

### HTTP RPC Server
```
cd http/server
go run main.go
```

### HTTP + JSON-RPC Server
```
cd http-jsonrpc/server
go run main.go
```

---

## 🧪 Running Clients

### Go Clients

```
cd tcp/client
go run main.go localhost:1234
```

```
cd jsonrpc/client
go run main.go localhost:1234
```

```
cd http/client
go run main.go localhost
```

---

### JavaScript Clients

```
node socket-js-rpc.js
```

```
node http-js-rpc.js
```

---

### PHP Clients

```
php socket-php-rpc.php
```

```
php http-php-rpc.php
```

---

## 🔌 API Methods

###  Arith.Multiply

```
{
  "method": "Arith.Multiply",
  "params": [{"A": 17, "B": 8}],
  "id": 1
}
```

---

### Arith.Divide

```
{
  "method": "Arith.Divide",
  "params": [{"A": 17, "B": 8}],
  "id": 1
}
```

Response:

```
{
  "Quo": 2,
  "Rem": 1
}
```

```
Multiply Result: 136

```

---

## ⚠️ Notes

- No external dependencies required

---

