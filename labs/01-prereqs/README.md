# Pré-requisitos
Voltar para: [Home](../../README.md)

Como instalar os pré-requisitos necessários para executar os exemplos neste repositório

- [Pré-requisitos](#pré-requisitos)
  - [Baixar e instalar](#baixar-e-instalar)
    - [Instalação do Go no Windows](#instalação-do-go-no-windows)
    - [Instalação do Go no Mac](#instalação-do-go-no-mac)
    - [Instalação do Go no Linux](#instalação-do-go-no-linux)
  - [Configurações](#configurações)
    - [Plugins útiis para o Visual Studio Code](#plugins-útiis-para-o-visual-studio-code)
    - [Outros plugins úteis para Visual Studio Code](#outros-plugins-úteis-para-visual-studio-code)
  - [Links auxiliares](#links-auxiliares)

## Baixar e instalar

* [Baixar a ultima versão do Go](https://go.dev/dl/)
* [Baixar IDE Visual Studio Code](https://code.visualstudio.com/download)
  
### Instalação do Go no Windows

1. Abra o arquivo MSI que você baixou e siga as instruções para instalar o **Go**.
> Por padrão, o instalador instalará **Go** to Program Files ou Program Files (x86). Você pode alterar o local conforme necessário. Após a instalação, você precisará fechar e reabrir qualquer prompt de comando aberto para que as alterações no ambiente feitas pelo instalador sejam refletidas no prompt de comando.

2. Verifique se você instalou o **Go**.
   1. No Windows, clique no menu Iniciar.
   2. Na caixa de pesquisa do menu, digite cmd e pressione a tecla Enter.
   3. Na janela do prompt de comando que aparece, digite o seguinte comando: `$ go version`
   4. Confirme se o comando imprime a versão instalada do **Go**.

### Instalação do Go no Mac

1. Abra o arquivo do pacote que você baixou e siga as instruções para instalar o **Go**.
>O pacote instala a distribuição **Go** em `/usr/local/go`. O pacote deve colocar o diretório `/usr/local/go/bin` em sua variável de ambiente **PATH**. Pode ser necessário reiniciar qualquer sessão aberta do Terminal para que a alteração tenha efeito.
2. Verifique se você instalou o Go abrindo um prompt de comando e digitando o seguinte comando: `$ go version`.
3. Confirme se o comando imprime a versão instalada do **Go**.
> Exemplo de saida: `go version go1.17.5 darwin/amd64`   

### Instalação do Go no Linux

1. Extraia o arquivo que você baixou em `/usr/local`, criando uma árvore **Go** em `/usr/local/go`.   

:warning: Importante: Esta etapa removerá uma instalação anterior em `/usr/local/go`, se houver, antes da extração. Faça backup de todos os dados antes de continuar.   

Por exemplo, execute o seguinte como root ou por meio de sudo:   
`rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.6.linux-amd64.tar.gz`   

2. Adicione `/usr/local/go/bin` à variável de ambiente **PATH**.

Você pode fazer isso adicionando a seguinte linha ao seu `$HOME/.profile` ou `/etc/profile` ou `~/.zshrc` (para uma instalação em todo o sistema):   

`export PATH=$PATH:/usr/local/go/bin`

> :exclamation: Observação: as alterações feitas em um arquivo de perfil podem não se aplicar até a próxima vez que você efetuar login no computador. Para aplicar as alterações imediatamente, basta executar os comandos do shell diretamente ou executá-los a partir do perfil usando um comando como    
`source $HOME/.profile` ou `source /etc/profile` ou `source ~/.zshrc`.

3. Verifique se você instalou o Go abrindo um prompt de comando e digitando o seguinte comando: `$ go version`.
4. Confirme se o comando imprime a versão instalada do **Go**.

## Configurações

### Plugins útiis para o Visual Studio Code

* [Go Extension Pack](https://marketplace.visualstudio.com/items?itemName=doggy8088.go-extension-pack)
* [Code Runner](https://marketplace.visualstudio.com/items?itemName=formulahendry.code-runner)

### Outros plugins úteis para Visual Studio Code

* [Dracula Theme Official](https://marketplace.visualstudio.com/items?itemName=dracula-theme.theme-dracula)
* [Docker](https://marketplace.visualstudio.com/items?itemName=ms-azuretools.vscode-docker)
* [Git Extension Pack](https://marketplace.visualstudio.com/items?itemName=donjayamanne.git-extension-pack)
* [Markdown Preview Github Styling](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-preview-github-styles)
* [Markdown Checkboxes](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-checkbox)
* [Markdown Emoji](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-emoji)

## Links auxiliares

- [Configure seu ambiente de desenvolvimento](https://www.practical-go-lessons.com/chap-4-setup-your-dev-environment)

---

Ir para próximo: [Hello World Go](../02-hello-world/README.md)