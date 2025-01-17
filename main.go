package main

import (
	"log"
	"os"
	"tarefa/banco"
	"tarefa/telas"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const LARGURA float32 = 360
const ALTURA float32 = 740

func main() {
	a := app.NewWithID("com.tarefas")
	if err := os.Chdir(a.Storage().RootURI().Path()); err != nil {
		log.Fatal(err)
	}
	banco.CriarBanco(a)
	telas.InicializaTelas(a)
	janela := a.NewWindow("Tarefas")
	janela.Resize(fyne.NewSize(LARGURA, ALTURA))

	telas.TelaLogin(janela)

	janela.ShowAndRun()
}
