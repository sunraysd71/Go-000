package mysql

import (
	"errors"
	"os"
	"bufio"
	"strings"
)

var rows = make(map[string]string, 10)

var ErrNoRows = errors.New("No rows found in mysql.")
var ErrCantWriteRows = errors.New("Can't write data to database.")
var ErrCantOpenDatabase = errors.New("Can't open database.")
var ErrCorruptedDatabase = errors.New("Database format corrupted.")

func MysqlQuery(s string) (string, error) {
	f, err := os.Open("./sqlfile.txt")
	if err != nil{
		return "", ErrCantOpenDatabase
	}
	sc := bufio.NewScanner(f)
	end := true
	for sc.Scan(){
		if s == strings.Split(sc.Text(), " ")[0]{
			end = false
			break
		}
	}
	if end {
		return "", ErrNoRows
	}
	if sc.Err() != nil{
		return "", sc.Err()
	}
	if len(strings.Split(sc.Text()," ")) == 2{
		return strings.Split(sc.Text()," ")[1], nil
	}else{
		return "", ErrCorruptedDatabase
	}
}

func MysqlAdd(k string, v string) error {
	if _, err := os.Stat("./sqlfile.txt"); err != nil{
		os.Create("./sqlfile.txt")
	}
	f, err := os.OpenFile("./sqlfile.txt", os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil{
		return ErrCantOpenDatabase
	}
	buf := k + " " + v + "\n"
	if _, err = f.Write([]byte(buf)); err != nil{
		return 	ErrCantWriteRows
	}
	defer f.Close()
	return nil
}
