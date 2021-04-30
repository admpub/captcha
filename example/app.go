package main

import (
	"fmt"
	"image/png"
	"log"
	"net/http"

	"github.com/admpub/captcha"
)

func generateCaptchaHandler(w http.ResponseWriter, r *http.Request) {
	cp := captcha.NewCaptcha(125, 40, 4)
	cp.SetFontPath("bin/")
	cp.SetFontName("Courier New")
	cp.SetMode(0) // 设置1为数学公式
	code, img := cp.OutPut()
	//备注：code 可以根据情况存储到session，并在使用时取出验证

	fmt.Println(code)              // DEBUG
	w.Header().Set("X-Code", code) // DEBUG

	w.Header().Set("Content-Type", "image/png; charset=utf-8")

	png.Encode(w, img)
}

func main() {
	http.HandleFunc("/captcha", generateCaptchaHandler)

	if err := http.ListenAndServe("localhost:8080", nil); err != nil {
		log.Fatal(err)
	}
}
