package main

import (
	"fmt"
	"internalapi/dbconnector"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
	conn, err := dbconnector.Get()
	if err != nil {
		fmt.Fprintln(ctx, err.Error())
		return
	}
	query, err := Asset("queries/example/example.sql")
	if err != nil {
		fmt.Println(err.Error())
	}
	rows, err := conn.Query(string(query))

	if err != nil {
		fmt.Println("Deu ruim")
		return
	}
	for rows.Next() {
		data := ""
		rows.Scan(&data)
		fmt.Println(data)
	}
	rows.Close()
	return
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}
