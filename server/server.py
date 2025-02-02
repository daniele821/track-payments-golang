#!/bin/python3

from database import Db

db = Db("test")
print(db.getAllCities())
print(db.getAllItems())
print(db.getAllShops())
print(db.getAllMethods())
print(db.getAllPayments())
print(db.getAllDetails())
print()
for line in db.getAll():
    print(line)
