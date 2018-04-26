package main

import (
    "net/http"
    "database/sql"
	"encoding/json"
    "fmt"
	"strings"
	"unicode"
    _ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func chaxun(w http.ResponseWriter, r *http.Request) {

  keys := r.URL.Query()
  
  w.Header().Set("Content-Type", "application/json;charset:utf-8;")
  biaoge := keys["neirong"][0]
  yaopinname := strings.TrimFunc(keys["yaopinname"][0],unicode.IsSpace)
    
  if biaoge == "zhongbiao" { 
    shengchanqiye := strings.TrimFunc(keys["shengchanqiye"][0],unicode.IsSpace)
	type Rst struct {
	  Tym string
      Jx string
      Gg string
      Spm string
      Jg string
      Jgsm string
	  Lsxj string
      Bz string
      Cglx string
      Sfjy string
      Scqy string
      Sy string
	}
	
	if yaopinname=="" && shengchanqiye=="" {
	  rst := []Rst {{"你未输入任何有效的字符串", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}}
	  r, _ := json.Marshal(rst)
	  fmt.Fprintf(w, string(r))
	} else {
	rsts := []Rst{}
    db, err := sql.Open("sqlite3", "./db/yiyaoshuju.db")
    checkErr(err)
    rows, err := db.Query("SELECT tym, jx, gg, spm, jg, jgsm, lsxj, bz, cglx, sfjy, scqy, sy FROM yaopinmulu WHERE tym LIKE ? AND scqy LIKE ?", 
	                      "%" + yaopinname + "%", "%" + shengchanqiye + "%")
    checkErr(err)
	defer rows.Close()
    for rows.Next() {
	  rst := Rst{}
      err = rows.Scan(&rst.Tym, &rst.Jx, &rst.Gg, &rst.Spm, &rst.Jg, &rst.Jgsm, &rst.Lsxj, &rst.Bz, &rst.Cglx, &rst.Sfjy, &rst.Scqy, &rst.Sy)
      checkErr(err)
	  rsts = append(rsts, rst)
    }
	r, _ := json.Marshal(rsts)
	fmt.Fprintf(w, string(r))
  }  
  }
  if biaoge == "yibao" { 
	type Rst struct {
	  Fldm string
      Fl1 string
      Fl3 string
      Yblx string
      Bh string
      Ypmc string
      Jx string
      Bz string
	}
	if yaopinname=="" {
	  rst := []Rst {{"你未输入任何有效的字符串", " ", " ", " ", " ", " ", " ", " "}}
	  r, _ := json.Marshal(rst)
	  fmt.Fprintf(w, string(r))
	} else {
	rsts := []Rst{}
    db, err := sql.Open("sqlite3", "./db/yiyaoshuju.db")
    checkErr(err)
    rows, err := db.Query("SELECT fldm, fl1, fl3, yblx, bh, ypmc, jx, bz FROM yibaomulu WHERE ypmc LIKE ?", "%" + yaopinname + "%")
    checkErr(err)
	defer rows.Close()
    for rows.Next() {
	  rst := Rst{}
      err = rows.Scan(&rst.Fldm, &rst.Fl1, &rst.Fl3, &rst.Yblx, &rst.Bh, &rst.Ypmc, &rst.Jx, &rst.Bz)
      checkErr(err)
	  rsts = append(rsts, rst)
    }
	r, _ := json.Marshal(rsts)
	fmt.Fprintf(w, string(r))
  }  
}
}
  
  if biaoge == "jcbc" { 
    shengchanqiye := strings.TrimFunc(keys["shengchanqiye"][0],unicode.IsSpace)
	type Rst struct {
	  Tym string
      Jx string
      Gg string
      Bz string
      Bzcl string
      Sccj string
	  Jg string
      Jgsm string
      Zdjyjg string
      Pjjyjg string
      Cglx string
      Xmmc string
	}
	
	if yaopinname=="" && shengchanqiye=="" {
	  rst := []Rst {{"你未输入任何有效的字符串", " ", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}}
	  r, _ := json.Marshal(rst)
	  fmt.Fprintf(w, string(r))
	} else {
	rsts := []Rst{}
    db, err := sql.Open("sqlite3", "./db/yiyaoshuju.db")
    checkErr(err)
    rows, err := db.Query("SELECT tym, jx, gg, bz, bzcl, sccj, jg, jgsm, zdjyjg, pjjyjg, cglx, xmmc FROM jicengbuchongmulu WHERE tym LIKE ? AND sccj LIKE ?", 
	                      "%" + yaopinname + "%", "%" + shengchanqiye + "%")
    checkErr(err)
	defer rows.Close()
    for rows.Next() {
	  rst := Rst{}
      err = rows.Scan(&rst.Tym, &rst.Jx, &rst.Gg, &rst.Bz, &rst.Bzcl, &rst.Sccj, &rst.Jg, &rst.Jgsm, &rst.Zdjyjg, &rst.Pjjyjg, &rst.Cglx, &rst.Xmmc)
      checkErr(err)
	  rsts = append(rsts, rst)
    }
	r, _ := json.Marshal(rsts)
	fmt.Fprintf(w, string(r))
  }  
  }

func main() {
  server := http.Server{
	  Addr: ":80",
  }
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
  })
  http.HandleFunc("/chaxun", chaxun)
  server.ListenAndServe()
}
