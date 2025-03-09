# API Registro de Clientes

Esse projeto tem como objetivo desenvolver um sistema simples que permite gerenciar pedidos de forma genérica, além de disponibilizar uma API para identificar a vogal em uma string seguindo parâmetros específicos.

## Tabela de Conteúdos

- [Sobre](#sobre)
- [Tecnologias](#tecnologias)
- [Requisitos](#requisitos)
- [Rodando a Aplicação](#uso)
- [Estrutura do Banco de Dados](#tabela)
- [Estrutura do Projeto](#estrutura)
- [Infraestrutura](#infraestrutura)

<div id='sobre'/>

## Sobre

Essa API foi desenvolvida para  realizar o cadastros, atualização e remoção de dados de cadastros de clientes.

<div id='tecnologias'/>

## Tecnologias

<div style="display: flex">

 <img align="center" alt="Golang" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/go/go-original.svg" />

 <img align="center" alt="Docker" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original-wordmark.svg" />
 <img align="center" alt="PostgreSQL" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original-wordmark.svg" />

 <img align="center" alt="Swagger" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/swagger/swagger-original.svg" />


</div>

<div id='requisitos'/>

## Requisitos

<ul>
  <li>Git</li>
  <li>Deve possuir o <a href="https://docs.docker.com/engine/install/">Docker</a> e também o <a href="https://docs.docker.com/compose/install/">Docker-compose</a> instalados em sua máquina.</li>
</ul>

<div id='uso'/>

## Rodando a Aplicação

Instruções para iniciar a aplicação.

```sh
# Clone o repositório
git clone https://github.com/Gileno29/clientes-API.git

# Navegue até o diretório do projeto
cd cd clientes-API/
```
Será necessário criar um arquivo de variáveis ".env" na raiz do projeto


```sh
vim .env
```

```yml
ENVIRONMENT=production
PROD_POSTGRES_USER=usercad
PROD_POSTGRES_PASSWORD=cad
PROD_POSTGRES_DB=dbcadclientes
PROD_DATABASE_HOST=db
PRO_DATABASE_PORT=5432


```

```sh
docker-compose up --build

  OU 

docker-compose up -d  --build #rodar em backgroud
```
*Obs:* Verifique se já possui serviços funcionando em sua máquina nas portas da aplicação, caso haja, desative-os.

Seguindo a ordem corretamente, a API vai estar acessível no endpoint: http://localhost:8080/clientes


## Utilização API de Clientes

Esta API gerencia clientes de uma empresa fictícia.


## Endpoints

### Documentação SWAGGER

Opcionalmente a  documentção interativa contida nesse arquivo é possivel utilizar a documentação criada com o swagger.
no seguinte end point:
- **URL**: `localhost:8080/swagger/index.html`

<img src="https://github.com/Gileno29/clientes-API/blob/main/doc_img/image.png"/>




### Cadastrar Cliente
- **Método**: `POST`
- **URL**: `/clientes`
- **Descrição**: Cadastra um novo cliente.
- **Parâmetros**:
  - `documento` (string): CPF/CNPJ do cliente.
  - `nome` (string): Nome ou razão social do cliente.
  - `blocklist` (boolean): Status de blocklist.
- **Respostas**:
  - `201 Created`: Cliente cadastrado com sucesso.
  - `400 Bad Request`: Dados inválidos.
  - `409 Conflict`: Cliente já cadastrado.

### Listar Clientes
- **Método**: `GET`
- **URL**: `/clientes`
- **Descrição**: Retorna uma lista de clientes com suporte a paginação e filtro por nome/razão social.
- **Parâmetros**:
  - `razao_social` (string, opcional): Filtro por nome/razão social.
  - `page` (int, opcional): Número da página (padrão: 1).
  - `limit` (int, opcional): Número de itens por página (padrão: 10).
- **Respostas**:
  - `200 OK`: Lista de clientes.
  - `404 Not Found`: Nenhum cliente encontrado.

### Verificar Cliente
- **Método**: `GET`
- **URL**: `/clientes/{documento}`
- **Descrição**: Verifica se um cliente com o documento fornecido está cadastrado.
- **Parâmetros**:
  - `documento` (string): CPF/CNPJ do cliente.
- **Respostas**:
  - `200 OK`: Cliente encontrado.
  - `404 Not Found`: Cliente não encontrado.

### Atualizar Cliente
- **Método**: `PUT`
- **URL**: `/clientes/{documento}`
- **Descrição**: Atualiza a razão social e/ou o status de blocklist de um cliente.
- **Parâmetros**:
  - `documento` (string): CPF/CNPJ do cliente.
  - `razaosocial` (string, opcional): Nova razão social.
  - `blocklist` (boolean, opcional): Novo status de blocklist.
- **Respostas**:
  - `200 OK`: Cliente atualizado com sucesso.
  - `400 Bad Request`: Dados inválidos.
  - `404 Not Found`: Cliente não encontrado.
  - `500 Internal Server Error`: Erro ao atualizar cliente.

### Deletar Cliente
- **Método**: `DELETE`
- **URL**: `/clientes/{documento}`
- **Descrição**: Deleta um cliente com base no documento fornecido.
- **Parâmetros**:
  - `documento` (string): CPF/CNPJ do cliente.
- **Respostas**:
  - `200 OK`: Cliente deletado com sucesso.
  - `400 Bad Request`: Documento inválido.
  - `404 Not Found`: Cliente não encontrado.
  - `500 Internal Server Error`: Erro ao deletar cliente.

### Status do Servidor
- **Método**: `GET`
- **URL**: `/status`
- **Descrição**: Retorna informações sobre o tempo de atividade e o número de requisições atendidas.
- **Respostas**:
  - `200 OK`: Status do servidor.
