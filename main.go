package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct{
	ID int `json:"id"`
	Nama string `json:"nama"`
	Deskripsi string `json:"deskripsi"`
} 

var category = []Category{
	{ID: 1, Nama: "Beach Club", Deskripsi: "Tempat hiburan tepi pantai dengan musik dan minuman"},
	{ID: 2, Nama: "Rentals", Deskripsi: "Penyewaan alat olahraga air dan perlengkapan pantai"},
	{ID: 3, Nama: "Guide", Deskripsi: "Pemandu wisata lokal untuk eksplorasi pantai dan sekitarnya"},
	{ID: 4, Nama: "Mall", Deskripsi: "Pusat perbelanjaan modern dengan beragam toko, restoran, dan hiburan keluarga"},
	{ID: 5, Nama: "Villa", Deskripsi: "Villa mewah pribadi dengan fasilitas lengkap dan pemandangan pantai eksklusif"},
}

//ini untuk category get by id
func getCategoryById(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path,"/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w,"Invalid kategori ID", http.StatusBadRequest)
		return
	}

	for _, p := range category{
		if p.ID == id{
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Kategori Tidak Ada", http.StatusNotFound)
}

func getAllCategory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(category)
}
func createCategory(w http.ResponseWriter, r *http.Request){
	var categoryBaru Category
	err := json.NewDecoder(r.Body,).Decode(&categoryBaru)
	if err != nil {
		http.Error(w, "Invalid Request",http.StatusBadRequest)
		return
	}

	categoryBaru.ID = len(category) +1
	category = append(category, categoryBaru)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(categoryBaru)
}
func updateCategory(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/category/")
	id, err := strconv.Atoi(idStr)
	if err!=nil{
		http.Error(w,"Invalid Kategori Id", http.StatusBadRequest)
		return
	}
	//Buat Body
	var categoryBaru Category
	err = json.NewDecoder(r.Body,).Decode(&categoryBaru)
	if err != nil {
		http.Error(w, "Invalid Request",http.StatusBadRequest)
		return
	}

	//Looping
	for i := range category {
		if category[i].ID == id{
		//set id agar id menu sebelumnya tidak berubah jadi 0
		categoryBaru.ID = id
		category[i] = categoryBaru
		w.Header().Set("Content-Type","application/json")
		json.NewEncoder(w).Encode(categoryBaru)
		return
		}
	}
	http.Error(w,"Belum Ada Kategori Tersebut", http.StatusNotFound)
}
func deleteCategory(w http.ResponseWriter, r *http.Request){
	idStr := strings.TrimPrefix(r.URL.Path, "/category/")

	id, err := strconv.Atoi(idStr)
	if err!=nil{
		http.Error(w,"Invalid Kategori Id", http.StatusBadRequest)
		return
	}
	
	for i, p:= range category{
		if p.ID == id{
			category = append(category[:i], category[i+1:]...)
			w.Header().Set("Content-Type","application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message":"Succesfully deleted",
			})
			return
		}
	}
	http.Error(w,"Belum Ada Kategori Tersebut", http.StatusNotFound)
}


func main (){
	http.HandleFunc("/category/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoryById(w, r)
		} else if r.Method == "PUT" {
			updateCategory(w, r)
		} else if r.Method == "DELETE" {
			deleteCategory(w, r)
		}
	})

	http.HandleFunc("/category", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getAllCategory(w, r)
		} else if r.Method == "POST" {
			createCategory(w, r)
		}
	})

	fmt.Println("Start listening on Port:8080")
	err:=http.ListenAndServe(":8080", nil)
	if err!= nil{
		fmt.Println("Port Error:", err)
	}
}