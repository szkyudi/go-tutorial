package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"szkyudi/go-tutorial/controller"
	"szkyudi/go-tutorial/model"
)

var todoModel = model.NewTodoModel()
var todoController = controller.NewTodoController(todoModel)
var router = controller.NewRouter(todoController)

func migrate() {
	sql := `INSERT INTO todos(id, name, status) VALUES
			('00000000000000000000000000', '買い物', '作業中'),
			('00000000000000000000000001', '皿洗い', '完了');
		`

	_, err := model.Db.Exec(sql)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Migration is success!")
}

func main() {
	f := flag.String("option", "", "migrate database or not")
	flag.Parse()
	if *f == "migrate" {
		migrate()
	}
	http.HandleFunc("/fetch-todos", router.FetchTodos)
	http.HandleFunc("/add-todo", router.AddTodo)
	http.HandleFunc("/change-todo", router.ChangeTodo)
	http.HandleFunc("/delete-todo", router.DeleteTodo)
	http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("APP_PORT")), nil)
}
