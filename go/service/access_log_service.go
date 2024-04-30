package service

import (
    "database/sql"
    "fmt"
    "time"
    "double_up/dao"
    "double_up/model"
    "double_up/utils"
)

func WriteAccessLog(db *sql.DB, ipAddress string, api string) error {
    accessLog := model.NewAccessLog(ipAddress, api)
    return dao.CreateAccessLog(db, accessLog)
}

func CheckAccess(db *sql.DB, ipAddress string) error {
    startTime := time.Now().Add(-1 * time.Minute)
    apiCountLastMinute, err := dao.CountRowsForLastMinute(db, ipAddress, startTime)
    if err != nil {
        return err
    }
    if apiCountLastMinute > utils.RequestLimitPerMinute {
		return fmt.Errorf("Too many requests! IP address: %s", ipAddress)
	}
    return nil
}

func CheckAccessAllowed(db *sql.DB, ipAddress string, api string) error {
    WriteAccessLog(db, ipAddress, api)
    CheckAccess(db, ipAddress)
    return nil
}
