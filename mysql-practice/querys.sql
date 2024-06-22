CREATE TABLE MemberTable(
    IdMember INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50),
    DNI VARCHAR(12)
)


CREATE TABLE ParentTable(
    IdParent INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50),
    Rel VARCHAR(20),
    IdMember INT,
    FOREIGN KEY (IdMember) REFERENCES MemberTable(IdMember)
)




SELECT * FROM MemberTable

SELECT * FROM ParentTable

DROP TABLE ParentTable

INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('familiar1', 'prima', 1)

INSERT INTO ParentTable (Name, Rel, IdMember) VALUES ('familiar3', 'PRIMssO', 1)