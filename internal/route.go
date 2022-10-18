package internal

import (
	"encoding/json"
	"net/http"
	"time"
)

type application struct {
	Name      string    `json: "name"`
	Surname   string    `json: "surname"`
	CreatedAt string `json: "created_at"`
}

func Router() {
	http.HandleFunc("/",getApplication)
	http.ListenAndServe(":8000",nil)
}

func getApplication(w http.ResponseWriter, r *http.Request){
	query := r.URL.Query()
	name := query.Get("name")
	surname := query.Get("surname")
	app := application{Name: name, Surname: surname, CreatedAt: time.Now().Format("15:04:05 02-01-2006")}

	w.Header().Add("content-type", "application/json")
	jsonByte,_ := json.Marshal(app)
	w.Write(jsonByte)
}
