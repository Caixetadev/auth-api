# :sparkles: Auth API

API de autenticação com JWT Token desenvolvida em Golang utilizando o framework Echo. Com a utilização dessa API, é possível garantir a segurança e eficiência no processo de autenticação de usuários, permitindo que apenas usuários autenticados tenham acesso aos recursos protegidos.

## Instalação

Para executar o projeto, siga os passos abaixo:

1. Certifique-se de ter o Golang instalado em sua máquina. Caso não tenha, acesse o site oficial do Golang e siga as instruções de instalação.

2. Clone o repositório do projeto em seu ambiente de desenvolvimento utilizando o comando abaixo:

```
git clone https://github.com/Caixetadev/auth-api.git
```

3. No diretório raiz do projeto, execute o seguinte comando para instalar as dependências do projeto:

```
go get -d ./...
```

4. Crie o arquivo .env na raiz do projeto, e defina as seguintes variáveis de ambiente:

```
JWT_SECRET=seu_secreto_aqui
DB_USER=seu_usuario_do_banco_de_dados
DB_PASSWORD=sua_senha_do_banco_de_dados
DB_NAME=seu_nome_do_banco_de_dados
DB_HOST=seu_host_do_banco_de_dados
DB_PORT=sua_porta_do_banco_de_dados
```

5. Para executar a aplicação localmente, execute o seguinte comando no diretório raiz do projeto:

```
go run cmd/api/main.go
```
O servidor será iniciado, e poderá ser acessado através do endereço http://localhost:3333.

6. Para executar a aplicação utilizando Docker Compose, certifique-se de ter o Docker e o Docker Compose instalados em sua máquina. Em seguida, execute o seguinte comando no diretório raiz do projeto:

```
docker-compose up
```
O servidor será iniciado dentro de um container Docker, e poderá ser acessado através do endereço http://localhost:3333.

Agora que a aplicação está rodando, você pode fazer requisições para a API utilizando ferramentas como o Postman ou o cURL. A documentação da API está disponível em http://localhost:3333/swagger/index.html.
