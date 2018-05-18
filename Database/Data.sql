USE Cafe_POS_V2;

SET FOREIGN_KEY_CHECKS = 0; 

TRUNCATE TABLE STAFF_ROLES;
TRUNCATE TABLE FILLINGS_ITEMS;
TRUNCATE TABLE ITEM_CATEGORY;
TRUNCATE TABLE ORDER_ITEMS;
TRUNCATE TABLE ORDERS_TABS;
TRUNCATE TABLE TRANSACTIONS_ORDERS;

TRUNCATE TABLE STAFF;
TRUNCATE TABLE CATEGORY;
TRUNCATE TABLE FILLINGS;
TRUNCATE TABLE ITEMS;
TRUNCATE TABLE ORDERS;
TRUNCATE TABLE ROLES;
TRUNCATE TABLE TABLES;
TRUNCATE TABLE TABS;
TRUNCATE TABLE TRANSACTIONS;

SET FOREIGN_KEY_CHECKS = 1;

/* STAFF */
SELECT 'STAFF TABLE' AS '';
INSERT INTO  STAFF VALUES('STA00001', 'Tom1', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
INSERT INTO  STAFF VALUES('STA00002', 'Tom2', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
INSERT INTO  STAFF VALUES('STA00003', 'Tom3', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
INSERT INTO  STAFF VALUES('STA00004', 'Tom4', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
INSERT INTO  STAFF VALUES('STA00005', 'Tom5', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
INSERT INTO  STAFF VALUES('STA00006', 'Tom6', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
INSERT INTO  STAFF VALUES('STA00007', 'Tom7', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
INSERT INTO  STAFF VALUES('STA00008', 'Tom8', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
INSERT INTO  STAFF VALUES('STA00009', 'Tom9', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
INSERT INTO  STAFF VALUES('STA00010', 'Tom10', 'Kelly', '012345678901234', '1996-11-16', '2012-01-01', 5.5 );
SELECT * FROM STAFF;

/* TABLES */
SELECT 'TABLES TABLE' AS '';
INSERT INTO TABLES VALUES('TBL00001', 'Window', 'Under the window');
INSERT INTO TABLES VALUES('TBL00002', 'Front', 'Front Of The Shop');
INSERT INTO TABLES VALUES('TBL00003', 'Back Right', 'Back Right');
INSERT INTO TABLES VALUES('TBL00004', 'Back Left', 'Back Left');
INSERT INTO TABLES VALUES('TBL00005', 'Middel Right', 'Middel Right');
INSERT INTO TABLES VALUES('TBL00006', 'Middel Left', 'Middel Left');
INSERT INTO TABLES VALUES('TBL00007', 'Take Out', 'Take Out');
INSERT INTO TABLES VALUES('TBL00000', 'No Table', 'No Table');
SELECT * FROM TABLES;

/* CATEGORY */
SELECT 'CATEGORY TABLE'AS '';
INSERT INTO CATEGORY VALUES('CAT00001', 'Can', 'Drink');
INSERT INTO CATEGORY VALUES('CAT00002', 'Coffee', 'Drink');
INSERT INTO CATEGORY VALUES('CAT00003', 'Sandwich', 'Food');
INSERT INTO CATEGORY VALUES('CAT00006', 'Cake', 'Cake');
INSERT INTO CATEGORY VALUES('CAT00007', 'Kids', 'Food');
INSERT INTO CATEGORY VALUES('CAT00008', 'Breakfast', 'Food');
INSERT INTO CATEGORY VALUES('CAT00009', 'Carton', 'Drink');
INSERT INTO CATEGORY VALUES('CAT00010', 'Tea', 'Drink');
INSERT INTO CATEGORY VALUES('CAT00011', 'Hot Choc', 'Drink');
INSERT INTO CATEGORY VALUES('CAT00012', 'Salard', 'Food');
INSERT INTO CATEGORY VALUES('CAT00013', 'Frozen', 'Food');
INSERT INTO CATEGORY VALUES('CAT00014', 'Panni', 'Food');
INSERT INTO CATEGORY VALUES('CAT00015', 'Toasti', 'Food');
INSERT INTO CATEGORY VALUES('CAT00016', 'Potato', 'Food');
SELECT * FROM CATEGORY;

/* ITEMS */
SELECT 'ITEMS TABLE' AS '';
INSERT INTO ITEMS VALUES('ITEM0001', 'Can', 'Canned Drink', 1.0, 50);
INSERT INTO ITEMS VALUES('ITEM0002', 'Coffee', 'Reg Coffee', 1.70, 1000000);
INSERT INTO ITEMS VALUES('ITEM0003', 'Sandwich', 'Sandwich', 3.25, 20000);
INSERT INTO ITEMS VALUES('ITEM0004', 'Cake', 'Cake', 1.75, 10);
INSERT INTO ITEMS VALUES('ITEM0005', 'Salard', 'Its A Salard', 3.0, 50);
INSERT INTO ITEMS VALUES('ITEM0006', 'Latte', 'Coffee With alot of milk', 2.20, 50);
INSERT INTO ITEMS VALUES('ITEM0007', 'Ice cream', 'Ice cream', 1.5, 50);
INSERT INTO ITEMS VALUES('ITEM0008', 'Carton', 'Kids Drink', 0.85, 50);
INSERT INTO ITEMS VALUES('ITEM0009', 'Panni', 'Panni', 3.25, 50);
INSERT INTO ITEMS VALUES('ITEM0010', 'Scram Egg', 'Scram Egg On ToasT', 3.25, 50);
INSERT INTO ITEMS VALUES('ITEM0011', 'Toast', 'Toast', 1.15, 30);
INSERT INTO ITEMS VALUES('ITEM0012', 'Tea', 'Tea', 1.50, 255);
INSERT INTO ITEMS VALUES('ITEM0013', 'Jacket Potato', 'Jacket Patato', 3.95, 255);
INSERT INTO ITEMS VALUES('ITEM0014', 'Hot Chocolate', 'Hot Choc', 2.00, 255);
SELECT * FROM ITEMS;

/* ITEM_CATEGORY */
SELECT 'ITEM_CATEGORY' AS'';
INSERT INTO ITEM_CATEGORY VALUES('ITEM0001', 'CAT00001');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0002', 'CAT00002');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0003', 'CAT00003');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0004', 'CAT00006');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0005', 'CAT00012');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0006', 'CAT00002');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0007', 'CAT00013');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0008', 'CAT00009');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0009', 'CAT00014');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0010', 'CAT00008');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0011', 'CAT00008');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0012', 'CAT00010');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0013', 'CAT00016');
INSERT INTO ITEM_CATEGORY VALUES('ITEM0014', 'CAT00011');
SELECT * FROM ITEM_CATEGORY;

/* ORDERS */
SELECT 'ORDERS TABLE' AS '';
INSERT INTO ORDERS VALUES('ORD00001', '14:50:56', 'TBL00001', FALSE, FALSE, 11.75);
INSERT INTO ORDERS VALUES('ORD00002', '15:55:53', 'TBL00002', FALSE, FALSE, 2.64);
INSERT INTO ORDERS VALUES('ORD00003', '13:01:25', 'TBL00003', TRUE, FALSE, 3.40);
INSERT INTO ORDERS VALUES('ORD00004', '09:45:12', 'TBL00004', TRUE, FALSE, 3.70);
INSERT INTO ORDERS VALUES('ORD00005', '14:55:56', 'TBL00005', FALSE, FALSE, 9.50);
INSERT INTO ORDERS VALUES('ORD00006', '16:50:56', 'TBL00006', TRUE, FALSE,  8.50);
SELECT * FROM ORDERS;

/* ORDER_ITEMS */
SELECT 'ORDER_ITEMS' AS '';
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00001', 'ITEM0001', 'NONE', 2);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00001', 'ITEM0009', 'NONE', 1);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00001', 'ITEM0003', 'TEST', 2);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00002', 'ITEM0012', 'NONE', 1);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00002', 'ITEM0011', 'NONE', 1);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00003', 'ITEM0002', 'NONE', 2);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00004', 'ITEM0006', 'TEST', 1);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00004', 'ITEM0012', 'NONE', 1);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00005', 'ITEM0012', 'NONE', 2);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00005', 'ITEM0004', 'NONE', 2);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00005', 'ITEM0007', 'NONE', 2);
INSERT INTO ORDER_ITEMS VALUES(NULL, 'ORD00006', 'ITEM0002', 'NONE', 5);
SELECT * FROM ORDER_ITEMS;

/* TRANSACTIONS */
SELECT 'TRANSACTIONS' AS '';
INSERT INTO TRANSACTIONS VALUES('T0000001', 3.50, '16:23:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000002', -3.50, '16:23:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000003', 8.5, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000004', 8.5, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000005', 8.5, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000006', 8.5, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000007', -3.0, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000008', 8.5, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000009', 18.5, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000010', -8.5, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000011', 10.5, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000012', 8.5, '12:12:12', '2017-05-19');
INSERT INTO TRANSACTIONS VALUES('T0000013', 1.5, '12:12:12', '2017-05-19');
SELECT * FROM TRANSACTIONS;

/*TRANSACTIONS_ORDERS*/
SELECT 'TRANSACTIONS_ORDERS' AS '';
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000001','ORD00001', 'A');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000002','ORD00001', 'R');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000003','ORD00001', 'A');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000004','ORD00002', 'A');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000005','ORD00003', 'A');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000006','ORD00004', 'A');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000007','ORD00004', 'R');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000008','ORD00005', 'A');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000009','ORD00005', 'A');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000010','ORD00005', 'R');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000011','ORD00005', 'A');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000012','ORD00006', 'A');
INSERT INTO TRANSACTIONS_ORDERS VALUES('T0000013','ORD00006', 'A');
SELECT * FROM TRANSACTIONS_ORDERS;

/*ROLES*/
SELECT 'ROLES' AS '';
INSERT INTO ROLES VALUES('ROLE0001', 'Barista', 'Makes Coffee and performs most tasks front of house');
INSERT INTO ROLES VALUES('ROLE0002', 'Manager', 'Manages Staff and Cooks Food');
SELECT * FROM ROLES;

/* STAFF_ROLES*/
SELECT 'STAFF_ROLES' AS '';
INSERT INTO STAFF_ROLES VALUES('STA00001', 'ROLE0001');
INSERT INTO STAFF_ROLES VALUES('STA00002', 'ROLE0001');
INSERT INTO STAFF_ROLES VALUES('STA00003', 'ROLE0001');
INSERT INTO STAFF_ROLES VALUES('STA00004', 'ROLE0001');
INSERT INTO STAFF_ROLES VALUES('STA00005', 'ROLE0001');
INSERT INTO STAFF_ROLES VALUES('STA00006', 'ROLE0002');
INSERT INTO STAFF_ROLES VALUES('STA00007', 'ROLE0001');
INSERT INTO STAFF_ROLES VALUES('STA00008', 'ROLE0001');
INSERT INTO STAFF_ROLES VALUES('STA00009', 'ROLE0002');
INSERT INTO STAFF_ROLES VALUES('STA00010', 'ROLE0002');
SELECT * FROM STAFF_ROLES;

/* TABS */
SELECT 'TABS' AS '';
INSERT INTO TABS VALUES('TAB00001', 'Art Group', '2017-07-20', FALSE);
SELECT * FROM TABS;

/* ORDERS_TABS */
SELECT 'ORDERS_TABS' AS '';
INSERT INTO ORDERS_TABS VALUES('TAB00001', 'ORD00001');
INSERT INTO ORDERS_TABS VALUES('TAB00001', 'ORD00002');
SELECT * FROM ORDERS_TABS;

/* FILLINGS */
SELECT 'FILLINGS' AS '';
INSERT INTO FILLINGS VALUES('FILL0001', 'TUNA', 'TUNA');
INSERT INTO FILLINGS VALUES('FILL0002', 'CHEESE AND BEANS', 'CHEESE AND BEANS');
INSERT INTO FILLINGS VALUES('FILL0003', 'BBQ CHICKEN AND MOZ', 'BBQ CHICKEN AND MOZ');
INSERT INTO FILLINGS VALUES('FILL0004', 'CHEESE AND PICKEL', 'CHESSE AND PICKEL');
SELECT * FROM FILLINGS;

/* FILLINGS_ITEMS */
SELECT 'FILLINGS_ITEMS' AS '';
INSERT INTO FILLINGS_ITEMS VALUES('FILL0001', 'ITEM0013');
INSERT INTO FILLINGS_ITEMS VALUES('FILL0001', 'ITEM0003');
INSERT INTO FILLINGS_ITEMS VALUES('FILL0001', 'ITEM0009');
INSERT INTO FILLINGS_ITEMS VALUES('FILL0002', 'ITEM0013');
INSERT INTO FILLINGS_ITEMS VALUES('FILL0003', 'ITEM0003');
INSERT INTO FILLINGS_ITEMS VALUES('FILL0003', 'ITEM0009');
INSERT INTO FILLINGS_ITEMS VALUES('FILL0004', 'ITEM0003');
SELECT * FROM FILLINGS_ITEMS;
