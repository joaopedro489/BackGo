# BackGo
## Trabalho Final PSI Tech Lead - realizado em GO
### Feito por: João Pedro Cavalcante
### Utilizando a framework [Gin](https://github.com/gin-gonic/gin) e [Gorm](https://github.com/go-gorm/gorm)

## Como Instalar e Executar o Projeto?
* Clone o repositório
* Antes crie um banco de dados no mysql, pode ser feito pelo phpmyadmin, com o nome "demo", se desejar outro nome se
dirija ao arquivo BackGo/Models/setup.go e na linha 9 troque o nome de "demo" para o desejado, importante lembrar
não pode conter senha o banco de dados, caso deseja na mesma linha 9 mude "root@" para "userDb@password"
* Após isso na pasta raiz do repositorio, com o Go já instalado `go run main.go`, este comando irá baixar as
dependencias do arquivo e já executar

## Como Testar?
#### Para o Teste recomenda-se utilizar o Postman

Após a aplicação já estar servindo com o comando `go run main.go` será possível testar as seguintes rotas:

* "localhost:8080/getAllUsers" - **GET**
* "localhost:8080/getUser/{{id}}" - **GET**
* "localhost:8080/getUserPosts/{{id}}" - **GET**
* "localhost:8080/userPost/{{id}}" - **POST**
* "localhost:8080/createUser" - **POST**
* "localhost:8080/updateUser/{{id}}" - **PUT**
* "localhost:8080/userUpdatePost/{{post_id}}/{{user_id}}" - **PUT**
* "localhost:8080/deleteUser/{{id}}" - **DELETE**
* "localhost:8080/userDeletePost/{{post_id}}/{{user_id}}" - **DELETE**
* "localhost:8080/getAllPosts" - **GET**
* "localhost:8080/getPost/{{id}}" - **GET**
* "localhost:8080/createPost" - **POST**
* "localhost:8080/updatePost/{{id}}" - **PUT**
* "localhost:8080/deletePost/{{id}} - **DELETE**

#### Quase todas rotas podem ser utilizadas normalmente pelo Postman, os dados que serão mandadados para rota PUT precisarão ser mandados por query, isto é usando o Params para mandar os dados pela URL. E o Post precisa ser mandado por Form-data
#### Não tem restrições para criar uma entidade, isto é se quiser poderá só mandar o nome para criação de um user

Atributos para serem considerados para o teste, usar em rotas de Post e Put:

Atributos de User
* name
* email
* age
* password

Atributos de Post
* title
* text
* user_id (necessário se usar a rota cretePost)
