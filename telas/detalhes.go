package telas

import (
	"tarefa/banco"
	"tarefa/modelos"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func TelaDetalhes(janela fyne.Window, tarefa *modelos.Tarefa, voltarParaPendentes bool) {
	NomeTarefa.Set(tarefa.Nome)
	DescricaoTarefa.Set(tarefa.Descricao)

	nome := widget.NewEntryWithData(NomeTarefa)
	desc := widget.NewEntryWithData(DescricaoTarefa)
	desc.MultiLine = true

	estado := "Pendente"
	if tarefa.Concluida {
		estado = "Concluída"
	}

	formulario := widget.NewForm()
	formulario.Append("Nome da tarefa:", nome)
	formulario.Append("Descrição da tarefa:", desc)
	formulario.Append("Estado: ", widget.NewLabel(estado))

	coluna := container.NewVBox(
		formulario,
		widget.NewButton("Salvar", func() {
			nome, _ := NomeTarefa.Get()
			descricao, _ := DescricaoTarefa.Get()
			tarefa.Nome = nome
			tarefa.Descricao = descricao
			go banco.AtualizarTarefa(*tarefa)
			TelaTarefas(janela, voltarParaPendentes)
		}),
		widget.NewButton("Voltar", func() {
			TelaTarefas(janela, voltarParaPendentes)
		}),
		widget.NewButton("Excluir", func() {
			go banco.DeletarTarefa(tarefa.Id)
			indice := 0
			for i := range Tarefas {
				if Tarefas[i].Id == tarefa.Id {
					indice = i
				}
			}

			Tarefas = append(Tarefas[:indice], Tarefas[indice+1:]...)
			TelaTarefas(janela, voltarParaPendentes)
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
