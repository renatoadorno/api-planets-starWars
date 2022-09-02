# Bem vindo ao repositório do projeto API Planets StarWars

Minha primeira APIRest em Golang, foi desenvolvido utilizando o Golang e mongoDB, seguindo os princípios REST, é uma API simples, com funcionalidades como:
- Adicionar um planeta (com nome, clima e terreno)
- Buscar por nome
- Buscar por ID
- Remover planeta

## Técnologias usadas

<section>
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="go">
  <img src="https://img.shields.io/badge/MongoDB-4EA94B?style=for-the-badge&logo=mongodb&logoColor=white" alt="mongoDB">
</section>

## Como executar o projeto localmente?

Execute
```bash
docker-compose up -d
```
na raiz do projeto.

Após executar o docker-compose é só iniciar o server com
```bash
 go run main.go
```

## Rotas

Para testar as rotas da API utilize o [Postman](https://www.postman.com/), [Insomnia](https://insomnia.rest/) ou qualquer outra ferramenta semelhante.

A API por padrão ira rodar na **porta 6000**, também será necessario a Authorization com o **token:OBIWANKENOBI**

Para adicionar novos planetas utilize a rota: **POST   /planets/**

#### Body

```json
  {
    "name": "Yavin IV",
    "climate": "temperate, tropical",
    "terrain": "jungle, rainforests"
  }
```

Para fazer uma Busca por nome utilize a rota: **GET /planets/get/:name**
Para fazer uma Busca por id utilize a rota: **GET /planets/:id**
Para deletar um planeta utilize a rota: **DELETE /planets/delete/:name**
