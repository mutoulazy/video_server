package session

import (
	"log"
	"sync"
	"time"
	"video_server/api/dbops"
	"video_server/api/defs"
	"video_server/api/utils"
)

// session cache
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	sessions, err := dbops.RetrieveAllSessions()
	if err != nil {
		log.Fatalf("LoadSessionsFromDB err:%v", err)
		return
	}

	sessions.Range(func(key, value interface{}) bool {
		session := value.(*defs.SimpleSession)
		sessionMap.Store(key, session)
		return true
	})
}

func GenerateNewSessionId(uname string) string {
	id, _ := utils.NewUUID()
	ctime := nowInMilli()
	ttl := ctime + 30*60*1000 // 30 min
	session := &defs.SimpleSession{Username: uname, TTL: ttl}
	sessionMap.Store(id, session)
	dbops.InsertSession(id, ttl, uname)
	return id
}

// 判断session是否超时
func IsSessionExpired(sid string) (string, bool) {
	session, ok := sessionMap.Load(sid)
	if ok {
		currentTime := nowInMilli()
		if session.(*defs.SimpleSession).TTL < currentTime {
			deleteExpiredSession(sid)
			return "", true
		} else {
			return session.(*defs.SimpleSession).Username, false
		}
	}

	return "", true
}
