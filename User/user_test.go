package User

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("开始测试:")
	t.Run("测试查询一条记录:", testUser_SelectUserByEmail)
}

func testUser_SelectUserByEmail(t *testing.T) {
	fmt.Println("测试查询一条记录")
	user := User{}
	user.InitMysql()
	prepare, err := user.Db.Prepare("select verify_code from users.user where email=?")
	if err != nil {
		return
	}
	row := prepare.QueryRow("1784929126@qq.com")
	var verify string
	row.Scan(&verify)
	fmt.Println("row", row)
	fmt.Println("Verify code is:", verify)
}
