package main

import (
	"net/http"
	"html/template" 
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"
)

func checkErr(err error) { 
	if err != nil {
		panic(err)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	 if r.Method == "GET" {
		t, _ := template.ParseFiles("template/mysql/create.gtpl")
		t.Execute(w, nil) 
		fmt.Println("create")
	 }else{
	 	db, err := sql.Open("mysql", "root:@/testgo?charset=utf8")
	 	checkErr(err)
	 	
	 	r.ParseMultipartForm(32 << 20)
	 	
	 	username := r.Form["username"]
	 	departname := r.Form["departname"]
	 	
	 	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	 	checkErr(err)
	 	
	 	res, err := stmt.Exec(username[0], departname[0], time.Now()) 
	 	checkErr(err)
	 	
		id, err := res.LastInsertId() 
		checkErr(err)
		
		fmt.Println(id)
	 	
	 	fmt.Println(username)
	 	
	 	fmt.Println(departname)
	 }
}

func read(w http.ResponseWriter, r *http.Request){
	db, err := sql.Open("mysql", "root:@/testgo?charset=utf8")
	checkErr(err)
	
	rows, err := db.Query("SELECT uid,username FROM userinfo")
	checkErr(err)
	
	var users []User
    var user User
    
	for rows.Next() {
		err = rows.Scan(&user.uid, &user.username)
		checkErr(err)
		users = append(users, user)
	}
	
	t, err := template.New("read.gtpl").ParseFiles("template/mysql/read.gtpl")
	fmt.Println(users)
	err = t.Execute(w, users)
	if err != nil {
		panic(err)
	}
}
