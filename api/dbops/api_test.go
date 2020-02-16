package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
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
	t.Run("prepareUser", testAddUser)
	t.Run("addVideo", testAddVideoInfo)
	t.Run("getVideo", testGetVideoInfo)
	t.Run("delVideo", testDeleteVideoInfo)
	t.Run("regetVideo", testRegetVideoInfo)
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

func TestComments(t *testing.T) {
	cleanTables()
	t.Run("addUser", testAddUser)
	t.Run("addCommnets", testAddComments)
	t.Run("listComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}