package service

import (
    "github.com/google/uuid"
    "double_up/dao"
    "double_up/model"
    "database/sql"
)

func GetPlayer(db *sql.DB, id uuid.UUID) (*model.Player, error) {
    return dao.FindById(db, id)
}
