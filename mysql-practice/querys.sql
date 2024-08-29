CREATE TABLE MemberTable(
    IdMember INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50),
    DNI VARCHAR(12),
    IdEnterprise INT,
    FOREIGN KEY (IdEnterprise) REFERENCES EnterpriseTable(IdEnterprise) ON DELETE SET NULL
)

CREATE TABLE ParentTable(
    IdParent INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50),
    Rel VARCHAR(20),
    IdMember INT,
    FOREIGN KEY (IdMember) REFERENCES MemberTable(IdMember) ON DELETE SET NULL
)

CREATE TABLE EnterpriseTable(
    IdEnterprise INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50),
    Address VARCHAR(50)
)


INSERT INTO EnterpriseTable (Name, Address) VALUES ('coto', 'valenzuela 223')


SELECT * FROM MemberTable

SELECT * FROM ParentTable


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

SELECT M.Name, E.Name FROM MemberTable M INNER JOIN EnterpriseTable E ON M.IdEnterprise = E.IdEnterprise WHERE M.IdEnterprise = 2