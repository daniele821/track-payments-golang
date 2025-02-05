#!/bin/python3

import sqlite3
import os
import enum
import configs
import json


class Db:
    connections = []

    def __init__(self, nameDb):
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
    def __runTransaction__(self, *queryData):
        try:
            acc = []
            self.__cursor__.execute("BEGIN TRANSACTION;")
            for query, data in zip(queryData[::2], queryData[1::2]):
                self.__cursor__.execute(query, data)
                dictRes = {
                    "result": self.__cursor__.fetchall(),
                    "attributes": ([d[0] for d in self.__cursor__.description] if self.__cursor__.description else None),
                    "id": self.__cursor__.lastrowid,
                    "rows": self.__cursor__.rowcount,
                }
                acc.append(dictRes)
            self.__conn__.commit()
        except BaseException as e:
            self.__conn__.rollback()
            raise
        # fmt: off
        match len(acc):
            case 0: return None
            case 1: return acc[0]
            case _: return acc
        # fmt: on

    # queries
    def insertCity(self, city):
        query = "INSERT INTO CITY(name) VALUES(?);"
        data = (city,)
        return self.__runTransaction__(query, data)

    def insertShop(self, shop):
        query = "INSERT INTO SHOP(name) VALUES(?);"
        data = (shop,)
        return self.__runTransaction__(query, data)

    def insertMethod(self, method):
        query = "INSERT INTO PAYMENT_METHOD(method) VALUES(?);"
        data = (method,)
        return self.__runTransaction__(query, data)

    def insertCategory(self, category):
        query = "INSERT INTO CATEGORY(name) VALUES(?);"
        data = (category,)
        return self.__runTransaction__(query, data)

    def insertItem(self, item, category):
        query = "INSERT INTO ITEM(name, category) VALUES(?, ?);"
        data = (item, category)
        return self.__runTransaction__(query, data)

    def insertPayment(self, date, city, shop, payment_method):
        query = """
        INSERT INTO PAYMENT(date, total_price, city, shop, payment_method)
        VALUES(?, 0, ?, ?, ?);
        """
        data = (date, city, shop, payment_method)
        return self.__runTransaction__(query, data)

    def insertDetailOrder(self, nameItem, paymentId, quantity, unitPrice):
        query1 = """
        INSERT INTO DETAIL_ORDER(nameItem, paymentId, quantity, unit_price)
        VALUES(?, ?, ?, ?);
        """
        query2 = """
        UPDATE PAYMENT
        SET total_price = (
            SELECT IFNULL( SUM(quantity * unit_price),0 ) from DETAIL_ORDER WHERE paymentId = ?
        )
        WHERE paymentId = ?;
        """
        data1 = (nameItem, paymentId, quantity, unitPrice)
        data2 = (paymentId, paymentId)
        return self.__runTransaction__(query1, data1, query2, data2)[0]

    def updateCity(self, city, newCity):
        query = "UPDATE CITY SET name = ? WHERE name = ?;"
        data = (newCity, city)
        return self.__runTransaction__(query, data)

    def updateShop(self, shop, newShop):
        query = "UPDATE SHOP SET name = ? WHERE name = ?;"
        data = (newShop, shop)
        return self.__runTransaction__(query, data)

    def updateMethod(self, method, newMethod):
        query = "UPDATE PAYMENT_METHOD SET method = ? WHERE method = ?;"
        data = (newMethod, method)
        return self.__runTransaction__(query, data)

    def updateItem(self, item, newItem):
        query = "UPDATE ITEM SET name = ? WHERE name = ?;"
        data = (newItem, item)
        return self.__runTransaction__(query, data)

    def updatePayment(self, paymentId, newDate, newCity, newShop, newPayment_method):
        query = "UPDATE PAYMENT SET date = ?, city = ?, shop = ?, method = ? WHERE paymentId = ?;"
        data = (date, city, shop, payment_method, paymentId)
        return self.__runTransaction__(query, data)

    def updateDetailOrder(self, nameItem, paymentId, newQuantity, newUnitPrice):
        query1 = """
        UPDATE DETAIL_ORDER SET quantity = ?, unit_price = ? WHERE nameItem = ? AND paymentId = ?
        """
        query2 = """
        UPDATE PAYMENT
        SET total_price = (
            SELECT IFNULL( SUM(quantity * unit_price),0 ) from DETAIL_ORDER WHERE paymentId = ?
        )
        WHERE paymentId = ?;
        """
        data1 = (newQuantity, newUnitPrice, nameItem, paymentId)
        data2 = (paymentId, paymentId)
        return self.__runTransaction__(query1, data1, query2, data2)[0]

    # interaction with server
    def __msg__(self, status_code, status, error=None, res=None):
        return status_code, json.dumps({"status": status, "error": error, "res": res})

    def __err_typeMissingInJson_msg(self, typeKey):
        return self.__msg__(400, f"missing '{typeKey}' in the json request!", error="invalid request")

    def __query_msg__(self, typesReq, requestData, method):
        acc = []
        for typeReq in typesReq:
            if typeReq not in requestData:
                return self.__err_typeMissingInJson_msg(f"data.{typeReq}")
            acc.append(requestData[typeReq])
        try:
            resDict = method(*acc)
        except sqlite3.DatabaseError as e:
            return self.__msg__(500, f"query failed: {type(e).__name__}: {e}", error=f"{e}")
        # except Exception as e:
        #     return self.__msg__(500, f"internal error: {e}", error="internal error")
        return self.__msg__(200, "query was successful!", res=resDict)

    def answerPostRequest(self, requestJson):
        try:
            request = json.loads(requestJson)
        except Exception as e:
            return self.__msg__(400, f"invalid json: {type(e).__name__}: {e}", "invalid request")

        if "type" not in request:
            return self.__err_typeMissingInJson_msg("type")

        if "data" not in request:
            return self.__err_typeMissingInJson_msg("data")

        requestType = request["type"]
        requestData = request["data"]

        # fmt: off
        match request["type"]:
            case "insert-city":         return self.__query_msg__(["city"], requestData, self.insertCity)
            case "insert-shop":         return self.__query_msg__(["shop"], requestData, self.insertShop)
            case "insert-method":       return self.__query_msg__(["method"], requestData, self.insertMethod)
            case "insert-category":     return self.__query_msg__(["category"], requestData, self.insertCategory)
            case "insert-item":         return self.__query_msg__(["item", "category"], requestData, self.insertItem)
            case "insert-payment":      return self.__query_msg__(["date", "city", "shop", "method"], requestData, self.insertPayment)
            case "insert-detailorder":  return self.__query_msg__(["nameItem", "paymentId", "quantity", "unitPrice"], requestData, self.insertDetailOrder)
            case "update-item":         return self.__query_msg__(["item", "newItem"], requestData, self.updateItem)
            case "update-city":         return self.__query_msg__(["city", "newCity"], requestData, self.updateCity)
            case "update-shop":         return self.__query_msg__(["shop", "newShop"], requestData, self.updateShop)
            case "update-method":       return self.__query_msg__(["method", "newMethod"], requestData, self.updateMethod)
            case "update-payment":      return self.__query_msg__(["paymentId", "newDate", "newCity", "newShop", "newMethod"], requestData, self.updatePayment)
            case "update-detailorder":  return self.__query_msg__(["nameItem", "paymentId", "newQuantity", "newUnitPrice"], requestData, self.updateDetailOrder)
            case _:                     return self.__msg__(400, "invalid 'type' value in json request!", error="invalid request")
        # fmt: on

        return self.__msg__(400, "this code is unreachable!", error="unreachable code")
