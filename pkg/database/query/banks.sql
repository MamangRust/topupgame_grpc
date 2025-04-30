-- Get Banks with Pagination and Total Count
-- name: GetBanks :many
SELECT *, COUNT(*) OVER () AS total_count
FROM banks
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR name ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3;

-- Get Active Banks with Pagination and Total Count
-- name: GetBanksActive :many
SELECT *, COUNT(*) OVER () AS total_count
FROM banks
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR name ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3;

-- Get Trashed Banks with Pagination and Total Count
-- name: GetBanksTrashed :many
SELECT *, COUNT(*) OVER () AS total_count
FROM banks
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR name ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3;

-- name: GetMonthAmountBankSuccess :many
WITH 
    date_ranges AS (
        SELECT 
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_banks AS (
        SELECT bank_id, name AS bank_name
        FROM banks
        WHERE deleted_at IS NULL
    ),
    report_months AS (
        SELECT 
            EXTRACT(YEAR FROM range1_start)::integer AS year,
            EXTRACT(MONTH FROM range1_start)::integer AS month
        FROM date_ranges
        UNION
        SELECT 
            EXTRACT(YEAR FROM range2_start)::integer AS year,
            EXTRACT(MONTH FROM range2_start)::integer AS month
        FROM date_ranges
    ),
    month_bank_combos AS (
        SELECT
            rm.year,
            rm.month,
            ab.bank_id,
            ab.bank_name
        FROM report_months rm
        CROSS JOIN active_banks ab
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            b.bank_id,
            b.name AS bank_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN banks b ON t.bank_id = b.bank_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            b.bank_id,
            b.name
    )
SELECT
    mbc.bank_id,
    mbc.bank_name,
    mbc.year::text,
    TO_CHAR(TO_DATE(mbc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_bank_combos mbc
LEFT JOIN monthly_transactions mt ON
    mbc.year = mt.year AND
    mbc.month = mt.month AND
    mbc.bank_id = mt.bank_id
ORDER BY 
    mbc.bank_name ASC,
    mbc.year DESC,
    mbc.month DESC;


-- name: GetYearAmountBankSuccess :many
WITH
    active_banks AS (
        SELECT bank_id, name AS bank_name
        FROM banks
        WHERE deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_bank_combos AS (
        SELECT
            ry.year,
            ab.bank_id,
            ab.bank_name
        FROM report_years ry
        CROSS JOIN active_banks ab
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            b.bank_id,
            b.name AS bank_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            b.bank_id,
            b.name
    )
SELECT
    ybc.bank_id,
    ybc.bank_name,
    ybc.year::text,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_bank_combos ybc
LEFT JOIN yearly_transactions yt ON
    ybc.year = yt.year AND
    ybc.bank_id = yt.bank_id
ORDER BY 
    ybc.bank_name ASC,
    ybc.year DESC;

-- name: GetMonthAmountBankFailed :many
WITH 
    date_ranges AS (
        SELECT 
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_banks AS (
        SELECT bank_id, name AS bank_name
        FROM banks
        WHERE deleted_at IS NULL
    ),
    report_months AS (
        SELECT 
            EXTRACT(YEAR FROM range1_start)::integer AS year,
            EXTRACT(MONTH FROM range1_start)::integer AS month
        FROM date_ranges
        UNION
        SELECT 
            EXTRACT(YEAR FROM range2_start)::integer AS year,
            EXTRACT(MONTH FROM range2_start)::integer AS month
        FROM date_ranges
    ),
    month_bank_combos AS (
        SELECT
            rm.year,
            rm.month,
            ab.bank_id,
            ab.bank_name
        FROM report_months rm
        CROSS JOIN active_banks ab
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            b.bank_id,
            b.name AS bank_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN banks b ON t.bank_id = b.bank_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            b.bank_id,
            b.name
    )
SELECT
    mbc.bank_id,
    mbc.bank_name,
    mbc.year::text,
    TO_CHAR(TO_DATE(mbc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_bank_combos mbc
LEFT JOIN monthly_transactions mt ON
    mbc.year = mt.year AND
    mbc.month = mt.month AND
    mbc.bank_id = mt.bank_id
ORDER BY 
    mbc.bank_name ASC,
    mbc.year DESC,
    mbc.month DESC;


-- name: GetYearAmountBankFailed :many
WITH
    active_banks AS (
        SELECT bank_id, name AS bank_name
        FROM banks
        WHERE deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_bank_combos AS (
        SELECT
            ry.year,
            ab.bank_id,
            ab.bank_name
        FROM report_years ry
        CROSS JOIN active_banks ab
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            b.bank_id,
            b.name AS bank_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            b.bank_id,
            b.name
    )
SELECT
    ybc.bank_id,
    ybc.bank_name,
    ybc.year::text,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_bank_combos ybc
LEFT JOIN yearly_transactions yt ON
    ybc.year = yt.year AND
    ybc.bank_id = yt.bank_id
ORDER BY 
    ybc.bank_name ASC,
    ybc.year DESC;

-- name: GetMonthBankMethodsSuccess :many
WITH
    date_range AS (
        SELECT date_trunc('month', $1::timestamp) AS start_date, date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    active_bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
    ),
    all_months AS (
        SELECT generate_series(
                (
                    SELECT start_date
                    FROM date_range
                ), (
                    SELECT end_date
                    FROM date_range
                ), interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT am.activity_month, abm.bank_id, abm.bank_name, abm.payment_method
        FROM
            all_months am
            CROSS JOIN active_bank_methods abm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(*) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND t.created_at BETWEEN (
                SELECT start_date
                FROM date_range
            ) AND (
                SELECT end_date
                FROM date_range
            )
        GROUP BY
            date_trunc('month', t.created_at),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN monthly_stats ms ON ac.activity_month = ms.activity_month
    AND ac.bank_id = ms.bank_id
    AND ac.payment_method = ms.payment_method
ORDER BY ac.activity_month, ac.bank_name, ac.payment_method;

-- name: GetYearBankMethodsSuccess :many
WITH
    year_range AS (
        SELECT EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int - 4 AS start_year, EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int AS end_year
    ),
    active_bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
    ),
    all_years AS (
        SELECT generate_series(
                (
                    SELECT start_year
                    FROM year_range
                ), (
                    SELECT end_year
                    FROM year_range
                )
            )::text AS year
    ),
    all_combinations AS (
        SELECT ay.year, abm.bank_id, abm.bank_name, abm.payment_method
        FROM
            all_years ay
            CROSS JOIN active_bank_methods abm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.created_at
            )::text AS year,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(*) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(
                YEAR
                FROM t.created_at
            ) BETWEEN (
                SELECT start_year
                FROM year_range
            ) AND (
                SELECT end_year
                FROM year_range
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.created_at
            ),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN yearly_stats ys ON ac.year = ys.year
    AND ac.bank_id = ys.bank_id
    AND ac.payment_method = ys.payment_method
ORDER BY ac.year DESC, ac.bank_name, ac.payment_method;

-- name: GetMonthBankMethodsFailed :many
WITH
    date_range AS (
        SELECT date_trunc('month', $1::timestamp) AS start_date, date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    active_bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
    ),
    all_months AS (
        SELECT generate_series(
                (
                    SELECT start_date
                    FROM date_range
                ), (
                    SELECT end_date
                    FROM date_range
                ), interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT am.activity_month, abm.bank_id, abm.bank_name, abm.payment_method
        FROM
            all_months am
            CROSS JOIN active_bank_methods abm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(*) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.created_at BETWEEN (
                SELECT start_date
                FROM date_range
            ) AND (
                SELECT end_date
                FROM date_range
            )
        GROUP BY
            date_trunc('month', t.created_at),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN monthly_stats ms ON ac.activity_month = ms.activity_month
    AND ac.bank_id = ms.bank_id
    AND ac.payment_method = ms.payment_method
ORDER BY ac.activity_month, ac.bank_name, ac.payment_method;

-- name: GetYearBankMethodsFailed :many
WITH
    -- Define the 5-year range
    year_range AS (
        SELECT EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int - 4 AS start_year, EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int AS end_year
    ),
    active_bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
    ),
    all_years AS (
        SELECT generate_series(
                (
                    SELECT start_year
                    FROM year_range
                ), (
                    SELECT end_year
                    FROM year_range
                )
            )::text AS year
    ),
    all_combinations AS (
        SELECT ay.year, abm.bank_id, abm.bank_name, abm.payment_method
        FROM
            all_years ay
            CROSS JOIN active_bank_methods abm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.created_at
            )::text AS year,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(*) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(
                YEAR
                FROM t.created_at
            ) BETWEEN (
                SELECT start_year
                FROM year_range
            ) AND (
                SELECT end_year
                FROM year_range
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.created_at
            ),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN yearly_stats ys ON ac.year = ys.year
    AND ac.bank_id = ys.bank_id
    AND ac.payment_method = ys.payment_method
ORDER BY ac.year DESC, ac.bank_name, ac.payment_method;

-- name: GetMonthAmountBankSuccessById :many
WITH
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_bank AS (
        SELECT b.bank_id, b.name AS bank_name
        FROM banks b
        WHERE
            b.bank_id = $5
            AND b.deleted_at IS NULL
    ),
    all_months AS (
        SELECT
            EXTRACT(YEAR FROM dr.range1_start)::integer AS year,
            EXTRACT(MONTH FROM dr.range1_start)::integer AS month
        FROM date_ranges dr
        UNION
        SELECT
            EXTRACT(YEAR FROM dr.range2_start)::integer AS year,
            EXTRACT(MONTH FROM dr.range2_start)::integer AS month
        FROM date_ranges dr
    ),
    month_bank_combos AS (
        SELECT
            am.year,
            am.month,
            tb.bank_id,
            tb.bank_name
        FROM all_months am
        CROSS JOIN target_bank tb
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            t.bank_id,
            b.name AS bank_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN banks b ON t.bank_id = b.bank_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end)
            OR (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND t.bank_id = $5
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            t.bank_id,
            b.name
    )
SELECT
    mbc.bank_id,
    mbc.bank_name,
    mbc.year::text AS year,
    TO_CHAR(TO_DATE(mbc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM
    month_bank_combos mbc
LEFT JOIN
    monthly_transactions mt
    ON mbc.year = mt.year
    AND mbc.month = mt.month
    AND mbc.bank_id = mt.bank_id
ORDER BY
    mbc.year DESC,
    mbc.month DESC;


-- name: GetYearAmountBankSuccessById :many
WITH
    target_bank AS (
        SELECT b.bank_id, b.name AS bank_name
        FROM banks b
        WHERE
            b.bank_id = $2
            AND b.deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT ($1::integer - 1) AS year
    ),
    year_bank_combos AS (
        SELECT
            ry.year,
            tb.bank_id,
            tb.bank_name
        FROM report_years ry
        CROSS JOIN target_bank tb
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            t.bank_id,
            b.name AS bank_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND t.bank_id = $2
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, ($1::integer - 1))
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            t.bank_id,
            b.name
    )
SELECT
    ybc.bank_id,
    ybc.bank_name,
    ybc.year::text AS year,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM
    year_bank_combos ybc
LEFT JOIN
    yearly_transactions yt
    ON ybc.year = yt.year
    AND ybc.bank_id = yt.bank_id
ORDER BY
    ybc.year DESC;



-- name: GetMonthAmountBankFailedById :many
WITH
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_bank AS (
        SELECT b.bank_id, b.name AS bank_name
        FROM banks b
        WHERE
            b.bank_id = $5
            AND b.deleted_at IS NULL
    ),
    all_months AS (
        SELECT
            EXTRACT(YEAR FROM dr.range1_start)::integer AS year,
            EXTRACT(MONTH FROM dr.range1_start)::integer AS month
        FROM date_ranges dr
        UNION
        SELECT
            EXTRACT(YEAR FROM dr.range2_start)::integer AS year,
            EXTRACT(MONTH FROM dr.range2_start)::integer AS month
        FROM date_ranges dr
    ),
    month_bank_combos AS (
        SELECT
            am.year,
            am.month,
            tb.bank_id,
            tb.bank_name
        FROM all_months am
        CROSS JOIN target_bank tb
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            t.bank_id,
            b.name AS bank_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN banks b ON t.bank_id = b.bank_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end)
            OR (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.bank_id = $5
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            t.bank_id,
            b.name
    )
SELECT
    mbc.bank_id,
    mbc.bank_name,
    mbc.year::text AS year,
    TO_CHAR(TO_DATE(mbc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM
    month_bank_combos mbc
LEFT JOIN
    monthly_transactions mt
    ON mbc.year = mt.year
    AND mbc.month = mt.month
    AND mbc.bank_id = mt.bank_id
ORDER BY
    mbc.year DESC,
    mbc.month DESC;


-- name: GetYearAmountBankFailedById :many
WITH
    target_bank AS (
        SELECT b.bank_id, b.name AS bank_name
        FROM banks b
        WHERE
            b.bank_id = $2
            AND b.deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT ($1::integer - 1) AS year
    ),
    year_bank_combos AS (
        SELECT
            ry.year,
            tb.bank_id,
            tb.bank_name
        FROM report_years ry
        CROSS JOIN target_bank tb
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            t.bank_id,
            b.name AS bank_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.bank_id = $2
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, ($1::integer - 1))
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            t.bank_id,
            b.name
    )
SELECT
    ybc.bank_id,
    ybc.bank_name,
    ybc.year::text AS year,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM
    year_bank_combos ybc
LEFT JOIN
    yearly_transactions yt
    ON ybc.year = yt.year
    AND ybc.bank_id = yt.bank_id
ORDER BY
    ybc.year DESC;


-- name: GetMonthBankMethodsSuccessById :many
WITH
    date_range AS (
        SELECT date_trunc('month', $1::timestamp) AS start_date, date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND b.bank_id = $2
    ),
    all_months AS (
        SELECT generate_series(
                (
                    SELECT start_date
                    FROM date_range
                ), (
                    SELECT end_date
                    FROM date_range
                ), interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT am.activity_month, bm.bank_id, bm.bank_name, bm.payment_method
        FROM all_months am
            CROSS JOIN bank_methods bm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND b.bank_id = $2
            AND t.created_at BETWEEN (
                SELECT start_date
                FROM date_range
            ) AND (
                SELECT end_date
                FROM date_range
            )
        GROUP BY
            date_trunc('month', t.created_at),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN monthly_transactions mt ON ac.activity_month = mt.activity_month
    AND ac.bank_id = mt.bank_id
    AND ac.payment_method = mt.payment_method
ORDER BY ac.activity_month, ac.payment_method;

-- name: GetYearBankMethodsSuccessById :many
WITH
    year_range AS (
        SELECT EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int - 4 AS start_year, EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int AS end_year
    ),
    bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND b.bank_id = $2
    ),
    all_years AS (
        SELECT generate_series(
                (
                    SELECT start_year
                    FROM year_range
                ), (
                    SELECT end_year
                    FROM year_range
                )
            )::text AS year
    ),
    all_combinations AS (
        SELECT ay.year, bm.bank_id, bm.bank_name, bm.payment_method
        FROM all_years ay
            CROSS JOIN bank_methods bm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.created_at
            )::text AS year,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND b.bank_id = $2
            AND EXTRACT(
                YEAR
                FROM t.created_at
            ) BETWEEN (
                SELECT start_year
                FROM year_range
            ) AND (
                SELECT end_year
                FROM year_range
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.created_at
            ),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN yearly_transactions yt ON ac.year = yt.year
    AND ac.bank_id = yt.bank_id
    AND ac.payment_method = yt.payment_method
ORDER BY ac.year DESC, ac.payment_method;

-- name: GetMonthBankMethodsFailedById :many
WITH
    date_range AS (
        SELECT date_trunc('month', $1::timestamp) AS start_date, date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND b.bank_id = $2
    ),
    all_months AS (
        SELECT generate_series(
                (
                    SELECT start_date
                    FROM date_range
                ), (
                    SELECT end_date
                    FROM date_range
                ), interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT am.activity_month, bm.bank_id, bm.bank_name, bm.payment_method
        FROM all_months am
            CROSS JOIN bank_methods bm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND b.bank_id = $2
            AND t.created_at BETWEEN (
                SELECT start_date
                FROM date_range
            ) AND (
                SELECT end_date
                FROM date_range
            )
        GROUP BY
            date_trunc('month', t.created_at),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN monthly_transactions mt ON ac.activity_month = mt.activity_month
    AND ac.bank_id = mt.bank_id
    AND ac.payment_method = mt.payment_method
ORDER BY ac.activity_month, ac.payment_method;

-- name: GetYearBankMethodsFailedById :many
WITH
    year_range AS (
        SELECT EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int - 4 AS start_year, EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int AS end_year
    ),
    bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND b.bank_id = $2
    ),
    all_years AS (
        SELECT generate_series(
                (
                    SELECT start_year
                    FROM year_range
                ), (
                    SELECT end_year
                    FROM year_range
                )
            )::text AS year
    ),
    all_combinations AS (
        SELECT ay.year, bm.bank_id, bm.bank_name, bm.payment_method
        FROM all_years ay
            CROSS JOIN bank_methods bm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.created_at
            )::text AS year,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND b.bank_id = $2
            AND EXTRACT(
                YEAR
                FROM t.created_at
            ) BETWEEN (
                SELECT start_year
                FROM year_range
            ) AND (
                SELECT end_year
                FROM year_range
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.created_at
            ),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN yearly_transactions yt ON ac.year = yt.year
    AND ac.bank_id = yt.bank_id
    AND ac.payment_method = yt.payment_method
ORDER BY ac.year DESC, ac.payment_method;

-- name: GetMonthAmountBankSuccessByMerchant :many
WITH
    date_ranges AS (
        SELECT
            $1::timestamp AS start1,
            $2::timestamp AS end1,
            $3::timestamp AS start2,
            $4::timestamp AS end2
    ),
    all_banks AS (
        SELECT b.bank_id, b.name AS bank_name
        FROM banks b
        WHERE
            b.deleted_at IS NULL
    ),
    all_months AS (
        SELECT EXTRACT(
                YEAR
                FROM start1
            )::integer AS year, EXTRACT(
                MONTH
                FROM start1
            )::integer AS month
        FROM date_ranges
        UNION
        SELECT EXTRACT(
                YEAR
                FROM start2
            )::integer AS year, EXTRACT(
                MONTH
                FROM start2
            )::integer AS month
        FROM date_ranges
    ),
    bank_month_combos AS (
        SELECT ab.bank_id, ab.bank_name, am.year, am.month
        FROM all_banks ab
            CROSS JOIN all_months am
    ),
    actual_data AS (
        SELECT
            b.bank_id,
            b.name AS bank_name,
            EXTRACT(
                YEAR
                FROM t.created_at
            )::integer AS year,
            EXTRACT(
                MONTH
                FROM t.created_at
            )::integer AS month,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM
            transactions t
            JOIN banks b ON t.bank_id = b.bank_id
            JOIN date_ranges dr ON (
                (
                    t.created_at BETWEEN dr.start1 AND dr.end1
                )
                OR (
                    t.created_at BETWEEN dr.start2 AND dr.end2
                )
            )
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $5
        GROUP BY
            b.bank_id,
            b.name,
            EXTRACT(
                YEAR
                FROM t.created_at
            ),
            EXTRACT(
                MONTH
                FROM t.created_at
            )
    )
SELECT
    bmc.bank_id,
    bmc.bank_name,
    bmc.year::text,
    TO_CHAR(
        TO_DATE(bmc.month::text, 'MM'),
        'Mon'
    ) AS month,
    COALESCE(ad.total_success, 0) AS total_success,
    COALESCE(ad.total_amount, 0) AS total_amount
FROM
    bank_month_combos bmc
    LEFT JOIN actual_data ad ON bmc.bank_id = ad.bank_id
    AND bmc.year = ad.year
    AND bmc.month = ad.month
ORDER BY bmc.year DESC, bmc.month DESC;

-- name: GetYearAmountBankSuccessByMerchant :many
WITH
    all_banks AS (
        SELECT b.bank_id, b.name AS bank_name
        FROM banks b
        WHERE
            b.deleted_at IS NULL
    ),
    all_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    bank_year_combos AS (
        SELECT ab.bank_id, ab.bank_name, ay.year
        FROM all_banks ab
            CROSS JOIN all_years ay
    ),
    actual_data AS (
        SELECT
            b.bank_id,
            b.name AS bank_name,
            EXTRACT(
                YEAR
                FROM t.created_at
            )::integer AS year,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $2
            AND EXTRACT(
                YEAR
                FROM t.created_at
            ) IN ($1::integer, $1::integer - 1)
        GROUP BY
            b.bank_id,
            b.name,
            EXTRACT(
                YEAR
                FROM t.created_at
            )
    )
SELECT
    byc.bank_id,
    byc.bank_name,
    byc.year::text,
    COALESCE(ad.total_success, 0) AS total_success,
    COALESCE(ad.total_amount, 0) AS total_amount
FROM
    bank_year_combos byc
    LEFT JOIN actual_data ad ON byc.bank_id = ad.bank_id
    AND byc.year = ad.year
ORDER BY byc.year DESC;


-- name: GetMonthAmountBankFailedByMerchant :many
WITH
    date_ranges AS (
        SELECT
            $1::timestamp AS start1,
            $2::timestamp AS end1,
            $3::timestamp AS start2,
            $4::timestamp AS end2
    ),
    all_banks AS (
        SELECT b.bank_id, b.name AS bank_name
        FROM banks b
        WHERE
            b.deleted_at IS NULL
    ),
    all_months AS (
        SELECT EXTRACT(
                YEAR
                FROM start1
            )::integer AS year, EXTRACT(
                MONTH
                FROM start1
            )::integer AS month
        FROM date_ranges
        UNION
        SELECT EXTRACT(
                YEAR
                FROM start2
            )::integer AS year, EXTRACT(
                MONTH
                FROM start2
            )::integer AS month
        FROM date_ranges
    ),
    bank_month_combos AS (
        SELECT ab.bank_id, ab.bank_name, am.year, am.month
        FROM all_banks ab
            CROSS JOIN all_months am
    ),
    actual_data AS (
        SELECT
            b.bank_id,
            b.name AS bank_name,
            EXTRACT(
                YEAR
                FROM t.created_at
            )::integer AS year,
            EXTRACT(
                MONTH
                FROM t.created_at
            )::integer AS month,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM
            transactions t
            JOIN banks b ON t.bank_id = b.bank_id
            JOIN date_ranges dr ON (
                (
                    t.created_at BETWEEN dr.start1 AND dr.end1
                )
                OR (
                    t.created_at BETWEEN dr.start2 AND dr.end2
                )
            )
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $5
        GROUP BY
            b.bank_id,
            b.name,
            EXTRACT(
                YEAR
                FROM t.created_at
            ),
            EXTRACT(
                MONTH
                FROM t.created_at
            )
    )
SELECT
    bmc.bank_id,
    bmc.bank_name,
    bmc.year::text,
    TO_CHAR(
        TO_DATE(bmc.month::text, 'MM'),
        'Mon'
    ) AS month,
    COALESCE(ad.total_failed, 0) AS total_failed,
    COALESCE(ad.total_amount, 0) AS total_amount
FROM
    bank_month_combos bmc
    LEFT JOIN actual_data ad ON bmc.bank_id = ad.bank_id
    AND bmc.year = ad.year
    AND bmc.month = ad.month
ORDER BY bmc.year DESC, bmc.month DESC;

-- name: GetYearAmountBankFailedByMerchant :many
WITH
    all_banks AS (
        SELECT b.bank_id, b.name AS bank_name
        FROM banks b
        WHERE
            b.deleted_at IS NULL
    ),
    all_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    bank_year_combos AS (
        SELECT ab.bank_id, ab.bank_name, ay.year
        FROM all_banks ab
            CROSS JOIN all_years ay
    ),
    actual_data AS (
        SELECT
            b.bank_id,
            b.name AS bank_name,
            EXTRACT(
                YEAR
                FROM t.created_at
            )::integer AS year,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $2
            AND EXTRACT(
                YEAR
                FROM t.created_at
            ) IN ($1::integer, $1::integer - 1)
        GROUP BY
            b.bank_id,
            b.name,
            EXTRACT(
                YEAR
                FROM t.created_at
            )
    )
SELECT
    byc.bank_id,
    byc.bank_name,
    byc.year::text,
    COALESCE(ad.total_failed, 0) AS total_failed,
    COALESCE(ad.total_amount, 0) AS total_amount
FROM
    bank_year_combos byc
    LEFT JOIN actual_data ad ON byc.bank_id = ad.bank_id
    AND byc.year = ad.year
ORDER BY byc.year DESC;

-- name: GetMonthBankMethodsSuccessByMerchant :many
WITH
    date_range AS (
        SELECT date_trunc('month', $1::timestamp) AS start_date, date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.merchant_id = $2
    ),
    all_months AS (
        SELECT generate_series(
                (
                    SELECT start_date
                    FROM date_range
                ), (
                    SELECT end_date
                    FROM date_range
                ), interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT am.activity_month, bm.bank_id, bm.bank_name, bm.payment_method
        FROM all_months am
            CROSS JOIN bank_methods bm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $2
            AND t.created_at BETWEEN (
                SELECT start_date
                FROM date_range
            ) AND (
                SELECT end_date
                FROM date_range
            )
        GROUP BY
            date_trunc('month', t.created_at),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN monthly_transactions mt ON ac.activity_month = mt.activity_month
    AND ac.bank_id = mt.bank_id
    AND ac.payment_method = mt.payment_method
ORDER BY ac.activity_month, ac.bank_name, ac.payment_method;

-- name: GetYearBankMethodsSuccessByMerchant :many
WITH
    year_range AS (
        SELECT EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int - 4 AS start_year, EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int AS end_year
    ),
    bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.merchant_id = $2
    ),
    all_years AS (
        SELECT generate_series(
                (
                    SELECT start_year
                    FROM year_range
                ), (
                    SELECT end_year
                    FROM year_range
                )
            )::text AS year
    ),
    all_combinations AS (
        SELECT ay.year, bm.bank_id, bm.bank_name, bm.payment_method
        FROM all_years ay
            CROSS JOIN bank_methods bm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.created_at
            )::text AS year,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $2
            AND EXTRACT(
                YEAR
                FROM t.created_at
            ) BETWEEN (
                SELECT start_year
                FROM year_range
            ) AND (
                SELECT end_year
                FROM year_range
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.created_at
            ),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN yearly_transactions yt ON ac.year = yt.year
    AND ac.bank_id = yt.bank_id
    AND ac.payment_method = yt.payment_method
ORDER BY ac.year DESC, ac.bank_name, ac.payment_method;

-- name: GetMonthBankMethodsFailedByMerchant :many
WITH
    date_range AS (
        SELECT date_trunc('month', $1::timestamp) AS start_date, date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.merchant_id = $2
    ),
    all_months AS (
        SELECT generate_series(
                (
                    SELECT start_date
                    FROM date_range
                ), (
                    SELECT end_date
                    FROM date_range
                ), interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT am.activity_month, bm.bank_id, bm.bank_name, bm.payment_method
        FROM all_months am
            CROSS JOIN bank_methods bm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $2
            AND t.created_at BETWEEN (
                SELECT start_date
                FROM date_range
            ) AND (
                SELECT end_date
                FROM date_range
            )
        GROUP BY
            date_trunc('month', t.created_at),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN monthly_transactions mt ON ac.activity_month = mt.activity_month
    AND ac.bank_id = mt.bank_id
    AND ac.payment_method = mt.payment_method
ORDER BY ac.activity_month, ac.bank_name, ac.payment_method;

-- name: GetYearBankMethodsFailedByMerchant :many
WITH
    year_range AS (
        SELECT EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int - 4 AS start_year, EXTRACT(
                YEAR
                FROM $1::timestamp
            )::int AS end_year
    ),
    bank_methods AS (
        SELECT DISTINCT
            b.bank_id,
            b.name AS bank_name,
            t.payment_method
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.merchant_id = $2
    ),
    all_years AS (
        SELECT generate_series(
                (
                    SELECT start_year
                    FROM year_range
                ), (
                    SELECT end_year
                    FROM year_range
                )
            )::text AS year
    ),
    all_combinations AS (
        SELECT ay.year, bm.bank_id, bm.bank_name, bm.payment_method
        FROM all_years ay
            CROSS JOIN bank_methods bm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(
                YEAR
                FROM t.created_at
            )::text AS year,
            b.bank_id,
            b.name AS bank_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
            JOIN banks b ON t.bank_id = b.bank_id
        WHERE
            t.deleted_at IS NULL
            AND b.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $2
            AND EXTRACT(
                YEAR
                FROM t.created_at
            ) BETWEEN (
                SELECT start_year
                FROM year_range
            ) AND (
                SELECT end_year
                FROM year_range
            )
        GROUP BY
            EXTRACT(
                YEAR
                FROM t.created_at
            ),
            b.bank_id,
            b.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.bank_id,
    ac.bank_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM
    all_combinations ac
    LEFT JOIN yearly_transactions yt ON ac.year = yt.year
    AND ac.bank_id = yt.bank_id
    AND ac.payment_method = yt.payment_method
ORDER BY ac.year DESC, ac.bank_name, ac.payment_method;

-- Create Bank
-- name: CreateBank :one
INSERT INTO banks (name) VALUES ($1) RETURNING *;

-- Get Bank by ID
-- name: GetBankByID :one
SELECT
    bank_id,
    name,
    created_at,
    updated_at,
    deleted_at
FROM banks
WHERE
    bank_id = $1
    AND deleted_at IS NULL;

-- Update Bank
-- name: UpdateBank :one
UPDATE banks
SET
    name = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE
    bank_id = $1
    AND deleted_at IS NULL
RETURNING
    *;

-- Trash Bank (Soft Delete)
-- name: TrashBank :one
UPDATE banks
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    bank_id = $1
    AND deleted_at IS NULL
RETURNING
    *;

-- Restore Trashed Bank
-- name: RestoreBank :one
UPDATE banks
SET
    deleted_at = NULL
WHERE
    bank_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    *;

-- Delete Bank Permanently
-- name: DeleteBankPermanently :exec
DELETE FROM banks WHERE bank_id = $1 AND deleted_at IS NOT NULL;

-- Restore All Trashed Banks
-- name: RestoreAllBanks :exec
UPDATE banks SET deleted_at = NULL WHERE deleted_at IS NOT NULL;

-- Delete All Trashed Banks Permanently
-- name: DeleteAllPermanentBanks :exec
DELETE FROM banks WHERE deleted_at IS NOT NULL;