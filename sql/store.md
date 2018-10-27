CREATE TABLE IF NOT EXISTS stores (
    `Id`  int NOT NULL AUTO_INCREMENT,
    `name`  varchar(100)  NOT NULL,
    `desc`  varchar(500),
    `regionID` int not null,
    PRIMARY KEY(`Id`)
) CHARSET=utf8;

INSERT INTO stores(name, regionID) VALUES('周巷大润发龙凤珠宝', 330200);

CREATE TABLE IF NOT EXISTS employees (
    `Id` INT NOT NULL AUTO_INCREMENT,
    `systemId` VARCHAR(11) NOT NULL,
    `psd` VARCHAR(100) NOT NULL,
    `name` VARCHAR(32) NOT NULL,
    `gender` INT NOT NULL,
    `IDCard` VARCHAR(16) NOT NULL,
    `cellphone` VARCHAR(12),
    `storeId` INT NOT NULL,
    PRIMARY KEY(`Id`),
     FOREIGN KEY(`storeId`) REFERENCES stores(`Id`)
)CHARSET=utf8;

INSERT INTO employees (systemId, psd, name, gender, IDCard, storeId)
    VALUES('admin', 'admin', '管理员', 1, '123456', 1);

