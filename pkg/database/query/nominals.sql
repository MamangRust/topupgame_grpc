-- Get Nominals with Pagination and Total Count
-- name: GetNominals :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM nominals
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Active Nominals with Pagination and Total Count
-- name: GetNominalsActive :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM nominals
WHERE deleted_at IS NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- Get Trashed Nominals with Pagination and Total Count
-- name: GetNominalsTrashed :many
SELECT
    *,
    COUNT(*) OVER() AS total_count
FROM nominals
WHERE deleted_at IS NOT NULL
AND ($1::TEXT IS NULL OR name ILIKE '%' || $1 || '%')
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;



-- name: GetMonthAmountNominalsSuccess :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_nominals AS (
        SELECT nominal_id, name AS nominal_name
        FROM nominals
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
    month_nominal_combos AS (
        SELECT
            rm.year,
            rm.month,
            an.nominal_id,
            an.nominal_name
        FROM report_months rm
        CROSS JOIN active_nominals an
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(n.price * n.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    mnc.nominal_id,
    mnc.nominal_name,
    mnc.year::text,
    TO_CHAR(TO_DATE(mnc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_nominal_combos mnc
LEFT JOIN monthly_transactions mt ON
    mnc.year = mt.year AND
    mnc.month = mt.month AND
    mnc.nominal_id = mt.nominal_id
ORDER BY 
    mnc.nominal_name ASC,
    mnc.year DESC,
    mnc.month DESC;

-- name: GetYearAmountNominalsSuccess :many
WITH
    active_nominals AS (
        SELECT nominal_id, name AS nominal_name
        FROM nominals
        WHERE deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_nominal_combos AS (
        SELECT
            ry.year,
            an.nominal_id,
            an.nominal_name
        FROM report_years ry
        CROSS JOIN active_nominals an
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(n.price * n.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    ync.nominal_id,
    ync.nominal_name,
    ync.year::text,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_nominal_combos ync
LEFT JOIN yearly_transactions yt ON
    ync.year = yt.year AND
    ync.nominal_id = yt.nominal_id
ORDER BY 
    ync.nominal_name ASC,
    ync.year DESC;

-- name: GetMonthAmountNominalsFailed :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_nominals AS (
        SELECT nominal_id, name AS nominal_name
        FROM nominals
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
    month_nominal_combos AS (
        SELECT
            rm.year,
            rm.month,
            an.nominal_id,
            an.nominal_name
        FROM report_months rm
        CROSS JOIN active_nominals an
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(n.price * n.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    mnc.nominal_id,
    mnc.nominal_name,
    mnc.year::text,
    TO_CHAR(TO_DATE(mnc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_nominal_combos mnc
LEFT JOIN monthly_transactions mt ON
    mnc.year = mt.year AND
    mnc.month = mt.month AND
    mnc.nominal_id = mt.nominal_id
ORDER BY 
    mnc.nominal_name ASC,
    mnc.year DESC,
    mnc.month DESC;

-- name: GetYearAmountNominalsFailed :many
WITH
    active_nominals AS (
        SELECT nominal_id, name AS nominal_name
        FROM nominals
        WHERE deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_nominal_combos AS (
        SELECT
            ry.year,
            an.nominal_id,
            an.nominal_name
        FROM report_years ry
        CROSS JOIN active_nominals an
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(n.price * n.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    ync.nominal_id,
    ync.nominal_name,
    ync.year::text,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_nominal_combos ync
LEFT JOIN yearly_transactions yt ON
    ync.year = yt.year AND
    ync.nominal_id = yt.nominal_id
ORDER BY 
    ync.nominal_name ASC,
    ync.year DESC;


-- name: GetMonthMethodNominalsSuccess :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    active_nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
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
            anm.nominal_id,
            anm.nominal_name,
            anm.payment_method
        FROM all_months am
        CROSS JOIN active_nominal_methods anm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.nominal_id = ms.nominal_id AND
    ac.payment_method = ms.payment_method
ORDER BY 
    ac.activity_month,
    ac.nominal_name,
    ac.payment_method;


-- name: GetYearMethodNominalsSuccess :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    active_nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
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
            anm.nominal_id,
            anm.nominal_name,
            anm.payment_method
        FROM all_years ay
        CROSS JOIN active_nominal_methods anm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.nominal_id = ys.nominal_id AND
    ac.payment_method = ys.payment_method
ORDER BY 
    ac.year DESC,
    ac.nominal_name,
    ac.payment_method;



-- name: GetMonthMethodNominalsFailed :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    active_nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
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
            anm.nominal_id,
            anm.nominal_name,
            anm.payment_method
        FROM all_months am
        CROSS JOIN active_nominal_methods anm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.nominal_id = ms.nominal_id AND
    ac.payment_method = ms.payment_method
ORDER BY 
    ac.activity_month,
    ac.nominal_name,
    ac.payment_method;


-- name: GetYearMethodNominalsFailed :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    active_nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
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
            anm.nominal_id,
            anm.nominal_name,
            anm.payment_method
        FROM all_years ay
        CROSS JOIN active_nominal_methods anm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.nominal_id = ys.nominal_id AND
    ac.payment_method = ys.payment_method
ORDER BY 
    ac.year DESC,
    ac.nominal_name,
    ac.payment_method;



-- name: GetMonthAmountNominalsSuccessById :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_nominal AS (
        SELECT 
            n.nominal_id,
            n.name AS nominal_name
        FROM nominals n
        WHERE 
            n.deleted_at IS NULL
            AND n.nominal_id = $5  
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
    month_nominal_combos AS (
        SELECT
            rm.year,
            rm.month,
            tn.nominal_id,
            tn.nominal_name
        FROM report_months rm
        CROSS JOIN target_nominal tn
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(n.price * n.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND n.nominal_id = $5  
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    mnc.nominal_id,
    mnc.nominal_name,
    mnc.year::text,
    TO_CHAR(TO_DATE(mnc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_nominal_combos mnc
LEFT JOIN monthly_transactions mt ON
    mnc.year = mt.year AND
    mnc.month = mt.month AND
    mnc.nominal_id = mt.nominal_id
ORDER BY 
    mnc.year DESC,
    mnc.month DESC;



-- name: GetYearAmountNominalsSuccessById :many
WITH
    target_nominal AS (
        SELECT nominal_id, name AS nominal_name
        FROM nominals
        WHERE 
            deleted_at IS NULL
            AND nominals.nominal_id = $2  
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_nominal_combos AS (
        SELECT
            ry.year,
            tn.nominal_id,
            tn.nominal_name
        FROM report_years ry
        CROSS JOIN target_nominal tn
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(n.price * n.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
            AND n.nominal_id = $2 
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    ync.nominal_id,
    ync.nominal_name,
    ync.year::text,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_nominal_combos ync
LEFT JOIN yearly_transactions yt ON
    ync.year = yt.year AND
    ync.nominal_id = yt.nominal_id
ORDER BY 
    ync.year DESC;

 

-- name: GetMonthAmountNominalsFailedById :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_nominal AS (
        SELECT nominal_id, name AS nominal_name
        FROM nominals
        WHERE 
            deleted_at IS NULL
            AND nominals.nominal_id = $5 
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
    month_nominal_combos AS (
        SELECT
            rm.year,
            rm.month,
            tn.nominal_id,
            tn.nominal_name
        FROM report_months rm
        CROSS JOIN target_nominal tn
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(n.price * n.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND n.nominal_id = $5  
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    mnc.nominal_id,
    mnc.nominal_name,
    mnc.year::text,
    TO_CHAR(TO_DATE(mnc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_nominal_combos mnc
LEFT JOIN monthly_transactions mt ON
    mnc.year = mt.year AND
    mnc.month = mt.month AND
    mnc.nominal_id = mt.nominal_id
ORDER BY 
    mnc.year DESC,
    mnc.month DESC;


-- name: GetYearAmountNominalsFailedById :many
WITH
    target_nominal AS (
        SELECT nominal_id, name AS nominal_name
        FROM nominals
        WHERE 
            deleted_at IS NULL
            AND nominals.nominal_id = $2  
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_nominal_combos AS (
        SELECT
            ry.year,
            tn.nominal_id,
            tn.nominal_name
        FROM report_years ry
        CROSS JOIN target_nominal tn
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(n.price * n.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
            AND n.nominal_id = $2  
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    ync.nominal_id,
    ync.nominal_name,
    ync.year::text,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_nominal_combos ync
LEFT JOIN yearly_transactions yt ON
    ync.year = yt.year AND
    ync.nominal_id = yt.nominal_id
ORDER BY 
    ync.year DESC;

 


-- name: GetMonthMethodNominalsSuccessById :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND n.nominal_id = $2 
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
            nm.nominal_id,
            nm.nominal_name,
            nm.payment_method
        FROM all_months am
        CROSS JOIN nominal_methods nm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND n.nominal_id = $2  
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month AND
    ac.nominal_id = mt.nominal_id AND
    ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month,
    ac.nominal_name,
    ac.payment_method;


-- name: GetYearMethodNominalsSuccessById :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND n.nominal_id = $2 
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
            nm.nominal_id,
            nm.nominal_name,
            nm.payment_method
        FROM all_years ay
        CROSS JOIN nominal_methods nm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND n.nominal_id = $2 
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year AND
    ac.nominal_id = yt.nominal_id AND
    ac.payment_method = yt.payment_method
ORDER BY 
    ac.year DESC,
    ac.nominal_name,
    ac.payment_method;




-- name: GetMonthMethodNominalsFailedById :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND n.nominal_id = $2 
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
            nm.nominal_id,
            nm.nominal_name,
            nm.payment_method
        FROM all_months am
        CROSS JOIN nominal_methods nm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND n.nominal_id = $2  
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month AND
    ac.nominal_id = mt.nominal_id AND
    ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month,
    ac.nominal_name,
    ac.payment_method;


-- name: GetYearMethodNominalsFailedById :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND n.nominal_id = $2 
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
            nm.nominal_id,
            nm.nominal_name,
            nm.payment_method
        FROM all_years ay
        CROSS JOIN nominal_methods nm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND n.nominal_id = $2 
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year AND
    ac.nominal_id = yt.nominal_id AND
    ac.payment_method = yt.payment_method
ORDER BY 
    ac.year DESC,
    ac.nominal_name,
    ac.payment_method;


-- name: GetMonthAmountNominalsSuccessByMerchant :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_nominals AS (
        SELECT DISTINCT n.nominal_id, n.name AS nominal_name
        FROM nominals n
        JOIN transactions t ON t.nominal_id = n.nominal_id
        WHERE 
            n.deleted_at IS NULL
            AND t.deleted_at IS NULL
            AND t.merchant_id = $5
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
    month_nominal_combos AS (
        SELECT
            rm.year,
            rm.month,
            tn.nominal_id,
            tn.nominal_name
        FROM report_months rm
        CROSS JOIN target_nominals tn
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end)
            OR (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $5
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    mnc.nominal_id,
    mnc.nominal_name,
    mnc.year::text AS year,
    TO_CHAR(TO_DATE(mnc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_nominal_combos mnc
LEFT JOIN monthly_transactions mt 
    ON mnc.year = mt.year
    AND mnc.month = mt.month
    AND mnc.nominal_id = mt.nominal_id
ORDER BY 
    mnc.year DESC,
    mnc.month DESC;



-- name: GetYearAmountNominalsSuccessByMerchant :many
WITH
    target_nominals AS (
        SELECT DISTINCT n.nominal_id, n.name AS nominal_name
        FROM nominals n
        JOIN transactions t ON t.nominal_id = n.nominal_id
        WHERE 
            n.deleted_at IS NULL
            AND t.deleted_at IS NULL
            AND t.merchant_id = $2
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT ($1::integer - 1) AS year
    ),
    year_nominal_combos AS (
        SELECT
            ry.year,
            tn.nominal_id,
            tn.nominal_name
        FROM report_years ry
        CROSS JOIN target_nominals tn
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $2
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, ($1::integer - 1))
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    ync.nominal_id,
    ync.nominal_name,
    ync.year::text AS year,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_nominal_combos ync
LEFT JOIN yearly_transactions yt
    ON ync.year = yt.year
    AND ync.nominal_id = yt.nominal_id
ORDER BY 
    ync.year DESC;


-- name: GetMonthAmountNominalsFailedByMerchant :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_nominals AS (
        SELECT DISTINCT n.nominal_id, n.name AS nominal_name
        FROM nominals n
        JOIN transactions t ON t.nominal_id = n.nominal_id
        WHERE 
            n.deleted_at IS NULL
            AND t.deleted_at IS NULL
            AND t.merchant_id = $5
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
    month_nominal_combos AS (
        SELECT
            rm.year,
            rm.month,
            tn.nominal_id,
            tn.nominal_name
        FROM report_months rm
        CROSS JOIN target_nominals tn
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end)
            OR (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $5
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    mnc.nominal_id,
    mnc.nominal_name,
    mnc.year::text AS year,
    TO_CHAR(TO_DATE(mnc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_nominal_combos mnc
LEFT JOIN monthly_transactions mt 
    ON mnc.year = mt.year
    AND mnc.month = mt.month
    AND mnc.nominal_id = mt.nominal_id
ORDER BY 
    mnc.year DESC,
    mnc.month DESC;



-- name: GetYearAmountNominalsFailedByMerchant :many
WITH
    target_nominals AS (
        SELECT DISTINCT n.nominal_id, n.name AS nominal_name
        FROM nominals n
        JOIN transactions t ON t.nominal_id = n.nominal_id
        WHERE 
            n.deleted_at IS NULL
            AND t.deleted_at IS NULL
            AND t.merchant_id = $2
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT ($1::integer - 1) AS year
    ),
    year_nominal_combos AS (
        SELECT
            ry.year,
            tn.nominal_id,
            tn.nominal_name
        FROM report_years ry
        CROSS JOIN target_nominals tn
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            n.nominal_id,
            n.name AS nominal_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $2
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, ($1::integer - 1))
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name
    )
SELECT
    ync.nominal_id,
    ync.nominal_name,
    ync.year::text AS year,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_nominal_combos ync
LEFT JOIN yearly_transactions yt
    ON ync.year = yt.year
    AND ync.nominal_id = yt.nominal_id
ORDER BY 
    ync.year DESC;





-- name: GetMonthMethodNominalsSuccessByMerchant :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND n.nominal_id = $2
            AND t.merchant_id = $3
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
            nm.nominal_id,
            nm.nominal_name,
            nm.payment_method
        FROM all_months am
        CROSS JOIN nominal_methods nm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND n.nominal_id = $2
            AND t.merchant_id = $3
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month AND
    ac.nominal_id = mt.nominal_id AND
    ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month,
    ac.nominal_name,
    ac.payment_method;


-- name: GetYearMethodNominalsSuccessByMerchant :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND n.nominal_id = $2
            AND t.merchant_id = $3
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
            nm.nominal_id,
            nm.nominal_name,
            nm.payment_method
        FROM all_years ay
        CROSS JOIN nominal_methods nm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'success'
            AND n.nominal_id = $2
            AND t.merchant_id = $3
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year AND
    ac.nominal_id = yt.nominal_id AND
    ac.payment_method = yt.payment_method
ORDER BY 
    ac.year DESC,
    ac.nominal_name,
    ac.payment_method;




-- name: GetMonthMethodNominalsFailedByMerchant :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND n.nominal_id = $2
            AND t.merchant_id = $3
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
            nm.nominal_id,
            nm.nominal_name,
            nm.payment_method
        FROM all_months am
        CROSS JOIN nominal_methods nm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id 
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND n.nominal_id = $2
            AND t.merchant_id = $3
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month AND
    ac.nominal_id = mt.nominal_id AND
    ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month,
    ac.nominal_name,
    ac.payment_method;


-- name: GetYearMethodNominalsFailedByMerchant :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    nominal_methods AS (
        SELECT DISTINCT
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND n.nominal_id = $2
            AND t.merchant_id = $3
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
            nm.nominal_id,
            nm.nominal_name,
            nm.payment_method
        FROM all_years ay
        CROSS JOIN nominal_methods nm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            n.nominal_id,
            n.name AS nominal_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(n.price * n.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN nominals n ON t.nominal_id = n.nominal_id  
        WHERE
            t.deleted_at IS NULL
            AND n.deleted_at IS NULL
            AND t.status = 'failed'
            AND n.nominal_id = $2
            AND t.merchant_id = $3
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            n.nominal_id,
            n.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.nominal_id,
    ac.nominal_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year AND
    ac.nominal_id = yt.nominal_id AND
    ac.payment_method = yt.payment_method
ORDER BY 
    ac.year DESC,
    ac.nominal_name,
    ac.payment_method;


-- Create Nominal
-- name: CreateNominal :one
INSERT INTO nominals (name, quantity, price, voucher_id)
VALUES ($1, $2, $3, $4)
  RETURNING *;


-- Get Nominal by ID
-- name: GetNominalByID :one
SELECT *
FROM nominals
WHERE nominal_id = $1
  AND deleted_at IS NULL;


-- Update Nominal
-- name: UpdateNominal :one
UPDATE nominals
SET name = $2,
    quantity = $3,
    price = $4,
    voucher_id = $5,
    updated_at = CURRENT_TIMESTAMP
WHERE nominal_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- name: DecreaseNominalQuantity :exec
UPDATE nominals 
SET quantity = quantity - $1 
WHERE nominal_id = $2 AND quantity >= $1;



-- Trash Nominal (Soft Delete)
-- name: TrashNominal :one
UPDATE nominals
SET deleted_at = CURRENT_TIMESTAMP
WHERE nominal_id = $1
  AND deleted_at IS NULL
  RETURNING *;


-- Restore Trashed Nominal
-- name: RestoreNominal :one
UPDATE nominals
SET deleted_at = NULL
WHERE nominal_id = $1
  AND deleted_at IS NOT NULL
  RETURNING *;


-- Delete Nominal Permanently
-- name: DeleteNominalPermanently :exec
DELETE FROM nominals WHERE nominal_id = $1 AND deleted_at IS NOT NULL;


-- Restore All Trashed Nominals
-- name: RestoreAllNominals :exec
UPDATE nominals
SET deleted_at = NULL
WHERE deleted_at IS NOT NULL;


-- Delete All Trashed Nominals Permanently
-- name: DeleteAllPermanentNominals :exec
DELETE FROM nominals WHERE deleted_at IS NOT NULL;
