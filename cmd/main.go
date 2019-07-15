package main

import (
	"fmt"
	"log"

	"github.com/tkachenkom/tm-stream/controller"

	"github.com/tkachenkom/tm-stream/cfg"
)

func main() {
	fmt.Println(`
	yra! 
	⊂_ヽ
	　 ＼＼ service
	　　 ＼( ͡° ͜ʖ ͡°)
	　　　 >　⌒ヽ
	　　　/ 　 へ＼
	　　 /　　/　＼＼nachal
	　　 ﾚ　ノ　　 ヽ_つ
	　　/　/
	　 /　/|
	　(　(ヽ
	　|　|、＼svoy
	　| 丿 ＼ ⌒)
	　| |　　) /
	ノ )　　Lﾉ
	(_／raboty`)

	_, err := cfg.New()
	if err != nil {
		log.Fatal(err)
	}

	c := &controller.Routes{}
	c.Create()
}
