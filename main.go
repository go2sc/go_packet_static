package main

import (
	assetfs "github.com/elazarl/go-bindata-assetfs"
	"io"
	"log"
	"net/http"
	"os"
	"packetHello/static"
)

func helloWorld( w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "hello my bindata")
}

func favico(w http.ResponseWriter, r *http.Request){
	f, err := os.Open("data/hello.html")
	if err != nil{
		log.Panicln(err)
		return
	}

	_, err = io.Copy(w, f)
	if err != nil{
		log.Panicln(err)
		return
	}
}

func main(){

	//http.Handle("/static2/", http.StripPrefix("/static2/", http.FileServer(assetFS())))
	//http.HandleFunc("/hello", http.HandlerFunc(helloWorld))
	//http.Handle("/", http.FileServer(assetFS()))
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(assetFS())))
	//http.Handle("/", http.FileServer(static.AssetFS()))
	//http.Handle("/", http.StripPrefix("/", http.FileServer(static.AssetFS())))
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(AssetFS())))

	//files := assetfs.AssetFS{
	//	Asset:     static.Asset,
	//	AssetDir:  static.AssetDir,
	//	Prefix:    "static",
	//}

	// 打包命令：go-bindata -pkg static -o static/static.go static/
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(&assetfs.AssetFS{Asset: static.Asset, AssetDir: static.AssetDir, Prefix: "static"})))
	//http.Handle("/",
	//	http.FileServer())

	log.Println("listern:", 10080)
	err := http.ListenAndServe(":10080", nil)
	if err != nil{
		log.Println("err:", err)
		return
	}
}