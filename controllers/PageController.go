package controllers

import (
	"fmt"
	"modules/models"
	"net/http"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

func HomePage(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles("./views/index.html", "./views/partials/_header.html", "./views/partials/_navbar.html", "./views/partials/_footer.html")
	if err != nil {
		fmt.Println(err)
	}

	data_photo := models.Photo{}.GetAll()

	view.ExecuteTemplate(w, "index", data_photo)

}

func Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles("./views/add.html", "./views/partials/_header.html", "./views/partials/_navbar.html", "./views/partials/_footer.html")

	if err != nil {
		fmt.Println(err)
	}

	view.ExecuteTemplate(w, "add", nil)

}

func About(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles("./views/about.html", "./views/partials/_header.html", "./views/partials/_navbar.html", "./views/partials/_footer.html")

	if err != nil {
		fmt.Println(err)
	}

	view.ExecuteTemplate(w, "about", nil)

}

func PhotoCatchID(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	view, err := template.ParseFiles("./views/photo.html", "./views/partials/_header.html", "./views/partials/_navbar.html", "./views/partials/_footer.html")

	if err != nil {
		fmt.Println(err)
	}

	id := params.ByName("id")

	photo := models.Photo{}.Get(id)

	view.ExecuteTemplate(w, "Photos", photo)
}

func PhotoEdit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	view, err := template.ParseFiles("./views/edit.html", "./views/partials/_header.html", "./views/partials/_navbar.html", "./views/partials/_footer.html")

	if err != nil {
		fmt.Println(err)
	}

	editId := params.ByName("editId")

	photo := models.Photo{}.Get(editId)

	view.ExecuteTemplate(w, "edit", photo)

}