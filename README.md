# Backend em GO

## Descrição
BackendContrat é um projeto em Go que gerencia usuários, incluindo candidatos e recrutadores, utilizando PostgreSQL como banco de dados. O projeto utiliza o framework Gin para a criação de APIs RESTful.

## Estrutura do Projeto
- `Application/Controllers`: Contém os controladores que lidam com as requisições HTTP.
- `Application/Routes`: Define as rotas da aplicação.
- `Data/Service`: Contém os serviços que interagem com o banco de dados.
- `Domain/Entities`: Define as entidades do domínio.
- `Domain/Factory`: Contém as fábricas para criar instâncias de usuários.
- `Utils`: Contém utilitários, como a geração de UUIDs.

## Pré-requisitos
- Go 1.16+
- Docker
- Docker Compose

## Configuração

### Clonar o Repositório
```sh
git clone https://github.com/seu-usuario/BackandContrat.git
cd BackandContrat
