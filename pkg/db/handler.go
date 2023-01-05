package db

import (
	"github.com/jackbalageru/MERN-CRUD/go-server/pkg/model"
	"github.com/pkg/errors"
)

func (h *DBHandler) CreateBoard(board *model.Board) (*model.Board, error) {
	result := h.gDB.Create(board)

	return board, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) GetBoardList() ([]*model.Board, error) {
	boardList := []*model.Board{}
	result := h.gDB.Find(&boardList)

	return boardList, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) GetBoardByID(id string) (*model.Board, error) {
	board := &model.Board{}
	result := h.gDB.First(board, id)  // 특정 id 목록 하나 추가해줌 

	return board, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) UpdateBoard(id uint, newBoard *model.Board) (*model.Board, error) {
	oldBoard := &model.Board{}
	result := h.gDB.Model(oldBoard).Where(id).Updates(newBoard)
	// where로 특정 id에 대한 데이터 수정 

	return oldBoard, errors.Wrap(result.Error, "db handler error")
}

func (h *DBHandler) DeleteBoardByID(id string) error {
	result := h.gDB.Delete(&model.Board{}, id)

	return errors.Wrap(result.Error, "db handler error")
}
