package telas

import (
	"fmt"
	"tarefa/banco"
	"tarefa/modelos"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var NomeTarefa binding.String
var DescricaoTarefa binding.String
var ConcluidaTarefa binding.Bool
var Concluida bool = false

func TelaAdicionarTarefa(janela fyne.Window, canal chan<- banco.ResultadoTarefa) {
	NomeTarefa.Set("")
	DescricaoTarefa.Set("")

	nome := widget.NewEntryWithData(NomeTarefa)
	desc := widget.NewEntryWithData(DescricaoTarefa)
	desc.MultiLine = true

	formulario := widget.NewForm()
	formulario.Append("Nome da tarefa:", nome)
	formulario.Append("Descrição da tarefa:", desc)

	coluna := container.NewVBox(
		formulario,
		widget.NewButton("Salvar", func() {
			nome, _ := NomeTarefa.Get()
			descricao, _ := DescricaoTarefa.Get()
			novaTarefa := modelos.Tarefa{
				Nome:      nome,
				Descricao: descricao,
				Concluida: false,
			}
			banco.CriarTarefa(novaTarefa, canal)
			App.SendNotification(fyne.NewNotification("Tarefa Criada", fmt.Sprintf("A tarefa '%s' foi criada.", novaTarefa.Nome)))
		}),
		widget.NewButton("Cancelar", func() {
			TelaTarefas(janela, true)
		}),
	)
	conteudoPrincipal := container.NewGridWithRows(
		3,
		layout.NewSpacer(),
		coluna,
		layout.NewSpacer(),
	)
	janela.SetContent(conteudoPrincipal)
}
