### Desafio Conductor de Seleção: RESTful API em Golang
O desafio consistiu no design e implementação de uma aplicação web (API RESTful) em golang.

Até a presente entrega do desafio (25/02/2021) as seguintes features já foram implementadas:

1. CICD completo com deploy em Heroku via container docker
  * CI com os estágios de Lint, Test, Swagger Generation e Trigger de CD
  * CD com Build local, Build da imagem docker e deploy do container em Heroku
2. Arquitetura base do webapp:
  * Roteamento e pipeline HTTP em Gorilla Mux
  * Middlewares de Authorização em header
  * Middleware de CORS
  * Middleware de logging
  * Middleware de recovery, com tratamento básico para status codes de erro como 404, 402, 500
  * Divisão em camadas (Control | Facade | DBContext)
  * Integração com um banco embarcado (SQLite)
3. Documentação da API via Swagger com a biblioteca go-swagger
4. CRUD Completo da entidade 'Account'

O que ficou faltando até a data de 25/02/2021:

1. CRUD da entidade 'Transactions'
2. CRUD da entidade 'Cards'
3. Endpoint de geração do PDF
4. Testes de unidade
5. Ajustes de segurança, como:
  * Paginação nas rotas de listagem
  * Autenticação via token de expiração
  * Tratamento + aprofundado de retornos de rotas (status codes + explicativos)
6. gRPC em k8s
7. Injeção de dependências e separação por Interface (I de SOLID)
7. Outros pormenores, como:
  * Extração de algumas constantes
  * Code-review
  * Separação de ambientes de Prod e Staging no CICD

### Sobre o andamento do desenvolvimento

Foi a primeira aplicação web que eu já fiz em go. Até então eu tinha feito somente duas lambdas simples, um software de investimento automatizado na bolsa que utilizada médias móveis (feito com gochannels) e um 'game of life' do Conway. 

#### SOLID:

I - Como era meu primeiro web-app em go, usando várias libs pela primeira vez, acredito que eu possa melhorar muito em alguns pontos SOLID como por exemplo Interface Segregation. Tentei aplicar Interface Segregation no tratamento de Erros, mas acredito que seja possível separar as camadas inteiramente por interfaces (Controller | Facade | DBContext), porém não soube exatamente como fazer e requeria + tempo e estudo. 

D - Acredito que eu tenha conseguido respeitar o D de Dependency Inversion na organização do meu pacote, como por exemplo ao entregar para os controllers a responsabilidade de se registrarem no roteador, e não o contrário, porém acredito que um framework de Injeção de Dependências facilitaria a manutenção dessa boa prática. 

L - Em termos de Liskov, não lidei muito com composição de objetos.

S - Respeitei o conceito de single responsability ao implementar os módulos pela ótica de sua utilização e não de suas funcionalidades, ou seja, um módulo deve ter uma responsabilidade de uso única, mesmo que tenha + de uma funcionalidade. Por exemplo os pacotes de Context, Controller e Facade lidam com + de uma entidade, porém são utilizados pela mesma razão.

O - Encapsulei algumas funções, estruturas e constantes para que não fossem expostas para modificação.

### Documentação 

O endereço base, bem como o token de autorização em produção será enviado por email para os avaliadores do desafio. Abaixo segue breve documentação de cada endpoint.

1. GET /accounts

Endpoint de listagem de accounts. 

Status codes esperados:
* 200: Listagem concluída.
* 204: Lista vazia.

Resposta esperada:
```
{
  "accounts": [
    {
    "created_at": "2021-02-25T05:13:07.319Z",
    "deleted_at": "2021-02-25T05:13:07.319Z",
    "id": 0,
    "status": "string",
    "updated_at": "2021-02-25T05:13:07.319Z"
    }
  ]
}
```

2. GET /accounts/{id}

Endpoint de recuperação de uma account por ID

Parâmetros de rota:
* id: ID do objeto account

Status codes esperados:
* 200: Objeto recuperado.
* 204: Objeto não existe.
* 400: Bad request (ID de tipo errado)

Resposta esperada:
```
{
  "account": {
   "created_at": "2021-02-25T06:23:17.081Z",
    "deleted_at": "2021-02-25T06:23:17.081Z",
    "id": 0,
    "status": "string",
    "updated_at": "2021-02-25T06:23:17.081Z"
  }
}
```

3. DELETE /accounts/{id}

Endpoint para deletar um objeto account por seu ID.

Parâmetros de rota:
* id: ID do objeto account

Status codes esperados:
* 200: Objeto deletado.
* 204: Objeto não existe.
* 400: Bad request (ID de tipo errado)

Resposta esperada:
```
{
  "account": {
   "created_at": "2021-02-25T06:23:17.081Z",
    "deleted_at": "2021-02-25T06:23:17.081Z",
    "id": 0,
    "status": "string",
    "updated_at": "2021-02-25T06:23:17.081Z"
  }
}
```

4. PUT /accounts/{id}

Endpoint para atualizar um objeto account por seu ID.

Parâmetros de rota:
* id: ID do objeto account
* status: Status novo do objeto account

Status codes esperados:
* 200: Objeto atualizado.
* 204: Objeto não existe.
* 400: Bad request (ID de tipo errado)

5. POST /accounts

Endpoint para adicionar um novo objeto account

Parâmetros de corpo:
* Objeto account
```
{
  "status": "string"
}
```

Status codes esperados:
* 200: Objeto adicionado.
* 422: Entidade não pode ser processada.

Resposta esperada:
```
{
  "account": {
   "created_at": "2021-02-25T06:23:17.081Z",
    "deleted_at": "2021-02-25T06:23:17.081Z",
    "id": 0,
    "status": "string",
    "updated_at": "2021-02-25T06:23:17.081Z"
  }
}
```

### URL's

Swagger: https://conductortask.herokuapp.com/swagger/

API: https://conductortask.herokuapp.com/api