package main

import (
	"tarefa/banco"
	"tarefa/telas"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

const LARGURA float32 = 360
const ALTURA float32 = 740

func main() {
	a := app.NewWithID("com.tarefas")
	telas.InicializaTelas(a)
	banco.CriarBanco(a)
	janela := a.NewWindow("Tarefas")
	janela.Resize(fyne.NewSize(LARGURA, ALTURA))

	telas.TelaLogin(janela)

	janela.ShowAndRun()
}
