DROP TABLE EnterpriseTable

DROP TABLE MemberTable

DROP TABLE ParentTable

DROP TABLE PaymentTable

DROP TABLE Users

SELECT * FROM UserTable

CREATE TABLE UserTable(
    IdUser INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Username VARCHAR(20),
    Hash VARCHAR(100),
    Role VARCHAR(20)
)

CREATE TABLE EnterpriseTable(
    IdEnterprise INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(255),
    EnterpriseNumber VARCHAR(50),
    Address VARCHAR(50),
    CUIT VARCHAR(50),
    District VARCHAR(50),
    PostalCode VARCHAR(10),
    Phone VARCHAR(50),
    Created_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Updated_At TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)

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
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (IdEnterprise) REFERENCES EnterpriseTable(IdEnterprise) ON DELETE SET NULL
    -- si le pones SET NULL se pone en NULL y si le podes cambiar el IdEnterprise a 1, con 0 no te dejaba
    -- si le pones NO ACTION no te deja eliminar ninguna empresa
    -- si le pones SET DEFAULT no te deja eliminar ninguna empresa
)

CREATE TABLE ParentTable(
    IdParent INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Name VARCHAR(50) NOT NULL,
    LastName VARCHAR(50) NOT NULL,
    Rel VARCHAR(50) NOT NULL,
    Birthday VARCHAR(50) NOT NULL,
    Gender VARCHAR(50) NOT NULL,
    CUIL VARCHAR(50) NOT NULL,
    IdMember INT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (IdMember) REFERENCES MemberTable(IdMember) ON DELETE CASCADE
)

CREATE TABLE PaymentTable(
    IdPayment INT PRIMARY KEY AUTO_INCREMENT NOT NULL,
    Month VARCHAR(2) NOT NULL,
    Year VARCHAR(4) NOT NULL,
    Status VARCHAR(6) NOT NULL,
    Amount INT,
    PaymentDate VARCHAR(20),
    Commentary VARCHAR(400),
    IdEnterprise INT NOT NULL,
    CreatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UpdatedAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (IdEnterprise) REFERENCES EnterpriseTable(IdEnterprise) ON DELETE CASCADE
)

SELECT * FROM EnterpriseTable

SELECT * FROM MemberTable

SELECT * FROM ParentTable

SELECT * FROM PaymentTable




UPDATE EnterpriseTable SET IdEnterprise = 0, Name = "SIN EMPRESA", Address = "POR DEFECTO" WHERE IdEnterprise = 149


INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('familiar3', 'PRIMssO', 1)

SELECT COUNT(IdMember) FROM MemberTable;

DELETE FROM MemberTable WHERE IdMember = '50'

DELETE FROM ParentTable WHERE IdParent = '50'

DELETE FROM EnterpriseTable WHERE IdEnterprise = '50'


INSERT INTO MemberTable (Name, DNI) VALUES ('memberprueba', '44343')


SELECT LAST_INSERT_ID();

SELECT M.Name, E.Name FROM MemberTable M INNER JOIN EnterpriseTable E ON M.IdEnterprise = E.IdEnterprise WHERE M.IdEnterprise = 4

SELECT P.Name, M.Name FROM ParentTable P INNER JOIN MemberTable M ON P.IdMember = M.IdMember WHERE M.IdMember = 3


-- SET FOREIGN_KEY_CHECKS=0;

SELECT * FROM MemberTable WHERE `LastName` = 'Alonso'

SELECT Year FROM PaymentTable WHERE IdEnterprise = 2 GROUP BY Year 