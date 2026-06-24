CREATE OR REPLACE PROCEDURE main.create_specialist(
    p_full_name             VARCHAR(100),
    p_post                  VARCHAR(100),
    p_phone_number          VARCHAR(15),
    p_managing_campaign_id  INTEGER,
    p_employment_date       DATE DEFAULT CURRENT_DATE
)
LANGUAGE plpgsql
AS $$
BEGIN

    -- Проверка управляющей компании
    IF NOT EXISTS (
        SELECT 1
        FROM main.managing_campaigns
        WHERE id = p_managing_campaign_id
    ) THEN
        RAISE EXCEPTION
            'Managing campaign with id % does not exist',
            p_managing_campaign_id;
    END IF;

    -- Добавление специалиста
    INSERT INTO main.specialists (
        full_name,
        post,
        phone_number,
        managing_campaign_id,
        employment_date
    )
    VALUES (
        p_full_name,
        p_post,
        p_phone_number,
        p_managing_campaign_id,
        p_employment_date
    );

END;
$$;

CREATE OR REPLACE PROCEDURE main.delete_specialist(
    p_specialist_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN

    -- Проверка существования специалиста
    IF NOT EXISTS (
        SELECT 1
        FROM main.specialists
        WHERE id = p_specialist_id
    ) THEN
        RAISE EXCEPTION
            'Specialist with id % does not exist',
            p_specialist_id;
    END IF;

    -- Удаление специалиста
    DELETE FROM main.specialists
    WHERE id = p_specialist_id;

END;
$$;