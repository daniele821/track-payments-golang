#!/bin/python3

import os

SCRIPT_FILE = os.path.normpath(__file__)
SCRIPT_DIR = os.path.normpath(os.path.dirname(SCRIPT_FILE))

DATA_DIR = os.path.normpath(os.path.join(SCRIPT_DIR, "data"))
WEBSITE_DIR = os.path.normpath(os.path.join(SCRIPT_DIR, "website"))
SQLGEN_DIR = os.path.normpath(os.path.join(SCRIPT_DIR, "db"))
SQLGEN_FILE = os.path.normpath(os.path.join(SQLGEN_DIR, "creation.sql"))
