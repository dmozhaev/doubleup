package dao

import (
    "database/sql"
)

func RunInTransaction(db *sql.DB, query func(dbParam *sql.DB, txParam *sql.Tx) error) error {
    // start transaction
    tx, err := db.Begin()
    if err != nil {
        return err
    }

    defer func() {
        if r := recover(); r != nil {
            tx.Rollback()
        }
    }()

    // perform the query
    err = query(db, tx)
    if err != nil {
        return err
    }

    // commit the transaction
    if err = tx.Commit(); err != nil {
        return err
    }

    return nil
}
