#!/bin/python3

import http.server
import threading


class CustomHTTPHandler(http.server.SimpleHTTPRequestHandler):
    pass


def run_server():
    server = http.server.HTTPServer(("localhost", 0), CustomHTTPHandler)
    threading.Thread(target=server.serve_forever).start()
    return str(server.server_address[0]) + ":" + str(server.server_address[1])
