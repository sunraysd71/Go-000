package service

import (
	"mysql"
	xerrors "github.com/pkg/errors"
)

func ServiceQuerysql(s string) (string, error){
	v, err := mysql.MysqlQuery(s)
	if err != nil{
		return "",xerrors.Wrap(err, "Querying key ["+ s +"] failed.")
	}
	return v, nil
}

func ServiceAddsql(k string, v string) error {
	return xerrors.Wrap(mysql.MysqlAdd(k,v), "Adding key "+ k +" failed.")
}