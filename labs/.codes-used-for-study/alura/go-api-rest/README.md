# Alura - Go API Rest

Projeto feito seguindo o curso da Alura de API Rest com Go

## Ferramentas

- [Gorilla Mux - Router](https://github.com/gorilla/mux)
- Docker / docker-compose
- Postgres / PgAdmin
- [GORM - The fantastic ORM library for Golang](https://gorm.io/)


## Desafio 1 - Json, Request, Response

- Criar uma API Rest de **Personalidades**
  - Implementar um endpoint `/api/personalidades`
  - Retornar uma lista de um tipo Personalidade serializado em `JSON`, essa lista neste momento deve ser inicializada em memória
  - Modelo Personalidade: `Nome`, `Historia`

**Estrutura do Projeto**
```
.
├── controllers
│   └── controllers.go
├── models
│   └── personalidades.go
├── routes
│   └── routes.go
├── main.go    
```

## Desafio 2 - Roteador, recursos por ID

- Adicionar o campo `Id` no modelo *Personalidade* 
- Implementar um *endpoint* para obter uma Personalidade pelo *ID* `/api/personalidades/1`
- Utilizar a lib `gorilla/mux` para facilitar o Roteamento e obter o `ID` passado por *PathParameter*

Instalando dependencia do `mux`:
```
go get -u github.com/gorilla/mux
```

## Desafio 3 - Docker, Conexão com banco e exibir dados do banco

- Criar um `docker-compose` para subir um Postgres e um PgAdmin e cria a tabela já populando com dados do Script SQL.
- Acessar o banco de dados usando a lib `gorm`

Instalando dependencias do `gorm` e `posgres`:
```
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
```
