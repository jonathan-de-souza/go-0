package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
}

func os10maioresEstadosDoBrasil() ([]Estado, error) {
	data, err := preencheArrayEstados()

	if err != nil {
		return nil, err
	}

	sort.Sort(sort.Reverse(byExtensao(data)))

	cont := len(data)
	var arrayTratado []Estado
	for cont > 10 {
		arrayTratado = data[:cont-1]
		cont--
	}

	return arrayTratado, nil
}

func preencheArrayEstados() ([]Estado, error) {
	file, err := os.Open("estados.txt")

	if err != nil {
		return nil, err
	}
	defer file.Close()

	var estados []Estado

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		linhaValores := strings.Split(scanner.Text(), ";")

		var e float64

		i, errParse := strconv.ParseFloat(linhaValores[1], 64)
		if errParse == nil {
			e = i
		}

		estados = append(estados, Estado{nome: strings.TrimSpace(linhaValores[0]), extensao: e})
	}
	return estados, scanner.Err()
}

type Estado struct {
	nome     string
	extensao float64
}

type byExtensao []Estado

func (s byExtensao) Len() int {
	return len(s)
}
func (s byExtensao) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byExtensao) Less(i, j int) bool {
	return s[i].extensao < s[j].extensao
}
