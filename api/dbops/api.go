package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	defer stmtIns.Close()
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	return pwd, nil
}

func DeleteUser(loginName, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	defer stmtDel.Close()
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	uuid, err := utils.NewUUID()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info (id, author_id, name, display_ctime) 
								VALUES (?, ?, ?, ?)`)
	defer stmtIns.Close()
	_, err = stmtIns.Exec(uuid, aid, name, ctime)
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	result := &defs.VideoInfo{Id: uuid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	return result, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")
	defer stmtOut.Close()
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}

	var (
		aid int
		displayCtime string
		name string
	)
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &displayCtime)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("%s", err)
		return nil, err
	}
	if err == sql.ErrNoRows {
		log.Printf("%s", err)
		return nil, nil
	}
	result := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: displayCtime}
	return result, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	defer stmtDel.Close()
	if err != nil {
		log.Printf("%s", err)
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	return nil
}