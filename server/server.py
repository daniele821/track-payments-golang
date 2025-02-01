#!/bin/python3

import sqlite3
import os

FILE_NAME = __file__
DIR_NAME = os.path.dirname(FILE_NAME)

sqlite3.connect(DIR_NAME + "/payments.db")
