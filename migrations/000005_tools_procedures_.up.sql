-- =========================================
-- 1. ДОБАВИТЬ ИНСТРУМЕНТ
-- =========================================

CREATE OR REPLACE PROCEDURE main.create_tool(
    p_name VARCHAR(100),
    p_managing_campaign_id INTEGER,
    p_quantity INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN

    IF p_quantity <= 0 THEN
        RAISE EXCEPTION
            'Quantity must be greater than zero';
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

    INSERT INTO main.tools (
        name,
        managing_campaign_id,
        quantity
    )
    VALUES (
        p_name,
        p_managing_campaign_id,
        p_quantity
    );

END;
$$;


-- =========================================
-- 2. ОБНОВИТЬ КОЛИЧЕСТВО ИНСТРУМЕНТА
-- =========================================

CREATE OR REPLACE PROCEDURE main.update_tool_quantity(
    p_tool_id INTEGER,
    p_quantity INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM main.tools
        WHERE id = p_tool_id
    ) THEN
        RAISE EXCEPTION
            'Tool with id % does not exist',
            p_tool_id;
    END IF;

    IF p_quantity <= 0 THEN
        RAISE EXCEPTION
            'Quantity must be greater than zero';
    END IF;

    UPDATE main.tools
    SET quantity = p_quantity
    WHERE id = p_tool_id;

END;
$$;


-- =========================================
-- 3. УДАЛИТЬ ИНСТРУМЕНТ
-- =========================================

CREATE OR REPLACE PROCEDURE main.delete_tool(
    p_tool_id INTEGER
)
LANGUAGE plpgsql
AS $$
BEGIN

    IF NOT EXISTS (
        SELECT 1
        FROM main.tools
        WHERE id = p_tool_id
    ) THEN
        RAISE EXCEPTION
            'Tool with id % does not exist',
            p_tool_id;
    END IF;

    DELETE FROM main.tools
    WHERE id = p_tool_id;

END;
$$;