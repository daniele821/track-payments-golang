#!/bin/python3

import sqlite3
import os
import enum
import configs

DATA_DIR = configs.DATA_DIR
SQLGEN_FILE = configs.SQLGEN_FILE


class Db:
    def __init__(self, nameDb):
        self.__dbpath__ = os.path.join(DATA_DIR, nameDb + ".db")
        os.makedirs(DATA_DIR, exist_ok=True)
        self.__conn__ = sqlite3.connect(self.__dbpath__)
        self.__conn__.execute("PRAGMA foreign_keys = ON")
        self.__cursor__ = self.__conn__.cursor()
        self.__cursor__.executescript(open(SQLGEN_FILE, "r").read())

    # utility functions
    def __select__(self, query, data=()):
        res = self.__cursor__.execute(query, data).fetchall()
        attr = [description[0] for description in self.__cursor__.description]
        return res, attr

    def __execute__(self, query, data=()):
        self.__cursor__.execute(query, data)
        self.__conn__.commit()
        return self.__cursor__.lastrowid, self.__cursor__.rowcount

    def __updateTotalPrice__(self, paymentId):
        query = """
        UPDATE PAYMENT
        SET total_price = (
            SELECT IFNULL( SUM(quantity * unit_price),0) from DETAIL_ORDER WHERE paymentId = ?
        )
        WHERE paymentId = ?
        """
        data = (paymentId, paymentId)
        return self.__execute__(query, data)

    # selector queries
    def getCity(self):
        query = "SELECT * FROM CITY"
        return self.__select__(query)

    def getShop(self):
        query = "SELECT * FROM SHOP"
        return self.__select__(query)

    def getMethod(self):
        query = "SELECT * FROM PAYMENT_METHOD"
        return self.__select__(query)

    def getItem(self):
        query = "SELECT * FROM ITEM"
        return self.__select__(query)

    def getDetail(self):
        query = "SELECT * FROM DETAIL_ORDER"
        return self.__select__(query)

    def getPayment(self):
        query = "SELECT * FROM PAYMENT"
        return self.__select__(query)

    def getFullDetail(self):
        query = """
        SELECT P.*, D.nameItem, D.quantity, D.unit_price
        FROM PAYMENT P, DETAIL_ORDER D, ITEM I
        WHERE P.paymentId = D.paymentId AND D.nameItem = I.name
        """
        return self.__select__(query)

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

    def insertDetail(self, item, paymentId, quantity, unitPrice):
        query = "INSERT INTO DETAIL_ORDER(nameItem, paymentId, quantity, unit_price) values(?, ?, ?, ?)"
        data = (item, paymentId, quantity, unitPrice)
        res = self.__execute__(query, data)
        self.__updateTotalPrice__(paymentId)
        return res

    def insertPayment(self, date, city, shop, method):
        query = "INSERT INTO PAYMENT(date, total_price, city, shop, payment_method) values(?,?,?,?,?)"
        data = (date, 0, city, shop, method)
        return self.__execute__(query, data)

    # update queries
    def updateCity(self, old, new):
        query = "UPDATE CITY SET name = ? WHERE name = ?"
        data = (new, old)
        return self.__execute__(query, data)

    def updateShop(self, old, new):
        query = "UPDATE SHOP SET name = ? WHERE name = ?"
        data = (new, old)
        return self.__execute__(query, data)

    def updateMethod(self, old, new):
        query = "UPDATE PAYMENT_METHOD SET method = ? WHERE method = ?"
        data = (new, old)
        return self.__execute__(query, data)

    def updateItem(self, old, new):
        query = "UPDATE ITEM SET name = ? WHERE name = ?"
        data = (new, old)
        return self.__execute__(query, data)

    def updateDetail(self, item, paymentId, newQuantity, newUnitPrice):
        query = "UPDATE DETAIL_ORDER SET quantity = ?, unit_price = ? WHERE nameItem = ? AND paymentId = ?"
        data = (newQuantity, newUnitPrice, item, paymentId)
        res = self.__execute__(query, data)
        self.__updateTotalPrice__(paymentId)
        return res

    def updatePayment(self, paymentId, newDate, newCity, newShop, newMethod):
        query = "UPDATE ITEM SET date = ?, city = ?, shop = ?, payment_method = ? WHERE paymentId = ?"
        data = (newDate, newCity, newShop, newMethod, paymentId)
        return self.__execute__(query, data)

    # deletion queries
    def deleteDetail(self, item, paymentId):
        query = "DELETE FROM DETAIL_ORDER WHERE nameItem = ? AND paymentId = ?"
        data = (item, paymentId)
        res = self.__execute__(query, data)
        self.__updateTotalPrice__(paymentId)
        return res

    def deletePayment(self, paymentId):
        query = "DELETE FROM PAYMENT WHERE paymentId = ?"
        data = (paymentId,)
        return self.__execute__(query, data)
