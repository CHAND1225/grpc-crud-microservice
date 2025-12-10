# gRPC CRUD Microservice (Golang + PostgreSQL)

This project implements a simple **User Service** using **Golang**, **gRPC**, and **PostgreSQL**.

## Features

- Create and fetch users over gRPC.
- PostgreSQL as persistent storage.
- Clean separation of API definitions (`proto/user.proto`) and server implementation (`server/main.go`).
- Database initialized via `db/init.sql` and Docker.

## Tech Stack

- Go
- gRPC
- Protocol Buffers
- PostgreSQL
- Docker

## High-Level Architecture

Client (gRPC) → UserService (Go server) → PostgreSQL (`users` table)

## Concepts Demonstrated

- gRPC-based RPC communication
- CRUD operations with PostgreSQL
- Service–database separation & layering
- Containerized database for local development
