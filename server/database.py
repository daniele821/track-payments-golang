#!/bin/python3

import sqlite3
import os
import enum
import configs
import json


class Db:
    connections = []

    def __init__(self, nameDb):
        # disallow multiple connection to same database!
        if nameDb in Db.connections:
            raise ValueError(f"database {nameDb} already opened!")
        else:
            Db.connections.append(nameDb)

        self.__dbpath__ = os.path.join(configs.DATA_DIR, nameDb + ".db")
        os.makedirs(configs.DATA_DIR, exist_ok=True)
        self.__conn__ = sqlite3.connect(self.__dbpath__)
        self.__conn__.execute("PRAGMA foreign_keys = ON")
        self.__cursor__ = self.__conn__.cursor()
        self.__cursor__.executescript(open(configs.SQLGEN_FILE, "r").read())

    # private utility functions
    def __execute__(self, queryAndData):
        acc = []
        self.__cursor__.execute("BEGIN TRANSACTION;")
        for query, data in queryAndData:
            self.__cursor__.execute(query, data)
            dictRes = {
                "result": self.__cursor__.fetchall(),
                "attributes": [desc[0] for desc in self.__cursor__.description],
                "id": self.__cursor__.lastrowid,
                "rows": self.__cursor__.rowcount,
            }
            acc.append(dictRes)
        self.__conn__.commit()
        return acc

    def __select__(self, query, data=()):
        res = self.__cursor__.execute(query, data).fetchall()
        attr = [description[0] for description in self.__cursor__.description]
        return {"query": res, "attributes": attr}

    def __executeOldFunc__(self, query, data=()):
        self.__cursor__.execute(query, data)
        self.__conn__.commit()
        return {"id": self.__cursor__.lastrowid, "rows": self.__cursor__.rowcount}

    def __updateTotalPrice__(self, paymentId):
        query = """
        UPDATE PAYMENT
        SET total_price = (
            SELECT IFNULL( SUM(quantity * unit_price),0) from DETAIL_ORDER WHERE paymentId = ?
        )
        WHERE paymentId = ?
        """
        data = (paymentId, paymentId)
        return self.__executeOldFunc__(query, data)

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
        return self.__executeOldFunc__(query, data)

    def insertShop(self, shop):
        query = "INSERT INTO SHOP(name) values(?)"
        data = (shop,)
        return self.__executeOldFunc__(query, data)

    def insertMethod(self, method):
        query = "INSERT INTO PAYMENT_METHOD(method) values(?)"
        data = (method,)
        return self.__executeOldFunc__(query, data)

    def insertItem(self, item):
        query = "INSERT INTO ITEM(name) values(?)"
        data = (item,)
        return self.__executeOldFunc__(query, data)

    def insertDetail(self, item, paymentId, quantity, unitPrice):
        query = "INSERT INTO DETAIL_ORDER(nameItem, paymentId, quantity, unit_price) values(?, ?, ?, ?)"
        data = (item, paymentId, quantity, unitPrice)
        res = self.__executeOldFunc__(query, data)
        self.__updateTotalPrice__(paymentId)
        return res

    def insertPayment(self, date, city, shop, method):
        query = "INSERT INTO PAYMENT(date, total_price, city, shop, payment_method) values(?,?,?,?,?)"
        data = (date, 0, city, shop, method)
        return self.__executeOldFunc__(query, data)

    # update queries
    def updateCity(self, old, new):
        query = "UPDATE CITY SET name = ? WHERE name = ?"
        data = (new, old)
        return self.__executeOldFunc__(query, data)

    def updateShop(self, old, new):
        query = "UPDATE SHOP SET name = ? WHERE name = ?"
        data = (new, old)
        return self.__executeOldFunc__(query, data)

    def updateMethod(self, old, new):
        query = "UPDATE PAYMENT_METHOD SET method = ? WHERE method = ?"
        data = (new, old)
        return self.__executeOldFunc__(query, data)

    def updateItem(self, old, new):
        query = "UPDATE ITEM SET name = ? WHERE name = ?"
        data = (new, old)
        return self.__executeOldFunc__(query, data)

    def updateDetail(self, item, paymentId, newQuantity, newUnitPrice):
        query = "UPDATE DETAIL_ORDER SET quantity = ?, unit_price = ? WHERE nameItem = ? AND paymentId = ?"
        data = (newQuantity, newUnitPrice, item, paymentId)
        res = self.__executeOldFunc__(query, data)
        self.__updateTotalPrice__(paymentId)
        return res

    def updatePayment(self, paymentId, newDate, newCity, newShop, newMethod):
        query = "UPDATE ITEM SET date = ?, city = ?, shop = ?, payment_method = ? WHERE paymentId = ?"
        data = (newDate, newCity, newShop, newMethod, paymentId)
        return self.__executeOldFunc__(query, data)

    # deletion queries
    def deleteDetail(self, item, paymentId):
        query = "DELETE FROM DETAIL_ORDER WHERE nameItem = ? AND paymentId = ?"
        data = (item, paymentId)
        res = self.__executeOldFunc__(query, data)
        self.__updateTotalPrice__(paymentId)
        return res

    def deletePayment(self, paymentId):
        query = "DELETE FROM PAYMENT WHERE paymentId = ?"
        data = (paymentId,)
        return self.__executeOldFunc__(query, data)

    # interaction with server
    def __err_msg__(self, msg, status_code=400):
        return status_code, json.dumps({"status": msg})

    def __err_typeMissingInJson_msg(self, typeKey):
        return self.__err_msg__(f"missing '{typeKey}' in the json request!")

    def __query_msg__(self, typesReq, requestData, method):
        acc = []
        for typeReq in typesReq:
            if typeReq not in requestData:
                return self.__err_typeMissingInJson_msg(f"data.{typeReq}")
            acc.append(requestData[typeReq])
        try:
            resDict = method(*acc)
        except Exception as e:
            return self.__err_msg__(f"query failed with error: {type(e).__name__}: {e}")
        return 200, json.dumps({"status": "query was successful!", "res": resDict})

    def answerPostRequest(self, requestJson):
        try:
            request = json.loads(requestJson)
        except:
            return self.__err_msg__("invalid json!")

        if "type" not in request:
            return self.__err_typeMissingInJson_msg("type")

        if "data" not in request:
            return self.__err_typeMissingInJson_msg("data")

        requestType = request["type"]
        requestData = request["data"]

        match request["type"]:
            case "insert-city":
                return self.__query_msg__(["city"], requestData, self.insertCity)
            case "insert-shop":
                return self.__query_msg__(["shop"], requestData, self.insertShop)
            case "insert-method":
                return self.__query_msg__(["method"], requestData, self.insertMethod)
            case "insert-item":
                return self.__query_msg__(["item"], requestData, self.insertItem)
            case "insert-detail":
                return self.__query_msg__(
                    ["item, paymentId, quantity, unitPrice"],
                    requestData,
                    self.insertDetail,
                )
            case "insert-payment":
                return self.__query_msg__(
                    ["date", "city", "shop", "method"], requestData, self.insertPayment
                )
            case "update-city":
                return self.__query_msg__(
                    ["city", "newCity"], requestData, self.updateCity
                )
            case "update-shop":
                return self.__query_msg__(
                    ["shop", "newShop"], requestData, self.updateShop
                )
            case "update-method":
                return self.__query_msg__(
                    ["method", "newMethod"], requestData, self.updateMethod
                )
            case "update-item":
                return self.__query_msg__(
                    ["item", "newItem"], requestData, self.updateItem
                )
            case "update-detail":
                return self.__query_msg__(
                    ["item", "paymentId", "newQuantity", "newUnitPrice"],
                    requestData,
                    self.updateDetail,
                )
            case "update-payment":
                return self.__query_msg__(
                    ["paymentId", "newDate", "newCity", "newShop", "newMethod"],
                    requestData,
                    self.updatePayment,
                )
            case "delete-detail":
                return self.__query_msg__(
                    ["item", "paymentId"], requestData, self.deleteDetail
                )
            case "delete-payment":
                return self.__query_msg__(
                    ["paymentId"], requestData, self.deletePayment
                )
            case "select-city":
                return self.__query_msg__([], requestData, self.getCity)
            case "select-shop":
                return self.__query_msg__([], requestData, self.getShop)
            case "select-method":
                return self.__query_msg__([], requestData, self.getMethod)
            case "select-item":
                return self.__query_msg__([], requestData, self.getItem)
            case "select-detail":
                return self.__query_msg__([], requestData, self.getDetail)
            case "select-payment":
                return self.__query_msg__([], requestData, self.getPayment)
            case "select-fulldetail":
                return self.__query_msg__([], requestData, self.getFullDetail)
            case _:
                return self.__err_msg__("invalid 'type' value in json request!")

        return 200, json.dumps({"status": "TODO!"})
