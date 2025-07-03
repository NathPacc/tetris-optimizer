package tetris

import (
	"math"
	"strings"
)

func Resolve(datas string) [][]rune {
	tetros := GetTetros(datas)

	// On récupère la taille minimum possible de la grille

	sizeGrid := int(math.Ceil(math.Sqrt(float64(4 * len(tetros)))))
	for {

		// On créé une grille vide

		grid := make([][]rune, sizeGrid)
		for i := range grid {
			grid[i] = []rune(strings.Repeat(".", sizeGrid))
		}
		result := solveRecursive(tetros, 0, grid)

		// Si il n'est pas possible de remplir la grille, on augmente sa taille et on recommence

		if result != nil {
			return result
		}
		sizeGrid++
	}
}

func solveRecursive(tetros []Tetromino, index int, grid [][]rune) [][]rune {

	// Si tous les tetrominos ont été placés, ont renvoie la grille

	if index == len(tetros) {
		return grid
	}

	tetro := normalizeTetro(tetros[index])

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			pt := Point{
				X: x,
				Y: y,
			}
			if canPlace(tetro, pt, grid) {
				grid = place(tetro, pt, grid)
				res := solveRecursive(tetros, index+1, grid)
				if res != nil {
					return res
				}

				// Si l'on atteint c'ette étape, c'est qu'il n'est pas possible de remplir la grille, il faut retirer le dernier tetromino placé
				// pour le replacer ailleurs
				grid = remove(tetro, pt, grid)
			}
		}
	}
	return nil
}

func canPlace(tetro Tetromino, point Point, grid [][]rune) bool {
	for _, coord := range tetro.Coord {
		x := coord.X + point.X
		y := coord.Y + point.Y

		// Si l'un de nos blocs sort de la grille, on ne peut pas le placer

		if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
			return false
		}

		// Si l'une des coordonnées est déjà remplie par un autre tetromino, ont ne peut pas le placer

		if grid[x][y] != '.' {
			return false
		}
	}
	return true
}

func place(tetro Tetromino, point Point, grid [][]rune) [][]rune {
	for _, coord := range tetro.Coord {
		x := coord.X + point.X
		y := coord.Y + point.Y
		grid[x][y] = tetro.ID
	}
	return grid
}

func remove(tetro Tetromino, pt Point, grid [][]rune) [][]rune {
	for _, coord := range tetro.Coord {
		x := coord.X + pt.X
		y := coord.Y + pt.Y
		grid[x][y] = '.'
	}
	return grid
}

// Le but de la fonction suivante est de placer le coin haut-gauche du tetromino au coordonnées (0;0) pour simplifier son placement
// dans la grille

func normalizeTetro(tetro Tetromino) Tetromino {
	answer := Tetromino{
		ID:    tetro.ID,
		Coord: []Point{},
	}
	minX := tetro.Coord[0].X
	minY := tetro.Coord[0].Y

	// On identifie les coordonnées du point haut-gauche du tetromino
	for _, pt := range tetro.Coord {
		if pt.X < minX {
			minX = pt.X
		}
		if pt.Y < minY {
			minY = pt.Y
		}
	}

	// On modifie les coordonnées de tous les points de sorte à ce que le point haut-gauche se retrouve en (0;0)
	for _, pt := range tetro.Coord {
		normalized := Point{
			X: pt.X - minX,
			Y: pt.Y - minY,
		}
		answer.Coord = append(answer.Coord, normalized)
	}

	return answer
}
