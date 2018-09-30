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
  
  switch biaobe {
  case "zhongbiao": 
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
			defer db.Close()
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
    case "yibao": 
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
			defer db.Close()
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
    case "jcbc":
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
			defer db.Close()
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
    case "factoryh":
	    factory := strings.TrimFunc(keys["factory"][0],unicode.IsSpace)
		city := strings.TrimFunc(keys["city"][0],unicode.IsSpace)
		county := strings.TrimFunc(keys["county"][0],unicode.IsSpace)
		variety := strings.TrimFunc(keys["variety"][0],unicode.IsSpace)
		hospital := strings.TrimFunc(keys["hospital"][0],unicode.IsSpace)
		type Rst struct {
		  Factory string
		  City string
		  County string
		  Hospital string
		  Variety string
		  Name string
		  Tel string
		  Remark string
		}
		
		if yaopinname=="" && factory=="" && city=="" && county=="" & hospital =="" {
		    rst := []Rst {{"你未输入任何有效的字符串", " ", " ", " ", " ", " ", " ", " "}}
		    r, _ := json.Marshal(rst)
		    fmt.Fprintf(w, string(r))
		} else {
			rsts := []Rst{}
			db, err := sql.Open("sqlite3", "./db/yiyaoshuju.db")
			defer db.Close()
			checkErr(err)
			rows, err := db.Query("SELECT factory, city, county, hospital, variety, name, tel, remark FROM factoryh WHERE variety LIKE ? AND factory LIKE ? And city LIKE ? And county LIKE ? And hospital LIKE ?", 
								  "%" + yaopinname + "%", "%" + factory + "%", "%" + city + "%","%" + county + "%","%" + hospital + "%")
			checkErr(err)
			defer rows.Close()
			for rows.Next() {
			  rst := Rst{}
			  err = rows.Scan(&rst.Factory, &rst.City, &rst.County, &rst.Hospital, &rst.Variety, &rst.Name, &rst.Tel, &rst.Remark)
			  checkErr(err)
			  rsts = append(rsts, rst)
			}
			r, _ := json.Marshal(rsts)
			fmt.Fprintf(w, string(r))
	    }
		
	case "terminalh":
	    factory := strings.TrimFunc(keys["factory"][0],unicode.IsSpace)
		city := strings.TrimFunc(keys["city"][0],unicode.IsSpace)
		county := strings.TrimFunc(keys["county"][0],unicode.IsSpace)
		variety := strings.TrimFunc(keys["variety"][0],unicode.IsSpace)
		hospital := strings.TrimFunc(keys["hospital"][0],unicode.IsSpace)
		type Rst struct {
		  Name string
		  Tel string
		  City string
		  County string
		  Hospital string
		  Variety string
		  Factory string
		  Remark string
		}
		
		if yaopinname=="" && factory=="" && city=="" && county=="" & hospital =="" {
		    rst := []Rst {{"你未输入任何有效的字符串", " ", " ", " ", " ", " ", " ", " "}}
		    r, _ := json.Marshal(rst)
		    fmt.Fprintf(w, string(r))
		} else {
			rsts := []Rst{}
			db, err := sql.Open("sqlite3", "./db/yiyaoshuju.db")
			defer db.Close()
			checkErr(err)
			rows, err := db.Query("SELECT name, tel, city, county, hospital, variety, factory, remark FROM terminalh WHERE variety LIKE ? AND factory LIKE ? And city LIKE ? And county LIKE ? And hospital LIKE ?", 
								  "%" + yaopinname + "%", "%" + factory + "%", "%" + city + "%","%" + county + "%","%" + hospital + "%")
			checkErr(err)
			defer rows.Close()
			for rows.Next() {
			  rst := Rst{}
			  err = rows.Scan(&rst.Name, &rst.Tel, &rst.City, &rst.County, &rst.Hospital, &rst.Variety, &rst.Factory, &rst.Remark)
			  checkErr(err)
			  rsts = append(rsts, rst)
			}
			r, _ := json.Marshal(rsts)
			fmt.Fprintf(w, string(r))
	    }
	case "channelh":
	    channel := strings.TrimFunc(keys["channel"][0],unicode.IsSpace)
		city := strings.TrimFunc(keys["city"][0],unicode.IsSpace)
		county := strings.TrimFunc(keys["county"][0],unicode.IsSpace)
		type Rst struct {
		  Channel string
		  Type string
		  Region string
		  Hospital string
		  Department string
		  Post string
		  Name string
		  Tel string
		  Remark string
		}
		
		if channel=="" && city=="" && county=="" {
		    rst := []Rst {{"你未输入任何有效的字符串", " ", " ", " ", " ", " ", " ", " ", " "}}
		    r, _ := json.Marshal(rst)
		    fmt.Fprintf(w, string(r))
		} else {
			rsts := []Rst{}
			db, err := sql.Open("sqlite3", "./db/yiyaoshuju.db")
			defer db.Close()
			checkErr(err)
			rows, err := db.Query("SELECT channel, type, region, hospital, department, post, name, tel, remark FROM channelh WHERE channel LIKE ? AND city LIKE ? And county LIKE ?", 
								  "%" + channel + "%", "%" + channel + "%", "%" + city + "%","%" + county + "%")
			checkErr(err)
			defer rows.Close()
			for rows.Next() {
			  rst := Rst{}
			  err = rows.Scan(&rst.Channel, &rst.Type, &rst.Region, &rst.Hospital, &rst.Department, &rst.Post, &rst.Name, &rst.Tel, &rst.Remark)
			  checkErr(err)
			  rsts = append(rsts, rst)
			}
			r, _ := json.Marshal(rsts)
			fmt.Fprintf(w, string(r))
	    }
	case "terminals":
	    hospital := strings.TrimFunc(keys["hospital"][0],unicode.IsSpace)
		city := strings.TrimFunc(keys["city"][0],unicode.IsSpace)
		county := strings.TrimFunc(keys["county"][0],unicode.IsSpace)
		type Rst struct {
		  Hospital string
		  Type string
		  Legalman string
		  Class string
		  Economy string
		  City string
		  County string
		  Address string
		  Works string
		  Beds string
		  Remark string
		}
		
		if hospital=="" && city=="" && county=="" {
		    rst := []Rst {{"你未输入任何有效的字符串", " ", " ", " ", " ", " ", " ", " ", " ", " ", " "}}
		    r, _ := json.Marshal(rst)
		    fmt.Fprintf(w, string(r))
		} else {
			rsts := []Rst{}
			db, err := sql.Open("sqlite3", "./db/yiyaoshuju.db")
			defer db.Close()
			checkErr(err)
			rows, err := db.Query("SELECT hospital, type, legalman, class, economy, city, county, address,  works, beds, remark FROM terminals WHERE hospital LIKE ? AND city LIKE ? And county LIKE ?", 
								  "%" + hospital + "%", "%" + city + "%","%" + county + "%")
			checkErr(err)
			defer rows.Close()
			for rows.Next() {
			  rst := Rst{}
			  err = rows.Scan(&rst.Hospital, &rst.Type, &rst.Legalman, &rst.Class, &rst.Economy, &rst.City, &rst.County, &rst.Address, &rst.Works, &rst.Beds, &rst.Remark)
			  checkErr(err)
			  rsts = append(rsts, rst)
			}
			r, _ := json.Marshal(rsts)
			fmt.Fprintf(w, string(r))
	    }
	    
    default:
	    type Rst struct {
		  Tym string
		}
	    rst := Rst{"你未输入任何有效的字符串"}
		r, _ := json.Marshal(rst)
		fmt.Fprintf(w, string(r))
	
	
  }
}
func main() {
  server := http.Server{
	  Addr: "127.0.0.1:8090",
  }
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, r.URL.Path[1:])
  })
  http.HandleFunc("/chaxun", chaxun)
  server.ListenAndServe()
}
