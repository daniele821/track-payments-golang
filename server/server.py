#!/bin/python3

import http.server
import threading
import configs
import json
import database

DB = database.Db("payments")

class CustomHTTPHandler(http.server.SimpleHTTPRequestHandler):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, directory=configs.WEBSITE_DIR, **kwargs)

    def do_POST(self):
        # parse received data
        content_length = int(self.headers.get("Content-Length", 0))
        post_data = self.rfile.read(content_length).decode("utf-8")

        # execute query requested by json        
        status_code, responseJson = DB.answerPostRequest(post_data)
        
        # Send a simple response
        self.send_response(status_code)
        self.send_header("Content-type", "text/json")
        self.end_headers()
        self.wfile.write(responseJson.encode('utf-8'))


def run_server():
    server = http.server.HTTPServer(("localhost", configs.FLAGS.port), CustomHTTPHandler)
    threading.Thread(target=server.serve_forever).start()
    # http:// is necessary for termux-open to work
    return "http://localhost:" + str(server.server_address[1])
