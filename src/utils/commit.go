package utils

import "github.com/jinzhu/gorm"

func FinishTx(tx *gorm.DB, succ *bool) {
	if succ != nil && *succ {
		tx.Commit()
	} else {
		tx.Rollback()
	}
}

func SetTrue(succ *bool) {
	*succ = true
}
