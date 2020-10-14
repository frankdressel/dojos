package engine

import (
	"de.moduliertersingvogel.dojo.board/board"
	"errors"
)

func NextMove(board board.Board, active string) error {
	if active != "black" {
		return errors.New("Currently only supporting black moves")
	}

	return nil
}
