# go-clean-arch todo list api
    Esse é um projeto feito para aprender Clean Arch com Go
## Como executar
    1. **Clone the repository**:
        ```bash
        git clone https://github.com/your-username/go-todolist-api.git
        cd go-todolist-api
        ```

    2. **Install dependencies**:
        ```bash
        go mod tidy
        ``` 

    3. **Run the project**:
        ```bash
        go run cmd/main.go
        ```
## Project Structure

    - **Entities**: As Entities encapsulam as regras essenciais do dominio.
    - **Use Cases**: Implementa a lógica de aplicação necessária para a regras de negocio.
    - **Interfaces**: Camada externa que manipula o input/output (HTTP,database).
    - **Repositories**: Intermede a comunicação com a infraestrutura de dados.