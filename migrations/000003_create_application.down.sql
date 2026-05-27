DROP FUNCTION IF EXISTS main.get_closed_applications();

DROP FUNCTION IF EXISTS main.get_applications_by_specialist(INTEGER);

DROP FUNCTION IF EXISTS main.get_applications_by_house(INTEGER);

DROP PROCEDURE IF EXISTS main.close_application(INTEGER);

DROP PROCEDURE IF EXISTS main.delete_application(INTEGER);

DROP PROCEDURE IF EXISTS main.create_application(
    TEXT,
    INTEGER,
    INTEGER,
    INTEGER,
    INTEGER,
    TIMESTAMP
);