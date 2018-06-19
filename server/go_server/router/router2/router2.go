package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"reflect"
	"time"
)

func main() {
	mux := http.NewServeMux()
	// 这样子写，不是很麻烦吗？
	addRoute(mux, &M{"jd", map[string]MethodOption{
		"Name":    MethodOption{"Name", http.MethodGet, []string{}, "JSON"},
		"SetName": MethodOption{"SetName", http.MethodPost, []string{"Name"}, "JSON"},
	}})
	s := &http.Server{
		Addr: ":8810", IdleTimeout: time.Minute, Handler: mux,
	}

	// Shutdown 时调用
	s.RegisterOnShutdown(func() {
		fmt.Println("Shutdown gracefully.")
	})
	// 监听系统中断，终止 server
	idleConnsClosed := make(chan struct{})
	go func() {
		defer func() {
			close(idleConnsClosed)
		}()
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt)
		for {
			select {
			case <-c:
				if err := s.Shutdown(context.Background()); err != nil {
					log.Fatal(err)
				}
				return
			}
		}
	}()

	// 启动 server
	fmt.Println(s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	<-idleConnsClosed
}

// T T 类型
type T interface{}

// 把下面的 M 类型绑定的每个方法都映射到一个HTTP 请求里去
func addRoute(mux *http.ServeMux, typ T) {
	if mux == nil {
		log.Fatalf("wrong mux, %v", mux)
	}
	structType := reflect.TypeOf(typ)
	if structType.Kind() != reflect.Ptr {
		log.Fatalf("type 参数请输入 struct 指针类型，现在输入的是 %v", structType)
	}
	if structType.Elem().Kind() != reflect.Struct {
		log.Fatalf("typ 参数请输入 struct 类型值，输入的类型值为 %v", structType)
	}
	structValue := reflect.ValueOf(typ)
	for i := 0; i < structType.NumMethod(); i++ {
		structMethod := structType.Method(i)
		methodValue := structValue.Method(i)
		// 路由地址
		u := url.URL{
			Path: "/" + structMethod.Name,
		}
		pattern := u.String()

		// 方法需要的参数 -- 只能拿到参数类型，拿不到参数名
		params := []reflect.Value{}
		numParam := structMethod.Type.NumIn()
		if numParam > 1 { // 首个参数是 receiver，忽略它
			for j := 1; j < numParam; j++ {
				param := structMethod.Type.In(j)
				_ = param
				// fmt.Println("=== param ", param)
			}
		}
		// fmt.Printf("%+v, %+v, %s\n", structMethod, structMethod.Type, pattern)
		mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Receive '" + pattern + "' call")

			// 执行该方法 TODO:根据 MethodOption 获取参数
			ret := methodValue.Call(params)

			fmt.Fprintln(w, time.Now(), "You call '", pattern, "' method ", "result is ", ret)
		})
	}
}

// MethodOption 方法选项
type MethodOption struct {
	name   string
	method string   // GET/POST 等方法
	params []string // 参数名，顺序与方法参数一致
	format string   // JSON 等返回格式
}

// M M 类型
type M struct {
	name    string
	options map[string]MethodOption // 方法的配置，键为方法名，值为方法选项
}

// Name 获取名字
func (m *M) Name() string {
	return m.name
}

// SetName 设置名字
func (m *M) SetName(name string) {
	m.name = name
}
