CREATE TABLE activities
(
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at DateTime NOT NULL,
    updated_at DateTime NOT NULL,
    PRIMARY KEY (id)
)ENGINE = InnoDB;