package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"cicada/server/go_server/model"
	"cicada/server/go_server/util"
)

type paramOption struct {
	Name    string       // 参数名
	Must    bool         // 是否必须：true 则必须传入；false 时如果有传则用，没传则忽略
	Default interface{}  // 默认值，非必传参数使用
	Kind    reflect.Kind // 参数类型
	IsPtr   bool         // 参数类型是否指针
	Range   []int        // 参数值范围，仅可用于整形参数。时间的话，先要转为时间戳
	Enum    []int        // 参数值枚举，如1、3、5等
	Regexp  string       // 正则匹配
}

// handleParam 处理参数
func handleParam(values url.Values, paramOptionMap map[string]paramOption, param map[string]interface{}) (err error) {
	valueMap := map[string][]string(values) // 转为map，判断没传，还是传了空值
	for k, po := range paramOptionMap {
		vs := valueMap[k]
		// 必须存在
		if len(vs) == 0 && po.Must {
			err = errors.New("必须输入参数: " + k)
			fmt.Printf("err is %+v\n", err)
			return
		}
		// 默认值
		if po.Default != nil {
			param[k] = po.Default
		}
		for _, v := range vs {
			var actualValue interface{} = v
			// 类型校验
			switch po.Kind {
			case reflect.Int:
				var iv int
				iv, err = strconv.Atoi(v)
				if err != nil {
					err = errors.New("参数类型不正确，请传入 int 类型值")
					return
				}
				// 范围判断
				if len(po.Range) == 2 {
					if iv < po.Range[0] ||
						iv > po.Range[1] {
						err = errors.New("参数值超出范围")
						return
					}
				}
				// 枚举检查
				if len(po.Enum) > 0 {
					var valid bool
					for _, e := range po.Enum {
						if e == iv { // 参数值必须在enum数组里
							valid = true
							break
						}
					}
					if !valid {
						err = errors.New("非法值")
						fmt.Printf("invalid value %d, enum is %v\n", iv, po.Enum)
						return
					}
				}
				if po.IsPtr {
					actualValue = &iv
				} else {
					actualValue = iv
				}
			case reflect.Float64:
				var fv float64
				fv, err = strconv.ParseFloat(v, 64)
				if err != nil {
					err = errors.New("参数类型不正确，请传入 float 类型值")
					return
				}
				if po.IsPtr {
					actualValue = &fv
				} else {
					actualValue = fv
				}
			case reflect.String:
				if po.Regexp != "" {
					// 是否满足正则表达式
					var reg *regexp.Regexp
					reg, err = regexp.Compile(po.Regexp)
					if err != nil {
						log.Printf("正则表达式有问题：%v\n", err)
						continue
					}
					if !reg.MatchString(v) {
						err = errors.New("参数不匹配")
						return
					}
				}
				if po.IsPtr {
					actualValue = &v
				} else {
					actualValue = v
				}
			}

			// 保存参数
			param[k] = actualValue
		}
	}

	return
}

// NewMux 路由注册
func NewMux() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/RegisterCode", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {

		var phone string
		if v, ok := param["Phone"]; ok {
			phone = v.(string)
		}
		code, err := model.PhoneRegisterCode(phone)
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
	}, JSON))

	mux.Handle("/Register", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {

		v, err = model.NameRegister(param["Name"].(string), param["Password"].(string))
		return

	}, http.MethodPost, map[string]paramOption{
		"Name": paramOption{
			Must: true,
		},
		"Password": paramOption{
			Must: true,
		},
	}, JSON))

	mux.Handle("/Login", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {

		u, err := model.NameLogin(param["Name"].(string), param["Password"].(string))
		if err != nil {
			return
		}

		// 设置登陆态
		maxAge := 3600 * 24
		jwt := util.NewJSONWebToken("HelloIamJD")
		jwt.Iss = "server"
		jwt.Iat = time.Now().Unix()
		jwt.Exp = time.Now().Add(24 * time.Hour).Unix()
		jwt.FromUser = u.ID
		session, err := jwt.Token()
		if err != nil {
			return
		}
		cookie := fmt.Sprintf("jd_session=%s; HttpOnly; max-age=%d", session, maxAge)
		headers = append(headers, customHeader{
			"Set-Cookie",
			cookie,
		})

		v = u

		return
	}, http.MethodPost, map[string]paramOption{
		"Name": paramOption{
			Must: true,
		},
		"Password": paramOption{
			Must: true,
		},
	}, JSON))

	mux.Handle("/Logout", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {

		// 怎么让 cookie 失效呢？ TODO
		// 1 后端不处理，前端把cookie去掉，不过会有cookie泄露的问题
		// 2 登出/重置密码后，在后端黑名单里记录cookie，以后每个cookie均检查是否在黑名单，直到该cookie过期

		return
	}, http.MethodPost, map[string]paramOption{}, JSON))

	mux.Handle("/ExpenseList", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {

		// if userID == 0 {
		// 	err = errors.New("please login")
		// 	return
		// }

		ep := model.ExpenseParam{}
		ep.Size = 10
		err = util.MapToStruct(param, &ep)
		if err != nil {
			return
		}

		v, err = model.ExpenseList(ep)
		return

	}, http.MethodGet, map[string]paramOption{
		"ID": paramOption{
			Kind:  reflect.Int,
			IsPtr: true,
		},
	}, JSON))

	mux.Handle("/Upload", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {
		return
	}, http.MethodPost, map[string]paramOption{}, JSON))

	mux.Handle("/AddNote", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {
		id, err := model.AddNote(model.Note{
			UserID: 1,
			Title:  param["Title"].(string),
			Detail: param["Detail"].(string),
		})
		if err != nil {
			return
		}
		v = map[string]interface{}{
			"ID": id,
		}

		return
	}, http.MethodPost, map[string]paramOption{
		"Title": paramOption{
			Must: true,
			Kind: reflect.String,
		},
		"Detail": paramOption{
			Must: true,
			Kind: reflect.String,
		},
	}, JSON))

	mux.Handle("/GetNoteList", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {
		note := model.Note{}
		if id , ok := param["ID"]; ok {
			note.ID = id.(int)
		}
		if title, ok := param["Title"]; ok {
			note.Title = title.(string)
		}
		v, err = model.GetNoteList(note, model.CommonParam{
			Size:   param["Size"].(int),
			Offset: param["Offset"].(int),
		})
		if err != nil {
			return
		}

		return
	}, http.MethodGet, map[string]paramOption{
		"ID": paramOption{
			Kind: reflect.Int,
		},
		"Title": paramOption{
			Kind: reflect.String,
		},
		"Size": paramOption{
			Default: 10,
			Kind:    reflect.Int,
		},
		"Offset": paramOption{
			Default: 0,
			Kind:    reflect.Int,
		},
	}, JSON))

	mux.Handle("/GetNote", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {
		v, err = model.GetNote(param["ID"].(int))
		if err != nil {
			return
		}

		return
	}, http.MethodGet, map[string]paramOption{
		"ID": paramOption{
			Must: true,
			Kind: reflect.Int,
		},
	}, JSON))

	mux.Handle("/ModifyNote", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, headers []customHeader, err error) {
		err = model.ModifyNote(model.Note{
			ID:     param["ID"].(int),
			Title:  param["Title"].(string),
			Detail: param["Detail"].(string),
		})
		if err != nil {
			return
		}

		return
	}, http.MethodPost, map[string]paramOption{
		"ID": paramOption{
			Must: true,
			Kind: reflect.Int,
		},
		"Title": paramOption{
			Must: true,
			Kind: reflect.String,
		},
		"Detail": paramOption{
			Must: true,
			Kind: reflect.String,
		},
	}, JSON))

	return mux
}

// Format 格式
type Format string

const (
	// JSON json 格式
	JSON Format = "JSON"
)

type customHeader struct {
	Key   string
	Value string
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
		[]customHeader,
		error,
	),
	method string,
	paramOptionMap map[string]paramOption,
	responseFormat Format,
) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("request method %s, request url %s\n", r.Method, r.URL.String())

		// 获取参数
		var values url.Values
		var param = make(map[string]interface{})
		if method != r.Method {
			w.Write([]byte("该接口不支持method: " + r.Method + "!"))
			return
		}
		switch method {
		case http.MethodGet:
			values = r.URL.Query()
		case http.MethodPost:
			err := r.ParseForm() // Content-Type must be application/x-www-form-urlencoded
			if err != nil {
				w.Write([]byte(err.Error()))
				return
			}
			values = r.PostForm
			fmt.Printf("%+v\n", values)
		default:
			w.Write([]byte("暂不支持 Get,Post 外的方法"))
			return
		}
		if err := handleParam(values, paramOptionMap, param); err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// 根据 cookie 获取 userID
		var userID int
		cookies := r.Cookies()
		for _, cookie := range cookies {
			if cookie.Name == "jd_session" {
				userID, _ = util.CookieUser(cookie.Value)
				jwt := util.NewJSONWebToken("HelloIamJD")
				ok, err := jwt.Verify(cookie.Value)
				if err != nil {
					w.Write([]byte(err.Error()))
					return
				}
				if ok {
					// 检查 cookie 是否已过期
					if jwt.Exp-time.Now().Unix() < 0 {
						w.Write([]byte("登陆态已过期，请重新登陆"))
						return
					}
					userID = jwt.FromUser
				}
			}
		}

		// 执行方法
		v, headers, err := f(userID, param)
		if err != nil {
			w.Write([]byte(err.Error()))
			return
		}

		// 设置 header
		// 文本格式
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 跨域
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// custom header
		for _, header := range headers {
			w.Header().Set(header.Key, header.Value)
		}

		// 返回
		switch responseFormat {
		case JSON:
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
