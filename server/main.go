package main

import (
	"encoding/json"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"time"

	"cicada/server/ao"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetOutput(os.Stdout)
	rand.Seed(time.Now().Unix())

	addr := ":8520"

	mux := newMux()

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

type paramOption struct {
	Must bool         // 是否必须
	Kind reflect.Kind // 参数类型
}

func checkParam(k, v string, po paramOption) (err error) {
	// 必须存在
	if v == "" && po.Must {
		err = errors.New("必须输入参数: " + k)
		return
	}
	// 类型校验
	switch po.Kind {
	case reflect.Int:
		_, err = strconv.Atoi(v)
		if err != nil {
			err = errors.New("参数类型不正确，请传入 int 类型值")
			return
		}
	case reflect.Float64:
		_, err = strconv.ParseFloat(v, 64)
		if err != nil {
			err = errors.New("参数类型不正确，请传入 float 类型值")
			return
		}
	}

	return
}

// 路由注册
func newMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/RegisterCode", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, err error) {

		// 校验权限 TODO
		// if userID != 0 {
		// roles := []string{}
		// if err := ao.CheckRole(userID, roles); err != nil {
		// w.Write([]byte("用户没有该权限"))
		// return
		// }
		// }

		var phone = param["Phone"].(string)
		code, err := ao.RegisterCode(phone)
		if err != nil {
			return
		}

		v = map[string]string{
			"code": code,
		}

		return
	}, http.MethodPost, map[string]paramOption{
		"Phone": paramOption{
			Must: true,
			Kind: reflect.String,
		},
	}, "JSON"))

	return mux
}

// 参数校验 -- method, paramOptionMap
// 执行方法 -- f
// 返回结果 -- responseFormat
func handlerWrapper(
	f func(
		userID int,
		param map[string]interface{},
	) (
		interface{},
		error,
	),
	method string,
	paramOptionMap map[string]paramOption,
	responseFormat string,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 获取参数
		var param = make(map[string]interface{})
		switch method {
		case http.MethodGet:
			values := r.URL.Query()
			for k, po := range paramOptionMap {
				v := values.Get(k)
				if err := checkParam(k, v, po); err != nil {
					w.Write([]byte(err.Error()))
					return
				}
				param[k] = v
			}
		case http.MethodPost:
			for k, po := range paramOptionMap {
				v := r.FormValue(k)
				if err := checkParam(k, v, po); err != nil {
					w.Write([]byte(err.Error()))
					return
				}
				param[k] = v
			}
		default:
			w.Write([]byte("暂不支持 Get,Post 外的方法"))
			return
		}

		// 根据 cookie 获取 userID TODO
		var userID int
		// cookie, _ := r.Cookie("login_session")
		// userID = CookieUser(cookie)

		// 执行方法
		v, err := f(userID, param)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// 返回
		switch responseFormat {
		case "JSON":
			r, err := json.Marshal(v)
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			w.Write(r)
		default:
			w.Write([]byte("暂不支持 JSON 外的返回格式"))
		}
	})
}
