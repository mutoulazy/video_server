package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video_server/api/defs"
)

func InsertSession(sid string, ttl int64, uname string) error {
	ttlStr := strconv.FormatInt(ttl, 10)
	stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id, TTL, login_name) VALUES (?, ?, ?)")
	defer stmtIns.Close()
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	_, err = stmtIns.Exec(sid, ttlStr, uname)
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	return nil
}

func RetrieveSession(sid string) (*defs.SimpleSession, error) {
	session := &defs.SimpleSession{}
	stmtOut, err := dbConn.Prepare("SELECT TTL, login_name FROM sessions WHERE session_id=?")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	var ttl string
	var uname string
	stmtOut.QueryRow(sid).Scan(&ttl, &uname)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s", err)
		return nil, err
	}

	if ttl64, err := strconv.ParseInt(ttl, 10, 64); err == nil {
		session.TTL = ttl64
		session.Username = uname
	} else {
		log.Printf("%s", err)
		return nil, err
	}
	return session, nil
}

func RetrieveAllSessions() (*sync.Map, error) {
	sessionMap := &sync.Map{}
	stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	rows, err := stmtOut.Query()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	for rows.Next() {
		var id string
		var ttlstr string
		var login_name string
		if err1 := rows.Scan(&id, &ttlstr, &login_name); err1 != nil {
			log.Printf("retrive sessions error: %s", err1)
			break
		}

		if ttl, err1 := strconv.ParseInt(ttlstr, 10, 64); err1 == nil {
			session := &defs.SimpleSession{Username: login_name, TTL: ttl}
			sessionMap.Store(id, session)
			log.Printf(" session id: %s, ttl: %d", id, session.TTL)
		} else {
			log.Printf("%s", err1)
			return nil, err1
		}
	}

	return sessionMap, nil
}

func DeleteSession(sid string) error {
	stmtOut, err := dbConn.Prepare("DELETE FROM sessions WHERE session_id = ?")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	if _, err := stmtOut.Query(sid); err != nil {
		log.Printf("%s", err)
		return err
	}
	return nil
}
