# Desenvolvendo uma RESTful API com _Go_  e _Gin_

Este tutorial apresenta os conceitos básicos de como escrever uma API de serviço Web RESTful com **Go** e **Gin Web Framework** (Gin).   
O **Gin** simplifica muitas tarefas de codificação associadas à criação de aplicativos da Web.   
Neste tutorial, vamos usar o Gin para:

- rotear solicitações;
- recuperar detalhes da solicitação;
- e empacotar JSON para respostas. 

Neste tutorial, vamos construir um servidor de API RESTful com dois endpoints.   
Seu projeto de exemplo será um repositório de dados sobre discos de jazz antigos.    
O tutorial inclui as seguintes seções: 

- Design API endpoints.
- Criação da estrutura de diretórios do projeto.
- Criação dos dados em memória.
- Escrever um handler para retornar todos os items.
- Escrever um handler para adicionar um novo item.
- Escrever um handler retornar um Item específico.

---

## Create project

```
mkdir web-service-gin
cd web-service-gin
go mod init example/web-service-gin
```

### Import gin

```go
import (
    "net/http"

    "github.com/gin-gonic/gin"
)
```


## Model

```go
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}
```

## Controller

```go
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
```

## Router

```go
func main() {
	fmt.Println("Tutorial: Developing a RESTful API with Go and Gin")
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)

	log.Fatal(router.Run("localhost:8080"))
}
```

## Run

> Na linha de comando, use go `get` para adicionar o módulo `github.com/gin-gonic/gin` como uma dependência para o seu módulo.   
> Use um argumento de ponto `.` para obter dependências do seu código no diretório atual.
```
go get .
```
output: `go get: added github.com/gin-gonic/gin v1.7.2`

depois de obter as dependencias rode: 
```
go run .
```

---

## Reference

- [Tutorial: Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)