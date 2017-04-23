package main

import (
	"fmt"
	"internalapi/dbconnector"
	"log"

	"github.com/buaazp/fasthttprouter"
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
	rows, err := conn.Query("SELECT getdate()")

	if err != nil {
		fmt.Println("Deu ruim")
		return
	}
	for rows.Next() {
		data := ""
		rows.Scan(&data)
	}
	rows.Close()
	return
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}
