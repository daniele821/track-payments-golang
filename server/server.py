#!/bin/python3

import sqlite3
import os
import tempfile

FILE_NAME = __file__
DIR_NAME = os.path.dirname(FILE_NAME)
DB_FILE = DIR_NAME + "/.payments.db"

sqlite3.connect(DB_FILE)
