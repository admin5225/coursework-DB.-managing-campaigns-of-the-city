DELETE FROM main.applications;
DELETE FROM main.flat;
DELETE FROM main.fundraising;

-- 2. Зависимые от managing_campaigns и houses
DELETE FROM main.tools;
DELETE FROM main.specialists;
DELETE FROM main.apartment_complexes;
DELETE FROM main.houses;

-- 3. Справочники (на которые ссылаются applications)
DELETE FROM main.work_types;
DELETE FROM main.applications_statuses;

-- 4. Главная таблица
DELETE FROM main.managing_campaigns;