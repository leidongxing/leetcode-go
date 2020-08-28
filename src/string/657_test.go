package string

//机器人能否返回原点
func judgeCircle(moves string) bool {
	rl := 0
	ud := 0
	for i := 0; i < len(moves); i++ {
		if moves[i] == 'R' {
			rl++
		} else if moves[i] == 'L' {
			rl--
		} else if moves[i] == 'U' {
			ud++
		} else if moves[i] == 'D' {
			ud--
		}
	}
	return rl == 0 && ud == 0
}
