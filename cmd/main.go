package main

import (
	"database/sql"
	"fmt"
	"log"

	"codeberg.org/Pardalis/pardalis-api/cmd/api"
	"codeberg.org/Pardalis/pardalis-api/configs"
	"codeberg.org/Pardalis/pardalis-api/db"
)

// main 🐄 – El punto de entrada donde todo comienza y nada funciona como debería.
// Aquí inicializamos la base de datos (¡porque nadie quiere empezar sin una!), creamos el servidor API,
// y finalmente, tratamos de iniciar el servidor. Si algo sale mal, simplemente logueamos el error y nos vamos a casa. 🏡
func main() {

	// Intentamos crear una conexión a la base de datos SQLite. Si esto falla, es probable que tu vida de desarrollador también falle. 😱
	db, err := db.NewSQLiteStorage("./pardalis.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	// Inicializamos el almacenamiento. Porque, claro, necesitas hacer más que solo conectar, ¿verdad? 🤔
	initStorage(db)

	// Creamos el servidor API con la grandiosa idea de que alguien realmente lo usará. 🚀
	server := api.NewAPIServer(fmt.Sprintf(":%s", configs.Envs.Port), db)

	// Intentamos iniciar el servidor. Si no funciona, al menos hicimos todo lo posible, ¿no? 🤷‍♂️
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}

// initStorage 🐄 – Porque no puedes simplemente conectar a la base de datos y esperar que todo funcione mágicamente.
// Verificamos la conexión a la base de datos, y si algo sale mal, simplemente hacemos un *log.Fatal* y nos rendimos. 🎯
func initStorage(db *sql.DB) {
	// Verificamos la conexión a la base de datos. Si no podemos conectar, ¿para qué estamos aquí? 🤷‍♂️
	err := db.Ping()
	if err != nil {
		// Otro *log.Fatal*. ¡Qué sorpresa! No es como si esto ocurriera cada vez que probamos algo nuevo. 🙃
		log.Fatal(err)
	}

	// Si todo va bien, ¡enhorabuena! La conexión a la base de datos está establecida. 🎉
	log.Println("DB: Successfully connected!")
}
