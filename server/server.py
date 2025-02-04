#!/bin/python3

import http.server
import threading
import configs
import json
import database
import sys
import os
import subprocess


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
        self.wfile.write(responseJson.encode("utf-8"))


def open_link(link):
    if sys.platform == "linux":
        if "TERMUX_VERSION" in os.environ:
            subprocess.run(["termux-open", link])
        else:
            subprocess.run(["xdg-open", link])
    elif sys.platform == "win32":
        # `shell=True` is required on Windows!
        subprocess.run(["start", link], shell=True)
    elif sys.platform == "darwin":
        subprocess.run(["open", link])
    else:
        raise Exception("unable to open link: unknown platform!")


def run():
    ip = "localhost"
    port = configs.FLAGS.port
    addr_str = f"http://{ip}:{port}"
    if not configs.FLAGS.noserver:
        # running server is a blocking call, thus if we want to avoid threads, it must be the 
        # last function we call before ending the program...
        if configs.FLAGS.gui:
            open_link(addr_str)
            print("launching browser to connect to the server...")
        else:
            print("not launching browser!")
        print(f"starting server on {addr_str}")
        http.server.HTTPServer((ip, port), CustomHTTPHandler).serve_forever()
