package main

import (
	"fmt"
	"html/template" 
	"log"
	"net/http"
	"strings"
	"io"
	"crypto/md5"
	"time"
	"strconv"
	"os"
)


func sayhelloName(w http.ResponseWriter, r *http.Request) { 
	r.ParseForm() //解析参数,默认是不会解析的 
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息 
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, "")) 
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到 w 的是输出到客户端的 
}


func login(w http.ResponseWriter, r *http.Request) {
	 fmt.Println("method:", r.Method) //获取请求的方法 
	 if r.Method == "GET" {
	 	crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10)) 
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("template/login.gtpl")
		t.Execute(w, token) 
	 } else {
	 	//默认 情况下,Handler 里面是不会自动解析 form 的,必须显式的调用 r.ParseForm()后,你才能 对这个表单数据进行操作
	 	r.ParseMultipartForm(32 << 20) //有文件上传调这个就行，参数表示限制大小
	 	token := r.Form.Get("token")
		//请求的是登陆数据,那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"]) 
		fmt.Println("password:", r.Form["password"])
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
		template.HTMLEscape(w,[]byte("<br>"))
		template.HTMLEscape(w, []byte(token))
		
		
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./upload/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666) 
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		
		if token != "" {
			//验证 token 的合法性
			//todo 正确输出到客户端的方式，string类型
			template.HTMLEscape(w,[]byte ("token 正确"))
		} else {
			//不存在 token 报错
			template.HTMLEscape(w,[]byte ("token 错误"))
		}
	}
}

func main(){
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login) //设置访问的路由
	http.HandleFunc("/create", create) //设置访问的路由
	http.HandleFunc("/read", read) //设置访问的路由
	err := http.ListenAndServe(":9092", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}