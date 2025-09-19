# Sistema de Ponto Eletônico - SPE




## Directory Structure

```
bpp-backend/
├── controllers/        # Funções que implementam a lógica dos endpoints da API.
├── docs/               # Arquivos de documentação gerados pelo Swagger.
├── middlewares/        # Middlewares, como autenticação e autorização.
├── models/             # Estruturas de dados e validações.
├── repository/         # Acesso e manipulação de dados (repositórios/banco de dados).
├── routes/             # Configuração e registro das rotas da API.
├── services/           # Regras de negócio e integração entre controllers e repositórios.
├── go.mod              # Definição do módulo Go e dependências.
├── go.sum              # Checksum das dependências para garantir integridade.
├── main.go             # Ponto de entrada da aplicação.
└── README.md           # Documentação do projeto e instruções de uso.
```



## Requirements

- Go 1.25 or upper.



## Project Setup

1. Clone this repository:

    ```sh
    git clone https://github.com/carlos-paz12/bpp-backend.git ~/spe
    ```

2. Navigate to the project dir?

    ```sh
    cd ~/spe
    ```

3. Install the dependencies:

    ```sh
    go mod download
    go mod tidy
    ```

4. Run the application:

    ```sh
    go run main.go
    ```



---

&copy; Departamento de Informática e Matemática Aplicada - DIMAp
