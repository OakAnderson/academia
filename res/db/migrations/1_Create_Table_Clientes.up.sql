CREATE TABLE IF NOT EXISTS `clientes` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `cpf` VARCHAR(15) NOT NULL,
    `sexo` CHAR(1) DEFAULT NULL,
    `peso_inicial` float NOT NULL,
    `peso_atual` float NOT NULL,
    `altura` float NOT NULL,
    `biceps` float NOT NULL,
    `peito` float NOT NULL,
    PRIMARY KEY(`id`)
)