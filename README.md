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

* /jwt-golang/auth POST <br />
Header<br />
 Content-Type: application/json<br />
Body: {"username":"user1", "password":"pass1"}<br />

* /jwt-golang/secured/offers POST<br />
Header<br />
 Authorization: token from authentication endpoint<br />
 Content-Type: application/json<br />
Body: {"name":"offer 1"}<br />

* /jwt-golang/offers GET
