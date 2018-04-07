package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"regexp"
	"strconv"

	"cicada/server/ao"
)

func main() {
	addr := ":8520"
	mux := newMux()

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

type paramOption struct {
	Name   string       // 参数名
	Must   bool         // 是否必须：true 则必须传入；false 时如果有传则用，没传则忽略
	Kind   reflect.Kind // 参数类型
	IsPtr  bool         // 参数类型是否指针
	Range  []int        // 参数值范围，仅可用于整形参数。时间的话，先要转为时间戳
	Regexp string       // 正则匹配
}

// handleParam 处理参数
func handleParam(values url.Values, paramOptionMap map[string]paramOption, param map[string]interface{}) (err error) {
	valueMap := map[string][]string(values) // 转为map，判断没传，还是传了空值
	for k, po := range paramOptionMap {
		vs := valueMap[k]
		// 必须存在
		if len(vs) == 0 && po.Must {
			err = errors.New("必须输入参数: " + k)
			return
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

		var phone string
		if v, ok := param["Phone"]; ok {
			phone = v.(string)
		}
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

	mux.Handle("/ExpenseList", handlerWrapper(func(userID int, param map[string]interface{}) (v interface{}, err error) {
		ep := ao.ExpenseParam{}
		err = mapToStruct(param, &ep)
		if err != nil {
			return
		}

		return ao.ExpenseList(ep)
	}, http.MethodGet, map[string]paramOption{
		"ID": paramOption{
			Kind:  reflect.Int,
			IsPtr: true,
		},
	}, "JSON"))

	return mux
}

// 参数转换
func mapToStruct(param map[string]interface{}, s interface{}) (err error) {
	sType := reflect.TypeOf(s)
	if sType.Kind() != reflect.Ptr {
		err = errors.New("参数s请传入struct指针")
		return
	}
	sType = sType.Elem()
	if sType.Kind() != reflect.Struct {
		err = errors.New("参数s请传入struct")
		return
	}
	sValue := reflect.ValueOf(s)
	for i := 0; i < sType.NumField(); i++ {
		field := sType.Field(i)
		fieldKind := field.Type.Kind()
		if fieldKind == reflect.Struct {
			innerFieldType := field.Type
			innerFieldValue := sValue.Elem().Field(i)
			for j := 0; j < innerFieldType.NumField(); j++ {
				innerField := innerFieldType.Field(j)
				innerFieldName := innerField.Name
				if innerField.Type.Kind() == reflect.Ptr {
					if v, ok := param[innerFieldName]; ok {
						vv := reflect.ValueOf(v)
						innerFieldValue.Field(j).Set(vv)
					}
				} else {
					if v, ok := param[innerFieldName]; ok {
						vv := reflect.ValueOf(v)
						innerFieldValue.Field(j).Set(vv)
					}
				}
			}
		} else {
			fieldName := field.Name
			if v, ok := param[fieldName]; ok {
				vv := reflect.ValueOf(v)
				sValue.Elem().Field(i).Set(vv)
			}
		}
	}
	return
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
		default:
			w.Write([]byte("暂不支持 Get,Post 外的方法"))
			return
		}
		if err := handleParam(values, paramOptionMap, param); err != nil {
			w.Write([]byte(err.Error()))
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

		// 设置 header
		// 文本格式
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// 跨域
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

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
