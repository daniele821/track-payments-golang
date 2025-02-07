#!/bin/python3

import os
import argparse

# paths
SCRIPT_DIR = os.path.normpath(os.path.dirname(__file__))
DATA_DIR = os.path.normpath(os.path.join(SCRIPT_DIR, "data"))
WEBSITE_DIR = os.path.normpath(os.path.join(SCRIPT_DIR, "website"))
SQLGEN_FILE = os.path.normpath(os.path.join(os.path.dirname(SCRIPT_DIR), "db", "TRACK_PAYMENTS2.sqlite.sql"))


# parse flags
def parse_args():
    p = argparse.ArgumentParser(description="A program to track payments")
    p.add_argument("-g", "--gui", dest="gui", action="store_true", help="open server in a browser")
    p.add_argument("-p", "--port", dest="port", type=int, default=8080, help="specify the port to use")
    p.add_argument("-s", "--no-server", dest="noserver", action="store_true", help="do not run the server")
    return p.parse_args()


FLAGS = parse_args()
