-- *********************************************
-- * SQL SQLite generation                     
-- *--------------------------------------------
-- * DB-MAIN version: 11.0.2              
-- * Generator date: Sep 20 2021              
-- * Generation date: Sun Feb  2 01:34:05 2025 
-- * LUN file: /personal/repos/daniele821/track-payments/db/TRACK_PAYMENTS.lun 
-- * Schema: track_payments/logic 
-- ********************************************* 

PRAGMA foreign_keys = ON;

create table CITY (
     name           TEXT    not null,
     constraint IDCITY primary key (name)
);

create table ITEM (
     name           TEXT    not null,
     constraint IDITEM primary key (name)
);

create table PAYMENT_METHOD (
     method         TEXT    not null,
     constraint IDPAYMENT_METHOD primary key (method)
);

create table SHOP (
     name           TEXT    not null,
     constraint IDSHOP primary key (name)
);

create table PAYMENT (
     paymentId      INTEGER not null,
     date           TEXT    not null,
     total_price    INTEGER not null,
     city           TEXT    not null,
     shop           TEXT    not null,
     payment_method TEXT    not null,
     constraint IDPAYMENT primary key (paymentId),
     foreign key (city) references CITY(name),
     foreign key (shop) references SHOP(name),
     foreign key (payment_method) references PAYMENT_METHOD(method)
);

create table DETAIL_ORDER (
     nameItem       TEXT    not null,
     paymentId      INTEGER not null,
     quantity       INTEGER not null,
     unit_price     INTEGER not null,
     constraint IDDETAIL_ORDER primary key (paymentId, nameItem),
     foreign key (paymentId) references PAYMENT(paymentId),
     foreign key (nameItem) references ITEM(name)
);

