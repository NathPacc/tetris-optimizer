package tetris

import (
	"strings"
)

type Point struct {
	X int
	Y int
}

type Tetromino struct {
	Coord []Point
	ID    rune
}

func GetTetros(datas string) []Tetromino {
	answer := []Tetromino{}
	datatab := strings.Fields(datas)
	datatab = removeEmpty(datatab)
	tetro := Tetromino{
		Coord: []Point{},
	}

	for i := range datatab {

		// On récupère les lettres de l'alphabet toutes les 4 lignes parcourues

		id := rune(i/4 + 65)
		tetro.ID = id

		line := datatab[i]
		for j := range line {
			if string(line[j]) == "#" {
				point := Point{
					X: i % 4,
					Y: j,
				}
				tetro.Coord = append(tetro.Coord, point)
			}
		}

		// Toutes les 4 lignes, on ajoute le tetromino, même si il n'est pas valable, à la liste.

		if i%4 == 3 {
			answer = append(answer, tetro)
			tetro.Coord = []Point{}
		}
	}
	return answer
}

func removeEmpty(tab []string) []string {
	copy := []string{}
	for _, str := range tab {
		if str != "" {
			copy = append(copy, str)
		}
	}
	return copy
}

func IsConnected(tetro Tetromino) bool {

	// On enregistre les blocs déjà visités

	visited := make(map[int]bool)
	queue := []int{0}
	visited[0] = true

	// On visite un par un les blocs du tetromino

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		for i := 0; i < 4; i++ {

			// Si notre bloc est voisin d'un bloc pas encore visité, on le marque et on l'ajoute à la liste des blocs à tester

			if !visited[i] && AreAdjacent(tetro.Coord[current], tetro.Coord[i]) {
				visited[i] = true
				queue = append(queue, i)
			}
		}
	}

	// Si nous n'avons pas 4 blocs de visités, alors un d'entre eux est séparé du reste.

	return len(visited) == 4
}

func AreAdjacent(a, b Point) bool {
	dx := abs(a.X - b.X)
	dy := abs(a.Y - b.Y)
	return dx+dy == 1
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
