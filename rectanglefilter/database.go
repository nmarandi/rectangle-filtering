package rectanglefilter

import (
	"database/sql"

	// sqlite3
	_ "github.com/mattn/go-sqlite3"
)

var db, nil = sql.Open("sqlite3", "./rectangle.db")

// DBInit function
func dbInit() {
	db.Exec("create table if not exists testTable (x integer,y integer,width integer,height integer,datetime text)")
}

// AddRectangleDb to database
func AddRectangleDb(rectangle Rectangle) {
	tx, _ := db.Begin()
	stmt, sterr := tx.Prepare("insert into testTable (x,y,width,height,datetime) values (?,?,?,?,DateTime('now'))")
	_, err := stmt.Exec(rectangle.X, rectangle.Y, rectangle.Width, rectangle.Height)
	print(err, sterr)
	tx.Commit()
}

// GetAllRectangleDb function
func GetAllRectangleDb() []Rectangle {
	rows, _ := db.Query("select * from testTable")
	var store = make([]Rectangle, 0)
	for rows.Next() {
		rect := Rectangle{}
		rows.Scan(&rect.X, &rect.Y, &rect.Width, &rect.Height, &rect.InsertTime)
		store = append(store, rect)
	}
	return store
}

// GetUniqueRectanglesDb function
func GetUniqueRectanglesDb() []Rectangle {
	rows, err := db.Query("select * from testTable group by x,y,width,height having count(*) = 1 order by datetime ASC")
	print(err)
	var store = make([]Rectangle, 0)
	for rows.Next() {
		rect := Rectangle{}
		rows.Scan(&rect.X, &rect.Y, &rect.Width, &rect.Height, &rect.InsertTime)
		store = append(store, rect)
	}
	return store
}
