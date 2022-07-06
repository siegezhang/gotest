package test

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var Db *gorm.DB

func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

func connectDB() *gorm.DB {
	var err error
	dsn := "root:123456@(127.0.0.1)/go_gin_api?parseTime=true&charset=utf8mb4&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // table name prefix, table for `User` would be `t_users`
			SingularTable: false, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})

	if err != nil {
		//	fmt.Println("Error connecting to database : error=%v", err)
		return nil
	}
	return db
}

type User struct {
	gorm.Model
	ID    int
	Name  string `gorm:"column:myname"`
	Email string
}

//指定表名
func (User) TableName() string {
	return "user_info"
}

//create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}

func TestMysqlCreate(t *testing.T) {
	db := InitDb()
	//表不存在，则创建表结构
	db.AutoMigrate(&User{})
	CreateUser(InitDb(), &User{
		Name:  "siege",
		Email: "z1178858918@qq.com",
	})
}

func TestMysqlGet(t *testing.T) {
	Convey("1 should equal 1", t, func() {
		So(1, ShouldEqual, 1)
	})
	db := InitDb()
	//表不存在，则创建表结构
	db.AutoMigrate(&User{})
	var user []User
	GetUsers(db, &user)
}

func TestMysqlGetOne(t *testing.T) {
	db := InitDb()
	var user User
	GetUser(db, &user, "1")
}

func TestMysqlUpdate(t *testing.T) {
	db := InitDb()
	user := User{
		ID:    1,
		Name:  "siege11",
		Email: "z11788589118@qq.com",
	}
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	UpdateUser(db, &user)
}
func TestMysqlDelete(t *testing.T) {
	user := User{
		ID:    1,
		Name:  "siege11",
		Email: "z11788589118@qq.com",
	}
	DeleteUser(InitDb(), &user, "1")
}
