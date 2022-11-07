package notepage

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

type note struct {
    noteName   string
    noteText string
    noteStatus string
    noteDelegation string
    noteSharedUsers string
}

func AddNewNote() {
	
	tmpl := template.Must(template.ParseFiles("create-note.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := note{
            noteName:   r.FormValue("Name"),
            noteText: r.FormValue("Text"),
            noteStatus: r.FormValue("Status"),
			noteDelegation: r.FormValue("Delegation"),
			noteSharedUsers: r.FormValue("sharedUsers"),
        }

        // do something with details
        fmt.Println(details.noteText)

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

		sqlQuery := `INSERT INTO notes (noteName, noteText, noteStatus, noteDelegation, noteSharedUsers) VALUES($1, $2, $3, $4, $5);`

		_, err = db.Exec(sqlQuery, details.noteName, details.noteText, details.noteStatus, details.noteDelegation, details.noteSharedUsers)
		if err != nil {
			log.Fatal(err)
		}
		

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}

func SelectNote() {
	
	tmpl := template.Must(template.ParseFiles("create-note.html"))

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method != http.MethodPost {
            tmpl.Execute(w, nil)
            return
        }

        details := note{
            noteName:   r.FormValue("Name"),
            noteText: r.FormValue("Text"),
            noteStatus: r.FormValue("Status"),
			noteDelegation: r.FormValue("Delegation"),
			noteSharedUsers: r.FormValue("sharedUsers"),
        }

        // do something with details
        fmt.Println(details.noteText)

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

		rows, err := db.Query("SELECT noteName, noteText, noteStatus, noteDelegation, noteSharedUsers FROM notes;")
		if err != nil {
   			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
    		var noteName string
			
    		if err := rows.Scan(&noteName); err != nil {
            	log.Fatal(err)
    		}
    		fmt.Println(noteName)
		}
		
		if err := rows.Err(); err != nil {
    		log.Fatal(err)
		}
		

        tmpl.Execute(w, struct{ Success bool }{true})
    })

    http.ListenAndServe(":8080", nil)
}

