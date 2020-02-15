package dbops

import (
	"fmt"
	"testing"
)

var tempvid string

func TestMain(m *testing.M) {
	cleanTables()
	m.Run()
	cleanTables()
}

func cleanTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")

}

func TestUsers(t *testing.T) {
	t.Run("addUser", testAddUser)
	t.Run("getUser", testGetUser)
	t.Run("deleteUser", testDeleteUser)
	t.Run("regetUser", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("mutou", "123")
	if err != nil {
		t.Errorf("Error of Add User %v", err)
	}
}

func testGetUser (t *testing.T) {
	pwd, err := GetUserCredential("mutou")
	if pwd != "123" || err != nil {
		t.Error("Error of Get User")
	}
	fmt.Println(pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("mutou", "123")
	if err != nil {
		t.Errorf("Error of Delete User %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("mutou")
	if err != nil {
		t.Errorf("Error of Reget User %v", err)
	}
	if pwd != "" {
		t.Error("delete User Test Failed")
	}
}

func TestVideoInfo(t *testing.T) {
	cleanTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}