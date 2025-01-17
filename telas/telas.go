package telas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
)

var App fyne.App

func InicializaTelas(app fyne.App) {
	App = app
}
func init() {
	Login = binding.NewString()
	Senha = binding.NewString()
}
