package main

import (
	_ "github.com/cqspirit/GoWebSSH/routers"
	"github.com/astaxie/beego"
	"github.com/cqspirit/GoWebSSH/models"
	_ "github.com/cqspirit/GoWebSSH/filter"
)

func main() {
	models.CreateTable()
	beego.Run()
}

