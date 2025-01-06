# Decode PDF Base64

Esta aplicação Go é responsável por gerar arquivos PDF a partir de dados codificados em Base64 armazenados em um banco de dados Oracle e disponibilizá-los via um endpoint HTTP.

## Funcionalidades

- Conecta-se a um banco de dados Oracle usando credenciais fornecidas em um arquivo `.env`.
- Consulta um registro específico no banco de dados que contém um PDF codificado em Base64.
- Decodifica o Base64 e salva o PDF em um arquivo.
- Disponibiliza o PDF gerado através de um endpoint HTTP.

## Requisitos

- Docker
- Docker Compose
- Go 1.16 ou superior

## Configuração

1. Clone este repositório:

   ```sh
   git clone https://github.com/seu-usuario/decode_pdf_base64.git
   cd decode_pdf_base64
   ```

2. Crie um arquivo [.env](http://_vscodecontentref_/1) na raiz do projeto com as seguintes variáveis de ambiente:

   ```env
   DB_USER=seu_usuario
   DB_PASSWORD=sua_senha
   CONNECT_STRING=seu_connect_string
   ```

3. Certifique-se de que o diretório [output](http://_vscodecontentref_/2) existe na raiz do projeto:

   ```sh
   mkdir -p output
   ```

## Uso

### Executando com Docker

1. Construa e inicie os serviços Docker:

   ```sh
   docker-compose up --build
   ```

2. Acesse o endpoint HTTP para visualizar o PDF gerado:

   ```sh
   http://localhost:8080/laudo/28543.pdf
   ```

### Executando localmente

1. Instale as dependências:

   ```sh
   go mod tidy
   ```

2. Execute a aplicação:

   ```sh
   go run main.go
   ```

3. Acesse o endpoint HTTP para visualizar o PDF gerado:

   ```sh
   http://localhost:8080/laudo/28543.pdf
   ```

## Estrutura do Projeto

- [main.go](http://_vscodecontentref_/3): Código principal da aplicação.
- [Dockerfile](http://_vscodecontentref_/4): Dockerfile para construir a imagem da aplicação.
- [docker-compose.yaml](http://_vscodecontentref_/5): Arquivo de configuração do Docker Compose.
- [output](http://_vscodecontentref_/6): Diretório onde os arquivos PDF gerados são armazenados.
- [.env](http://_vscodecontentref_/7): Arquivo de configuração das variáveis de ambiente (não incluído no repositório).

## Licença

Este projeto está licenciado sob os termos da licença MIT.
