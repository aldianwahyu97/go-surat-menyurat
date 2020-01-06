package model

import (
	"fmt"
	"net/http"
	"html/template"
	_ "github.com/go-sql-driver/mysql"
	"encoding/json"
)

type suratmasuk struct{
	Id_suratmasuk string
	Nomorsurat string
	Tanggalmasuk string
	Pengirim string
	Penerima string
	Prihal string
	File string
	Keterangan string
}

var each = suratmasuk{}
var datasuratmasuk []suratmasuk

func GetDataSuratMasuk(){
	db, err := connect()
	if err != nil{
		fmt.Println(err.Error())
		return 
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM suratmasuk")

	if err != nil{
		fmt.Println(err.Error())
		return 
	}
	defer rows.Close()

	for rows.Next(){
		var err = rows.Scan(&each.Id_suratmasuk,&each.Nomorsurat, &each.Tanggalmasuk, &each.Pengirim, &each.Penerima, &each.Prihal, &each.File, &each.Keterangan)

		if err != nil{
			fmt.Println(err.Error())
			return
		}
		datasuratmasuk = append(datasuratmasuk,each)

		// fmt.Println(datasuratmasuk)
	}

	if err = rows.Err(); err != nil{
		fmt.Println(err.Error())
		return
	}

	// fmt.Println(datasuratmasuk)
}

func HandleSuratMasuk(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{

		var tmpl = template.Must(template.New("suratmasuk").ParseFiles("views/surat_masuk.html"))
		var result, err_data = json.Marshal(datasuratmasuk)
		var err_srtmasuk = json.Unmarshal([]byte(result),&datasuratmasuk)
		jumlahdata := len(datasuratmasuk)

		if err_srtmasuk != nil{
			fmt.Println(err_srtmasuk.Error())
			return
		}

		if err_data != nil{
			http.Error(w,err_data.Error(),http.StatusInternalServerError)
			return
		}

		var data = map[string]interface{}{
			"title" : "Data Surat Masuk",
			"Data" : datasuratmasuk,
			"jumlahdata" : jumlahdata,
		}

		var err = tmpl.Execute(w,data)
		if err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}

		// w.Write(result)
        return
	}
	http.Error(w,"",http.StatusBadRequest)
}

func HandleTambahSuratMasuk(w http.ResponseWriter,r *http.Request){
	if r.Method == "GET"{
		var tmpl = template.Must(template.New("tambahsuratmasuk").ParseFiles("views/tambah_suratmasuk.html"))
		var err = tmpl.Execute(w,nil)
		if err != nil{
			http.Error(w, err.Error(),http.StatusInternalServerError)
		}
		return
	}
	http.Error(w,"",http.StatusBadRequest)
}

func AddSuratMasuk(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		var tmpl = template.Must(template.New("tambahsuratmasuk").ParseFiles("views/tambah_suratmasuk.html"))

		db, err := connect()
		if err != nil{
			fmt.Println(err.Error())
			return 
		}
		defer db.Close()

		var nomorsurat = r.FormValue("nomorsurat")
		var tanggalmasuk = r.FormValue("tanggalmasuk")
		var pengirim = r.FormValue("pengirim")
		var penerima = r.FormValue("penerima")
		var prihal = r.FormValue("prihal")
		var file = r.FormValue("file")
		var keterangan = r.FormValue("keterangan")

		rows, err := db.Query("INSERT INTO `suratmasuk`(`id_suratmasuk`, `nomorsurat`, `tanggalmasuk`, `pengirim`, `penerima`, `prihal`, `file`, `keterangan`) VALUES (null,?,?,?,?,?,?,?)",nomorsurat,tanggalmasuk,pengirim,penerima,prihal,file,keterangan)

		if err != nil{
			fmt.Println(err.Error())
			return 
		}
		defer rows.Close()

		// notif := "Berhasil Tambah Data"

		var data = map[string]string{
			"notif" : "Berhasil Menambahkan Data",
		}

		if err := tmpl.Execute(w,data); err != nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
		}
		return
	}

	http.Error(w,"",http.StatusBadRequest)
}