CREATE OR REPLACE PROCEDURE main.add_application(
    p_description      TEXT,
    p_house_id         INTEGER,
    p_specialist_id    INTEGER,
    p_work_type_id     INTEGER,
    p_status_id        INTEGER,
    p_created_date     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)
LANGUAGE plpgsql
AS $$
BEGIN
    -- Проверка существования house
    IF NOT EXISTS (SELECT 1 FROM main.houses WHERE id = p_house_id) THEN
        RAISE EXCEPTION 'House with id % does not exist', p_house_id;
    END IF;

    -- Проверка specialist
    IF NOT EXISTS (SELECT 1 FROM main.specialists WHERE id = p_specialist_id) THEN
        RAISE EXCEPTION 'Specialist with id % does not exist', p_specialist_id;
    END IF;

    -- Проверка work_type
    IF NOT EXISTS (SELECT 1 FROM main.work_types WHERE id = p_work_type_id) THEN
        RAISE EXCEPTION 'Work type with id % does not exist', p_work_type_id;
    END IF;

    -- Проверка status
    IF NOT EXISTS (SELECT 1 FROM main.applications_statuses WHERE id = p_status_id) THEN
        RAISE EXCEPTION 'Status with id % does not exist', p_status_id;
    END IF;

    -- Вставка
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