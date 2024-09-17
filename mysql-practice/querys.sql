CREATE TABLE MemberTable(
    IdMember INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(50),
    DNI VARCHAR(12),
    IdEnterprise INT,
    FOREIGN KEY (IdEnterprise) REFERENCES EnterpriseTable(IdEnterprise) ON DELETE SET NULL
)

CREATE TABLE ParentTable(
    IdParent INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(50),
    Rel VARCHAR(20),
    IdMember INT,
    FOREIGN KEY (IdMember) REFERENCES MemberTable(IdMember) ON DELETE CASCADE
)

CREATE TABLE EnterpriseTable(
    IdEnterprise INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(50),
    Address VARCHAR(50)
)


INSERT INTO EnterpriseTable (Name, Address) VALUES ('coto', 'valenzuela 223')


SELECT * FROM MemberTable

SELECT * FROM ParentTable

SELECT * FROM EnterpriseTable

UPDATE EnterpriseTable SET IdEnterprise = 0, Name = "SIN EMPRESA", Address = "POR DEFECTO" WHERE IdEnterprise = 149


DROP TABLE MemberTable

DROP TABLE ParentTable

DROP TABLE EnterpriseTable

INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('familiar1', 'prima', 1)

INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('familiar3', 'PRIMssO', 1)

SELECT COUNT(IdMember) FROM MemberTable;

DELETE FROM MemberTable WHERE IdMember = '50'

DELETE FROM ParentTable WHERE IdParent = '50'

DELETE FROM EnterpriseTable WHERE IdEnterprise = '50'


INSERT INTO MemberTable (Name, DNI) VALUES ('memberprueba', '44343')


SELECT LAST_INSERT_ID();

SELECT M.Name, E.Name FROM MemberTable M INNER JOIN EnterpriseTable E ON M.IdEnterprise = E.IdEnterprise WHERE M.IdEnterprise = 4

SELECT P.Name, M.Name FROM ParentTable P INNER JOIN MemberTable M ON P.IdMember = M.IdMember WHERE M.IdMember = 3
-- 29 79 156 399



CREATE TABLE MemberTable(
    IdMember INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(50),
    LastName VARCHAR(50),
    DNI VARCHAR(12),
    Birthday DATE,
    Gender VARCHAR(25),
    MaritalStatus VARCHAR(15),
    Phone VARCHAR(30),
    Email VARCHAR(50),
    Address VARCHAR(50),
    FlatNumber VARCHAR(20),
    PostalCode INT,
    District VARCHAR(50),
    Town VARCHAR(50),
    MemberNumber BIGINT,
    CUIL VARCHAR(50),
    IdEnterprise INT,
    EntryDate DATE,
    IdCategory INT,
    LastCardEmition DATE,
    FOREIGN KEY (IdEnterprise) REFERENCES EnterpriseTable(IdEnterprise) ON DELETE SET NULL,
    FOREIGN KEY (IdCategory) REFERENCES CategoryTable(IdCategory)
)

CREATE TABLE ParentTable(
    IdParent INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(50),
    Rel VARCHAR(20),
    IdMember INT,
    FOREIGN KEY (IdMember) REFERENCES MemberTable(IdMember) ON DELETE CASCADE
)

CREATE TABLE CategoryTable(
    IdCategory PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Level INT,
    Name VARCHAR(100)
)

CREATE TABLE EnterpriseTable(
    IdEnterprise INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(50),
    Address VARCHAR(50),
    CUIT VARCHAR(50),
    State VARCHAR(50),
    District VARCHAR(50),
    Town VARCHAR(50),
    PostalCode INT,
    Phone VARCHAR(50),
    StartDate DATE
)