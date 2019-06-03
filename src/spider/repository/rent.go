package repository

import (
	"fmt"
	"log"
)

type Rent struct {
	Id     int
	Title  string
	Author string
	Url    string
}

func CheckExist(id string) bool {
	var count int
	rows, err := DB.Query("SELECT COUNT(Id) FROM Rent WHERE Id = ? LIMIT 1", id)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}

	if count > 0 {
		return true
	} else {
		return false
	}
}

func Insert(rent *Rent) bool {
	tx, err := DB.Begin()
	if err != nil {
		fmt.Println("tx fail")
		return false
	}

	stmt, err := tx.Prepare("INSERT INTO Rent (`Id`, `Title`, `Author`, `Url`) VALUES (?, ?, ?, ?)")
	if err != nil {
		fmt.Println("Prepare fail")
		return false
	}
	res, err := stmt.Exec(rent.Id, rent.Title, rent.Author, rent.Url)
	if err != nil {
		fmt.Println("Exec fail")
		return false
	}

	tx.Commit()

	fmt.Println(res.LastInsertId())

	return true
}
