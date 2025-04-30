-- Get Categories with Pagination and Total Count
-- name: GetCategories :many
SELECT *, COUNT(*) OVER () AS total_count
FROM categories
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

-- Get Active Categories with Pagination and Total Count
-- name: GetCategoriesActive :many
SELECT *, COUNT(*) OVER () AS total_count
FROM categories
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

-- Get Trashed Categories with Pagination and Total Count
-- name: GetCategoriesTrashed :many
SELECT *, COUNT(*) OVER () AS total_count
FROM categories
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


-- name: GetMonthAmountCategorySuccess :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_categories AS (
        SELECT category_id, name AS category_name
        FROM categories
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
    month_category_combos AS (
        SELECT
            rm.year,
            rm.month,
            ac.category_id,
            ac.category_name
        FROM report_months rm
        CROSS JOIN active_categories ac
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    mcc.category_id,
    mcc.category_name,
    mcc.year::text,
    TO_CHAR(TO_DATE(mcc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_category_combos mcc
LEFT JOIN monthly_transactions mt ON
    mcc.year = mt.year AND
    mcc.month = mt.month AND
    mcc.category_id = mt.category_id
ORDER BY 
    mcc.category_name ASC,
    mcc.year DESC,
    mcc.month DESC;

-- name: GetYearAmountCategorySuccess :many
WITH
    active_categories AS (
        SELECT category_id, name AS category_name
        FROM categories
        WHERE deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_category_combos AS (
        SELECT
            ry.year,
            ac.category_id,
            ac.category_name
        FROM report_years ry
        CROSS JOIN active_categories ac
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    ycc.category_id,
    ycc.category_name,
    ycc.year::text,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_category_combos ycc
LEFT JOIN yearly_transactions yt ON
    ycc.year = yt.year AND
    ycc.category_id = yt.category_id
ORDER BY 
    ycc.category_name ASC,
    ycc.year DESC;


-- name: GetMonthAmountCategoryFailed :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_categories AS (
        SELECT category_id, name AS category_name
        FROM categories
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
    month_category_combos AS (
        SELECT
            rm.year,
            rm.month,
            ac.category_id,
            ac.category_name
        FROM report_months rm
        CROSS JOIN active_categories ac
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    mcc.category_id,
    mcc.category_name,
    mcc.year::text,
    TO_CHAR(TO_DATE(mcc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_category_combos mcc
LEFT JOIN monthly_transactions mt ON
    mcc.year = mt.year AND
    mcc.month = mt.month AND
    mcc.category_id = mt.category_id
ORDER BY 
    mcc.category_name ASC,
    mcc.year DESC,
    mcc.month DESC;

-- name: GetYearAmountCategoryFailed :many
WITH
    active_categories AS (
        SELECT category_id, name AS category_name
        FROM categories
        WHERE deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_category_combos AS (
        SELECT
            ry.year,
            ac.category_id,
            ac.category_name
        FROM report_years ry
        CROSS JOIN active_categories ac
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    ycc.category_id,
    ycc.category_name,
    ycc.year::text,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_category_combos ycc
LEFT JOIN yearly_transactions yt ON
    ycc.year = yt.year AND
    ycc.category_id = yt.category_id
ORDER BY 
    ycc.category_name ASC,
    ycc.year DESC;


-- name: GetMonthMethodCategoriesSuccess :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    active_category_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
    ),
    all_months AS (
        SELECT 
            generate_series(
                (SELECT start_date FROM date_range),
                (SELECT end_date FROM date_range),
                interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT 
            am.activity_month,
            acm.category_id,
            acm.category_name,
            acm.payment_method
        FROM all_months am
        CROSS JOIN active_category_methods acm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.category_id = ms.category_id AND
    ac.payment_method = ms.payment_method
WHERE 
    EXISTS (
        SELECT 1 FROM active_category_methods 
        WHERE category_id = ac.category_id 
        AND payment_method = ac.payment_method
    )
ORDER BY 
    ac.activity_month,
    ac.category_name,
    ac.payment_method;

-- name: GetYearMethodCategoriesSuccess :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    active_category_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
    ),
    all_years AS (
        SELECT 
            generate_series(
                (SELECT start_year FROM year_range),
                (SELECT end_year FROM year_range)
            )::text AS year
    ),
    all_combinations AS (
        SELECT 
            ay.year,
            acm.category_id,
            acm.category_name,
            acm.payment_method
        FROM all_years ay
        CROSS JOIN active_category_methods acm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.category_id = ys.category_id AND
    ac.payment_method = ys.payment_method
WHERE 
    EXISTS (
        SELECT 1 FROM active_category_methods 
        WHERE category_id = ac.category_id 
        AND payment_method = ac.payment_method
    )
ORDER BY 
    ac.year DESC,
    ac.category_name,
    ac.payment_method;



-- name: GetMonthMethodCategoriesFailed :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    active_category_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
    ),
    all_months AS (
        SELECT 
            generate_series(
                (SELECT start_date FROM date_range),
                (SELECT end_date FROM date_range),
                interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT 
            am.activity_month,
            acm.category_id,
            acm.category_name,
            acm.payment_method
        FROM all_months am
        CROSS JOIN active_category_methods acm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.category_id = ms.category_id AND
    ac.payment_method = ms.payment_method
WHERE 
    EXISTS (
        SELECT 1 FROM active_category_methods 
        WHERE category_id = ac.category_id 
        AND payment_method = ac.payment_method
    )
ORDER BY 
    ac.activity_month,
    ac.category_name,
    ac.payment_method;

-- name: GetYearMethodCategoriesFailed :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    active_category_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
    ),
    all_years AS (
        SELECT 
            generate_series(
                (SELECT start_year FROM year_range),
                (SELECT end_year FROM year_range)
            )::text AS year
    ),
    all_combinations AS (
        SELECT 
            ay.year,
            acm.category_id,
            acm.category_name,
            acm.payment_method
        FROM all_years ay
        CROSS JOIN active_category_methods acm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.category_id = ys.category_id AND
    ac.payment_method = ys.payment_method
WHERE 
    EXISTS (
        SELECT 1 FROM active_category_methods 
        WHERE category_id = ac.category_id 
        AND payment_method = ac.payment_method
    )
ORDER BY 
    ac.year DESC,
    ac.category_name,
    ac.payment_method;



-- name: GetMonthAmountCategorySuccessById :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_categories AS (
        SELECT 
            c.category_id,
            c.name AS category_name
        FROM categories c
        WHERE c.deleted_at IS NULL
        AND c.category_id = $5 
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
    month_category_combos AS (
        SELECT
            rm.year,
            rm.month,
            ac.category_id,
            ac.category_name
        FROM report_months rm
        CROSS JOIN active_categories ac
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND c.category_id = $5 
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    mcc.category_id,
    mcc.category_name,
    mcc.year::text,
    TO_CHAR(TO_DATE(mcc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_category_combos mcc
LEFT JOIN monthly_transactions mt ON
    mcc.year = mt.year AND
    mcc.month = mt.month AND
    mcc.category_id = mt.category_id
ORDER BY 
    mcc.category_name ASC,
    mcc.year DESC,
    mcc.month DESC;



-- name: GetYearAmountCategorySuccessById :many
WITH
    active_categories AS (
        SELECT 
            c.category_id,
            c.name AS category_name
        FROM categories c
        WHERE c.deleted_at IS NULL
        AND c.category_id = $2  
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_category_combos AS (
        SELECT
            ry.year,
            ac.category_id,
            ac.category_name
        FROM report_years ry
        CROSS JOIN active_categories ac
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
            AND c.category_id = $2  
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    ycc.category_id,
    ycc.category_name,
    ycc.year::text,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_category_combos ycc
LEFT JOIN yearly_transactions yt ON
    ycc.year = yt.year AND
    ycc.category_id = yt.category_id
ORDER BY 
    ycc.category_name ASC,
    ycc.year DESC;




-- name: GetMonthAmountCategoryFailedById :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_categories AS (
        SELECT 
            c.category_id,
            c.name AS category_name
        FROM categories c
        WHERE c.deleted_at IS NULL
        AND c.category_id = $5 
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
    month_category_combos AS (
        SELECT
            rm.year,
            rm.month,
            ac.category_id,
            ac.category_name
        FROM report_months rm
        CROSS JOIN active_categories ac
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
            AND c.category_id = $5 
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    mcc.category_id,
    mcc.category_name,
    mcc.year::text,
    TO_CHAR(TO_DATE(mcc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_category_combos mcc
LEFT JOIN monthly_transactions mt ON
    mcc.year = mt.year AND
    mcc.month = mt.month AND
    mcc.category_id = mt.category_id
ORDER BY 
    mcc.category_name ASC,
    mcc.year DESC,
    mcc.month DESC;


-- name: GetYearAmountCategoryFailedById :many
WITH
    active_categories AS (
        SELECT 
            c.category_id,
            c.name AS category_name
        FROM categories c
        WHERE c.deleted_at IS NULL
        AND c.category_id = $2  
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_category_combos AS (
        SELECT
            ry.year,
            ac.category_id,
            ac.category_name
        FROM report_years ry
        CROSS JOIN active_categories ac
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
            AND c.category_id = $2  
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    ycc.category_id,
    ycc.category_name,
    ycc.year::text,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_category_combos ycc
LEFT JOIN yearly_transactions yt ON
    ycc.year = yt.year AND
    ycc.category_id = yt.category_id
ORDER BY 
    ycc.category_name ASC,
    ycc.year DESC;

-- name: GetMonthMethodCategoriesSuccessById :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    category_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND c.category_id = $2  
    ),
    all_months AS (
        SELECT 
            generate_series(
                (SELECT start_date FROM date_range),
                (SELECT end_date FROM date_range),
                interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT 
            am.activity_month,
            cm.category_id,
            cm.category_name,
            cm.payment_method
        FROM all_months am
        CROSS JOIN category_methods cm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND c.category_id = $2  
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.category_id = ms.category_id AND
    ac.payment_method = ms.payment_method
ORDER BY 
    ac.activity_month,
    ac.payment_method;

-- name: GetYearMethodCategoriesSuccessById :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    category_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND c.category_id = $2 
    ),
    all_years AS (
        SELECT 
            generate_series(
                (SELECT start_year FROM year_range),
                (SELECT end_year FROM year_range)
            )::text AS year
    ),
    all_combinations AS (
        SELECT 
            ay.year,
            cm.category_id,
            cm.category_name,
            cm.payment_method
        FROM all_years ay
        CROSS JOIN category_methods cm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND c.category_id = $2  
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.category_id = ys.category_id AND
    ac.payment_method = ys.payment_method
ORDER BY 
    ac.year DESC,
    ac.payment_method;


-- name: GetMonthMethodCategoriesFailedById :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    category_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND c.category_id = $2  
    ),
    all_months AS (
        SELECT 
            generate_series(
                (SELECT start_date FROM date_range),
                (SELECT end_date FROM date_range),
                interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT 
            am.activity_month,
            cm.category_id,
            cm.category_name,
            cm.payment_method
        FROM all_months am
        CROSS JOIN category_methods cm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
            AND c.category_id = $2  
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.category_id = ms.category_id AND
    ac.payment_method = ms.payment_method
ORDER BY 
    ac.activity_month,
    ac.payment_method;

-- name: GetYearMethodCategoriesFailedById :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    category_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND c.category_id = $2 
    ),
    all_years AS (
        SELECT 
            generate_series(
                (SELECT start_year FROM year_range),
                (SELECT end_year FROM year_range)
            )::text AS year
    ),
    all_combinations AS (
        SELECT 
            ay.year,
            cm.category_id,
            cm.category_name,
            cm.payment_method
        FROM all_years ay
        CROSS JOIN category_methods cm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
            AND c.category_id = $2  
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.category_id = ys.category_id AND
    ac.payment_method = ys.payment_method
ORDER BY 
    ac.year DESC,
    ac.payment_method;



-- name: GetMonthAmountCategorySuccessByMerchant :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    merchant_categories AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name
        FROM vouchers v
        JOIN categories c ON v.category_id = c.category_id
        WHERE 
            v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND v.merchant_id = $5
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
    month_category_combos AS (
        SELECT
            rm.year,
            rm.month,
            mc.category_id,
            mc.category_name
        FROM report_months rm
        CROSS JOIN merchant_categories mc
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND v.merchant_id = $5
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    mcc.category_id,
    mcc.category_name,
    mcc.year::text,
    TO_CHAR(TO_DATE(mcc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_category_combos mcc
LEFT JOIN monthly_transactions mt ON
    mcc.year = mt.year AND
    mcc.month = mt.month AND
    mcc.category_id = mt.category_id
ORDER BY 
    mcc.category_name ASC,
    mcc.year DESC,
    mcc.month DESC;

-- name: GetYearAmountCategorySuccessByMerchant :many
WITH
    merchant_categories AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name
        FROM vouchers v
        JOIN categories c ON v.category_id = c.category_id
        WHERE 
            v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND v.merchant_id = $2
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_category_combos AS (
        SELECT
            ry.year,
            mc.category_id,
            mc.category_name
        FROM report_years ry
        CROSS JOIN merchant_categories mc
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
            AND v.merchant_id = $2
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    ycc.category_id,
    ycc.category_name,
    ycc.year::text,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_category_combos ycc
LEFT JOIN yearly_transactions yt ON
    ycc.year = yt.year AND
    ycc.category_id = yt.category_id
ORDER BY 
    ycc.category_name ASC,
    ycc.year DESC;

-- name: GetMonthAmountCategoryFailedByMerchant :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    merchant_categories AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name
        FROM vouchers v
        JOIN categories c ON v.category_id = c.category_id
        WHERE 
            v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND v.merchant_id = $5
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
    month_category_combos AS (
        SELECT
            rm.year,
            rm.month,
            mc.category_id,
            mc.category_name
        FROM report_months rm
        CROSS JOIN merchant_categories mc
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
            AND v.merchant_id = $5
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    mcc.category_id,
    mcc.category_name,
    mcc.year::text,
    TO_CHAR(TO_DATE(mcc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_category_combos mcc
LEFT JOIN monthly_transactions mt ON
    mcc.year = mt.year AND
    mcc.month = mt.month AND
    mcc.category_id = mt.category_id
ORDER BY 
    mcc.category_name ASC,
    mcc.year DESC,
    mcc.month DESC;

-- name: GetYearAmountCategoryFailedByMerchant :many
WITH
    merchant_categories AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name
        FROM vouchers v
        JOIN categories c ON v.category_id = c.category_id
        WHERE 
            v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND v.merchant_id = $2
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_category_combos AS (
        SELECT
            ry.year,
            mc.category_id,
            mc.category_name
        FROM report_years ry
        CROSS JOIN merchant_categories mc
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            c.category_id,
            c.name AS category_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
            AND v.merchant_id = $2
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name
    )
SELECT
    ycc.category_id,
    ycc.category_name,
    ycc.year::text,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_category_combos ycc
LEFT JOIN yearly_transactions yt ON
    ycc.year = yt.year AND
    ycc.category_id = yt.category_id
ORDER BY 
    ycc.category_name ASC,
    ycc.year DESC;

-- name: GetMonthMethodCategoriesSuccessByMerchant :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    merchant_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.merchant_id = $2  
    ),
    all_months AS (
        SELECT 
            generate_series(
                (SELECT start_date FROM date_range),
                (SELECT end_date FROM date_range),
                interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT 
            am.activity_month,
            mm.category_id,
            mm.category_name,
            mm.payment_method
        FROM all_months am
        CROSS JOIN merchant_methods mm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $2  
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.category_id = ms.category_id AND
    ac.payment_method = ms.payment_method
ORDER BY 
    ac.activity_month,
    ac.category_name,
    ac.payment_method;

-- name: GetYearMethodCategoriesSuccessByMerchant :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    merchant_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.merchant_id = $2 
    ),
    all_years AS (
        SELECT 
            generate_series(
                (SELECT start_year FROM year_range),
                (SELECT end_year FROM year_range)
            )::text AS year
    ),
    all_combinations AS (
        SELECT 
            ay.year,
            mm.category_id,
            mm.category_name,
            mm.payment_method
        FROM all_years ay
        CROSS JOIN merchant_methods mm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $2  
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.category_id = ys.category_id AND
    ac.payment_method = ys.payment_method
ORDER BY 
    ac.year DESC,
    ac.category_name,
    ac.payment_method;



-- name: GetMonthMethodCategoriesFailedByMerchant :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    merchant_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.merchant_id = $2  
    ),
    all_months AS (
        SELECT 
            generate_series(
                (SELECT start_date FROM date_range),
                (SELECT end_date FROM date_range),
                interval '1 month'
            )::date AS activity_month
    ),
    all_combinations AS (
        SELECT 
            am.activity_month,
            mm.category_id,
            mm.category_name,
            mm.payment_method
        FROM all_months am
        CROSS JOIN merchant_methods mm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $2  
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.category_id = ms.category_id AND
    ac.payment_method = ms.payment_method
ORDER BY 
    ac.activity_month,
    ac.category_name,
    ac.payment_method;

-- name: GetYearMethodCategoriesFailedByMerchant :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    merchant_methods AS (
        SELECT DISTINCT
            c.category_id,
            c.name AS category_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.merchant_id = $2 
    ),
    all_years AS (
        SELECT 
            generate_series(
                (SELECT start_year FROM year_range),
                (SELECT end_year FROM year_range)
            )::text AS year
    ),
    all_combinations AS (
        SELECT 
            ay.year,
            mm.category_id,
            mm.category_name,
            mm.payment_method
        FROM all_years ay
        CROSS JOIN merchant_methods mm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            c.category_id,
            c.name AS category_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN categories c ON v.category_id = c.category_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND c.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $2  
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            c.category_id,
            c.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.category_id,
    ac.category_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.category_id = ys.category_id AND
    ac.payment_method = ys.payment_method
ORDER BY 
    ac.year DESC,
    ac.category_name,
    ac.payment_method;

-- Create Category
-- name: CreateCategory :one
INSERT INTO categories (name) VALUES ($1) RETURNING *;

-- Get Category by ID
-- name: GetCategoryByID :one
SELECT
    category_id,
    name,
    created_at,
    updated_at,
    deleted_at
FROM categories
WHERE
    category_id = $1
    AND deleted_at IS NULL;

-- Update Category
-- name: UpdateCategory :one
UPDATE categories
SET
    name = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE
    category_id = $1
    AND deleted_at IS NULL
RETURNING
    *;

-- Trash Category (Soft Delete)
-- name: TrashCategory :one
UPDATE categories
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    category_id = $1
    AND deleted_at IS NULL
RETURNING
    *;

-- Restore Trashed Category
-- name: RestoreCategory :one
UPDATE categories
SET
    deleted_at = NULL
WHERE
    category_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    *;

-- Delete Category Permanently
-- name: DeleteCategoryPermanently :exec
DELETE FROM categories
WHERE
    category_id = $1
    AND deleted_at IS NOT NULL;

-- Restore All Trashed Categories
-- name: RestoreAllCategories :exec
UPDATE categories SET deleted_at = NULL WHERE deleted_at IS NOT NULL;

-- Delete All Trashed Categories Permanently
-- name: DeleteAllPermanentCategories :exec
DELETE FROM categories WHERE deleted_at IS NOT NULL;