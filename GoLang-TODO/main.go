package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var db *mgo.Database
var rnd *renderer.Render

const (
	hostName       string = "localhost:27017"
	dbName         string = "demo_todo"
	collectionName string = "todo"
	port           string = ":9000"
)

type (
	todoModel struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		Title     string        `bson:"title"`
		Completed bool          `bson:"completed"`
		CreatedAt time.Time     `bson:"createdAt"`
	}
	todo struct {
		ID        string    `json:"id"`
		Title     string    `json:"title"`
		Completed bool      `json:"completed"`
		CreatedAt time.Time `json:"created_at"`
	}
)

func init() {
	rnd = renderer.New()

	sess, err := mgo.Dial(hostName)
	checkErr(err)

	sess.SetMode(mgo.Monotonic, true)
	db = sess.DB(dbName)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	err := rnd.Template(w, http.StatusOK, []string{"/static/home.tpl"}, nil)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Mount("/todo", todoHandler())

	servers := &http.Server{
		Addr:         port,
		Handler:      r,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		log.Println("Listening of port: ", port)
		err := servers.ListenAndServe()
		if err != nil {
			log.Printf("listen:%s\n", err)
		}
	}()

}

func todoHandler() http.Handler {
	rg := chi.NewRouter()
	rg.Group(func(r chi.Router) {
		r.Get("/", fetchTodos)
		r.Post("/", createTodo)
		r.Put("/{id}", updateTodo)
		r.Delete("/{id}", deleteTodo)
	})
	return rg
}

func fetchTodos(w http.ResponseWriter, req *http.Request) {

	todos := []todoModel{}
	if err := db.C(collectionName).Find(bson.M{}).All(&todos); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "failed to fetch todo",
			"error":   err,
		})
		return
	}

	todoList := []todo{}

	for _, t := range todos {
		todoList = append(todoList, todo{
			ID:        t.ID.Hex(),
			Title:     t.Title,
			Completed: t.Completed,
			CreatedAt: t.CreatedAt,
		})
	}
	rnd.JSON(w, http.StatusOK, renderer.M{
		"data": todoList,
	})

}

func createTodo(w http.ResponseWriter, req *http.Request) {
	var t todo

	if err := json.NewDecoder(req.Body).Decode(&t); err != nil {
		rnd.JSON(w, http.StatusProcessing, err)
		return
	}

	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The title is required",
		})
		return
	}

	tm := todoModel{
		ID:        bson.NewObjectId(),
		Title:     t.Title,
		Completed: false,
		CreatedAt: time.Now(),
	}

	if err := db.C(collectionName).Insert(&tm); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to save todo",
		})
		return
	}

	rnd.JSON(w, http.StatusCreated, renderer.M{
		"message": "todo created successfully",
		"todo_id": tm.ID.Hex(),
	})

}

func deleteTodo(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimSpace(chi.URLParam(req, "id"))

	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "This id is invalid",
		})
		return
	}

	if err := db.C(collectionName).RemoveId(bson.ObjectIdHex(id)); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "Failed to delete todo",
			"error":   err,
		})
		return
	}

	rnd.JSON(w, http.StatusOK, renderer.M{
		"message": "todo",
	})
}

func updateTodo(w http.ResponseWriter, req *http.Request) {
	id := strings.TrimSpace(chi.URLParam(req, "id"))

	if !bson.IsObjectIdHex(id) {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "id is invalid",
		})
		return
	}

	var t todo

	if err := json.NewDecoder(req.Body).Decode(&t); err != nil {
		rnd.Render(w, http.StatusProcessing, err)
	}

	if t.Title == "" {
		rnd.JSON(w, http.StatusBadRequest, renderer.M{
			"message": "The title field is required",
		})
		return
	}

	if err := db.C(collectionName).Update(
		bson.M{"_id": bson.ObjectIdHex(id)},
		bson.M{"title": t.Title, "completed": t.Completed},
	); err != nil {
		rnd.JSON(w, http.StatusProcessing, renderer.M{
			"message": "failed to update todo",
			"error":   err,
		})
	}

}
