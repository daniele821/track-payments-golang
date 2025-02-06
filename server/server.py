#!/bin/python3

import http.server
import threading
import configs
import json
import database
import sys
import os
import subprocess


THREAD_LOCAL = threading.local()


class CustomHTTPHandler(http.server.SimpleHTTPRequestHandler):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, directory=configs.WEBSITE_DIR, **kwargs)

    def do_POST(self):
        # parse received data
        content_length = int(self.headers.get("Content-Length", 0))
        post_data = self.rfile.read(content_length).decode("utf-8")

        # execute query requested by json
        status_code, responseJson = THREAD_LOCAL.db.answerPostRequest(post_data)

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


def server_worker(nameDb, ip, port):
    THREAD_LOCAL.db = database.Db(nameDb)
    http.server.HTTPServer((ip, port), CustomHTTPHandler).serve_forever()


def run(nameDb, ip, port, noServer=False, openGui=False):
    addr_str = f"http://{ip}:{port}"

    if not noServer:
        print(f"starting server on {addr_str}")
        threading.Thread(target=server_worker, args=[nameDb, ip, port]).start()

        if openGui:
            open_link(addr_str)
            print("launching browser to connect to the server...")
        else:
            print("not launching browser!")
