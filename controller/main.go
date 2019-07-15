package controller

import (
	"log"
	"net/http"
)

type Routes struct{}

func (r *Routes) Create() {
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(`
		yra! 
		⊂_ヽ
		　 ＼＼ service
		　　 ＼( ͡° ͜ʖ ͡°)
		　　　 >　⌒ヽ
		　　　/ 　 へ＼
		　　 /　　/　＼＼mojet
		　　 ﾚ　ノ　　 ヽ_つ
		　　/　/
		　 /　/|
		　(　(ヽ
		　|　|、＼normal'no
		　| 丿 ＼ ⌒)
		　| |　　) /
		ノ )　　Lﾉ
		(_／rabotat`))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
