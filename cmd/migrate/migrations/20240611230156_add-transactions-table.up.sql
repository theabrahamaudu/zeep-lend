CREATE TABLE IF NOT EXISTS transactions (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `lenderId` INT UNSIGNED NOT NULL,
    `borrowerId` INT UNSIGNED NOT NULL,
    `bookId` INT UNSIGNED NOT NULL, 
    `status` INT NOT NULL,
    `createdAt` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    PRIMARY KEY (id),
    FOREIGN KEY (lenderId) REFERENCES users(`id`),
    FOREIGN KEY (borrowerId) REFERENCES users(`id`),
    FOREIGN KEY (bookId) REFERENCES books(`id`)
);