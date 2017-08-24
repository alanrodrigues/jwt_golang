### O que �? ###

Exemplo de autentica��o em GO utilizando Json Web Token (JWT)

### Requisitos ###

* GO instalado na m�quina
* MongoDB instalado na m�quina
* Banco de Dados nomeado como jwt_golang_db

### Execu��o ###

* Instalar deped�ncias GO: $ go get
* Build: $ go build
* Executar: $ jwt_golang

* Incluir um usu�rio na collection "users" > db.users.insert({"username":"user1", "password":"pass1"})

### Funcionamento ###

* Acesso � aplica��o

/jwt-golang/auth POST 
application/json
{"username":"user1", "password":"pass1"}

/jwt-golang/secured/offers POST 
application/json
{"name":"offer 1"}

/jwt-golang/offers GET