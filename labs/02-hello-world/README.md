# Hello World de Go 
Voltar para: [Pré-requisitos](../01-prereqs/README.md)

Um exemplo simples de **Hello World**  com Go

- [Hello World de Go](#hello-world-de-go)
  - [Pré-requisitos](#pré-requisitos)
  - [Código](#código)
  - [Referencia](#referencia)

## Pré-requisitos

* Instalação do **Go** -> [veja aqui](../01-prereqs/README.md)
* **Alguma experiência em programação.** O código aqui é bem simples, mas ajuda saber algo sobre funções.
* **Uma ferramenta para editar seu código.** Qualquer editor de texto que você tenha funcionará bem. A maioria dos editores de texto tem um bom suporte para **Go**. Os mais populares são VSCode (gratuito), GoLand (pago) e Vim (gratuito).
* **Um terminal de comando.** Go funciona bem usando qualquer terminal no __Linux__ e __Mac__ e no __PowerShell__ ou cmd no __Windows__.

## Código

1. Abra um prompt de comando e digite `cd` para ir ao seu diretório inicial.
> No Windows use `cd %HOMEPATH%`
2. Crie um diretório `hello` para seu primeiro código-fonte **Go**.
```shell
mkdir hello
cd hello
```
3. Habilite o rastreamento de dependência para seu código.
> Quando seu código importa pacotes contidos em outros módulos, você gerencia essas dependências por meio do próprio módulo do seu código. Esse módulo é definido por um arquivo `go.mod` que rastreia os módulos que fornecem esses pacotes. Esse arquivo `go.mod` permanece com seu código, inclusive em seu repositório de código-fonte. Para habilitar o rastreamento de dependência para seu código criando um arquivo `go.mod`, execute o comando `go mod init`, dando a ele o nome do módulo em que seu código estará. O nome é o caminho do módulo do módulo. No desenvolvimento real, o caminho do módulo normalmente será o local do repositório onde seu código-fonte será mantido. Por exemplo, o caminho do módulo pode ser `github.com/mymodule`. Se você planeja publicar seu módulo para outros usarem, o caminho do módulo deve ser um local de onde as ferramentas **Go** possam baixar seu módulo. Para saber mais sobre como nomear um módulo com um caminho de módulo, consulte [Gerenciando dependências](https://go.dev/doc/modules/managing-dependencies#naming_module).

Para os propósitos deste tutorial, basta usar `example/hello`.   
```shell
$ go mod init example/hello
go: creating new go.mod: module example/hello
```
4. Em seu editor de texto, dentro do diretório `hello` crie um arquivo `hello.go` para escrever seu código.
5. Cole o código a seguir em seu arquivo `hello.go` e salve o arquivo.
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

6. No terminal execute seu código para ver o resultado.
```shell
$ go run .
Hello, World!
```

> :exclamation: Se temos um código já pronto que contem imports de libs externas, precisamos primeiro roda um camando para instalar todas as dependencias antes de rodar e testar o código: `go mod tidy`

## Referencia

* [Tutorial: Get started with Go](https://go.dev/doc/tutorial/getting-started).
* [Para outros tutoriais, consulte Tutoriais.](https://go.dev/doc/tutorial/)
   
---

Ir para próximo: [Hello World Go para Desenvolvimento Web](../03-hello-world-web/README.md)   