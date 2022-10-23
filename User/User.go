package User

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

// User 结构体
type User struct {
	ID         int
	Username   string
	Password   string
	email      string
	VerifyCode string
	IsVerify   int
}

func InitMysql() {
	Db, err = sql.Open("mysql", "root:660967@tcp(192.168.1.103:3306)/users")
	if err != nil {
		fmt.Println("打开失败:", err)
	}
}

// AddUser 添加user
func (user *User) AddUser() error {
	InitMysql()
	sqlStr := "insert into user(username, password, email) values (?,?,?)"
	inStmt, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现错误:", err)
		return err
	}
	_, err = inStmt.Exec(user.Username, user.Password, user.email)
	if err != nil {
		fmt.Println("执行出现异常:", err)
	}
	return nil
}

// SelectUserByEmail 查询是否存在此用户，并验证密码
func (user *User) SelectUserByEmail() error {
	InitMysql()

	sqlStr := "select username,password,email from user where email=?"

	prepare, err := Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出错:", err)
		return err
	}
	row := prepare.QueryRow(user.email)

	// 声明
	var username string
	var password string
	var email string
	err = row.Scan(&username, &password, &email)
	if err != nil {
		fmt.Println("查询错误:", err)
		return err
	}
	fmt.Println(username, password, email)
	return nil
}
