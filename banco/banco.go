package banco

import (
	"database/sql"
	"fmt" // joga isso fora quando for compilar para android
	"os"

	"fyne.io/fyne/v2"
	_ "github.com/mattn/go-sqlite3"
)

var conn *sql.DB

// deve ser chamado depois de ter criado o app
func CriarBanco(a fyne.App) {
	arquivo, erro := a.Storage().Open("banco.db")

	if erro != nil {
		if arquivo != nil {
			arquivo.Close()
		}

		f, _ := a.Storage().Create("banco.db")
		defer f.Close()
		conn, erro = sql.Open("sqlite3", f.URI().String())
		if erro != nil {
			fmt.Println(erro)
			os.Exit(1)
		}
		conn.Exec(`
			create table usuario(
				id integer primary key,
				login varchar(50) not null,
				senha varchar(50) not null
			)
		`)
		conn.Exec(`
			insert into usuario(
				login,
				senha 
			)
			values(
				'admin',
				'admin'
			)
		`)
		conn.Exec(`
			create table tarefa(
				id integer primary key,
				nome varchar(50) not null,
				descricao varchar(150) not null,
				concluida boolean not null
			)
		`)
	} else {
		conn, erro = sql.Open("sqlite3", arquivo.URI().String())
		if erro != nil {
			fmt.Println(erro)
			os.Exit(1)
		}
		if arquivo != nil {
			arquivo.Close()
		}

	}
}

func PegarConexao() *sql.DB {
	return conn
}
