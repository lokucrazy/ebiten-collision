package main

func hitsTop(player *Player, platform *Platform) (bool, int) {
	px0 := player.x
	px1 := player.x + player.width
	py0 := player.y
	py1 := player.y + player.height
	plx0 := platform.x
	plx1 := platform.x + platform.width
	ply0 := platform.y
	ply1 := platform.y + platform.height

	if inBetween(px0, plx0, plx1) || inBetween(px1, plx0, plx1) {
		if (py1 > ply0) && (py0 < ply0) {
			return false, ply0 - py1
		} else if (ply1 > py0) && (py1 > ply1) {
			return false, ply1 - py0
		}
	}
	if inBetween(py0, ply0, ply1) || inBetween(py1, ply0, ply1) {
		if (px1 > plx0) && (px0 < plx0) {
			return true, plx0 - px1
		} else if (plx1 > px0) && (px1 > plx1) {
			return true, plx1 - px0
		}
	}
	return false, 0
}
