package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	route := mux.NewRouter()

	// 	Route inisialisasi Folder Public
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/", home).Methods("GET")
	route.HandleFunc("/form-article", formProject).Methods("GET")
	route.HandleFunc("/contact", contact).Methods("GET")
	route.HandleFunc("/project-detail/{id}", projectDetail).Methods("GET")
	route.HandleFunc("/addProject", addProject).Methods("POST")

	fmt.Println("Server On Running in port 5000")
	http.ListenAndServe("localhost:5000", route)
}

type Project struct {
	Id        int
	Title     string
	DateStart string
	DateEnd   string
	Content   string
}

var projects = []Project{
	{
		Title:     "Aplikasi web dumbways",
		DateStart: "11 november 2022",
		DateEnd:   "12 desember 2022",
		Content:   "lorem ipsum dolor si amet",
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text-html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/index.html")

	if err != nil {
		w.Write([]byte("Massage : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}
func addProject(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	title := r.PostForm.Get("title")
	content := r.PostForm.Get("content")
	dateStart := r.PostForm.Get("date-start")
	dateEnd := r.PostForm.Get("date-end")

	var newProject = Project{
		Title:     title,
		Content:   content,
		DateStart: dateStart,
		DateEnd:   dateEnd,
	}
	projects = append(projects, newProject)

	fmt.Println(projects)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
func formProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text-html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/addproject.html")

	if err != nil {
		w.Write([]byte("Massage : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}
func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text-html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/contact.html")

	if err != nil {
		w.Write([]byte("Massage : " + err.Error()))
		return
	}

	tmpt.Execute(w, nil)
}
func projectDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "text/html; charset=utf-8")
	tmpt, err := template.ParseFiles("views/project-detail.html")

	if err != nil {
		w.Write([]byte("Message :" + err.Error()))
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var ProjectDetail = Project{}

	for index, data := range projects {
		if index == id {
			ProjectDetail = Project{
				Title:     data.Title,
				Content:   data.Content,
				DateStart: data.DateStart,
				DateEnd:   data.DateEnd,
			}
		}
	}

	dataDetail := map[string]interface{}{
		"Project": ProjectDetail,
	}

	tmpt.Execute(w, dataDetail)
}
