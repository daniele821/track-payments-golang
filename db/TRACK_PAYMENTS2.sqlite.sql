-- *********************************************
-- * SQL SQLite generation                     
-- *--------------------------------------------
-- * DB-MAIN version: 11.0.2              
-- * Generator date: Sep 20 2021              
-- * Generation date: Wed Feb  5 23:26:30 2025 
-- * LUN file: /personal/repos/daniele821/track-payments/db/TRACK_PAYMENTS.lun 
-- * Schema: track_payments/logic-1 
-- ********************************************* 


-- Database Section
-- ________________ 


-- Tables Section
-- _____________ 

create table CATEGORY (
     name char(1) not null,
     constraint IDCATEGORY primary key (name));

create table CITY (
     name char(1) not null,
     constraint IDCITY primary key (name));

create table DETAIL_ORDER (
     item char(1) not null,
     paymentId char(1) not null,
     quantity char(1) not null,
     unit_price char(1) not null,
     constraint IDDETAIL_ORDER primary key (paymentId, item),
     foreign key (paymentId) references PAYMENT,
     foreign key (item) references ITEM);

create table ITEM (
     name char(1) not null,
     category char(1) not null,
     constraint IDITEM primary key (name),
     foreign key (category) references CATEGORY);

create table PAYMENT (
     paymentId char(1) not null,
     date char(1) not null,
     total_price char(1) not null,
     city char(1) not null,
     shop char(1) not null,
     method char(1) not null,
     constraint IDPAYMENT primary key (paymentId),
     foreign key (city) references CITY,
     foreign key (shop) references SHOP,
     foreign key (method) references PAYMENT_METHOD);

create table PAYMENT_METHOD (
     method char(1) not null,
     constraint IDPAYMENT_METHOD primary key (method));

create table SHOP (
     name char(1) not null,
     constraint IDSHOP primary key (name));


-- Index Section
-- _____________ 

