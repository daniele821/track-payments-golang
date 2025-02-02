import sqlite3
import os
import enum

DATA_DIR = os.path.normpath(os.path.dirname(__file__) + "/data")
SQL_CREATION_FILE = os.path.normpath(DATA_DIR + "/../../db/TRACK_PAYMENTS.sqlite.sql")


class Db:
    def __init__(self, nameDb, resAsDict=True):
        self.__dbpath__ = DATA_DIR + "/" + nameDb + ".db"
        os.makedirs(DATA_DIR, exist_ok=True)
        self.__conn__ = sqlite3.connect(self.__dbpath__)
        self.__conn__.execute("PRAGMA foreign_keys = ON")
        self.__cursor__ = self.__conn__.cursor()
        self.__cursor__.executescript(open(SQL_CREATION_FILE, "r").read())

    # selector queries
    def getCity(self):
        return self.__cursor__.execute("SELECT name FROM CITY").fetchall()

    def getShop(self):
        return self.__cursor__.execute("SELECT name FROM SHOP").fetchall()

    def getMethod(self):
        return self.__cursor__.execute("SELECT method FROM PAYMENT_METHOD").fetchall()

    def getItem(self):
        return self.__cursor__.execute("SELECT name FROM ITEM").fetchall()

    def getDetail(self):
        query = "SELECT nameItem, paymentId, quantity, unit_price FROM DETAIL_ORDER"
        return self.__cursor__.execute(query).fetchall()

    def getPayment(self):
        query = "SELECT paymentId, date, total_price, city, shop, payment_method FROM PAYMENT"
        return self.__cursor__.execute(query).fetchall()

    def getFullDetail(self):
        query = """
        SELECT P.paymentId, D.nameItem, D.quantity, D.unit_price, P.date, P.total_price, 
            P.city, P.shop, P.payment_method 
        FROM PAYMENT P, DETAIL_ORDER D, ITEM I
        WHERE P.paymentId = D.paymentId AND D.nameItem = I.name
        """
        return self.__cursor__.execute(query).fetchall()

    # insertion queries
    def __insert__(self, query, data):
        self.__cursor__.execute(query, data)
        self.__conn__.commit()

    def insertCity(self, city):
        self.__insert__("INSERT INTO CITY(name) values(?)", (city,))

    def insertShop(self, shop):
        self.__insert__("INSERT INTO SHOP(name) values(?)", (shop,))

    def insertMethod(self, method):
        self.__insert__("INSERT INTO PAYMENT_METHOD(method) values(?)", (method,))

    def insertItem(self, item):
        self.__insert__("INSERT INTO ITEM(name) values(?)", (item,))
