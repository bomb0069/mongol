package mongo

import (
	"bitbucket.org/4fit/mongol/primitives"
	"bitbucket.org/4fit/mongol/primitives/custom_error"
)

type mongoLock struct {
	/*db *mgo.Database*/
}

func Lock( /*db *mgo.Database*/ ) (primitives.SyncLock, custom_error.CustomError) {
	/*if db == nil {
		return nil, custom_error.MakeErrorf("nil db-object provided")
	}
	locks := db.C("mongol_migration_lock")
	if locks == nil {
		return nil, custom_error.MakeErrorf("failed to get collection of locks")
	}
	lockCount, err := locks.Count()
	if err != nil {
		return nil, custom_error.MakeErrorf("failed to get lock-count. Error: %v", err)
	}
	if lockCount > 0 {
		return nil, custom_error.MakeErrorf("mongo seems to be locked")
	}*/
	return nil, custom_error.MakeErrorf("Not implemented")
}
