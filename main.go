package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request) {
		fileName := r.URL.Query().Get("name")

		// 漏洞：直接使用用户输入的文件路径，存在路径遍历风险
		data, err := ioutil.ReadFile("/var/www/" + fileName)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Write(data)

		// 漏洞：直接使用用户输入的文件路径，存在路径遍历风险
		data, err = ioutil.ReadFile("/var/www/" + fileName)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Write(data)

		data, err = ioutil.ReadFile("/var/www1111/" + fileName)
		if err != nil {
			http.Error(w, "File not found", http.StatusNotFound)
			return
		}
		w.Write(data)
	})

	abs(2)
	http.ListenAndServe(":8080", nil)
}

func abs(x int) int {
	if x >= 0 {
		return x
	} else {
		return x
	}
}

// package main

// import (
// 	"flag"

// 	"github.com/nrbackback/a-dog-a-day/config"
// 	"github.com/nrbackback/a-dog-a-day/email"
// 	"github.com/nrbackback/a-dog-a-day/log"
// 	"github.com/nrbackback/a-dog-a-day/picture"
// 	"github.com/nrbackback/a-dog-a-day/runner"
// 	"github.com/nrbackback/a-dog-a-day/title"
// )

// func main() {
// 	log.Init()
// 	var configFile string
// 	flag.StringVar(&configFile, "config", "config.yml", "load config")
// 	flag.Parse()

// 	if err := config.Load(configFile); err != nil {
// 		panic(err)
// 	}
// 	if err := picture.Init(config.Config.Picture); err != nil {
// 		panic(err)
// 	}
// 	email.WeiboSender = email.EmailSender{
// 		Conf: config.Config.Email,
// 	}
// 	title.Init(config.Config.Title)
// 	runner.Start(config.Config.Runner)
// }
