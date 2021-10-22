package dbs

import (
	"fmt"
	"log"
)

type User struct {
	Id int64 `db:"id"`
	UserName string `db:"username"`
	Email string `db:"email"`
	Password string `db:"password"`
	CreatedAt string `db:"createdAt"`
	Disabled int8 `db:"disableFlag"`
}

// 查询数据，指定字段名
func StructQueryField() {
	user := new(User)
	row := mysqlDB.QueryRow("select id, username, email, password, createdAt, disabled from user where id=?", 1)
	if err := row.Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.Disabled); err != nil {
		fmt.Printf("scan failed, err: %v", err)
		return
	}
	fmt.Println(user.Id, user.UserName, user.Email)
}

// 查询数据，取所有字段
func StructQueryAllField() {
	// 通过切片存储
	users := make([]User, 0)
	rows, _:= mysqlDB.Query("SELECT * FROM `user` limit ?", 100)
	// 遍历
	var user User
	for rows.Next(){
		rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.Disabled)
		users=append(users,user)
	}
	fmt.Println(users)
}

// 插入数据
func StructInsert() {
	ret, _ := mysqlDB.Exec("insert INTO user(username, email, password) values(?,?,?)", "user", "user@email.com", "user")

	//插入数据的主键id
	lastInsertID,_ := ret.LastInsertId()
	fmt.Println("LastInsertID:",lastInsertID)

	//影响行数
	rowsaffected,_ := ret.RowsAffected()
	fmt.Println("RowsAffected:",rowsaffected)
}

// 更新数据
func StructUpdate() {
	ret,_ := mysqlDB.Exec("UPDATE user set email=? where id=?","updateUser@email.com", 1)
	updates, _ := ret.RowsAffected()

	fmt.Println("RowsAffected:", updates)
}

// 删除数据
func StructDel() {
	ret,_ := mysqlDB.Exec("delete from user where id=?", 1)
	deletes,_ := ret.RowsAffected()

	fmt.Println("RowsAffected:", deletes)
}

func Count(all bool) int64 {
	sql := "select count(id) from user where disableFlag=0"
	if (all) {
		sql = "select count(id) from user"
	}
	var count int64
	mysqlDB.QueryRow(sql).Scan(&count)
	fmt.Println("user count: ", count)
	return count
}

func Find(offset int64, limit int64, all bool) []User {
	// 通过切片存储
	users := make([]User, 0)
	sql := "SELECT * FROM user where disableFlag=0 limit ? offset ?"
	if all {
		sql = "SELECT * FROM user where limit ? offset ?"
	}
	rows, _:= mysqlDB.Query(sql, limit, offset)
	// 遍历
	var user User
	for rows.Next() {
		rows.Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.Disabled)
		users=append(users,user)
	}
	fmt.Println("found users: ", users)
	return users
}

func FindOneById(id int64) *User {
	user := new(User)
	row := mysqlDB.QueryRow("select * from user where id=?", id);
	if err := row.Scan(&user.Id, &user.UserName, &user.Email, &user.Password, &user.CreatedAt, &user.Disabled); err != nil {
		fmt.Printf("scan failed, err: %v", err)
		return nil;
	}
	fmt.Println("found user: ", user)
	return user;
}

func CreateUser(username string, email string, password string) *User {
	ret, err := mysqlDB.Exec("insert INTO user(username, email, password) values(?,?,?)", username, email, password)

	if err != nil {
		log.Println("create user error: ", err.Error())
		return nil
	}

	//插入数据的主键id
	lastInsertID,_ := ret.LastInsertId()
	return FindOneById(lastInsertID)
}

func update(user *User) {
	ret, err := mysqlDB.Exec("UPDATE user set username=?, email=?, password=? where id=?", user.UserName, user.Email, user.Password)
	if err != nil {
		fmt.Printf("update user failed, err: %v", err)
	}
	updates, _ := ret.RowsAffected()
	if updates == 1 {
		fmt.Println("update success.")
	} else {
		fmt.Println("update failed.")
	}
}

func delete(user *User) {
	ret, err := mysqlDB.Exec("delete from user where id=?", user.Id)

	if err != nil {
		fmt.Printf("delete user failed, err: %v", err)
	}

	deletes, _ := ret.RowsAffected()

	if deletes == 1 {
		fmt.Println("delete success.")
	} else {
		fmt.Println("delete failed.")
	}
}
