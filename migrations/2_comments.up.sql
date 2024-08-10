CREATE TABLE if NOT EXISTS comments
(
    id
    SERIAL
    PRIMARY
    KEY,
    comment
    VARCHAR
    NOT
    NULL,
    user_id
    BIGINT
    REFERENCES
    users
(
    id
)
    );

INSERT INTO users(name)
VALUES ('yuji');
INSERT INTO comments (comment, user_id)
VALUES ('first test comment 1', 1);
