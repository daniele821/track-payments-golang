#!/bin/python3

import http.server
import threading
import configs


WEBSITE_DIR = configs.WEBSITE_DIR
FLAGS = configs.FLAGS


class CustomHTTPHandler(http.server.SimpleHTTPRequestHandler):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, directory=WEBSITE_DIR, **kwargs)


def run_server():
    server = http.server.HTTPServer(("localhost", FLAGS.port), CustomHTTPHandler)
    threading.Thread(target=server.serve_forever).start()
    # http:// is necessary for termux-open to work
    return "http://localhost:" + str(server.server_address[1])
