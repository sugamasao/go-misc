package main

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//func (m *Blog) TableName() string {
//	return "blog"
//}

type Blog struct {
	//  gorm.Model
	ID          int `gorm:"primary_key"`
	Title       string
	PublishedAt time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Entries     []Entry
}

type Entry struct {
	ID        int `gorm:"primary_key"`
	BlogID    int
	Title     string
	body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func main() {
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:13306)/example?parseTime=true")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	// 単数形テーブルの場合はこのフラグでいい感じにしてくれる
	db.SingularTable(true)

	// Read
	var blog Blog
	db.First(&blog, 1)
	fmt.Println("[1] SELECT : ", blog)

	// Update
	db.Model(&blog).Update("published_at", time.Now().Add(time.Duration(3)*time.Hour))
	// db.First(&blog, 1)
	fmt.Println("[1] UPDATE : ", blog)

	db.First(&blog, "id > ?", "2")
	fmt.Println("[> 2]", blog)

	var blogs []Blog
	db.Table("blog").Joins("join entry on blog.id = entry.blog_id").Where("entry.id > 2").Find(&blogs)
	fmt.Println("[join]", blogs)
	for i := range blogs {
		fmt.Printf("blog[%d] => %v\n", i, blogs[i])
	}

	fmt.Println("-----")
	rows, err := db.Table("blog").Joins("join entry on blog.id = entry.blog_id").Preload("Entries").Rows()
	if err != nil {
		fmt.Println(err)
		panic("join miss")
	}
	for rows.Next() {
		var b Blog
		db.ScanRows(rows, &b)
		fmt.Printf("blog=> %v * %v\n", b, rows)
	}
}
