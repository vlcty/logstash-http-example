#!/usr/bin/env python3
# Adapted from: https://pythonbasics.org/webserver/
from http.server import BaseHTTPRequestHandler, HTTPServer
# import time

class MyServer(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.send_header("Content-type", "text/html")
        self.end_headers()

        # Stolen from https://stackoverflow.com/a/5976905)
        post_body = self.rfile.read(int(self.headers.get('Content-Length'))).decode('utf8')
        self.wfile.write(bytes(f"Received: {str(post_body)}", "utf8"))

if __name__ == "__main__":        
    webServer = HTTPServer(("0.0.0.0", 8090), MyServer)
    webServer.serve_forever()