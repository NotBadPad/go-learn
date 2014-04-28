package main

func main() {
	"github.com/astaxie/beego"
}

func main() {
	log := beego.logs.NewLogger(10000)
	log.SetLogger("file", `{"filename":"buffetbox.log"}`)

	beego.Info("aaaaaaaaaaaaa")
}
