#!/bin/python3

import sqlite3
import os
import tempfile

DATA_DIR = os.path.normpath(os.path.dirname(__file__) + "/data")
DB_FILE = os.path.normpath(DATA_DIR + "/payments.db")
SQL_FILE = os.path.normpath(DATA_DIR + "/../../db/TRACK_PAYMENTS.sqlite.sql")

os.makedirs(DATA_DIR, exist_ok=True)

conn = sqlite3.connect(DB_FILE)
cursor = conn.cursor()
cursor.executescript(open(SQL_FILE, "r").read())
