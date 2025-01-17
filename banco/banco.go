package banco

import (
	"database/sql"
	"fmt" // joga isso fora quando for compilar para android
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/storage"
	_ "github.com/mattn/go-sqlite3"
)

var conn *sql.DB

// deve ser chamado depois de ter criado o app
func CriarBanco(a fyne.App) {
	arquivos := a.Storage()

	uriArquivoDB := storage.NewFileURI(filepath.Join(arquivos.RootURI().Path(), "banco.db"))
	existe, erro := storage.Exists(uriArquivoDB)
	if erro != nil {
		fmt.Println(erro)
		os.Exit(1)
	}

	if existe == false {
		fmt.Println("criando em:", filepath.Join(arquivos.RootURI().Path(), "banco.db"))
		db, erro := sql.Open("sqlite3", filepath.Join(arquivos.RootURI().Path(), "banco.db"))
		if erro != nil {
			fmt.Println(erro)
			os.Exit(1)
		}
		defer db.Close()
	} else {
		fmt.Println("criado em:", uriArquivoDB.String())
	}
}

func PegarConexao() *sql.DB {
	return conn
}
