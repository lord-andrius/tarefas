package telas

import (
	"fmt"
	"tarefa/banco"
	"tarefa/modelos"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var Tarefas []modelos.Tarefa

func desenhaTarefaPendente(janela fyne.Window, t *modelos.Tarefa) fyne.CanvasObject {
	return container.NewGridWithColumns(
		3,
		widget.NewLabel(t.Nome),
		widget.NewButton("Detalhe", func() {
			TelaDetalhes(janela, t, true)
		}),
		widget.NewButton("Concluir", func() {
			t.Concluida = true
			go banco.AtualizarTarefa(*t)
			TelaTarefas(janela, true)
		}),
	)
}

func desenhaTarefaConcluida(janela fyne.Window, t *modelos.Tarefa) fyne.CanvasObject {
	return container.NewGridWithColumns(
		3,
		widget.NewLabel(t.Nome),
		widget.NewButton("Detalhe", func() {
			TelaDetalhes(janela, t, false)
		}),
		widget.NewButton("Restaurar", func() {
			t.Concluida = false
			go banco.AtualizarTarefa(*t)
			TelaTarefas(janela, true)
		}),
	)
}

func TelaTarefas(janela fyne.Window, mostrarPendentes bool) {
	login, _ := Login.Get()
	adicinouNasPendentes := false
	var parteCentralPendentes *fyne.Container
	if len(Tarefas) != 0 {
		parteCentralPendentes = container.NewVBox()
		for i := range Tarefas {
			if !Tarefas[i].Concluida {
				adicinouNasPendentes = true
				parteCentralPendentes.Add(desenhaTarefaPendente(janela, &Tarefas[i]))
			}
		}
	}
	if !adicinouNasPendentes {
		parteCentralPendentes = container.NewCenter(widget.NewLabel("Não há tarefas pendentes."))
	}

	parteDeBaixoPendentes := container.NewVBox(widget.NewButton("Adicionar Tarefa", func() {
		canal := make(chan banco.ResultadoTarefa)
		go func() {
			resultado := <-canal
			if resultado.Erro == nil {
				Tarefas = append(Tarefas, resultado.Tarefa)
				TelaTarefas(janela, true)
			} else {
				fmt.Println(resultado.Erro)
			}
		}()
		TelaAdicionarTarefa(janela, canal)
	}))

	var parteCentralConcluidas *fyne.Container

	adicionouNasConcluidas := false
	if len(Tarefas) != 0 {
		parteCentralConcluidas = container.NewVBox()
		for i := range Tarefas {
			if Tarefas[i].Concluida {
				adicionouNasConcluidas = true
				parteCentralConcluidas.Add(desenhaTarefaConcluida(janela, &Tarefas[i]))
			}
		}
	}
	if !adicionouNasConcluidas {
		parteCentralConcluidas = container.NewCenter(widget.NewLabel("Não há tarefas concluídas"))
	}

	pendentes := container.NewBorder(
		widget.NewLabel(fmt.Sprintf("Usuário Logador: %s", login)),
		parteDeBaixoPendentes,
		nil,
		nil,
		container.NewVScroll(parteCentralPendentes),
	)

	concluidas := container.NewBorder(
		widget.NewLabel(fmt.Sprintf("Usuário Logador: %s", login)),
		nil,
		nil,
		nil,
		container.NewVScroll(parteCentralConcluidas),
	)

	conteudoPrincipal := container.NewAppTabs(
		container.NewTabItem("Pendentes", pendentes),
		container.NewTabItem("Concluidas", concluidas),
	)
	conteudoPrincipal.SetTabLocation(container.TabLocationBottom)
	if !mostrarPendentes {
		conteudoPrincipal.SelectIndex(1)

	}

	janela.SetContent(conteudoPrincipal)
}
