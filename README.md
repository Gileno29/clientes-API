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