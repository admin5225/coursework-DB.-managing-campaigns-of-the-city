CREATE OR REPLACE PROCEDURE main.create_house(
    p_street VARCHAR(100),
    p_house_number INTEGER,
    p_entrances_number INTEGER,
    p_managing_campaign_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN

    IF p_entrances_number <= 0 THEN
        RAISE EXCEPTION
            'Entrances number must be greater than zero';
    END IF;

    IF EXISTS (
        SELECT 1
        FROM main.houses
        WHERE house_number = p_house_number
    ) THEN
        RAISE EXCEPTION
            'House with number % already exists',
            p_house_number;
    END IF;

    IF NOT EXISTS (
        SELECT 1
        FROM main.managing_campaigns
        WHERE id = p_managing_campaign_id
    ) THEN
        RAISE EXCEPTION
            'Managing campaign with id % does not exist',
            p_managing_campaign_id;
    END IF;

    INSERT INTO main.houses (
        street,
        house_number,
        entrances_number,
        managing_campaign_id
    )
    VALUES (
        p_street,
        p_house_number,
        p_entrances_number,
        p_managing_campaign_id
    );

END;
$$;

CREATE OR REPLACE PROCEDURE main.delete_house(
    p_house_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM main.houses
        WHERE id = p_house_id
    ) THEN
        RAISE EXCEPTION
            'House with id % does not exist',
            p_house_id;
    END IF;

    DELETE FROM main.houses
    WHERE id = p_house_id;

END;
$$;