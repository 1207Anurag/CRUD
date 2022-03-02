//simple server which can handle HTTP request.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Creating a Rest Api that allow us to 'Create','Delete','Update' 'delete'
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
	fmt.Println("endpoint hit:homepage")
}

// func handlerequest(){
// 	http.HandleFunc("/",homePage)
// 	log.Fatal(http.ListenAndServe(":3000",nil))
// }
// func main(){
// 	handlerequest()
// }
type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

//defining an array of article
var Articles []Article

func main() {
	//For sending the post request..
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Desc: "This is a good article", Content: "Content of the Article"},
		Article{Id: "2", Title: "Hello2", Desc: "This is better article", Content: "Content of the second article"},
	}
	handleRequests()
}

//create an end point which will return all the article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	//task of this is to return newly populated Article variable encode in JSON format
	json.NewEncoder(w).Encode(Articles)
}
func updateUser(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, _ := strconv.Atoi((vars["id"]))
    
    var article Article
    reqBody, e := ioutil.ReadAll(r.Body)
    if e != nil {
        fmt.Fprintf(w, "Enter correct data")
    }
    json.Unmarshal(reqBody, &article)

    for i, usr := range Articles {
        if strconv.Atoi(usr.Id) == id {
            usr.Title = article.Title
            usr.Desc = article.Desc
            usr.Content= article.Content
            Articles = append(Articles[:i], usr)
            json.NewEncoder(w).Encode(usr)

        }
    }
}

// func handleRquest(){
// 	http.HandleFunc("/",homePage)
// 	http.HandleFunc("/articles",returnAllArticles)
// 	log.Fatal(http.ListenAndServe(":10000",nil))
// }

//In The next part we are going to update the API to use gorilla/mux

//Getting started With Routers...
// Build our own router

//updating the handleRequest

func handleRequests() {
	//created a new instance of mux router
	myRouter := mux.NewRouter().StrictSlash(true)

	//replaced http.HandleFunction with myRouter.handleFunc
	//we just want to call this function if the  incoming request is a post request
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article", CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article", DeleteArticlesMux).Methods("DELETE")
	myRouter.HandleFunc("/article", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", returnSingleArticleMux).Methods("GET")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
	//The only real difference is that we now have a gorilla/mux router which will allow us to easily do things such as retrieve path and query parameters later on in this tutorial.
}

//Path Variables
//what is we want to see only one article...
//using gorilla/mux-> we can add variables to our paths and then pick and choose what articles we want to return

func returnSingleArticleMux(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	Key := vars["id"]
	//Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range Articles {
		if article.Id == Key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

//Crud Operation

//create new Article  Post Request....

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	//get the body of the post request
	//unmarshal this into a new Article array
	//append this our article array
	reqBody, _ := ioutil.ReadAll(r.Body)
	// fmt.Fprintf(w,"%+v",string(reqBody))
	var article Article
	json.Unmarshal(reqBody, &article) //We are first converting the bytes of data into struct
	//update our global articles array to include our new article
	Articles = append(Articles, article)
	json.NewEncoder(w).Encode(article)
}

//Delete
//for deleting purpose we need to use the mux for choosing a specific id to delete

func DeleteArticlesMux(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	//loop through all the articles..
	for index, article := range Articles {
		if article.Id == id {
			//Update our Articles array to remove the article....
			Articles = append(Articles[:index], Articles[index+1])
		}
	}
}
