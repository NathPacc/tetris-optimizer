### Description

This program arranges tetrominos in order to be contained in the smallest square possible (the tetrominos are not rotated).

### Authors

Nathan PACCOUD - Program created during my formation in Zone01 Rouen.

### How yo use 

The program use the data.txt as entry point. It needs to contain a formated text file of 4X4 squares where '.' represents an empty space and '#' represent a part of the tetromino, and each square muste be separated by a empty line. The file must end with an empty line.

Ex :

....\n
.##.\n
.##.\n
....\n
\n
.#..\n
.##.\n
.#..\n
....\n
\n
....\n
..##\n
.##.\n
....\n
\n
....\n
.##.\n
.##.\n
....\n
\n
....\n
..#.\n
.##.\n
.#..\n
\n

It will print in the terminal the square with uppercase letters to identify each tetrominos and '.' to represent the empty spaces.

### Algorithm Details

The tetrominos are recovered in the form of a struct comprising a slice of 4 Points and an ID. The tetrominos are each verified to make sure they are composed of exactly 4 bloks all linked. Then, it creates a square grid and try placing each tetrominos (which are normalized) one by one in it. If the program can't place one, it tries somewhere else in the grid. If it reaches the end of the grid with an unplaced tetromino, it restarts again with a bigger grid.
