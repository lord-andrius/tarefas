package banco

import (
	"fmt"
	modelos "tarefa/modelos"
)

func PegarUsuarioPeloLogin(login string) (modelos.Usuario, bool) {

	conexao := PegarConexao()
	textQuery := "select id, login, senha from usuario where login = ?"
	var usuarioEncontrado modelos.Usuario
	if erro := conexao.QueryRow(textQuery, login).Scan(
		&usuarioEncontrado.Id,
		&usuarioEncontrado.Login,
		&usuarioEncontrado.Senha,
	); erro != nil {
		fmt.Println(erro)
		return usuarioEncontrado, false
	}
	fmt.Println(login)
	return usuarioEncontrado, true
}
