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

func createReminderList(rows *sql.Rows) ([]note, error) {
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

		err := rows.Scan(&id, &text, &dateUnix)
		if err != nil {
			fmt.Print(err.Error())
			return nil, err
		}
		reminders = append(reminders,
			note{id, text, time.Unix(dateUnix, 0)})
	}
	return reminders, nil
}

func dash (length int) {
	for i:=0; i<length; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func displayReminders(reminders []note) {
	for _, reminder := range reminders {
		var nDash int
		if len(reminder.text) > len(reminder.creation.String()){
			nDash = len(reminder.text)
		} else {
			nDash = len(reminder.creation.String())
		}
		dash(nDash+2)
		fmt.Printf("|%s|\n", reminder.text)
		fmt.Printf("|%s|\n",
			reminder.creation.String())
		dash(nDash+2)
	}
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
	minDate := time.Now().Unix()-int64(time.Hour.Seconds())*12*7
	res, err := db.Query(
		"SELECT * FROM reminders WHERE creation > ?",
		minDate)
	reminders, err := createReminderList(res)
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}
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
	reminders, err := createReminderList(res)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
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
	if err != nil {
		return
	}
	displayReminders(myReminders)
}
