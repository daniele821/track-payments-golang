#!/bin/python3

import http.server
import threading
import configs
import json


WEBSITE_DIR = configs.WEBSITE_DIR
FLAGS = configs.FLAGS


class CustomHTTPHandler(http.server.SimpleHTTPRequestHandler):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, directory=WEBSITE_DIR, **kwargs)

    def do_POST(self):
        # parse received data
        content_length = int(self.headers.get("Content-Length", 0))
        post_data = self.rfile.read(content_length).decode("utf-8")

        response = json.loads(post_data)
        
        # TODO: THIS IS JUST A PLACEHOLDER FOR NOW, LOGIC IS STILL MISSING!

        # Send a simple response
        self.send_response(200)
        self.send_header("Content-type", "text/json")
        self.end_headers()
        self.wfile.write(json.dumps(response).encode('utf-8'))


def run_server():
    server = http.server.HTTPServer(("localhost", FLAGS.port), CustomHTTPHandler)
    threading.Thread(target=server.serve_forever).start()
    # http:// is necessary for termux-open to work
    return "http://localhost:" + str(server.server_address[1])
