package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
)

func main() {
	log.SetOutput(os.Stdout)

	l, err := net.Listen("tcp", ":5550")
	if err != nil {
		panic(err)
	}
	defer l.Close()

	log.Println(l.Addr())

	db, err := bolt.Open("packet.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	c := make(chan []byte, 100)
	startWorker(8, func(c chan []byte) {
		for b := range c {
			//  保存包
			if err := db.Update(func(tx *bolt.Tx) error {
				bucket, err := tx.CreateBucketIfNotExists([]byte("Bucket"))
				if err != nil {
					return err
				}
				id, err := bucket.NextSequence()
				if err != nil {
					return err
				}
				return bucket.Put([]byte("key"+strconv.Itoa(int(id))), b)
			}); err != nil {
				log.Println(err)
				continue
			}
		}
	}, c)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handleConn(conn, c)
	}
}

func handleConn(conn net.Conn, c chan []byte) {
	defer conn.Close()

	// 建立与后端服务的连接
	serConn, err := net.Dial("tcp", ":8520")
	if err != nil {
		log.Println(err)
		return
	}
	defer serConn.Close()

	// 绑定buf，当连接有数据时，保存一份副本到buf
	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(serConn, &buf1)
	w2 := io.MultiWriter(conn, &buf2)
	go func() {
		for {
			if buf1.Len() > 0 {
				data := make([]byte, buf1.Cap())
				_, err := buf1.Read(data)
				if err != nil {
					fmt.Println(err)
					continue
				}
				buf1.Reset()

				c <- data
			}
		}
	}()
	go func() {
		for {
			if buf2.Len() > 0 {
				data := make([]byte, buf2.Cap())
				_, err := buf2.Read(data)
				if err != nil {
					fmt.Println(err)
					continue
				}
				buf2.Reset()

				c <- data
			}
		}
	}()

	// 请求
	go func() {
		if _, err := io.Copy(w, conn); err != nil {
			fmt.Println(err)
		}
	}()

	// 返回
	if _, err := io.Copy(w2, serConn); err != nil {
		fmt.Println(err)
	}
}

func startWorker(num int, f func(chan []byte), c chan []byte) {
	for i := 0; i < num; i++ {
		go f(c)
	}
}
