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

    def getAllCities(self):
        return self.__cursor__.execute("SELECT name FROM CITY").fetchall()

    def getAllShops(self):
        return self.__cursor__.execute("SELECT name FROM SHOP").fetchall()

    def getAllMethods(self):
        return self.__cursor__.execute("SELECT method FROM PAYMENT_METHOD").fetchall()

    def getAllItems(self):
        return self.__cursor__.execute("SELECT name FROM ITEM").fetchall()

    def getAllDetails(self):
        query = "SELECT nameItem, paymentId, quantity, unit_price FROM DETAIL_ORDER"
        return self.__cursor__.execute(query).fetchall()

    def getAllPayments(self):
        query = "SELECT paymentId, date, total_price, city, shop, payment_method FROM PAYMENT"
        return self.__cursor__.execute(query).fetchall()

    def getAll(self):
        query = """
        SELECT P.paymentId, D.nameItem, D.quantity, D.unit_price, P.date, P.total_price, 
            P.city, P.shop, P.payment_method 
        FROM PAYMENT P, DETAIL_ORDER D, ITEM I
        WHERE P.paymentId = D.paymentId AND D.nameItem = I.name
        """
        return self.__cursor__.execute(query).fetchall()
