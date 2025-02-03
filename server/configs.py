#!/bin/python3

import os
import utils
import argparse

# paths
SCRIPT_FILE = os.path.normpath(__file__)
SCRIPT_DIR = os.path.normpath(os.path.dirname(SCRIPT_FILE))
DATA_DIR = os.path.normpath(os.path.join(SCRIPT_DIR, "data"))
WEBSITE_DIR = os.path.normpath(os.path.join(SCRIPT_DIR, "website"))
SQLGEN_DIR = os.path.normpath(os.path.join(SCRIPT_DIR, "db"))
SQLGEN_FILE = os.path.normpath(os.path.join(SQLGEN_DIR, "creation.sql"))


# parse flags
def parse_args():
    p = argparse.ArgumentParser(description="A program to track payments")
    p.add_argument(
        "-g", "--gui", dest="gui", action="store_true", help="open server in a browser"
    )
    p.add_argument(
        "-p", "--port", dest="port", type=int, default=8080, help="specify the port to use"
    )
    return p.parse_args()


FLAGS = parse_args()
