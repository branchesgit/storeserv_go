CREATE TABLE IF NOT EXISTS `region` (
    id INT  NOT NULL AUTO_INCREMENT,
    regionID INT NOT NULL,
    name VARCHAR(32) NOT NULL,
    pid INT,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO region(regionID, name) VALUES('330000', '浙江省');
INSERT INTO region(regionID, name, pid) VALUES(330200, '宁波市', 330000);
INSERT INTO region(name, regionID,  pid)  VALUES ('余姚市', '330281', '330200');
INSERT INTO region(name, regionID, pid)  VALUES ('慈溪市', '330282', '330200');

