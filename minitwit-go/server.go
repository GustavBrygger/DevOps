package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"log"
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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

type User struct {
	User_id         int `json:"id"`
	Username      string `json:"username"`
	Email    string `json:"email"`
	Pw_hash string `json:"pw_hash"`
}

type Message struct {
	Message_id         int `json:"message_id"`
	Author_id      int `json:"author_id"`
	Text    string `json:"text"`
	Pub_date int `json:"pub_date"`
	Flagged int `json:"flagged"`
}

func public_timeline(c *gin.Context){
	//var ids []int

	messages := make([]Message, 0);
	users := make([]User, 0);

	rows, err := DB.Query("SELECT message.*, user.* from message, user where message.flagged = 0 and message.author_id = user.user_id order by message.pub_date desc LIMIT 30")
	
	if err != nil{
		log.Fatal(err)
	}
	//fmt.Println(rows)
	
	defer rows.Close()
	for rows.Next() {
		//var _id int
		msg := Message{}
		usr := User{}
		err := rows.Scan(&msg.Message_id, &msg.Author_id, &msg.Text, &msg.Pub_date, &msg.Flagged, &usr.User_id, &usr.Username, &usr.Email, &usr.Pw_hash )

		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		messages = append(messages, msg)
		users = append(users, usr)
	}

	c.JSON(http.StatusOK, gin.H{
		"messages":messages,
		"users":users,
	})

	/*c.HTML(http.StatusOK, "timeline.html", gin.H{
		"messages": messages,
	})*/

	//var tpl = gonja.Must(gonja.FromFile("timeline.html"))*/

	
    /*return render_template('timeline.html', messages=query_db('''
        select message.*, user.* from message, user
        where message.flagged = 0 and message.author_id = user.user_id
        order by message.pub_date desc limit ?''', [PER_PAGE]))*/
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
		fmt.Println("WHAT JHJKHKJ")
		c.Redirect(http.StatusFound, "/timeline")
		return
	}

	fmt.Println(session);
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
		
		fmt.Println(usr)
		if err != nil{
			log.Fatal(err);
		}
		if user == nil {
			fmt.Println("Invalid username");
		} else if !CheckPasswordHash("secret", usr.Pw_hash) {
			fmt.Println("Invalid password");
		} else {
			session.Set("flash", "You were logged in")
			session.Set("user_id", usr.User_id)
			session.Save()
			fmt.Println(":)")
			c.Redirect(http.StatusFound, "/public")
			return
		}
	//}

	c.JSON(http.StatusOK, gin.H{"error":error, "user": usr});
	//c.HTML(http.StatusOK, "login.html", gin.H{"error": error})
}

func register(c *gin.Context){
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

func main() {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.LoadHTMLFiles("./template/timeline.html")

	ConnectDatabase()

	r.GET("/public", public_timeline)
	r.GET("/login", login)
	r.POST("/login", login)
	r.GET("/register", register)

	r.Run()
}

