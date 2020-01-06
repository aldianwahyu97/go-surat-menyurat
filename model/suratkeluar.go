package model

import (
	"fmt"
	"net/http"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

type suratkeluar struct{
	Id_suratkeluar string
	Nomorsurat string
	Tanggalkeluar string
	Institusipenerima string
	Penerima string
	Prihal string
	File string
	Keterangan string
}

var each_suratkeluar = suratkeluar{}
var datasuratkeluar []suratkeluar

func GetDataSuratKeluar(){
	db, err := connect()
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM suratkeluar")

	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	for rows.Next(){
		var err = rows.Scan(&each_suratkeluar.Id_suratkeluar,&each_suratkeluar.Nomorsurat,&each_suratkeluar.Tanggalkeluar,&each_suratkeluar.Institusipenerima,&each_suratkeluar.Penerima,&each_suratkeluar.Prihal,&each_suratkeluar.File,&each_suratkeluar.Keterangan)

		if err != nil{
			fmt.Println(err.Error())
			return
		}
		datasuratkeluar = append(datasuratkeluar,each_suratkeluar)
	}
	if err = rows.Err(); err != nil{
		fmt.Println(err.Error())
		return
	}
}

func HandleSuratKeluar(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		var tmpl = template.Must(template.New("suratkeluar").ParseFiles("views/surat_keluar.html"))
		var result, err_data = json.Marshal(datasuratkeluar)
		var err_srtkeluar = json.Unmarshal([]byte(result),&datasuratkeluar)
		jumlahdata := len(datasuratkeluar)

		if err_srtkeluar != nil{
			fmt.Println(err_srtkeluar.Error())
			return
		}

		if err_data != nil{
			http.Error(w,err_data.Error(),http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title" : "Data Surat Keluar",
			"Data" : datasuratkeluar,
			"jumlahdata" : jumlahdata,
		}

		var err = tmpl.Execute(w,data)
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		return
	}
	http.Error(w,"",http.StatusBadRequest)
}

func HandleTambahSuratKeluar(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
		var tmpl = template.Must(template.New("tambahsuratkeluar").ParseFiles("views/tambah_suratkeluar.html"))
		var err = tmpl.Execute(w,nil)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
		}
		return
	}
	http.Error(w,"",http.StatusBadRequest)
}

func AddSuratKeluar(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		db, err := connect()
		if err != nil{
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		var nomorsurat = r.FormValue("nomorsurat")
		var tanggalkeluar = r.FormValue("tanggalkeluar")
		var institusipenerima = r.FormValue("institusipenerima")
		var penerima = r.FormValue("penerima")
		var prihal = r.FormValue("prihal")
		var file = r.FormValue("file")
		var keterangan = r.FormValue("keterangan")

		rows, err := db.Query("INSERT INTO `suratkeluar`(`id_suratkeluar`, `nomorsurat`, `tanggalkeluar`, `institusipenerima`, `penerima`, `prihal`, `file`, `keterangan`) VALUES (null,?,?,?,?,?,?,?)",nomorsurat,tanggalkeluar,institusipenerima,penerima,prihal,file,keterangan)

		if err != nil{
			fmt.Println(err.Error())
			return
		}
		defer rows.Close()
	}
}
