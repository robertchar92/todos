CREATE TABLE todos
(
    id BIGINT(20) NOT NULL AUTO_INCREMENT,
    activity_group_id BIGINT(20) NOT NULL,
    title VARCHAR(255) NOT NULL,
    priority VARCHAR(255) NOT NULL,
    is_active TINYINT(1) NOT NULL Default true,
    created_at DateTime NOT NULL,
    updated_at DateTime NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (activity_group_id) REFERENCES activities(id)
)ENGINE = InnoDB;