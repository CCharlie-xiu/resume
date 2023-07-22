package DB

import (
	"fmt"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type User struct {
	gorm.Model
	UserID    uint `gorm:"primaryKey"`
	Username  string
	Password  string
	Avatar    string
	Status    string
	Email     string
	Telephone string
	UUID      string `gorm:"type:char(36);unique"`
}

type UserResume struct {
	gorm.Model
	ResumeID string `gorm:"foreignKey:ResumeDataID"`
	UserID   string `gorm:"foreignKey:UUID"`
}

type Resume struct {
	gorm.Model
	ResumeDataID string `gorm:"primaryKey"`
	Title        string
	Username     string
	Subtitle     string
	Skills       string
	School       string
	Project      string
	Competition  string
	Age          int
	Avatar       string
}

func (user *User) BeforeCreate(*gorm.DB) error {
	user.UUID = uuid.New().String()
	return nil
}

var (
	DB *gorm.DB
)

func Construct() {
	dsn := "root:xmfls200405@tcp(127.0.0.1:3306)/resume?charset=utf8mb4&parseTime=True&loc=Local"
	DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	fmt.Println("DB initialized in : ", DB)
	errDbAuto := DB.AutoMigrate(&User{}, &Resume{}, &UserResume{})
	if errDbAuto != nil {
		return
	}
}
