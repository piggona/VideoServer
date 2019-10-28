package dbops

import (
	"database/sql"
	"log"
	"time"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"

	"video_server/api/defs"
	"video_server/api/utils"
)

// AddUserCredential 添加一个用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES (?, ?)")
	if err != nil {
		return err
	}
	if _, err := stmtIns.Exec(loginName, pwd); err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

// GetUserCredential 获取用户的密码的哈希
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}

	var pwd string
	// stmtOut.QueryRow(loginName).Scan(&pwd)
	if err := stmtOut.QueryRow(loginName).Scan(&pwd); err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

// DeleteUser 删除一个用户
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("DeleteUser error: %s \n", err)
		return err
	}
	if _, err := stmtDel.Exec(loginName, pwd); err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

// AlterUser 为用户修改密码：之前点击修改时会发送一个邮箱验证码，使用验证码，用户名及新密码的哈希修改
func AlterUser(loginName string, pwd string, verifyCode string) error {
	return nil
}

// AddNewVideo 添加一个新视频，给它分配一个uuid并返回相应的存储在数据库中的数据结构，利于视频上传
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	// createtime->db->
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info 
		(id,author_id, name,display_ctime) VALUES (?,?,?,?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, nil
}

// GetVideoInfo 输入视频id，得到数据库中视频的信息
func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

// DeleteVideoInfo 输入视频id，删除这个视频
func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

// ListAllVideoinfo 输出所有视频信息列表
func ListAllVideoinfo() ([]*defs.VideoInfo, error) {
	return []*defs.VideoInfo{}, nil
}

// AddNewComments 添加一个评论
func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?,?,?,?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

// ListComments 将某个视频下的评论按时间段from-to列举出来
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.Login_name, comments.content FROM comments
	INNER JOIN users ON comments.author_id = users.id
	WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
