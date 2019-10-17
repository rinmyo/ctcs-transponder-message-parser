package ETCSPacks

type ShuntingDangerousPack struct {
	Part1 struct {
		NID_PACKET uint16
		Q_DIR      uint16
		L_PACKET   uint16
	}

	Part2 struct {
		Q_ASPECT uint16
	}
}
