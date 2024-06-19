CREATE TABLE MemberTable(
    IdMember INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50),
    DNI VARCHAR(12)
)


CREATE TABLE ParentTable(
    IdParent INT PRIMARY KEY AUTO_INCREMENT;
    Name VARCHAR(50);
    Rel VARCHAR(3);
    IdMember INT;
    FOREIGN KEY (IdMember) REFERENCES MemberTable(IdMember)
)




SELECT * FROM MemberTable