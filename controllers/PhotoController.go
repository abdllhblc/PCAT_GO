package controllers

import (
	"fmt"
	"io"
	"modules/models"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func PhotoCatch(w http.ResponseWriter, r *http.Request, param httprouter.Params) {

	title := r.FormValue("title")
	des := r.FormValue("description")

	file, handler, err := r.FormFile("image")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	uploadDir := "./assets/img/"

	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		err := os.Mkdir(uploadDir, os.ModePerm)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	filepath := uploadDir + handler.Filename
	out, err := os.Create(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	path := "/assets/img/" + handler.Filename

	models.Photo{
		Title:       title,
		Description: des,
		Image:       path,
	}.Create()

	http.Redirect(w, r, "/", http.StatusFound)

}


func UpdatePhoto(w http.ResponseWriter, r *http.Request,params httprouter.Params){

	id:= params.ByName("id");
	title:= r.FormValue("title")
	des := r.FormValue("description")
	photo := models.Photo{}.Get(id);
  
	photo.Updates(models.Photo{Title:title,Description:des})
    


	http.Redirect(w,r,"/photo/"+id,http.StatusFound);
}


func DeletePhoto(w http.ResponseWriter, r *http.Request,params httprouter.Params){
	
	id := params.ByName("id");


	photo := models.Photo{}.Get(id);

	Filename := "." + photo.Image
	err := os.Remove(Filename);

	if err != nil{
		fmt.Println("dosya silinemedi bir hata oluştu:",err);
		return
	}
	photo.Delete()
	http.Redirect(w,r,"/",http.StatusFound);
}