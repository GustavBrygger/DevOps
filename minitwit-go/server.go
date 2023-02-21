package main

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/noirbizarre/gonja"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func ConnectDatabase() error {
	db, err := sql.Open("sqlite3", "./database/minitwit.db")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

type CookieInfo struct {
	Username string `json:"username"`
	UserID   int    `json:"user_id"`
}

type User struct {
	User_id  int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Pw_hash  string `json:"pw_hash"`
}

type Message struct {
	Message_id     int    `json:"message_id"`
	Author_id      int    `json:"author_id"`
	Text           string `json:"text"`
	Pub_date       int    `json:"pub_date"`
	Flagged        int    `json:"flagged"`
	Formatted_date string `json:"formatted_date"`
	Username       string `json:"username"`
	Gravatar       string `json:"gravatar"`
}

func flashMessage(c *gin.Context, message string) {
	session := sessions.Default(c)
	session.AddFlash(message)
	if err := session.Save(); err != nil {
		log.Printf("error in flashMessage saving session: %s", err)
	}
}

func public_timeline(c *gin.Context) {

	messages := make([]Message, 0)
	users := make([]User, 0)
	session := sessions.Default(c)
	userID := session.Get("user_id")
	username := session.Get("username")

	cookie_info := CookieInfo{}

	if userID != nil {
		cookie_info.UserID = userID.(int)
		cookie_info.Username = username.(string)

	}

	fmt.Println(userID)
	fmt.Println(username)

	rows, err := DB.Query("SELECT message.*, user.* from message, user where message.flagged = 0 and message.author_id = user.user_id order by message.pub_date desc LIMIT 30")

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	for rows.Next() {
		//var _id int
		msg := Message{}
		usr := User{}
		err := rows.Scan(&msg.Message_id, &msg.Author_id, &msg.Text, &msg.Pub_date, &msg.Flagged, &usr.User_id, &usr.Username, &usr.Email, &usr.Pw_hash)

		str := "Mon, 02 Jan 2006 15:04:05"
		msg.Formatted_date = time.Unix(int64(msg.Pub_date), 0).Format(str)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		msg.Username = usr.Username
		msg.Gravatar = geturl(usr.Email)

		messages = append(messages, msg)
		users = append(users, usr)
	}

	var tpl = gonja.Must(gonja.FromFile("template/timeline_go.html"))

	user, err := DB.Query("SELECT * FROM user WHERE username = ? LIMIT 1", userID)
	usr := User{}
	user.Scan(&usr.User_id, &usr.Username, &usr.Email, &usr.Pw_hash)

	req := struct {
		Endpoint string
	}{
		Endpoint: c.Request.URL.Path,
	}
	println(req.Endpoint)

	out, err := tpl.Execute(gonja.Context{
		"messages":     messages,
		"g":            cookie_info,
		"request":      req,
		"profile_user": usr,
	})
	if err != nil {
		fmt.Println(err)
	}

	c.Writer.WriteString(out)
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func login(c *gin.Context) {

	/*rows, err := DB.QueryRow("SELECT * FROM user WHERE username = ?", "hello")
	defer rows.Close()
	for rows.Next() {
		//var _id int
		usr := User{}
		err := rows.Scan(&usr.User_id, &usr.Username, &usr.Email, &usr.Pw_hash )
		fmt.Println(usr)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

	}*/

	session := sessions.Default(c)
	if userID := session.Get("user_id"); userID != nil {
		fmt.Println("already logged in")
		c.Redirect(http.StatusFound, "/")
		return
	}

	var error string
	usr := User{}

	//if c.Request.Method == http.MethodPost {
	//username := c.PostForm("username")
	username := "hello"
	user, err := DB.Query("SELECT * FROM user WHERE username = ? LIMIT 1", username)

	defer user.Close()

	for user.Next() {
		user.Scan(&usr.User_id, &usr.Username, &usr.Email, &usr.Pw_hash)
		break
	}

	if err != nil {
		fmt.Println(err)
	}
	if user == nil {
		fmt.Println("Invalid username")
	} else if !CheckPasswordHash("secret", usr.Pw_hash) {
		fmt.Println("Invalid password")
	} else {
		fmt.Println("We logged the fuck in")
		session.Set("flash", "You were logged in")
		session.Set("user_id", usr.User_id)
		session.Set("username", usr.Username)
		session.Save()
		c.Redirect(http.StatusFound, "/")
		return
	}
	//}

	c.JSON(http.StatusOK, gin.H{"error": error, "user": usr})
	//c.HTML(http.StatusOK, "login.html", gin.H{"error": error})
}

func register(c *gin.Context) {
	password := "secret"
	hash, _ := HashPassword(password)
	email := "hellohellocom"
	username := "hello"

	stmt, err := DB.Prepare(fmt.Sprintf("INSERT INTO user (username, email, pw_hash) VALUES('%s', '%s', '%s')", username, email, hash))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	// Execute the insert statement with the desired values
	result, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the ID of the inserted record
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Record inserted with ID:", id)

}

func geturl(email string) string {
	email_md5 := fmt.Sprintf("%s", md5.Sum([]byte(email)))
	hex_md5_email := hex.EncodeToString([]byte(email_md5))

	//format string
	url := fmt.Sprintf("https://www.gravatar.com/avatar/%s?d=identicon&s=48", hex_md5_email)
	return url
}

func add_message(c *gin.Context) {
	session := sessions.Default(c)
	if userID := session.Get("user_id"); userID == nil {
		c.Redirect(http.StatusFound, "/")
		return
	}

	cur_date := time.Now().Unix()
	text := c.PostForm("text")
	stmt, err := DB.Prepare(fmt.Sprintf("INSERT INTO message (author_id, text, pub_date, flagged) VALUES('%d', '%s', '%d', 0)", session.Get("user_id"), text, cur_date))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()

	// Execute the insert statement with the desired values
	result, err := stmt.Exec()
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the ID of the inserted record
	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Record inserted with ID:", id)

	c.Redirect(http.StatusFound, "/")
}

func follow_user(c *gin.Context){
	session := sessions.Default(c)
	if userID := session.Get("user_id"); userID == nil {
		c.Abort()
		return
	}

	username := c.Param("username")
	profile_user := User{}

	user, err := DB.Query("SELECT * FROM user WHERE username = ? LIMIT 1", username)

	if err != nil{
		fmt.Println(err)
	}

	for user.Next() {
		user.Scan(&profile_user.User_id, &profile_user.Username, &profile_user.Email, &profile_user.Pw_hash)
		break
	}

	user.Close()

	stmt, err := DB.Prepare(fmt.Sprintf("INSERT INTO follower (who_id, whom_id) values (%d, %d)", session.Get("user_id"), profile_user.User_id))
	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	
	// Execute the insert statement with the desired values
	result, err := stmt.Exec()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	flashMessage(c, fmt.Sprintf("You are now following %s", username))
	c.Redirect(http.StatusFound, "/")
}

func unfollow_user(c *gin.Context){
	session := sessions.Default(c)
	if userID := session.Get("user_id"); userID == nil {
		c.Abort()
		return
	}

	username := c.Param("username")
	profile_user := User{}

	user, err := DB.Query("SELECT * FROM user WHERE username = ? LIMIT 1", username)

	if err != nil{
		fmt.Println(err)
	}

	for user.Next() {
		user.Scan(&profile_user.User_id, &profile_user.Username, &profile_user.Email, &profile_user.Pw_hash)
		break
	}

	user.Close()

	stmt, err := DB.Prepare(fmt.Sprintf("DELETE FROM follower WHERE who_id = %d AND whom_id = %d", session.Get("user_id"), profile_user.User_id))
	
	if err != nil {
		fmt.Println(err)
		return
	}
	defer stmt.Close()
	
	// Execute the insert statement with the desired values
	result, err := stmt.Exec()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)

	flashMessage(c, fmt.Sprintf("You are no longer following %s", username))
	c.Redirect(http.StatusFound, "/")

}
/*
func user_timeline(c *gin.Context) {
	messages := make([]Message, 0)
	users := make([]User, 0)
	session := sessions.Default(c)
	userID := session.Get("user_id")
	session_username := session.Get("username")

	cookie_info := CookieInfo{}

	// get username
	username := c.Param("username")
	profile_user := User{}


	user, err := DB.Query("SELECT * FROM user WHERE username = ? LIMIT 1", username)
	if err != nil {
		c.Abort()
		return
	}
	
	for user.Next() {
		user.Scan(&profile_user.User_id, &profile_user.Username, &profile_user.Email, &profile_user.Pw_hash)
		break
	}
	user.Close()

	followed := struct {
		Followed bool
	}{ 
		Followed: false,
	}

	if userID := session.Get("user_id"); userID != nil {

		res, err  := DB.Query(`select 1 from follower where
		follower.who_id = ? and follower.whom_id = ?`, userID, profile_user.Username)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}


		for res.Next() {
			res.Scan(&followed.Followed)
			break
		}
		res.Close()
	}

	if userID != nil {
		cookie_info.UserID = userID.(int)
		cookie_info.Username = session_username.(string)

	}
	
	// get messages
	rows, err := DB.Query(`select message.*, user.* from message, user where
	user.user_id = message.author_id and user.user_id = ?
	order by message.pub_date desc limit 30`, profile_user.User_id)

	for rows.Next() {
		//var _id int
		msg := Message{}
		usr := User{}
		err := rows.Scan(&msg.Message_id, &msg.Author_id, &msg.Text, &msg.Pub_date, &msg.Flagged, &usr.User_id, &usr.Username, &usr.Email, &usr.Pw_hash)

		str := "Mon, 02 Jan 2006 15:04:05"
		msg.Formatted_date = time.Unix(int64(msg.Pub_date), 0).Format(str)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		msg.Username = usr.Username
		msg.Gravatar = geturl(usr.Email)

		messages = append(messages, msg)
		users = append(users, usr)
	}

	var tpl = gonja.Must(gonja.FromFile("template/timeline_go.html"))

	req := struct {
		Endpoint string
	}{
		Endpoint: "user_timeline",
	}

	out, err := tpl.Execute(gonja.Context{
		"messages":     messages,
		"g":            cookie_info,
		"request":      req,
		"profile_user": profile_user,
		"followed":     followed,
	})
	if err != nil {
		fmt.Println(err)
	}

	c.Writer.WriteString(out)

}*/

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLFiles("./template/timeline_go.html")
	r.Static("/static", "./static/")

	ConnectDatabase()

	r.GET("/public", public_timeline)
	r.GET("/", public_timeline)
	r.GET("/login", login)
	r.POST("/login", login)
	r.GET("/register", register)
	r.POST("/add_message", add_message)
	//r.GET("/:username", user_timeline)

	r.GET("/:username/follow", follow_user)
	r.GET("/:username/unfollow", unfollow_user)

	r.Run()
}
