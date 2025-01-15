-- CREATE CLIENTS TABLE
CREATE TABLE `clients` (
    `id` INT AUTO_INCREMENT PRIMARY KEY
    `name` VARCHAR(100) NOT NULL
    `test` VARCHAR(100) NOT NULL
    `weight_kg` DECIMAL(10, 2) NOT NULL
    `address` VARCHAR(255) NOT NULL
    `street` VARCHAR(150) NOT NULL
    `number` VARCHAR(10) NOT NULL
    `neighborhood` VARCHAR(100) NOT NULL
    `complement` VARCHAR(150)
    `city` VARCHAR(100) NOT NULL
    `state` VARCHAR(50) NOT NULL
    `country` VARCHAR(50) NOT NULL
    `latitude` DECIMAL(10, 8) NOT NULL
    `longitude` DECIMAL(11, 8) NOT NULL
    `created_at` TIMESTAMP NOT NULL
    `updated_at` TIMESTAMP
)

-- TODO: remember -> run docker cp create-table.sql go-mysql:/create-table.sql