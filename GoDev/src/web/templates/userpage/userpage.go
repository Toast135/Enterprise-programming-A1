package userpage

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = ""
)

type userStruct struct {
    userName string
	userPassword string
}

func AddNewUser() {
	
	tmpl := template.Must(template.ParseFiles("create-user.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := userStruct{
            userName:   r.FormValue("username"),
            userPassword: r.FormValue("password"),
        }

        // do something with details
        fmt.Println(details.userPassword)

		//Connect to the database
		const dbname = "enterprisenotes"
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		// Ping the database for connectivity
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		sqlQuery := `INSERT INTO users (userName, userPassword) VALUES($1, $2);`

		_, err = db.Exec(sqlQuery, details.userName, details.userPassword)
		if err != nil {
			log.Fatal(err)
		}
		

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}

func SelectUser() {
	
	tmpl := template.Must(template.ParseFiles("create-user.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := userStruct{
            userName:   r.FormValue("username"),
            userPassword: r.FormValue("password"),
        }

        // do something with details
		fmt.Println(details.userName)

		//Connect to the database
		const dbname = "enterprisenotes"
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

		// Ping the database for connectivity
		db, err := sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()
		err = db.Ping()
		if err != nil {
			log.Fatal(err)
		}

		rows, err := db.Query("SELECT userName FROM users;")
		if err != nil {
   			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
    		var userName string
			
    		if err := rows.Scan(&userName); err != nil {
            	log.Fatal(err)
    		}
    		fmt.Println(userName)
		}
		
		if err := rows.Err(); err != nil {
    		log.Fatal(err)
		}
		

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}

