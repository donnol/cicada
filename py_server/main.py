from http import server
import json


class HelloHandler(server.BaseHTTPRequestHandler):
    def do_GET(self):
        print("hello")
        msg = json.dumps(["hello, get"])
        self.protocal_version = "HTTP/1.1"
        self.send_response(200)
        self.send_header("Hello", "Header")
        self.send_header("Content-Type", "application/json; charset=utf-8")
        self.end_headers()
        self.wfile.write(msg.encode())

    def do_POST(self):
        msg = json.dumps(["hello, post"])
        self.protocal_version = "HTTP/1.1"
        self.send_response(200)
        self.send_header("Hello", "Header")
        self.send_header("Content-Type", "application/json; charset=utf-8")
        self.end_headers()
        self.wfile.write(msg.encode())


if __name__ == '__main__':
    httpd = server.HTTPServer(('127.0.0.1', 8860), HelloHandler)
    print(8860)
    httpd.serve_forever()
