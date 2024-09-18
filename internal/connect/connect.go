package connect

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Persons struct {
	Id           int `gorm:"primaryKey"`
	Name         string
	Gender       string
	Age          int
	Address      string
	Phone_number string
}

func Connect() (*gorm.DB, error) {
	connString := "Server=LAPTOP-7CAHEI3Q\\HATHANHHAO;Database=Gotest;User Id=sa;Password=hao123;"

	// Mở kết nối
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal("Không thể mở kết nối:", err)
		return nil, err
	}

	fmt.Println("Kết nối thành công với SQL Server!")
	return db, nil
}

func AddUser(db *gorm.DB, per Persons) error {
	if err := db.Create(&per).Error; err != nil {
		return fmt.Errorf("thêm dữ liệu thất bại: %v", err)
	}
	fmt.Println("Thêm dữ liệu thành công!")
	return nil
}

func UpdateUser(db *gorm.DB, per Persons) error {
	if err := db.Model(&Persons{}).Where("id = ?", per.Id).Updates(per).Error; err != nil {
		return fmt.Errorf("cập nhật dữ liệu thất bại: %v", err)
	}
	fmt.Println("Cập nhật dữ liệu thành công!")
	return nil
}

func DeleteUser(db *gorm.DB, id int) error {
	if err := db.Delete(&Persons{}, id).Error; err != nil {
		return fmt.Errorf("xóa dữ liệu thất bại: %v", err)
	}
	fmt.Println("Xóa dữ liệu thành công!")
	return nil
}

func GetAllUsers(db *gorm.DB) ([]Persons, error) {
	var users []Persons
	if err := db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("lấy dữ liệu thất bại: %v", err)
	}
	return users, nil
}

func FindUserByID(db *gorm.DB, id int) (Persons, error) {
	var per Persons
	if err := db.First(&per, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return per, fmt.Errorf("không tìm thấy người dùng với ID: %d", id)
		}
		return per, fmt.Errorf("lỗi khi tìm kiếm người dùng: %v", err)
	}
	return per, nil
}
