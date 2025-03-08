# Order Manager

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

Esse software foi desenvolvido visando realizar tarefas simples como gerar pedidos, atualizar esses pedidos, e realizar o cadastro e remoção dos clientes e produtos que serão utilizados nos pedidos.

<div id='tecnologias'/>

## Tecnologias

<div style="display: flex">

 <img align="center" alt="Django-Rest" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/djangorest/djangorest-original.svg" />

 <img align="center" alt="Python" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/django/django-plain.svg" />
 <img align="center" alt="Docker" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/docker/docker-original-wordmark.svg" />
 <img align="center" alt="PostgreSQL" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/postgresql/postgresql-original-wordmark.svg" />
 <img align="center" alt="HTML" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/html5/html5-original-wordmark.svg" />
 <img align="center" alt="CSS" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/css3/css3-original.svg" />
 <img align="center" alt="Javascript" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/javascript/javascript-original.svg" />
 <img align="center" alt="Actions" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/githubactions/githubactions-original.svg"/>
 <img align="center" alt="bootstrap" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/bootstrap/bootstrap-original.svg" />
 <img align="center" alt="nginx" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/nginx/nginx-original.svg" />
 <img align="center" alt="nginx" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/ssh/ssh-original-wordmark.svg" />
 <img align="center" alt="nginx" height="50" width="100" src="https://cdn.jsdelivr.net/gh/devicons/devicon@latest/icons/ubuntu/ubuntu-original.svg" />

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
git clone https://github.com/Gileno29/orderManager.git

# Navegue até o diretório do projeto
cd ordermanager
```
Será necessário criar dois arquivos de variáveis ".env": um na raiz do projeto, onde encontra-se o docker-compose, e outro dentro de ```app/ordermanager```.

O primeiro arquivo:

```sh
vim .env
```

```yml
POSTGRES_USER=userorder
POSTGRES_PASSWORD=orderpass
POSTGRES_DB=order

```

O segundo, navegue até o diretório ```app/ordermanager ``` e crie o arquivo ".env":

```yml
SECRET_KEY='django-insecure-nj=w6(caq-bzxcb_0xiw^7@s7oxvfn@k7f&)+bple_8qiw%@pr'
DB_USER='userorder'
DB_HOST='db'
DB_PASSS='orderpass'
DB_NAME='order'
```

retorne para a raiz do projeto e rode

```sh
docker-compose up --build

  OU 

docker-compose up -d  --build #rodar em backgroud
```
*Obs:* Verifique se já possui serviços funcionando em sua máquina nas portas da aplicação, caso haja, desative-os.

Seguindo a ordem corretamente, o sistema deve iniciar e estar acessível no endereço: http://localhost:8080/

## Utilização

O sistema consiste em um software que vai gerenciar pedidos. Antes de poder criar um pedido, será preciso realizar o cadastro de um cliente e o cadastro dos produtos que poderão ser vendidos pelo software.



### interface sistema

<img src="https://github.com/Gileno29/orderManager/blob/main/ordermanager/doc/imagens/Imagem_tela_inicial.png"/>

- Opção de cadatro do cliente:

Funcionalidade para cadastro simples de um cliente. (Nome, Telefone, descrição do endereço e CPF/CNPJ)
 <img src="https://github.com/Gileno29/orderManager/blob/main/ordermanager/doc/imagens/cadastrar_clientes.png"/>

 Formulário de Cadastro:
 
 <img src="https://github.com/Gileno29/orderManager/blob/main/ordermanager/doc/imagens/CadastroCliente.png"/>

- Opção de cadastro de produtos: 
 Funcionalidade para cadastro simples de um produto. (Nome, descrição, valor)
  <img src="https://github.com/Gileno29/orderManager/blob/main/ordermanager/doc/imagens/tela_cadstro_produtos.png"/>

 Formlário de produtos
  <img src="https://github.com/Gileno29/orderManager/blob/main/ordermanager/doc/imagens/formulario_de_cadastro_de_produtos.png"/>

 
- Opção de cadastro de usuário:

  <img src="https://github.com/Gileno29/orderManager/blob/main/ordermanager/doc/imagens/cadstro_de_usuarios.png"/>

  formulario para cadastro de usuários:

  <img src="https://github.com/Gileno29/orderManager/blob/main/ordermanager/doc/imagens/formulario_de_cadastro_de_usuarios.png"/>

- Opção de relatórios detalhados:
 Permite a extração de relatórios em CSV com filtros específicos.

<img src="https://github.com/Gileno29/orderManager/blob/main/ordermanager/doc/imagens/relatorios_detalhamento.png"/>

- Opção de relatórios de vendas:
  Permite ao usuário visualizar dados gerais sobre as vendas.

  <img src="https://github.com/Gileno29/orderManager/blob/main/ordermanager/doc/imagens/relatorio_vendas.png"/>


O sistema vai abrir na tela inicial, onde ficam os pedidos cadastrados nas últimas 24 horas.




<div id='tabela'/>

## Estrutura do Database

 Abaixo segue um resumo dos models atuais do sistema que geram o banco de dados.

```bash
  order=# \d
                         List of relations
 Schema |               Name                |   Type   |   Owner   
--------+-----------------------------------+----------+-----------
 public | auth_group                        | table    | userorder
 public | auth_group_id_seq                 | sequence | userorder
 public | auth_group_permissions            | table    | userorder
 public | auth_group_permissions_id_seq     | sequence | userorder
 public | auth_permission                   | table    | userorder
 public | auth_permission_id_seq            | sequence | userorder
 public | auth_user                         | table    | userorder
 public | auth_user_groups                  | table    | userorder
 public | auth_user_groups_id_seq           | sequence | userorder
 public | auth_user_id_seq                  | sequence | userorder
 public | auth_user_user_permissions        | table    | userorder
 public | auth_user_user_permissions_id_seq | sequence | userorder
 public | django_admin_log                  | table    | userorder
 public | django_admin_log_id_seq           | sequence | userorder
 public | django_content_type               | table    | userorder
 public | django_content_type_id_seq        | sequence | userorder
 public | django_migrations                 | table    | userorder
 public | django_migrations_id_seq          | sequence | userorder
 public | django_session                    | table    | userorder
 public | orders_cliente                    | table    | userorder
 public | orders_cliente_id_seq             | sequence | userorder
 public | orders_itempedido                 | table    | userorder
 public | orders_itempedido_id_seq          | sequence | userorder
 public | orders_pedido                     | table    | userorder
 public | orders_pedido_id_seq              | sequence | userorder
 public | orders_produto                    | table    | userorder
 public | orders_produto_id_seq             | sequence | userorde
```

Models do Sistema:

```py

from django.db import models
from django.db.models import Sum
from django.shortcuts import render, redirect
from decimal import Decimal
from .constants import STATUS_CHOICES
# Create your models here.

class Cliente(models.Model):

    nome = models.CharField(max_length=100, verbose_name="Nome do Cliente")
    email = models.EmailField(blank=True, null=True, verbose_name="E-mail")
    telefone = models.CharField(max_length=15, blank=True, null=True, verbose_name="Telefone")
    endereco = models.TextField(blank=True, null=True, verbose_name="Endereço")
    cpfcnpj = models.CharField(max_length=20, blank=False, null=False, verbose_name="CPF/CNPJ")

    def __str__(self):
        return self.nome

    class Meta:
        verbose_name = "Cliente"
        verbose_name_plural = "Clientes"




class Produto(models.Model):
    
    nome = models.CharField(max_length=100, verbose_name="Nome do Produto")
    descricao = models.TextField(blank=True, null=True, verbose_name="Descrição")
    preco = models.DecimalField(max_digits=10, decimal_places=2, verbose_name="Preço Unitário")

    def __str__(self):
        return f"{self.nome} (R$ {self.preco})"

    class Meta:
        verbose_name = "Produto"
        verbose_name_plural = "Produtos"



class ItemPedido(models.Model):
    produto = models.ForeignKey(Produto, on_delete=models.SET_NULL, null=True, blank=True, verbose_name="Produto")
    quantidade = models.PositiveIntegerField(default=1, verbose_name="Quantidade")
    pedido = models.ForeignKey('Pedido', on_delete=models.CASCADE, related_name='itens', verbose_name="Pedido")

    def __str__(self):
         return f"{self.quantidade} x {self.produto.nome if self.produto else 'Produto Removido'}"

    def preco_item(self):
        return Decimal(self.quantidade) * Decimal(self.produto.preco)

    class Meta:
        verbose_name = "Item do Pedido"
        verbose_name_plural = "Itens do Pedido"

class Pedido(models.Model):
    cliente = models.ForeignKey(Cliente, on_delete=models.CASCADE, verbose_name="Cliente")
    data_pedido = models.DateTimeField(auto_now_add=True, verbose_name="Data do Pedido")
    status = models.CharField(max_length=20, choices=STATUS_CHOICES, default='em_andamento',  null=True, blank=True,verbose_name="Status")
    valor_total = models.DecimalField(max_digits=10, decimal_places=2, default=0, verbose_name="Valor Total")

    def __str__(self):
        return f"Pedido #{self.id} - {self.cliente.nome}"

    def preco_total(self):
        total = self.itens.aggregate(total=Sum(models.F('quantidade') * models.F('produto__preco')))['total']
        return Decimal(total) if total is not None else Decimal(0)

    def save(self, *args, **kwargs):

        super().save(*args, **kwargs)

    def atualizar_valor_total(self):

        self.valor_total = self.preco_total()
        self.save(update_fields=['valor_total'])

```




<div id='estrutura'/>

## Estrutura do projeto

O projeto possui a seguinte estrutura:

```sh
  .
├── app
│   ├── create_user.py
│   ├── db.sqlite3
│   ├── manage.py
│   ├── ordermanager
│   │   ├── asgi.py
│   │   ├── __init__.py
│   │   ├── settings.py
│   │   ├── urls.py
│   │   └── wsgi.py
│   ├── orders
│   │   ├── admin.py
│   │   ├── apps.py
│   │   ├── constants.py
│   │   ├── forms.py
│   │   ├── __init__.py
│   │   ├── middleware.py
│   │   ├── models.py
│   │   ├── tests.py
│   │   ├── views_api.py
│   │   └── views.py
│   └── templates
│       ├── cad-clientes.html
│       ├── cad-produtos.html
│       ├── criar-usuario.html
│       ├── editar-cliente.html
│       ├── editar-pedido.html
│       ├── editar-produto.html
│       ├── editar-usuario.html
│       ├── list-clientes.html
│       ├── list-produtos.html
│       ├── list-usuarios.html
│       ├── logged_out.html
│       ├── login.html
│       ├── main.html
│       ├── model-footer.html
│       ├── model-header.html
│       ├── model-page.html
│       ├── relatorio-pedidos.html
│       ├── relatorios.html
│       └── restante.html
├── docker-compose.yml
├── Dockerfile
├── entrypoint.sh
├── nginx
│   ├── Dockerfile
│   └── nginx.conf
└── requirements.txt
```
O core do aplicativo encontra-se no diretório ```app```. Nesse diretório, podem ser encontradas as configurações principais do sistema dentro do diretório ````ordermanager```. E também o APP ```orders```, que possui os models do sistema e as views para acesso.

Externamente, podem ser encontrados o arquivo de docker-compose e os arquivos de configuração do proxy reverso do NGINX.

Ainda nesse nível, encontram-se os arquivos para deploy e configuração da infraestrutura da aplicação.

<div id='infraestrutura'/>

## Infraestrutura
A infraestrutura para deploy consiste em 3 partes:

- Aplicação: se trata do sistema em si, que é conteinerizado dentro de um container do Python
- Banco de dados: container à parte com o database do sistema.
- Proxy Reverso: container com o serviço do NGINX que será responsável por receber as requisições e encaminhá-las ao serviço.



### Docker file

```sh
FROM python:3.12-slim

WORKDIR /app
RUN pip install --upgrade pip

COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt


COPY ./app .
COPY ./entrypoint.sh /

ENTRYPOINT ["sh", "/entrypoint.sh"]



```
O Dockerfile consiste em uma imagem criada a partir da imagem ```python:3.8-slim```. Ele vai:

- Criar o workdir da aplicação.
- Enviar o arquivo de requirements e instalar os mesmos.
- Copiar os arquivos da aplicação e enviar para o container.
- Expor a porta da aplicação.
- Por último, vai chamar o Gunicorn para subir o serviço.

*OBS*: Caso seja alterado algo do código da aplicação, da forma que está, esse container precisará ser buildado novamente. Execute:
```sh
   docker-compose down -v
   docker-compose up --build```
```

### Docker compose file
```yml
  version: '3.8'

services:
  web:
    volumes:
      - static:/app/static
    build:
      context: .
    ports:
      - "8001:8000"
    depends_on:
      db:
        condition: service_healthy
    networks:
      - database
      - webnet

  db:
    image: postgres:13
    command: -c 'max_connections=5000'
    env_file: .env
    environment:
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
      POSTGRES_DB: ${POSTGRES_DB}
    volumes:
      - postgres_data:/var/lib/postgresql/data
    networks:
      - database
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U \"$POSTGRES_USER\" -d \"$POSTGRES_DB\""]
      interval: 5s
      timeout: 5s
      retries: 5
  nginx:
    build: ./nginx
    volumes:
      - static:/static
    ports:
      - "8080:80"
    depends_on:
      - web
    networks:
      - webnet

volumes:
  postgres_data:
  static:

networks:
  webnet:
  database:
  ```
O ```docker-compose vai``` definir 3 serviços em sua estrutura: ```web ``` (aplicação), ```db``` (banco de dados) e nginx (proxy).
O serviço ```web``` está tanto na rede do database quanto na do ```proxy```, devido à necessidade de comunicação com ambos os serviços, enquanto o ```proxy``` e o ```database``` encontram-se em suas respectivas redes apenas.  

### Proxy web:
 ```sh
  events {
      worker_connections 1024;
  }

 

upstream web {
	server web:8000;
}

	server {
		listen 80;

		location / {
			proxy_pass http://web;
			proxy_set_header Host $host;
			proxy_set_header X-Real-IP $remote_addr;
			proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
			proxy_set_header X-Forwarded-Proto $scheme;
			proxy_connect_timeout 3600s;
			proxy_send_timeout 3600s;
			proxy_read_timeout 3600s;
			send_timeout 3600s;
		}

		location /static/ {
			
			alias /static/;  #FUNCIONA
		}
		client_max_body_size 16M;
	}

```
O arquivo de configuração do NGINX define uma configuração de proxy simples. O timeout pode ser ajustado para menos, dependendo da situação.

## API
aqui também: Foi desenvolvido um  end point de uma API com unico propósito de recer uma string e retornar a primeria vogal única após uma consoante que é prescedida por uma vogal. A requisição é do tipo POST

Exemplo com curl:
  

```sh

## GIHUB ACTIONS
	curl -X POST http://localhost:8080/api/encontrar-vogal/ -H "Content-Type: application/json" -d '{"string": "outraStringDeTeste"}'

```



# API de Clientes

Esta API gerencia clientes de uma empresa fictícia.

## Endpoints

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
