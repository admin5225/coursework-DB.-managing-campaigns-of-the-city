CREATE SCHEMA main;

CREATE TABLE main.managing_campaigns (
    id                SERIAL                      PRIMARY KEY,
    name              VARCHAR(100)    NOT NULL    UNIQUE,
    legal_address     VARCHAR(100)    NOT NULL,
    director          VARCHAR(100)    NOT NULL,
    phone_number      VARCHAR(15)     NOT NULL    CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    ),
    registration_date DATE            NOT NULL    DEFAULT CURRENT_DATE
);

CREATE TABLE main.specialists (
    id                   SERIAL                      PRIMARY KEY,
    full_name            VARCHAR(100)    NOT NULL,
    post                 VARCHAR(100)    NOT NULL,
    phone_number         VARCHAR(15) CHECK (
        phone_number ~ '^\+[0-9]+$'
        AND
        char_length(phone_number) BETWEEN 10 AND 15
    ),
    managing_campaign_id INTEGER         NOT NULL    REFERENCES main.managing_campaigns(id),
    employment_date      DATE            NOT NULL    DEFAULT CURRENT_DATE
);

CREATE TABLE main.tools (
    ID                      SERIAL                      PRIMARY KEY,
    name                    VARCHAR(100)    NOT NULL,
    managing_campaign_id    INTEGER         NOT NULL    REFERENCES main.managing_campaigns(id),
    quantity                INTEGER         NOT NULL    CHECK (quantity > 0)
);

CREATE TABLE main.apartment_complexes (
    id                      SERIAL                      PRIMARY KEY,
    name                    VARCHAR(100)    NOT NULL,
    organization_name       VARCHAR(100)    NOT NULL,
    director                VARCHAR(100)    NOT NULL,
    managing_campaign_id    INTEGER         NOT NULL    REFERENCES main.managing_campaigns(id)
);

CREATE TABLE main.houses (
    id                      SERIAL                      PRIMARY KEY,
    street                  VARCHAR(100)    NOT NULL,
    house_number            INTEGER         NOT NULL    UNIQUE,
    entrances_number        INTEGER         NOT NULL    CHECK (entrances_number > 0),
    managing_campaign_id    INTEGER         NOT NULL    REFERENCES main.managing_campaigns(id)
);

CREATE TABLE main.work_types (
    id      SERIAL                      PRIMARY KEY,
    name    VARCHAR(100)    NOT NULL    UNIQUE
);

CREATE TABLE main.applications_statuses (
    id      SERIAL                      PRIMARY KEY,
    name    VARCHAR(100)    NOT NULL    UNIQUE
);

CREATE TABLE main.applications (
    id               SERIAL                   PRIMARY KEY,
    description      TEXT         NOT NULL,
    created_date     TIMESTAMP    NOT NULL    DEFAULT CURRENT_TIMESTAMP,
    house_id         INTEGER      NOT NULL    REFERENCES main.houses(id),
    specialist_id    INTEGER      NOT NULL    REFERENCES main.specialists(id),
    work_type_id     INTEGER      NOT NULL    REFERENCES main.work_types(id),
    status_id        INTEGER      NOT NULL    REFERENCES main.applications_statuses(id)
);

CREATE TABLE main.flat (
    id          SERIAL                 PRIMARY KEY,
    house_id    INTEGER    NOT NULL    REFERENCES main.houses(id),
    floor       INTEGER    NOT NULL    CHECK (floor > 0),
    number      INTEGER    NOT NULL    UNIQUE CHECK (number > 0)
);

CREATE TABLE main.fundraising (
    id          SERIAL                        PRIMARY KEY,
    house_id    INTEGER           NOT NULL    REFERENCES main.houses(id),
    target      VARCHAR(100)      NOT NULL,
    amount      NUMERIC(10, 2)    NOT NULL    CHECK (amount > 0)
);
