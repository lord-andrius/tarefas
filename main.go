package main

import "fyne.io/fyne/v2/app"
import "fyne.io/fyne/v2"
import "tarefa/telas"
import "tarefa/banco"
const LARGURA float32 = 360
const ALTURA  float32 = 740

func main() {
	a := app.NewWithID("com.tarefas")
	banco.CriarBanco(a)
	janela := a.NewWindow("Tarefas")
	janela.Resize(fyne.NewSize(LARGURA, ALTURA))
	
	telas.TelaLogin(janela)	

	janela.ShowAndRun()
}
