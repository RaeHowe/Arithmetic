package main

func main()  {

}

type Point struct {
	x int
	y int
}

func (p *Point) walk(x, y int) {
	p.x += x
	p.y += y
}

func isPathCrossing(path string) bool {
	rule := map[rune][2]int{ //运行规则（固定不变的）
		'E': {1, 0},
		'W': {-1, 0},
		'N': {0, 1},
		'S': {0, -1},
	}

	visited := map[Point]bool{}
	point := Point{0, 0} //原点
	visited[point] = true //经过的点都放置到map中，如果从map中取指定的点取到的话，就返回true

	for _, p := range path {
		point.walk(rule[p][0], rule[p][1])
		if visited[point] {
			return true
		}
		visited[point] = true
	}

	return false
}