package main

import (
	"fmt"
	"net/http"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"ApsGolang_ArsipSurat/model"
)

func HandleIndex(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		var tmpl = template.Must(template.New("index").ParseFiles("views/index.html"))
		var err = tmpl.Execute(w,nil)
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		return
	}
	http.Error(w,"",http.StatusBadRequest)
}

func HandleLogin(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		var tmpl = template.Must(template.New("login").ParseFiles("views/login.html"))
		var err = tmpl.Execute(w,nil)
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		return
	}
	http.Error(w,"",http.StatusBadRequest)
}

func main(){
	model.GetDataSuratMasuk()
	model.GetDataSuratKeluar()
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("assets"))))

	// Routing 
	http.HandleFunc("/",HandleIndex)
	http.HandleFunc("/login",HandleLogin)
	http.HandleFunc("/suratmasuk",model.HandleSuratMasuk)
	http.HandleFunc("/suratkeluar",model.HandleSuratKeluar)
	http.HandleFunc("/suratmasuk/tambah",model.HandleTambahSuratMasuk)
	http.HandleFunc("/suratkeluar/tambah",model.HandleTambahSuratKeluar)
	http.HandleFunc("/suratmasuk/tambah/add",model.AddSuratMasuk)
	http.HandleFunc("/suratkeluar/tambah/add",model.AddSuratKeluar)

	fmt.Println("Memulai Server pada Port: 9000...")
	http.ListenAndServe(":9000",nil)
}