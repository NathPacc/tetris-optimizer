package main

import (
	"fmt"
	"os"
	"strings"
	tetris "tetris-optimization/functions"
)

func main() {

	// Vérification des arguments entrés

	if len(os.Args) < 2 {
		fmt.Println("Error : Need a file")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Error : Too many files, only need one")
		return
	}
	datas, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error : Unable to read file")
		return
	}

	// Vérification du contenu du fichier

	if len(tetris.GetTetros(string(datas))) == 0 {
		fmt.Println("Error : The file does not contain a single tetromino")
		return
	}
	if !checkFormat(string(datas)) {
		fmt.Println("Error : The file is not correctly formated")
		return
	}
	hasError, badtetros := checkErrors(string(datas))
	if hasError {
		fmt.Println("Error : The following tetrominos are incorrects :")
		printTetros(badtetros)
		return
	}

	// Résolution

	printSquare(tetris.Resolve(string(datas)))
}

func checkFormat(datas string) bool {
	lines := strings.Split(datas, "\r\n")
	i := 0
	for i < len(lines) {

		// Si la ligne est vide, on vérifie la ligne suivante

		if lines[i] == "" {
			i++
			continue
		}

		// On vérifie que l'on a pas un bloc incomplet entre notre ligne et la fin du document

		if i+4 > len(lines) {
			return false
		}

		// On vérifie bloc par bloc

		for j := 0; j < 4; j++ {
			line := lines[i+j]

			// On vérifie que chaque ligne est de la bonne longueur

			if len(line) != 4 {
				return false
			}

			// On vérifie qu'il n'y que les caracters attendus

			for _, c := range line {
				if c != '.' && c != '#' {
					return false
				}
			}
		}

		// On passe au bloc suivant

		i += 4

		// Si l'on n'est pas arrivé à la fin du tableau et que l'on ne retombe pas sur une ligne vide à l'étape suivante, alors un des blocs ne fait pas 4 lignes

		if i < len(lines) && lines[i] != "" {
			return false
		}
	}

	return true
}

func checkErrors(datas string) (bool, []tetris.Tetromino) {
	tetros := tetris.GetTetros(datas)
	hasError := false
	badtetros := []tetris.Tetromino{}
	for _, tetro := range tetros {

		// On vérifie bien qu'il y a quatres points

		if len(tetro.Coord) != 4 {
			hasError = true
			badtetros = append(badtetros, tetro)
			continue
		}

		// On vérifie bien qu'il n'y a pas de points détaché des autres.

		if !tetris.IsConnected(tetro) {
			hasError = true
			badtetros = append(badtetros, tetro)
		}
	}
	return hasError, badtetros
}

func printSquare(grid [][]rune) {
	for i := range grid {
		for j := range grid[i] {
			fmt.Print(string(grid[i][j]))
		}
		fmt.Print("\n")
	}
}

func printTetros(tetros []tetris.Tetromino) {
	for _, tetro := range tetros {
		printed := ""
		cursor := tetris.Point{
			X: 0,
			Y: 0,
		}

		for !(cursor.X == 4 && cursor.Y == 0) {

			// Choix du caractère imprimé

			if !isInTetro(cursor, tetro) {
				printed += "."
			} else if tetro.ID != 0 {
				printed += string(tetro.ID)
			} else {
				printed += "#"
			}

			// On passe au caractère suivant : si on est à la fin d'une ligne, on l'imprime et on passe à la suivante

			if cursor.Y != 3 {
				cursor.Y++
			} else {
				cursor.Y = 0
				cursor.X++
				fmt.Println(printed)
				printed = ""
			}
		}

		fmt.Print("\n")
	}
}

func isInTetro(cursor tetris.Point, tetro tetris.Tetromino) bool {
	for i := 0; i < len(tetro.Coord); i++ {
		if cursor == tetro.Coord[i] {
			return true
		}
	}
	return false
}
