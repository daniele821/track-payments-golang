import sqlite3
import os
import enum

DATA_DIR = os.path.normpath(os.path.dirname(__file__) + "/data")
SQL_CREATION_FILE = os.path.normpath(DATA_DIR + "/../../db/TRACK_PAYMENTS.sqlite.sql")


class Db:
    def __init__(self, nameDb):
        self.__dbpath__ = DATA_DIR + "/" + nameDb + ".db"
        os.makedirs(DATA_DIR, exist_ok=True)
        self.__conn__ = sqlite3.connect(self.__dbpath__)
        self.__conn__.execute("PRAGMA foreign_keys = ON")
        self.__cursor__ = self.__conn__.cursor()
        self.__cursor__.executescript(open(SQL_CREATION_FILE, "r").read())

    # utility functions
    def __execute__(self, query, data):
        self.__cursor__.execute(query, data)
        self.__conn__.commit()
        return self.__cursor__.lastrowid, self.__cursor__.rowcount

    def __updateTotalPrice__(self, paymentId):
        query = """
        SELECT IFNULL( SUM(quantity * unit_price),0) from DETAIL_ORDER WHERE paymentId = ?
        """
        data = (paymentId,)
        return self.__execute__(query, data)

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
    def insertCity(self, city):
        query = "INSERT INTO CITY(name) values(?)"
        data = (city,)
        return self.__execute__(query, data)

    def insertShop(self, shop):
        query = "INSERT INTO SHOP(name) values(?)"
        data = (shop,)
        return self.__execute__(query, data)

    def insertMethod(self, method):
        query = "INSERT INTO PAYMENT_METHOD(method) values(?)"
        data = (method,)
        return self.__execute__(query, data)

    def insertItem(self, item):
        query = "INSERT INTO ITEM(name) values(?)"
        data = (item,)
        return self.__execute__(query, data)

    def insertDetail(self, item, paymentId, quantity, unit_price):
        query = "INSERT INTO DETAIL_ORDER(nameItem, paymentId, quantity, unit_price) values(?, ?, ?, ?)"
        data = (item, paymentId, quantity, unit_price)
        return self.__execute__(query, data)

    def insertPayment(self, date, city, shop, method):
        query = "INSERT INTO PAYMENT(date, total_price, city, shop, payment_method) values(?,?,?,?,?)"
        data = (date, 0, city, shop, method)
        return self.__execute__(query, data)
