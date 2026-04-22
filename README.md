# 🛒 E-Commerce Order Payment Platform

A microservices-based backend system modeling core e-commerce workflows,
specifically Order creation and Payment processing, built using Go and
gRPC.

## 🏗️ Architecture Overview

This project follows Hexagonal Architecture (Ports & Adapters) to ensure
separation of concerns:

microservices/ ├── order/ \# Handles order creation and orchestration
├── payment/ \# Handles payment processing

### Flow

Client → Order Service → Payment Service

-   Order service manages order lifecycle\
-   Payment service processes payments\
-   Services communicate via gRPC (Protocol Buffers)

## ⚙️ Tech Stack

-   Go\
-   gRPC + Protocol Buffers\
-   Hexagonal Architecture\
-   Docker

## 🚀 Setup

### Install dependencies

go mod tidy

### Generate gRPC stubs

protoc --proto_path=proto\
--go_out=.\
--go-grpc_out=.\
--go_opt=module=github.com/vteja-code/microservices\
--go-grpc_opt=module=github.com/vteja-code/microservices\
proto/order/order.proto\
proto/payment/payment.proto

### Run services

cd order go run cmd/main.go

cd payment go run cmd/main.go

## 🐳 Docker

### Build

docker build -t order-service ./order docker build -t payment-service
./payment

### Run

docker run -p 50051:50051 order-service docker run -p 50052:50052
payment-service

## 📌 Features

-   Microservices-based design\
-   gRPC communication\
-   Clean architecture\
-   Scalable structure

## 🔮 Future Work

-   Inventory service\
-   Shipping service\
-   API Gateway
