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
    Name VARCHAR(100) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    DNI VARCHAR(50) NOT NULL,
    Birthday VARCHAR(50) NOT NULL,
    Gender VARCHAR(50) NOT NULL,
    MaritalStatus VARCHAR(50) NOT NULL,
    Phone VARCHAR(50) NOT NULL,
    Email VARCHAR(50),
    Address VARCHAR(50) NOT NULL,
    PostalCode VARCHAR(10) NOT NULL,
    District VARCHAR(50) NOT NULL,
    MemberNumber VARCHAR(50) NOT NULL,
    CUIL VARCHAR(50) NOT NULL,
    IdEnterprise INT,
    -- aca va sin NOT NULL, por si borras la empresa
    Category VARCHAR(100) NOT NULL,
    EntryDate VARCHAR(50) NOT NULL,
    FOREIGN KEY (IdEnterprise) REFERENCES EnterpriseTable(IdEnterprise) ON DELETE SET NULL
)

CREATE TABLE ParentTable(
    IdParent INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    Rel VARCHAR(50) NOT NULL,
    Birthday VARCHAR(50) NOT NULL,
    Gender VARCHAR(50) NOT NULL,
    CUIL VARCHAR(50) NOT NULL,
    IdMember INT,
    FOREIGN KEY (IdMember) REFERENCES MemberTable(IdMember) ON DELETE CASCADE
)


CREATE TABLE EnterpriseTable(
    IdEnterprise INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(255),
    Address VARCHAR(50),
    CUIT VARCHAR(50),
    District VARCHAR(50),
    PostalCode VARCHAR(10),
    Phone VARCHAR(50)
)

SELECT * FROM MemberTable WHERE CAST(Birthday AS DATE) > '2022-01-01'