#!/bin/python3

from database import Db

db = Db("payments")
print(db.getAllCities())
print(db.getAllDetails())
print(db.getAllItems())
print(db.getAllMethods())
print(db.getAllPayments())
print(db.getAllShops())
print()
for line in db.getAll():
    print(line)
