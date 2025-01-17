package banco

import (
	"database/sql"
	"fmt" // joga isso fora quando for compilar para android
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	_ "github.com/mattn/go-sqlite3"
)

var conn *sql.DB

// deve ser chamado depois de ter criado o app
func CriarBanco(a fyne.App) {
	_, erro := os.Stat("./banco.db")
	if os.IsNotExist(erro) {
		f, _ := os.Create("./banco.db")
		f.Close()
		conn, erro = sql.Open("sqlite3", filepath.Join("./banco.db"))
		if erro != nil {
			fmt.Println(erro)
			os.Exit(1)
		}
	} else {
		conn, erro = sql.Open("sqlite3", "./banco.db")
		if erro != nil {
			fmt.Println(erro)
			os.Exit(1)
		}
	}
}

func PegarConexao() *sql.DB {
	return conn
}
