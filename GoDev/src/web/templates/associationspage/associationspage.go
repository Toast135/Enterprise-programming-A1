package associationspage

import (
	"fmt"
	"html/template"
	"net/http"
	"database/sql"
	"log"
	"strconv"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = ""
)

type association struct {
    userID string
	noteID string
	associationPerm string
}

func AddNewAssociation() {
	
	tmpl := template.Must(template.ParseFiles("create-association.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := association{
            userID:   r.FormValue("user" ),
            noteID: r.FormValue("note"),
            associationPerm: r.FormValue("permission"),
        }

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

		sqlQuery := `INSERT INTO associations (UserID, NoteID, associationPerm) VALUES($1, $2, $3);`
		
		//convert strings from form to int
		userIDConv, err := strconv.Atoi(details.userID)
		if err != nil {
			log.Fatal(err)
		}

		noteIDConv, err := strconv.Atoi(details.noteID)
		if err != nil {
			log.Fatal(err)
		}

		//add data to database
		_, err = db.Exec(sqlQuery, userIDConv, noteIDConv, details.associationPerm)
		if err != nil {
			log.Fatal(err)
		}
		

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}

func SelectAssociation() {
	
	tmpl := template.Must(template.ParseFiles("create-note.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := association{
            userID:   r.FormValue("user" ),
            noteID: r.FormValue("note"),
            associationPerm: r.FormValue("permission"),
        }

        // do something with details
        fmt.Println(details.userID)

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

		rows, err := db.Query("SELECT UserID, NoteID, associationPerm FROM notes;")
		if err != nil {
   			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
    		var associationPerm string
			
    		if err := rows.Scan(&associationPerm); err != nil {
            	log.Fatal(err)
    		}
    		fmt.Println(associationPerm)
		}
		
		if err := rows.Err(); err != nil {
    		log.Fatal(err)
		}
		

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}