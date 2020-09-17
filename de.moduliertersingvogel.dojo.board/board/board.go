package board

type Stone struct {
	Type string
}

type Field struct {
	Color *string
	Stone *Stone
}

type Board struct {
	Fields [][]Field
}

func NewBoard(size int, stonefactory func() Stone, fieldfactory func(x int, y int) Field) Board {
	board := Board{}
	board.Fields = make([][]Field, size)
	fields := make([]Field, size*size)
	for i := range board.Fields {
		board.Fields[i], fields = fields[:size], fields[size:]
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			board.Fields[i][j] = fieldfactory(i, j)
		}
	}

	return board
}
