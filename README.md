# go-clean-arch todo list api
Esse é um projeto feito para aprender Clean Arch com Go
## Como executar
    1. **Clonar o repositorio**:
        ```bash
        git clone https://github.com/your-username/go-todolist-api.git
        cd go-todolist-api
        ```

    2. **Instalar dependencias**:
        ```bash
        go mod tidy
        ``` 

    3. **Executar o projeto**:
        ```bash
        go run cmd/main.go
        ```
## Estrutura do Projeto

    - **Entities**: As Entities encapsulam as regras essenciais do dominio.
    - **Use Cases**: Implementa a lógica de aplicação necessária para a regras de negocio.
    - **Interfaces**: Camada externa que manipula o input/output (HTTP,database).
    - **Repositories**: Intermede a comunicação com a infraestrutura de dados.