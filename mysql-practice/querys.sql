CREATE TABLE MemberTable(
    IdMember INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50),
    DNI INT
)

CREATE TABLE FamilyMemberTable(
    IdFamilyMember INT PRIMARY KEY AUTO_INCREMENT,
    Name VARCHAR(50),
    IdMember INT,
    FOREIGN KEY (IdMember) REFERENCES MemberTable(IdMember)
)