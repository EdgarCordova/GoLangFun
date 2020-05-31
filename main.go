package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"./datastore"
	"./repository"
	"./entities"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage endpoint Hit")
}

//#region TODO Controller
func getAllTodos(w http.ResponseWriter, r *http.Request) {
	var todoRepository repository.TodoRepository
	todos, _ := todoRepository.GetAll()
	json.NewEncoder(w).Encode(todos)
}

func getTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]

	var todoRepository repository.TodoRepository
	todos, err := todoRepository.GetById(id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return 
	}
	json.NewEncoder(w).Encode(todos)
}

func deleteTodoById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    id := vars["id"]
	
	var todoRepository repository.TodoRepository
	IsDeleted, err := todoRepository.DeleteById(&id)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return 
	}

	if IsDeleted == false {
		fmt.Fprintf(w, id + " was not deleted")
		return
	}

	fmt.Fprintf(w, id + " was successfully deleted")
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo entities.Todo

    err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	var todoRepository repository.TodoRepository

	isCreated := todoRepository.Create(&todo)
	if isCreated {
		fmt.Fprintf(w, "Todo item was created")
		return
	}

	fmt.Fprintf(w, "Something went wrong")
}

func updateTodoById(w http.ResponseWriter, r *http.Request) {
	var todo entities.Todo

    err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	vars := mux.Vars(r)
	id := vars["id"]
	

	var todoRepository repository.TodoRepository

	isCreated := todoRepository.UpdateById(&id, &todo)
	if isCreated {
		fmt.Fprintf(w, "Todo item was updated")
		return
	}

	fmt.Fprintf(w, "Something went wrong")
}	
//#endregion 

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)

	//#region todo Routes 
	myRouter.HandleFunc("/Todo", getAllTodos).Methods("GET")
	myRouter.HandleFunc("/Todo", createTodo).Methods("POST")
	myRouter.HandleFunc("/Todo/{id}", getTodoById).Methods("GET")
	myRouter.HandleFunc("/Todo/{id}", updateTodoById).Methods("POST")
	myRouter.HandleFunc("/Todo/{id}", deleteTodoById).Methods("DELETE")
	//#endregion 

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}


func main() {
	datastore.SetUpDatabase()
	handleRequests()
}
