package main

import "encoding/json"
import "net/http"
import "fmt"

// type student struct {
//     ID    string
//     Name  string
//     Grade int
// }

// var data = []student{
//     student{"E001", "ethan", 21},
//     student{"W001", "wick", 22},
//     student{"B001", "bourne", 23},
//     student{"B002", "bond", 23},
// }

type suratmasuk struct{
	Id_suratmasuk int
	Nomorsurat string
	Tanggalmasuk string
	// pengirim string
	// penerima string
	// prihal string
	// file string
	// keterangan string
}

var datasuratmasuk2 = []suratmasuk{
	// suratmasuk{1,"123/III/2020/DPRD-Kota","2020-01-02", "DPRD Kota Jakarta", "Chaerus Sulton","Undangan Rapat Koordinasi","surat.jpg","DIterima"},
	// suratmasuk{2,"123/III/2020/DPRD-Kota","2020-01-02", "DPRD Kota Jakarta", "Ilham Bintang","Undangan Rapat Koordinasi","surat.jpg","DIterima"},
	suratmasuk{1,"123/III/2020/DPRD-Kota","2020-01-02"},
	suratmasuk{2,"123/III/2020/DPRD-Kota","2020-01-02"},
}

func users(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var result, err = json.Marshal(datasuratmasuk2)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

// func user(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")

//     if r.Method == "POST" {
//         var id = r.FormValue("id")
//         var result []byte
//         var err error

//         for _, each := range datasuratmasuk2 {
//             if each.ID == id {
//                 result, err = json.Marshal(each)

//                 if err != nil {
//                     http.Error(w, err.Error(), http.StatusInternalServerError)
//                     return
//                 }

//                 w.Write(result)
//                 return
//             }
//         }

//         http.Error(w, "User not found", http.StatusBadRequest)
//         return
//     }

//     http.Error(w, "", http.StatusBadRequest)
// }

func main() {
	fmt.Println(datasuratmasuk2)
    http.HandleFunc("/users", users)
    // http.HandleFunc("/user", user)

    fmt.Println("starting web server at http://localhost:8081/")
    http.ListenAndServe(":8081", nil)
}