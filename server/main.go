package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//  NAME        TYPE                    REP

type User struct {
	ID          primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string                 `json:"name,omitempty"`
	Hash  	    string                 `json:"hash,omitempty"`
}

type Comment struct {
	ID          primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Message 	string                 `json:"message,omitempty"`
	Upvotes     []User
	Downvotes   []User
}

type Message struct {
	ID 			primitive.ObjectID     `json:"_id,omitempty" bson:"_id,omitempty"`
	Message 	string                 `json:"message,omitempty"`
	UserCode    string                 `json:"usercode,omitempty"`
	User        User                   `json:"user,omitempty"`
	Upvotes     []User				   `json:"upvotes,omitempty"`    // these may just be strings for the username
	Downvotes   []User				   `json:"downvotes,omitempty"` 
	Comments    []Message			   `json:"comments,omitempty"`
	// ^ Recursive messages
}

func NewUser(username string, hash string) (user User) {
	user.Name = username
	user.Hash = hash
	return
}

func (message *Message) ToJSONString() (string, error) {
	json, err := json.Marshal(message)
	if err != nil { return "", err }
	return string(json), nil
}

func (message *Message) ResolveUser() {
	// Todo: build this
	fmt.Printf("Resolving user for %s\n", message.UserCode)

	message.User = NewUser("Demo-User", "0xB00B1E55")
}

type Wrapper struct {
	messages []Message
}

func (wrapper *Wrapper) PostMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var message Message
	_ = json.NewDecoder(r.Body).Decode(&message)
	messageJSONString, err := message.ToJSONString()
	if err != nil { return }
	fmt.Println(message, messageJSONString)
	message.ResolveUser()
	//insertOneTask(task)
	//json.NewEncoder(w).Encode(task)
	fmt.Fprint(w, "Success")
}

func (wrapper *Wrapper) GetMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(wrapper.messages)
}

func ConfigureRouter() *mux.Router {

	router := mux.NewRouter()
	wrapper := Wrapper{}

	router.HandleFunc("/api/post_message", wrapper.PostMessage).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/post_message", wrapper.GetMessages).Methods("GET", "OPTIONS")

	// Host the react app build
	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./public/"))))

	return router
}

func main() {
	r := ConfigureRouter()
	// fs := http.FileServer(http.Dir("build"))
	// http.Handle("/", fs)



	fmt.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
