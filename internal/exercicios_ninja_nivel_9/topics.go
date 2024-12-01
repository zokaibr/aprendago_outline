package exercicios_ninja_nivel_9

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

const rootDir = "internal/exercicios_ninja_nivel_9"

func ExerciciosNinjaNivel9() {
	fmt.Printf("\n\nCapítulo 20: Exercícios Ninja Nível 9\n")

	executeSection("Na prática: Exercício #1")
	executeSection("Na prática: Exercício #2")
	executeSection("Na prática: Exercício #3")
	executeSection("Na prática: Exercício #4")
	executeSection("Na prática: Exercício #5")
	executeSection("Na prática: Exercício #6")
	executeSection("Na prática: Exercício #7")
}

func MenuExerciciosNinjaNivel9([]string) []format.MenuOptions {
	return []format.MenuOptions{
		{Options: "--exercicio=1 --nivel=9", ExecFunc: func() { executeSection("Na prática: Exercício #1") }},
		{Options: "--exercicio=1 --nivel=9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio1() }},
		{Options: "--exercicio=2 --nivel=9", ExecFunc: func() { executeSection("Na prática: Exercício #2") }},
		{Options: "--exercicio=2 --nivel=9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio2() }},
		{Options: "--exercicio=3 --nivel=9", ExecFunc: func() { executeSection("Na prática: Exercício #3") }},
		{Options: "--exercicio=3 --nivel=9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio3() }},
		{Options: "--exercicio=4 --nivel=9", ExecFunc: func() { executeSection("Na prática: Exercício #4") }},
		{Options: "--exercicio=4 --nivel=9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio4() }},
		{Options: "--exercicio=5 --nivel=9", ExecFunc: func() { executeSection("Na prática: Exercício #5") }},
		{Options: "--exercicio=5 --nivel=9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio5() }},
		{Options: "--exercicio=6 --nivel=9", ExecFunc: func() { executeSection("Na prática: Exercício #6") }},
		{Options: "--exercicio=6 --nivel=9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio6() }},
		{Options: "--exercicio=7 --nivel=9", ExecFunc: func() { executeSection("Na prática: Exercício #7") }},
		{Options: "--exercicio=7 --nivel=9 --resolucao", ExecFunc: func() { ResolucaoNaPraticaExercicio7() }},
	}
}

func HelpMeExerciciosNinjaNivel9() {
	hlp := []format.HelpMe{
		{Flag: "--exercicio=1 --nivel=9", Description: "Exibe o Exercício 1 do capítulo 20"},
		{Flag: "--exercicio=1 --nivel=9 --resolucao", Description: "Exibe a resolução do Exercício 1 do capítulo 20"},
		{Flag: "--exercicio=2 --nivel=9", Description: "Exibe o Exercício 2 do capítulo 20"},
		{Flag: "--exercicio=2 --nivel=9 --resolucao", Description: "Exibe a resolução do Exercício 2 do capítulo 20"},
		{Flag: "--exercicio=3 --nivel=9", Description: "Exibe o Exercício 3 do capítulo 20"},
		{Flag: "--exercicio=3 --nivel=9 --resolucao", Description: "Exibe a resolução do Exercício 3 do capítulo 20"},
		{Flag: "--exercicio=4 --nivel=9", Description: "Exibe o Exercício 4 do capítulo 20"},
		{Flag: "--exercicio=4 --nivel=9 --resolucao", Description: "Exibe a resolução do Exercício 4 do capítulo 20"},
		{Flag: "--exercicio=5 --nivel=9", Description: "Exibe o Exercício 5 do capítulo 20"},
		{Flag: "--exercicio=5 --nivel=9 --resolucao", Description: "Exibe a resolução do Exercício 5 do capítulo 20"},
		{Flag: "--exercicio=6 --nivel=9", Description: "Exibe o Exercício 6 do capítulo 20"},
		{Flag: "--exercicio=6 --nivel=9 --resolucao", Description: "Exibe a resolução do Exercício 6 do capítulo 20"},
		{Flag: "--exercicio=7 --nivel=9", Description: "Exibe o Exercício 7 do capítulo 20"},
		{Flag: "--exercicio=7 --nivel=9 --resolucao", Description: "Exibe a resolução do Exercício 7 do capítulo 20"},
	}

	fmt.Printf("\nCapítulo 20: Exercícios Ninja Nível 9\n")
	format.PrintHelpMe(hlp)
}

func executeSection(section string) {
	format.FormatSection(rootDir, section)
}
