# Products API

API de produtos construida utilizando Golang, Echo Framework, Docker e MySQL


### Executando a API

Usando o 'git clone' faça uma cópia do projeto e em seguida entre no diretório /products-api
O banco de dados MySQL foi configurado para executar no Docker, então após configurar o .env conforme o arquivo .env-example, execute:
```bash
$ docker compose up -d
```

Com o banco de dados inicializado, execute a aplicação:
```bash
$ go mod tidy
```
```bash
$ go run main.go
```
Com isso a API já estará pronta para uso e rodando localmente na porta definida no .env
Agora a API pode ser testada usando o [Postman](https://www.postman.com/) ou outra ferramenta de sua preferência
### As rotas implementadas foram:
```
- POST   /api/v1/product
- GET    /api/v1/product
- GET    /api/v1/product/{id}
- PUT    /api/v1/product/{id}
- DELETE /api/v1/product/{id}
```

## Sobre mim

Sou Henrique Caires, desenvolvedor de software. Estou a disposição para dúvidas, esclarecimentos e sugestões. Me encontre no linkedin: [Henrique Caires](https://www.linkedin.com/in/henrique-caires)
