package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"time"
)

type note struct {
	id       int
	text     string
	creation time.Time
}


/*
* TODO:
*  √ Load reminders from DB
*  * Allow reminders to have specified alert dates
*  * Display latest reminders
 */

func createReminderList(rows *sql.Rows) []note {
	/*
	* Creates a note list from the database
	* The query must have already been performed
	*/
	var reminders []note
	defer rows.Close()
	for rows.Next() {
		var id int
		var text string
		var dateUnix int64

		rows.Scan(&id, &text, &dateUnix)
		reminders = append(reminders,
			note{id, text, time.Unix(dateUnix, 0)})
	}
	return reminders
}

func loadToPeriod(/*limit time.Time*/) ([]note, error) {
	/*
	* TODO add reminder deadlines
	*/
	db, err := sql.Open("sqlite3", "reminders.db")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	res, err := db.Query("SELECT * FROM reminders WHERE creation > ?", time.Now().Unix(), limit.Unix())
	reminders := createReminderList(res)
	db.Close()
	return reminders, nil
}

func loadAllReminders() ([]note, error) {
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
	reminders := createReminderList(res)
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
	if err != nil {
		return
	}
	fmt.Println(myReminders)
}
