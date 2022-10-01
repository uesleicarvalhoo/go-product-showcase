# Go Product Showcase

<p align="center">
<img src="https://img.shields.io/static/v1?label=Version&message=0.0.0&color=7159c1&plastic"/>
</p>

## Tabela de conteúdos

- [Product Showcase](#go-product-showcase)
- [Tabela de conteúdos](#tabela-de-conteúdos)
- [Pré-requisitos](#pré-requisitos)
- [🎲 Rodando a aplicação](#-rodando-a-aplicação)
- [🎲 Contribuindo com o projeto](#-contribuindo-com-o-projeto)
- [Testes](#testes)
- [Release](#release)

## Instalação

### Pré-requisitos

Antes de começar, você vai precisar ter instalado em sua máquina as seguintes ferramentas:

### Docker

O [Docker](https://www.docker.com/) é usado para executar os containers na sua máquina, você pode usar uma ferramenta similar

### GNU Make

O [Makefile](https://www.gnu.org/software/make/) contem alguns scripts para iniciar a aplicação, suas dependencias, executar os testes e etc, você pode instalar com o linux com `sudo apt install make`, no Mac `brew install make` e no windows pode usar o [chocollatey](https://chocolatey.org/) com `choco install make`

Para saber quais sãos os comandos disponíveis, execute `make help` no terminal, você verá algo como:

``` bash
Usage:
  make [target]

Targets:
help        Display this help
run         Run app
swagger     Generate Swagger content
compose     Init containers with dev dependencies
release     Create a new release
test        Run tests of project
coverage    Run tests, make report and open into browser
clean       Remove Cache files
```

## 🎲 Rodando a aplicação

```bash
# Clone este repositório
$ git clone <https://github.com/uesleicarvalhoo/go-product-showcase>

# Acesse a pasta do projeto no terminal/cmd
$ cd go-product-showcase

# Você pode facilmente iniciar as dependencias de desenvolvimento com o comando
$ make compose
# Isso vai iniciar alguns containers:
# PostgreSQL    localhost:5432  -> Banco de dados da aplicação
# RabbitMQ      localhost:15672 -> Messageria
# Jaeger        localhost:16686 -> Monitoramento padrão Opentelemetry

# E para iniciar a aplicação em si, é só executar:
$ make run
# O servidor inciará na porta:8000 acesse <http://localhost:8000/>
# O servidor já conta com documentação integrada, disponível no path /swagger/index.html
```

## Contribuindo com o projeto

```bash
# Clone este repositório
$ git clone <https://github.com/uesleicarvalhoo/go-product-showcase>

# Acesse a pasta do projeto no terminal/cmd
$ cd go-product-showcase

# A aplicação está configurada para rodar com o PostgreSQL, você pode subir uma instancia com o Docker com o comando
$ docker-compose up database -d

# Faça suas alterações

# Formate o código
$ make format

# Garanta que os testes estão passando
$ make test

# Atualize o CHANGLOG.md

# Abra uma pull request e ela será analisada
```

## Liberando uma nova versão

Você pode liberar uma nova versão do serviço realizando uma nova release, para criar e publicar uma nova tag use o comando

```bash $ make release```

## Testes

A aplicação possui testes automatizados, para roda-los é bem simples, apenas execute o comando

```bash
# Executa os testes
$ make test
```

E caso queira um reporte dos testes, você pode rodar o comando

```bash
# Gera o report sobre os testes
$ make coverage
```
