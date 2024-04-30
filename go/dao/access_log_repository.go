package dao

import (
    "fmt"
    "time"
    "database/sql"
    "double_up/model"
)

func CreateAccessLog(db *sql.DB, accessLog *model.AccessLog) error {
    query := "INSERT INTO access_log (id, created_at, ip_address, api) VALUES ($1, $2, $3, $4)"
    _, err := db.Exec(query, accessLog.ID, accessLog.CreatedAt, accessLog.IpAddress, accessLog.Api)
    if err != nil {
        return fmt.Errorf("AccessLog cannot be created, id: %s. Error: %s", accessLog.ID, err)
    }
    return nil
}

func CountRowsForLastMinute(db *sql.DB, ipAddress string, startTime time.Time) (int16, error) {
    query := "SELECT COUNT(id) FROM access_log WHERE created_at >= $1 and ip_address = $2"
    row := db.QueryRow(query, startTime, ipAddress)

    var result int16
	err := row.Scan(&result)
	if err != nil {
	    return 0, fmt.Errorf("Error in CountRowsForLastMinute func: %s", err)
	}
    return result, nil
}
