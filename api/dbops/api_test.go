package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// 过程：init(dblogin, truncate tables)->run tests->clear data(truncate tables)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("haohao", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("haohao")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
	fmt.Println(pwd)
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("haohao", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("haohao")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

// 视频功能测试
var tempvid string

func TestVideoWorkFlow(t *testing.T) {
	clearTables()
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
	info, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of getVideoInfo: %v", err)
	}
	if info.Name != "my-video" {
		t.Error("Error of getVideoInfo: content not correct")
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	info, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if info != nil {
		t.Errorf("Deleting user test failed")
	}
}

// 评论功能测试

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	err := AddNewComments(tempvid, 1, "respect this video")
	if err != nil {
		t.Errorf("Error of AddNewComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := tempvid
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
