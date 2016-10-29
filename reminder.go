package main

import (
	"fmt"
	"time"
<<<<<<< HEAD
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type note struct {
	id int
	text string
	creation time.Time
}
=======
	"github.com/clagraff/argparse"
	"os"
)

type Reminder struct {
	Id int
	Data string
	Creation Time
}

/*
* TODO:
*  √ Load reminders from DB
*  * Allow reminders to have specified allert dates
*  * Display latest reminders
*/

func loadAllReminders() ([]note ,error) {
	var reminders []note
	db, err := sql.Open("sqlite3", "reminders.db")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	res, err := db.Query("SELECT * FROM reminders")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer res.Close()
	for res.Next(){
		var id int
		var text string
		var dateUnix int64

		err = res.Scan(&id, &text, &dateUnix)
		reminders = append(reminders,
			note{id, text, time.Unix(dateUnix, 0)})
	}
	db.Close()

	return reminders, nil
}

func saveReminder(reminder note) {
	db, err := sql.Open("sqlite3", "reminders.db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Exec("INSERT INTO reminders VALUES (?, ?, ?)",
		nil, reminder.text, reminder.creation.Unix())
}

func main() {
	testRem := note{01, "This is my test note", time.Now()}
	saveReminder(testRem)
	myReminders, err := loadAllReminders()
	if err != nil{
		return
	}
	fmt.Println(myReminders)
}
