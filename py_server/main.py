from http import server
import json


class HelloHandler(server.BaseHTTPRequestHandler):
    # def __init__(self, r, addr, s):
    #     print(self, r, addr, s)
    #
    def do_GET(self):
        print("hello")
        msg = json.dumps(["hello"])
        self.protocal_version = "HTTP/1.1"
        self.send_response(200)
        self.send_header("Hello", "Header")
        self.end_headers()
        self.wfile.write(msg.encode())


if __name__ == '__main__':
    httpd = server.HTTPServer(('127.0.0.1', 8860), HelloHandler)
    print(8860)
    httpd.serve_forever()
