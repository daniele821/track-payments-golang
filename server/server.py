#!/bin/python3

import sqlite3
import os
import tempfile

FILE_NAME = __file__
DIR_NAME = os.path.dirname(FILE_NAME)
DB_FILE = DIR_NAME + "/.payments.db"
SQL_FILE = os.path.dirname(DIR_NAME) + "/db/TRACK_PAYMENTS.sqlite.sql"

conn = sqlite3.connect(DB_FILE)
cursor = conn.cursor()
cursor.executescript(open(SQL_FILE,"r").read())

