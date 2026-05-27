-- =========================================
-- 1. ДОБАВИТЬ ЗАЯВКУ
-- =========================================

CREATE OR REPLACE PROCEDURE main.create_application(
    p_description      TEXT,
    p_house_id         INTEGER,
    p_specialist_id    INTEGER,
    p_work_type_id     INTEGER,
    p_status_id        INTEGER DEFAULT 1,
    p_created_date     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM main.houses
        WHERE id = p_house_id
    ) THEN
        RAISE EXCEPTION 'House with id % does not exist', p_house_id;
    END IF;

    IF NOT EXISTS (
        SELECT 1
        FROM main.specialists
        WHERE id = p_specialist_id
    ) THEN
        RAISE EXCEPTION 'Specialist with id % does not exist', p_specialist_id;
    END IF;

    IF NOT EXISTS (
        SELECT 1
        FROM main.work_types
        WHERE id = p_work_type_id
    ) THEN
        RAISE EXCEPTION 'Work type with id % does not exist', p_work_type_id;
    END IF;

    IF NOT EXISTS (
        SELECT 1
        FROM main.applications_statuses
        WHERE id = p_status_id
    ) THEN
        RAISE EXCEPTION 'Status with id % does not exist', p_status_id;
    END IF;

    INSERT INTO main.applications (
        description,
        created_date,
        house_id,
        specialist_id,
        work_type_id,
        status_id
    )
    VALUES (
        p_description,
        p_created_date,
        p_house_id,
        p_specialist_id,
        p_work_type_id,
        p_status_id
    );

END;
$$;


-- =========================================
-- 2. УДАЛИТЬ ЗАЯВКУ
-- =========================================

CREATE OR REPLACE PROCEDURE main.delete_application(
    p_application_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM main.applications
        WHERE id = p_application_id
    ) THEN
        RAISE EXCEPTION 'Application with id % does not exist', p_application_id;
    END IF;

    DELETE FROM main.applications
    WHERE id = p_application_id;

END;
$$;


-- =========================================
-- 3. ЗАКРЫТЬ ЗАЯВКУ
-- status_id = 3 -> "Завершена"
-- =========================================

CREATE OR REPLACE PROCEDURE main.close_application(
    p_application_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM main.applications
        WHERE id = p_application_id
    ) THEN
        RAISE EXCEPTION 'Application with id % does not exist', p_application_id;
    END IF;

    UPDATE main.applications
    SET status_id = 3
    WHERE id = p_application_id;

END;
$$;


-- =========================================
-- 4. ПОЛУЧИТЬ ЗАЯВКИ ПО ДОМУ
-- =========================================

CREATE OR REPLACE FUNCTION main.get_applications_by_house(
    p_house_id INTEGER
)
RETURNS TABLE (
    application_id     INTEGER,
    description        TEXT,
    created_date       TIMESTAMP,
    street             VARCHAR(100),
    house_number       INTEGER,
    specialist_name    VARCHAR(100),
    specialist_post    VARCHAR(100),
    work_type          VARCHAR(100),
    application_status VARCHAR(100)
)
LANGUAGE plpgsql
AS $$
BEGIN

    RETURN QUERY
    SELECT
        a.id,
        a.description,
        a.created_date,
        h.street,
        h.house_number,
        s.full_name,
        s.post,
        wt.name,
        aps.name
    FROM main.applications a
    JOIN main.houses h
        ON a.house_id = h.id
    JOIN main.specialists s
        ON a.specialist_id = s.id
    JOIN main.work_types wt
        ON a.work_type_id = wt.id
    JOIN main.applications_statuses aps
        ON a.status_id = aps.id
    WHERE a.house_id = p_house_id
    ORDER BY a.created_date DESC;

END;
$$;


-- =========================================
-- 5. ПОЛУЧИТЬ ЗАЯВКИ ПО СПЕЦИАЛИСТУ
-- =========================================

CREATE OR REPLACE FUNCTION main.get_applications_by_specialist(
    p_specialist_id INTEGER
)
RETURNS TABLE (
    application_id     INTEGER,
    description        TEXT,
    created_date       TIMESTAMP,
    street             VARCHAR(100),
    house_number       INTEGER,
    specialist_name    VARCHAR(100),
    specialist_post    VARCHAR(100),
    work_type          VARCHAR(100),
    application_status VARCHAR(100)
)
LANGUAGE plpgsql
AS $$
BEGIN

    RETURN QUERY
    SELECT
        a.id,
        a.description,
        a.created_date,
        h.street,
        h.house_number,
        s.full_name,
        s.post,
        wt.name,
        aps.name
    FROM main.applications a
    JOIN main.houses h
        ON a.house_id = h.id
    JOIN main.specialists s
        ON a.specialist_id = s.id
    JOIN main.work_types wt
        ON a.work_type_id = wt.id
    JOIN main.applications_statuses aps
        ON a.status_id = aps.id
    WHERE a.specialist_id = p_specialist_id
    ORDER BY a.created_date DESC;

END;
$$;


-- =========================================
-- 6. ПОЛУЧИТЬ ЗАКРЫТЫЕ ЗАЯВКИ
-- =========================================

CREATE OR REPLACE FUNCTION main.get_closed_applications()
RETURNS TABLE (
    application_id     INTEGER,
    description        TEXT,
    created_date       TIMESTAMP,
    street             VARCHAR(100),
    house_number       INTEGER,
    specialist_name    VARCHAR(100),
    specialist_post    VARCHAR(100),
    work_type          VARCHAR(100),
    application_status VARCHAR(100)
)
LANGUAGE plpgsql
AS $$
BEGIN

    RETURN QUERY
    SELECT
        a.id,
        a.description,
        a.created_date,
        h.street,
        h.house_number,
        s.full_name,
        s.post,
        wt.name,
        aps.name
    FROM main.applications a
    JOIN main.houses h
        ON a.house_id = h.id
    JOIN main.specialists s
        ON a.specialist_id = s.id
    JOIN main.work_types wt
        ON a.work_type_id = wt.id
    JOIN main.applications_statuses aps
        ON a.status_id = aps.id
    WHERE a.status_id = 3
    ORDER BY a.created_date DESC;

END;
$$;

