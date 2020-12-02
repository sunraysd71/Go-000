package main

import (
	"service"
	"fmt"
	errors "github.com/pkg/errors"
	"mysql"
)

func QueryKeyRequest(s string){
	v, err := service.ServiceQuerysql(s)
	if err != nil{
		switch errors.Cause(err){
			case mysql.ErrNoRows: fmt.Println("Opereration failed by ErrNoRows")
			case mysql.ErrCantOpenDatabase: fmt.Println("Opereration failed by ErrCantOpenDatabase")
			case mysql.ErrCorruptedDatabase: fmt.Println("Opereration failed by ErrCorruptedDatabase")
			default: fmt.Println("Opereration failed by ErrUnknown")
		}
			fmt.Print("Details:\n")
			fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
			fmt.Printf("stack trace:\n%+v\n", err)
			fmt.Println("####################End####################")
	}else{
		fmt.Println("Operation success")
		fmt.Printf("The value of key [%s] is [%s]\n",s, v)
		fmt.Println("####################End####################")
	}
}

func AddKeyRequest(k string,v string){
	if err := service.ServiceAddsql(k,v); err != nil{
		switch errors.Cause(err){
			case mysql.ErrCantWriteRows: fmt.Println("Opereration failed by ErrCantWriteRows")
			case mysql.ErrCantOpenDatabase: fmt.Println("Opereration failed by ErrCantOpenDatabase")
			default: fmt.Println("Opereration failed by ErrUnknown")
		}
		fmt.Print("Details:\n")
		fmt.Printf("original error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace:\n%+v\n", err)
		fmt.Println("####################End####################")
	}else{
		fmt.Println("Operation success")
		fmt.Printf("Key:Value [%s]:%s successfully added to the database.\n",k, v)
		fmt.Println("####################End####################")
	}
}

func main(){
	AddKeyRequest("abc","def")
	QueryKeyRequest("abc")
	QueryKeyRequest("abf")
}