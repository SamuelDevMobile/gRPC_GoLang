# Comunicação de Sistemas com gRPC em Golang

## Introdução
Este projeto consiste no entendimento dos conceitos de uma comunicação de sistemas com gRPC. Este framework criado pela Google, possibilita o cliente solicitar dados através de um contrato, onde os dados são transacionados em binário junto com a ultilização do Protocol Buffer e HTTP2. 

## Configuração do projeto
Ele pode ser baixado conforme abaixo:
```sh
$ git clone https://github.com/SamuelDevMobile/gRPC_GoLang.git
```
## AVISO!!!
> para conseguir executar este projeto, é preciso ter golang instalado ou ultilize um container com docker, sqlite3 e, do Evans instalado em sua maquina.
> crie uma tabela com o sqlite3.
```sh
create table categories (id string, name string, description string);
```

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

# Readme em construção
