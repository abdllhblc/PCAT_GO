package main

import (
	"fmt"
	"modules/controllers"
	"modules/models"
	"net/http"
	"github.com/julienschmidt/httprouter"
)




func main() {

	models.Photo{}.AutoMigrate()
	
	router:= httprouter.New()
       

	

	router.GET("/",controllers.HomePage);
	router.GET("/add",controllers.Add);
	router.GET("/about",controllers.About);
	router.POST("/photos",controllers.PhotoCatch);
	router.GET("/photo/:id",controllers.PhotoCatchID);// tekil foto sayfasını getiriyor.
	router.GET("/photos/edit/:editId",controllers.PhotoEdit); // tekil fotoğraf sayfasında update butonuna basınca edit sayfasını getiriyor.
	router.POST("/photos/:id",controllers.UpdatePhoto); // edit ekranında submit edince güncelleme işini yapıp tekrardan tekil foto sayfasına yönlendirecek endpoint.
	router.GET("/photos/delete/:id",controllers.DeletePhoto); 

	router.ServeFiles("/assets/*filepath",http.Dir("./assets")) // assets dosyalarını servis etmek için kullanılır
	fmt.Println("http://localhost:8000")
	http.ListenAndServe(":8000",router)

}



