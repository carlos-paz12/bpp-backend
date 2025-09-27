# Sistema de Ponto Eletônico - SPE




## Estrutura de diretórios

```
bpp-backend/
├── controllers/      # Implementação da lógica dos endpoints da API.
├── database/         # Configuração de conexão e inicialização do banco de dados.
├── docs/             # Documentação gerada automaticamente pelo Swagger.
├── middlewares/      # Middlewares de autenticação, autorização e outros.
├── migrations/       # Scripts de migração do banco de dados.
├── models/           # Estruturas de dados e validações.
├── repositories/     # Camada de acesso e manipulação de dados (DAO/repositórios).
├── routes/           # Registro e configuração das rotas da API.
├── seeders/          # Dados iniciais para popular o banco (seeders).
├── services/         # Regras de negócio e integração entre controllers e repositórios.
├── .env.example      # Exemplo de variáveis de ambiente.
├── .gitignore        # Arquivos e diretórios ignorados pelo Git.
├── go.mod            # Definição do módulo Go e dependências.
├── go.sum            # Checksums das dependências.
├── main.go           # Ponto de entrada da aplicação.
└── README.md         # Documentação do projeto.
```



## Requisitos

- Go 1.25 ou superior.



## Configuração do projeto

1. Clone este repositório:

    ```sh
    git clone https://github.com/carlos-paz12/bpp-backend.git ~/spe
    ```

2. Acesse o diretório do projeto:

    ```sh
    cd ~/spe
    ```

3. Configure as variáveis de ambiente:

    ```sh
    cp .env.example .env
    ```

    > Edite o arquivo .env com as credenciais do banco e demais configurações necessárias.

4. Instale as dependências:

    ```sh
    go mod download
    go mod tidy
    ```

4. Rode a aplicação:

    ```sh
    go run main.go
    ```



---

&copy; Departamento de Informática e Matemática Aplicada - DIMAp
