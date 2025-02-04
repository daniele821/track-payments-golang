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
    def __runTransaction__(self, queryAndData):
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

    # query = """
    # UPDATE PAYMENT
    # SET total_price = (
    #     SELECT IFNULL( SUM(quantity * unit_price),0) from DETAIL_ORDER WHERE paymentId = ?
    # )
    # WHERE paymentId = ?
    # """

    # query = """
    # SELECT P.*, D.nameItem, D.quantity, D.unit_price
    # FROM PAYMENT P, DETAIL_ORDER D, ITEM I
    # WHERE P.paymentId = D.paymentId AND D.nameItem = I.name
    # """

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
        except Exception as e:
            return self.__err_msg__(f"invalid json: {type(e).__name__}: {e}")

        if "type" not in request:
            return self.__err_typeMissingInJson_msg("type")

        if "data" not in request:
            return self.__err_typeMissingInJson_msg("data")

        requestType = request["type"]
        requestData = request["data"]

        # match request["type"]:
        #     case "insert-city":
        #         return self.__query_msg__(["city"], requestData, self.insertCity)
        #     case _:
        #         return self.__err_msg__("invalid 'type' value in json request!")

        return self.__err_msg__("nothing to do!")
