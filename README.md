# go-fundamentals

Projeto de estudo dos fundamentos de GoLang, baseado no curso de **Imersão Técnica** - Go Fundamentals, promovido pela academia (<https://platform.rocks/>).

## Descrição

Este projeto é uma aplicação CLI (Command Line Interface) em Go que permite gerenciar configurações de servidores e bancos de dados através de arquivos YAML ou JSON. Utiliza a biblioteca [Cobra](https://github.com/spf13/cobra) para construção da interface de linha de comando.

## Requisitos

- Ter o Go instalado na máquina (<https://go.dev/doc/install>).
- Docker (opcional, para execução em container)

## Estrutura do Projeto

```text
├── cmd/                   # Comandos da CLI
│   ├── root.go            # Comando raiz
│   ├── parse.go           # Comando para parsing de configurações
│   └── server.go          # Comando para gerenciar servidores
├── config/                # Pacote de configuração
│   ├── config.go          # Lógica de carregamento de configurações
│   ├── types.go           # Tipos de dados das configurações
│   └── config_test.go     # Testes do pacote config
├── bin/                   # Diretório para binários compilados
├── example_config.yaml    # Exemplo de arquivo de configuração
├── Dockerfile             # Configuração do Docker
├── Makefile               # Comandos de build e execução
└── main.go                # Ponto de entrada da aplicação
```

## Instalação

### Compilar a aplicação

```bash
make build
```

### Executar diretamente com Go

```bash
go run main.go [comando] [flags]
```

## Comandos Disponíveis

### Comando `parse`

Lê e valida um arquivo de configuração em formato YAML ou JSON.

```bash
# Usando o binário compilado
./bin/configparser parse --file example_config.yaml

# Usando go run
go run main.go parse --file example_config.yaml

# Usando Makefile
make run
```

### Comando `server`

Gerencia configurações de servidores.

```bash
# Usando o binário compilado
./bin/configparser server --file example_config.yaml

# Usando go run
go run main.go server --file example_config.yaml
```

### Ajuda

Para ver todas as opções disponíveis:

```bash
go run main.go --help
go run main.go parse --help
go run main.go server --help
```

## Exemplo de Arquivo de Configuração

O arquivo `example_config.yaml` contém um exemplo de configuração:

```yaml
servers:
  - name: server1
    host: 127.0.0.1
    port: 8080
    replicas: 3
database:
  host: localhost
  port: 5432
  user: admin
  password: secret
```

## Comandos Make Disponíveis

| Comando             | Descrição                                    |
|---------------------|----------------------------------------------|
| `make build`        | Compila a aplicação                          |
| `make run`          | Executa a aplicação com o arquivo de exemplo |
| `make test`         | Executa os testes                            |
| `make clean`        | Remove arquivos compilados                   |
| `make docker-build` | Constrói a imagem Docker                     |
| `make docker-run`   | Executa a aplicação no Docker                |

## Execução com Docker

### Construir a imagem

```bash
make docker-build
```

### Executar no container

```bash
make docker-run
```

## Testes

Para executar os testes:

```bash
make test
```

## To Do

- [X] Criar um parser de arquivos de configurações do tipo `JSON` e `YAML` utilizando a lib [Cobra](https://github.com/spf13/cobra) para construção de um CLI.
- [X] Implementar validação de configurações
- [X] Adicionar suporte a Docker
- [X] Criar comandos separados para parse e server

## Licença

Este projeto é licenciado sob a licença MIT - veja o arquivo [LICENSE](LICENSE) para detalhes.
