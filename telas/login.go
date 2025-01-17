package telas

import (
	"tarefa/banco"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// inicializados no telas.go
var Login binding.String
var Senha binding.String

const (
	LoginFoiAceito = iota
	LoginFoiRecusado
	LoginNaoFoiFeito
)

var estadoLogin = LoginNaoFoiFeito

func TelaLogin(janela fyne.Window) {
	entradaLogin := widget.NewEntryWithData(Login)
	entradaLogin.SetPlaceHolder("login")

	entradaSenha := widget.NewEntryWithData(Senha)
	entradaSenha.Password = true
	entradaSenha.SetPlaceHolder("senha")

	formularioLogin := widget.NewForm()
	formularioLogin.Append("Login:", entradaLogin)
	formularioLogin.Append("Senha:", entradaSenha)

	containerLogin := container.New(layout.NewVBoxLayout(), formularioLogin)

	conteudoPrincipal := container.New(layout.NewGridLayoutWithRows(3), layout.NewSpacer(), containerLogin, layout.NewSpacer())

	if estadoLogin == LoginFoiRecusado {
		containerLogin.Add(widget.NewLabel("Login ou Senha incorretos"))
	}

	containerLogin.Add(widget.NewButton("Logar", func() {
		textoLogin, _ := Login.Get()
		textoSenha, _ := Senha.Get()
		logar(textoLogin, textoSenha)
		if estadoLogin == LoginFoiAceito {
			TelaTarefas(janela, true)
		} else {
			TelaLogin(janela)
		}
	}))

	containerLogin.Add(widget.NewLabel(""))
	janela.SetContent(conteudoPrincipal)
}

func logar(login, senha string) {
	u, achou := banco.PegarUsuarioPeloLogin(login)

	if !achou {
		estadoLogin = LoginFoiRecusado
	}

	if u.Senha == senha {
		estadoLogin = LoginFoiAceito
	}

}
