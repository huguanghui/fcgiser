package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

var fv1Log = &ProcessURL{
	Process: make(map[string]*EleMethod),
}

type LogInfo struct {
	Ty   string `json:"type"`
	Time int    `json:"time"`
	Log  string `json:"log"`
}

type LogSt struct {
	ID      string  `json:"id"`
	Loginfo LogInfo `json:"loginfo"`
}

func init() {
	fv1Log.AddMethod("POST", &EleMethod{Opt: logPost})
	CmdRuner.Use("/api/fv1/log", fv1Log)
}

func logPost(resp http.ResponseWriter, req *http.Request) int {
	var res = Result{
		Reusult:     -1,
		Description: "Unknown",
	}
	fmt.Println(req.URL.RequestURI())
	body, err := ioutil.ReadAll(req.Body)
	if err == nil {
		var linfo LogSt
		json.Unmarshal(body, &linfo)
		dir, _ := os.Getwd()
		nowDate := time.Now()
		year, month, day := nowDate.Date()
		logfile := fmt.Sprintf("%s/log/%04d%02d%02d/%s.log", dir, year, month, day, linfo.ID)
		targetdir := filepath.Dir(logfile)
		// 文件夹不存在创建文件夹
		_, err := os.Stat(targetdir)
		if os.IsNotExist(err) {
			err := os.MkdirAll(targetdir, os.ModePerm)
			if err != nil {
				fmt.Println(err.Error())
			}
		}
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		// 2. 添加文件
		lg := fmt.Sprintf("[%d]-[%s] %s", linfo.Loginfo.Time, linfo.Loginfo.Ty, linfo.Loginfo.Log)
		f.Write([]byte(lg))
		f.Write([]byte{'\n'})
		res.Reusult = 0
		res.Description = "OK"
	}
	b, _ := json.Marshal(res)
	resp.Write(b)
	return 0
}
