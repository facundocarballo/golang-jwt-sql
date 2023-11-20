package types

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/facundocarballo/golang-mysql-connection/crypto"
	"github.com/facundocarballo/golang-mysql-connection/db"
)

type Task struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Owner       int    `json:"owner"`
}

func BodyToTask(body []byte) *Task {
	if len(body) == 0 {
		return nil
	}

	var task Task
	err := json.Unmarshal(body, &task)
	if err != nil {
		return nil
	}

	return &task
}

func CreateTask(
	w http.ResponseWriter,
	r *http.Request,
	database *sql.DB,
) bool {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the body of request.", http.StatusBadRequest)
		return false
	}
	defer r.Body.Close()

	task := BodyToTask(body)
	if task == nil {
		http.Error(w, "Error wrapping the body to task.", http.StatusBadRequest)
		return false
	}

	tokenString := crypto.GetJWTFromRequest(w, r)
	if tokenString == nil {
		return false
	}

	id := crypto.GetIdFromJWT(*tokenString)
	if id == nil {
		http.Error(w, "Error JWT not Valid", http.StatusBadRequest)
		return false
	}

	_, err = database.Exec(
		db.INSERT_TASK_STATEMENT,
		task.Name,
		task.Description,
		*id,
	)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating the task in the database. " + err.Error()))
		return false
	}

	resData := ResponseData{
		Message: "SUCCESFULL POST REQUEST",
	}
	resJSON := GetResponseDataJSON(resData)

	if resJSON == nil {
		http.Error(w, "Error converting the response data to JSON. ", http.StatusInternalServerError)
		return false
	}

	w.WriteHeader(http.StatusOK)
	w.Write(*resJSON)

	return true
}

func GetTasks(
	w http.ResponseWriter,
	r *http.Request,
	database *sql.DB,
) bool {
	tokenString := crypto.GetJWTFromRequest(w, r)
	if tokenString == nil {
		return false
	}

	id := crypto.GetIdFromJWT(*tokenString)
	if id == nil {
		http.Error(w, "Error JWT not Valid", http.StatusBadRequest)
		return false
	}

	rows, err := database.Query("SELECT id, name, description, owner FROM Task WHERE owner = (?)", id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Iterate Rows
	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.Owner)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return false
		}
		tasks = append(tasks, task)
	}

	// Check Error on Rows
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return false
	}

	// Send response to the client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)

	return true
}

func HandleTask(w http.ResponseWriter, r *http.Request, database *sql.DB) {
	if r.Method == http.MethodPost {
		CreateTask(w, r, database)
		return
	}

	if r.Method == http.MethodGet {
		GetTasks(w, r, database)
		return
	}

	http.Error(w, "Method not allowed to /user", http.StatusMethodNotAllowed)
}
