# Comunicação de Sistemas com gRPC em Golang

## Introdução
Este projeto consiste no entendimento dos conceitos de uma comunicação de sistemas com gRPC. Este framework criado pela Google, possibilita o cliente solicitar dados através de um contrato, onde os dados são transacionados em binário junto com a ultilização do Protocol Buffer e HTTP2. 

## Configuração do projeto
Ele pode ser baixado conforme abaixo:
```sh
$ git clone https://github.com/SamuelDevMobile/gRPC_GoLang.git
```
## AVISO!!!
> para conseguir executar este projeto, é preciso ter golang instalado (ou ultilize um container com docker), protobuf, protoc, sqlite3, e o Evans.

## Protobuf - Linux
```sh
$ apt install -y protobuf-compiler
$ protoc --version  # Ensure compiler version is 3+
```

## Protobuf - macOs
```sh
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```

## Protoc
```sh
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

Crie uma tabela com o sqlite3. (para instala-lo é bem simples tem um tutorialzinho na internet :) )
```sh
create table categories (id string, name string, description string);
```

Para instalar o Evans.

## macOs
```sh
brew tap ktr0731/evans
brew install evans
```

## Go v1.16 ou posterior é necessário.
```sh
go install github.com/ktr0731/evans@latest
```

Abra-o no seu VSC (Visual Studio Code)
![Captura de Tela 2022-11-02 às 00 48 46](https://user-images.githubusercontent.com/26841238/199391008-6f05552d-ef78-4e67-8206-2d5e123e1950.png)

Comando para iniciar um projeto com GoLang.
```sh
go mod init NOME_DO_SEU_MODULO
```

Criando seu arquivo protofile.
<img width="670" alt="Captura de Tela 2022-11-15 às 00 45 58" src="https://user-images.githubusercontent.com/26841238/201821540-3b0d53df-43fc-40fa-85d0-230fcdcfb8ca.png">

Execute o seguinte comando para gerar código com protoc.
```sh
protoc --go_out=. --go-grpc_out=. proto/NOME_DO_ARQUIVO.proto
```

Em seguide rode o seguinte comando para baixar as dependências.
```sh
go mod tidy
```

Com o banco de dados criado anteriormente e o Evans instalado, rode o seguinte comando.
```sh
evans -r repl
```
<img width="429" alt="Captura de Tela 2022-11-15 às 00 55 56" src="https://user-images.githubusercontent.com/26841238/201822678-8ecc7ae3-45e6-42e0-a3cb-5bd700a7b2e4.png">

Com este comando você consegue executar qualquer comando para executar uma chamada.
```sh
call O_QUE_DESEJA_EXECUTAR
```
<img width="495" alt="Captura de Tela 2022-11-15 às 00 58 17" src="https://user-images.githubusercontent.com/26841238/201822989-7e4982d9-63ee-4d32-99ab-a3b0d0264f66.png">


# Readme em construção
