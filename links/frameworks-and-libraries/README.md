# Frameworks e Bibliotecas

Voltar para: [Livros](../books/README.md)

- [Frameworks e Bibliotecas](#frameworks-e-bibliotecas)
  - [Web](#web)
    - [Gin Web Framework](#gin-web-framework)
    - [gorilla/mux](#gorillamux)
    - [swag - Swagger](#swag---swagger)
    - [jwt-go](#jwt-go)
  - [Dependency Injetion](#dependency-injetion)
    - [Wire: Automated Initialization in Go](#wire-automated-initialization-in-go)
  - [Database](#database)
    - [GORM](#gorm)
  - [Kafka](#kafka)
    - [gogen-avro](#gogen-avro)

## Web

### [Gin Web Framework](https://github.com/gin-gonic/gin)

> Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter. If you need performance and good productivity, you will love Gin.

### [gorilla/mux](https://github.com/gorilla/mux)

> Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.   
> The name mux stands for "HTTP request multiplexer". Like the standard http.ServeMux, mux.Router matches incoming requests against a list of registered routes and calls a handler for the route that matches the URL or other conditions. 

### [swag - Swagger](https://github.com/swaggo/swag)

> Swag converts Go annotations to Swagger Documentation 2.0. We've created a variety of plugins for popular [Go web frameworks](https://github.com/swaggo/swag#supported-web-frameworks). This allows you to quickly integrate with an existing Go project (using Swagger UI).

### [jwt-go](https://github.com/golang-jwt/jwt)

> A go (or 'golang' for search engine friendliness) implementation of JSON Web Tokens.

## Dependency Injetion

### [Wire: Automated Initialization in Go](https://github.com/google/wire)

> Wire is a code generation tool that automates connecting components using dependency injection. Dependencies between components are represented in Wire as function parameters, encouraging explicit initialization instead of global variables. Because Wire operates without runtime state or reflection, code written to be used with Wire is useful even for hand-written initialization.   
> For an overview, see the [introductory blog post](https://go.dev/blog/wire).

## Database

### [GORM](https://gorm.io/)

> The fantastic ORM library for Golang
> * Full-Featured ORM
> * Associations (has one, has many, belongs to, many to many, polymorphism, single-table inheritance)
> * Hooks (before/after create/save/update/delete/find)
> * Eager loading with Preload, Joins
> * Transactions, Nested Transactions, Save Point, RollbackTo to Saved Point
> * Context, Prepared Statement Mode, DryRun Mode
> * Batch Insert, FindInBatches, Find/Create with Map, CRUD with SQL Expr and Context Valuer
> * SQL Builder, Upsert, Locking, Optimizer/Index/Comment Hints, Named Argument, SubQuery
> * Composite Primary Key, Indexes, Constraints
> * Auto Migrations
> * Logger
> * Extendable, flexible plugin API: Database Resolver (multiple databases, read/write splitting) / Prometheus…
> * Every feature comes with tests
> * Developer Friendly

## Kafka

### [gogen-avro](https://github.com/actgardner/gogen-avro)

> Generates type-safe Go code based on your Avro schemas, including serializers and deserializers that support Avro's schema evolution rules. Also supports deserializing generic Avro data (in beta).

---

Ir para próximo: [Links Diversos](../miscellaneous/README.md)


