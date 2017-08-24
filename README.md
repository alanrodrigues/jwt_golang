### O que é? ###

Exemplo de autenticação em GO utilizando Json Web Token (JWT)

### Requisitos ###

* GO instalado na máquina
* MongoDB instalado na máquina
* Banco de Dados nomeado como jwt_golang_db

### Execução ###

* Instalar depedências GO: $ go get
* Build: $ go build
* Executar: $ jwt_golang

* Incluir um usuário na collection "users" > db.users.insert({"username":"user1", "password":"pass1"})

### Funcionamento ###

* Acesso à aplicação

/jwt-golang/auth POST 
application/json
{"username":"user1", "password":"pass1"}

/jwt-golang/secured/offers POST 
application/json
{"name":"offer 1"}

/jwt-golang/offers GET