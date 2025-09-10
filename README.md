# Sistema de Ponto Eletônico - SPE




## Directory Structure

```
bpp-backend/
├── controllers/        # Lógica de controle das rotas.
├── middlewares/        # Middlewares de autenticação e autorização.
├── models/             # Modelos de dados e lógica de negócios.
├── routes/             # Definição das rotas da API.
├── go.mod              # Módulo Go e dependências.
├── go.sum              # Checksum das dependências.
└── main.go             # Ponto de entrada da aplicação.
```



## Requirements

- Go 1.22 or upper.



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
    go mod tidy
    ```

4. Run the application:

    ```sh
    go run main.go
    ```



---

&copy; Departamento de Informática e Matemática Aplicada - DIMAp
