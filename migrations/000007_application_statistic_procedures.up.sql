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


CREATE OR REPLACE FUNCTION main.get_applications_statistics()
RETURNS TABLE (
    total_applications BIGINT,
    open_applications BIGINT,
    closed_applications BIGINT
)
LANGUAGE plpgsql
AS $$
BEGIN

    RETURN QUERY
    SELECT
        COUNT(*) AS total_applications,

        COUNT(*) FILTER (
            WHERE aps.name IN ('Создана', 'В работе')
        ) AS open_applications,

        COUNT(*) FILTER (
            WHERE aps.name = 'Завершена'
        ) AS closed_applications

    FROM main.applications a

    JOIN main.applications_statuses aps
        ON a.status_id = aps.id;

END;
$$;

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

    WHERE aps.name = 'Завершена'

    ORDER BY a.created_date DESC;

END;
$$;
