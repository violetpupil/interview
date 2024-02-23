package program

import "unicode"

const (
	Left = iota
	Top
	Right
	Bottom
)

// 机器人坐标问题
// https://github.com/lifei6671/interview-go/blob/master/question/q006.md
func RobotMove(cmd string, x, y, z int) (int, int, int) {
	repeat := 0
	repeatCmd := ""
	for _, c := range cmd {
		switch {
		case unicode.IsNumber(c):
			repeat = int(c) - '0' + repeat*10
		case repeat > 0 && c != '(' && c != ')':
			repeatCmd += string(c)
		case c == ')':
			for i := 0; i < repeat; i++ {
				x, y, z = RobotMove(repeatCmd, x, y, z)
			}
			repeat = 0
			repeatCmd = ""
		case c == 'F':
			switch z {
			case Right:
				x++
			case Left:
				x--
			case Top:
				y++
			case Bottom:
				y--
			}
		case c == 'B':
			switch z {
			case Right:
				x--
			case Left:
				x++
			case Top:
				y--
			case Bottom:
				y++
			}
		case c == 'L':
			z = (z - 1 + 4) % 4
		case c == 'R':
			z = (z + 1) % 4
		}
	}

	return x, y, z
}
