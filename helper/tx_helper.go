package helper

import "database/sql"

func TxCommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRollback := tx.Rollback()
		if errRollback != nil {
			return
		}
		panic(err)
	} else {
		errCommit := tx.Commit()
		if errCommit != nil {
			return
		}

	}

}
