package banco

import (
	"fmt"
	"os"
	"tarefa/modelos"
)

type ResultadoTarefa struct {
	Tarefa modelos.Tarefa
	Erro   error
}

func CriarTarefa(tarefa modelos.Tarefa, canal chan<- ResultadoTarefa) {
	conexao := PegarConexao()
	textoQuery := "insert into tarefa(nome, descricao, concluida) values(?, ? , ?)"
	_, erro := conexao.Exec(textoQuery, tarefa.Nome, tarefa.Descricao, tarefa.Concluida)
	if erro != nil {
		fmt.Println(erro)
		canal <- ResultadoTarefa{tarefa, erro}
	}
	textoQuery = "SELECT rowid from tarefa order by ROWID DESC limit 1"
	if erro := conexao.QueryRow(textoQuery).Scan(&tarefa.Id); erro != nil {

		fmt.Println(erro)
		canal <- ResultadoTarefa{tarefa, erro}
	}

	canal <- ResultadoTarefa{tarefa, nil}
}

func AtualizarTarefa(tarefa modelos.Tarefa) bool {
	conexao := PegarConexao()
	textoQuery := "update tarefa set nome = ?, descricao = ? , concluida = ? where id = ?"
	_, erro := conexao.Exec(textoQuery, tarefa.Nome, tarefa.Descricao, tarefa.Concluida, tarefa.Id)
	if erro != nil {
		fmt.Println(erro)
		return false
	}
	return true
}

func DeletarTarefa(id int) bool {
	conexao := PegarConexao()
	textoQuery := "delete from tarefa where id = ?"
	_, erro := conexao.Exec(textoQuery, id)
	if erro != nil {
		fmt.Println(erro)
		return false
	}
	return true
}

func PegarTodasAsTarefas() []modelos.Tarefa {
	conexao := PegarConexao()
	textoQuery := "select id, nome, descricao, concluida from tarefa"
	linhas, erro := conexao.Query(textoQuery)
	if erro != nil {
		fmt.Println(erro)
		os.Exit(1)
	}
	tarefasEncontradas := make([]modelos.Tarefa, 0)
	for linhas.Next() {
		var tarefaTemporaria modelos.Tarefa
		if erro := linhas.Scan(
			&tarefaTemporaria.Id,
			&tarefaTemporaria.Nome,
			&tarefaTemporaria.Descricao,
			&tarefaTemporaria.Concluida,
		); erro != nil {
			fmt.Println(erro)
			return []modelos.Tarefa{}
		}
		tarefasEncontradas = append(tarefasEncontradas, tarefaTemporaria)
	}
	return tarefasEncontradas
}
