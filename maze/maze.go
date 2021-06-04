package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func dumpMap(board [][]int) {
	for _, row := range board {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}

func readMaze(fileName string) (m [][]int, in, out point) {
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = file.Close()
	}()

	rd := bufio.NewReader(file)
	line, le := rd.ReadString('\n')
	if le != nil {
		return nil, point{}, point{}
	}
	val := strings.Fields(line)
	if len(val) != 6 {
		return nil, point{}, point{}
	}

	row, _ := strconv.Atoi(val[0])
	col, _ := strconv.Atoi(val[1])
	inRow, _ := strconv.Atoi(val[2])
	inCol, _ := strconv.Atoi(val[3])
	outRow, _ := strconv.Atoi(val[4])
	outCol, _ := strconv.Atoi(val[5])

	fmt.Println(val)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		line, le = rd.ReadString('\n')
		if le != nil || io.EOF == le {
			break
		}

		val = strings.Fields(line)
		if len(val) < col {
			break
		}

		for j := range maze[i] {
			maze[i][j], _ = strconv.Atoi(val[j])
		}
	}

	return maze, point{inRow, inCol}, point{outRow, outCol}
}

type point struct {
	i, j   int
}

func (p *point)add(d *point) point {
	return point{p.i + d.i, p.j + d.j}
}

func (p *point)mark(board [][]int) {
	board[p.i][p.j] = 88
}

func (p *point)at(board [][]int) (int, bool){
	// fmt.Printf("(%d, %d)\n", p.i, p.j)

	if p.i < 0 || p.i > len(board) - 1 {
		return 0, false
	}

	if p.j < 0 || p.j > len(board[p.i]) - 1 {
		return 0, false
	}

	return board[p.i][p.j], true
}

var dirs = []point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for s := range steps {
		steps[s] = make([]int, len(maze[s]))
	}

	// fmt.Println("steps")
	// dumpMap(steps)

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		for _, n := range dirs {
			next := cur.add(&n)
			// fmt.Printf("next:(%d, %d)\n", next.i, next.j)

			if cur.i > end.i && cur.j > end.j {
				break
			}

			if next == start {
				continue
			}

			val, ok := next.at(maze)
			if !ok || val != 0 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			steps[next.i][next.j] = steps[cur.i][cur.j] + 1

			if next == end {
				return steps
			}

			Q = append(Q, next)
		}
	}

	return steps
}

func walkBack(steps [][]int, start, end point) []point {

	Q := []point{end}

	for {
		cur := Q[len(Q) - 1]
		if cur == start {
			break
		}
		for _, d := range dirs {
			prev := cur.add(&d)
			pv, _ := prev.at(steps)
			cv, _ := cur.at(steps)
			//fmt.Printf("pv:%d cv:%d\n", pv, cv)
			if pv == cv - 1 {
				Q = append(Q, prev)
				break
			}
		}
	}

	return Q
}

func main() {
	maze, start, end:= readMaze("maze/maze.in")

	fmt.Println("Maze:")
	dumpMap(maze)

	steps := walk(maze, start, end)

	fmt.Println("steps:")
	dumpMap(steps)

	stepPath := walkBack(steps, start, end)
	for _, step := range stepPath {
		//fmt.Printf("(%d, %d) ", step.i, step.j)
		step.mark(maze)
	}
	fmt.Println()

	fmt.Println("Maze Marked:")
	dumpMap(maze)
}
