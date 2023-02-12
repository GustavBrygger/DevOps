package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
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

func main() {
	r := gin.Default()
	
	r.LoadHTMLFiles("./template/timeline.html")

	ConnectDatabase()

	r.GET("/public", public_timeline)
	
	r.Run()
}

