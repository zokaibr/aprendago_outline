package exercicios_ninja_nivel_1

import (
	"fmt"

	"github.com/fabianoflorentino/aprendago/pkg/format"
)

func ResolucaoNaPraticaExercicio1() {
	x := 42
	y := "James Bond"
	z := true

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	format.FormatResolucaoExercicios(resolucao)
}

func ResolucaoNaPraticaExercicio2() {
	var x int
	var y string
	var z bool

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	format.FormatResolucaoExercicios(resolucao)
}

func ResolucaoNaPraticaExercicio3() {
	x := 42
	y := "James Bond"
	z := true

	resolucao := fmt.Sprintf("%v %v %v", x, y, z)

	format.FormatResolucaoExercicios(resolucao)
}

func ResolucaoNaPraticaExercicio4() {
	type ninja int

	x := ninja(42)

	resolucao := fmt.Sprintf("%v", x)
	format.FormatResolucaoExercicios(resolucao)

	fmt.Printf("\n%T\n", x)

}

func ResolucaoNaPraticaExercicio5() {
	type ninja int

	var x ninja
	var y int

	x = 42
	y = int(x)

	resolucao := fmt.Sprintf("%T\n%T", x, y)

	format.FormatResolucaoExercicios(resolucao)
}

func RespondaAProva() {
	questionario := []format.Questionario{
		{Numero: "1.", Pergunta: "Qual o menor elemento em um programa que expressa uma ação a ser executada?", Opcoes: "[1] Uma declaração (Statement) [2] Uma expressão: "},
		{Numero: "2.", Pergunta: "A combinação de um ou mais valores, constantes, variáveis, operadores e funções que a linguagem interpreta e usa para produzir outro valor é?", Opcoes: "[1] Uma declaração (Statement) [2] Uma expressão: "},
		{Numero: "3.", Pergunta: "Quais são parênteses?", Opcoes: "[1] () [2] {} [3] []: "},
		{Numero: "4.", Pergunta: "Quais são colchetes?", Opcoes: "[1] () [2] {} [3] []: "},
		{Numero: "5.", Pergunta: "Quais são chaves?", Opcoes: "[1] () [2] {} [3] []: "},
		{Numero: "6.", Pergunta: "A abrangência de uma variável designa onde no código você pode acessar essa variável, e atribuir ou ler valores dela", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "7.", Pergunta: "Um tipo de dados primitivo é um tipo composto, criado a partir de outros tipos básicos que ja vem de fabrica na linguagem", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "8.", Pergunta: "O tipo int é um tipo de dado primitivo", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "9.", Pergunta: "O tipo string é um tipo de dado composto", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "10.", Pergunta: "Um tipo de dado composto permite que você crie estruturas formadas de outros tipos de dados", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "11.", Pergunta: "Quando declaramos uma variável com a palavra chave var e não atribuimos nenhum valor a esta variável, o compilador designa para esta um valor padrão, chamado de valor zero.", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "12.", Pergunta: "Palavras chaves servem a propositos especificos, mas fora isso podem ser usadas livremente ao longo do programa", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "13.", Pergunta: "Palavra chave e palavra reservada designam coisas diferentes", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "14.", Pergunta: "Uma palavra chave somente pode ser usada para seu proposito especifico e para declarar variaveis", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "15.", Pergunta: "Em 2 + 2, os numeros 2 são operadores", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "16.", Pergunta: "O termo package é uma palavra chave", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "17.", Pergunta: "O termo variable é uma palavra chave", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "18.", Pergunta: "O ponto de entrada para todos os programas é a função main(), que deve ficar dentro do package main", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "19.", Pergunta: "O operador curto de declaração pode ser usado ao invés de var em todas as situações", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "20.", Pergunta: "Quando vemos fmt.Println(), isto esta chamando a função Println() que pertence ao package fmt", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "21.", Pergunta: "Identificador é o nome atribuido a uma variável, função ou constante", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "22.", Pergunta: "Para utilizar uma função, variável ou constante de outro package utiliza-se o formato package-ponto-identificador. Por exemplo fmt.Println()", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "23.", Pergunta: "Qual é o caracter que permite jogar fora um valor? Ou seja, qual caracter permite que voce diga ao programa que não vai utilzar o valor retornado por uma função?", Opcoes: "[1] # [2] @ [3] _ [4] $ [5] isso é pegadinha...: "},
		{Numero: "24.", Pergunta: "Uma função cujo parâmetro é '... interface{}' é uma função variatica. Isso significa que você pode passar à função um número pré determinado de valores", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "25.", Pergunta: "Todo valor em Go pertence também ao tipo interface vazia, representado pela notação interface{}", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "26.", Pergunta: "2+3 é uma declaração/statement, não uma expressão", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "27.", Pergunta: "2+3 é uma expressão não uma declaração/statemente", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "28.", Pergunta: "Se eu quiser salvar o resultado de um format printing em uma variável, posso usar a função fmt.Sprintf()", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "29.", Pergunta: "Em Go podemos criar nossos proprios tipos?", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "30.", Pergunta: "Falando de tipos, em Go utilizamos o termo 'coerção' diferentemente de Java, por exemplo, onde se utiliza o termo 'conversão'", Opcoes: "[1] Verdade [2] Mentira: "},
		{Numero: "31.", Pergunta: "Todo tipo criado pelo programador tem um tipo subjacente", Opcoes: "[1] Verdade [2] Mentira: "},
	}

	gabarito := []format.Resposta{
		{Numero: "1.", Resposta: "1"}, {Numero: "2.", Resposta: "2"}, {Numero: "3.", Resposta: "1"}, {Numero: "4.", Resposta: "3"}, {Numero: "5.", Resposta: "2"},
		{Numero: "6.", Resposta: "1"}, {Numero: "7.", Resposta: "2"}, {Numero: "8.", Resposta: "1"}, {Numero: "9.", Resposta: "2"}, {Numero: "10.", Resposta: "1"},
		{Numero: "11.", Resposta: "1"}, {Numero: "12.", Resposta: "2"}, {Numero: "13.", Resposta: "1"}, {Numero: "14.", Resposta: "2"}, {Numero: "15.", Resposta: "2"},
		{Numero: "16.", Resposta: "1"}, {Numero: "17.", Resposta: "2"}, {Numero: "18.", Resposta: "1"}, {Numero: "19.", Resposta: "2"}, {Numero: "20.", Resposta: "1"},
		{Numero: "21.", Resposta: "1"}, {Numero: "22.", Resposta: "1"}, {Numero: "23.", Resposta: "3"}, {Numero: "24.", Resposta: "1"}, {Numero: "25.", Resposta: "1"},
		{Numero: "26.", Resposta: "2"}, {Numero: "27.", Resposta: "1"}, {Numero: "28.", Resposta: "1"}, {Numero: "29.", Resposta: "1"}, {Numero: "30.", Resposta: "1"}, {Numero: "31.", Resposta: "1"},
	}

	// Cria o questionário
	sq := new(format.ServicoQuestionario)
	sr := new(format.ServicoResposta)

	prova := sq.Cria(questionario)
	respostas, err := sr.Coleta(prova)
	if err != nil {
		fmt.Println("Erro ao coletar respostas:", err)
	}

	sr.Valida(respostas, gabarito)
}