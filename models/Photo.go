package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


type Photo struct{
	gorm.Model
	Title string
	Description string
	Image string
}



func (p Photo) AutoMigrate(){
	db,err:= gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	if err !=nil{
		fmt.Println(err)
		return
	}

	db.AutoMigrate(&p)
}


func (p Photo) Create()  {

	db,err:= gorm.Open(mysql.Open(dsn), &gorm.Config{}) 
	 
	if err !=nil{
		fmt.Println(err)
		return
	}

	db.Create(&p)
}


func (p Photo) Get(where ...interface{}) Photo {

	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return p
	}

	db.First(&p,where...)
	return p	
}


func (p Photo) GetAll(where ...interface{}) []Photo{
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil{
		fmt.Println(err)
		return nil
	}

	var photos []Photo
	db.Find(&photos,where...)
	return photos
}


func (p Photo) Update(column string , value interface{}){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil{
		fmt.Println(err)
		return 
	}


	db.Model(&p).Update(column,value)
}


func (p Photo) Updates(data Photo){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil{
		fmt.Println(err)
		return 
	}


	db.Model(&p).Updates(data)

}


func (p Photo) Delete (){
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil{
		fmt.Println(err)
		return 
	}


	db.Delete(&p,p.ID)

}



	// *****KULLANIM ÖRNEKLERİ*****


	// 	********CREATE********

	// 	models.Photo{
	// 		Title: title,
	// 		Description: des,
	// 	}.Create()


	// 	*********GET********
	// 	models.Photo{}.Get(veri tabanından id gir)


	// 	********GET ALL********
	// 	models.Photo{}.GetAll()



	// 	********UPDATE&UPDATES********
	// 	post:= models.Photo{}.Get(veri tabanından id gir)
	// 	post.Update(column,value) 
	// 	post.Updates(models.Photo{Title:"yeni title",Description:"yeni description"})



	// 	********DELETE********
	// 	post:= models.Photo{}.Get(veri tabanından id gir)
	// 	post.Delete()




