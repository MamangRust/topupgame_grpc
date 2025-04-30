-- Get Vouchers with Pagination and Total Count
-- name: GetVouchers :many
SELECT *, COUNT(*) OVER () AS total_count
FROM vouchers
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

-- Get Active Vouchers with Pagination and Total Count
-- name: GetVouchersActive :many
SELECT *, COUNT(*) OVER () AS total_count
FROM vouchers
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

-- Get Trashed Vouchers with Pagination and Total Count
-- name: GetVouchersTrashed :many
SELECT *, COUNT(*) OVER () AS total_count
FROM vouchers
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

-- name: GetMonthAmountVouchersSuccess :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_vouchers AS (
        SELECT voucher_id, name AS voucher_name
        FROM vouchers
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
    month_voucher_combos AS (
        SELECT
            rm.year,
            rm.month,
            av.voucher_id,
            av.voucher_name
        FROM report_months rm
        CROSS JOIN active_vouchers av
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(v.price * v.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    mvc.voucher_id,
    mvc.voucher_name,
    mvc.year::text,
    TO_CHAR(TO_DATE(mvc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_voucher_combos mvc
LEFT JOIN monthly_transactions mt ON
    mvc.year = mt.year AND
    mvc.month = mt.month AND
    mvc.voucher_id = mt.voucher_id
ORDER BY 
    mvc.voucher_name ASC,
    mvc.year DESC,
    mvc.month DESC;



-- name: GetYearAmountVouchersSuccess :many
WITH
    active_vouchers AS (
        SELECT voucher_id, name AS voucher_name
        FROM vouchers
        WHERE deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_voucher_combos AS (
        SELECT
            ry.year,
            av.voucher_id,
            av.voucher_name
        FROM report_years ry
        CROSS JOIN active_vouchers av
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(v.price * v.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    yvc.voucher_id,
    yvc.voucher_name,
    yvc.year::text,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_voucher_combos yvc
LEFT JOIN yearly_transactions yt ON
    yvc.year = yt.year AND
    yvc.voucher_id = yt.voucher_id
ORDER BY 
    yvc.voucher_name ASC,
    yvc.year DESC;


-- name: GetMonthAmountVouchersFailed :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    active_vouchers AS (
        SELECT voucher_id, name AS voucher_name
        FROM vouchers
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
    month_voucher_combos AS (
        SELECT
            rm.year,
            rm.month,
            av.voucher_id,
            av.voucher_name
        FROM report_months rm
        CROSS JOIN active_vouchers av
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(v.price * v.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    mvc.voucher_id,
    mvc.voucher_name,
    mvc.year::text,
    TO_CHAR(TO_DATE(mvc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_voucher_combos mvc
LEFT JOIN monthly_transactions mt ON
    mvc.year = mt.year AND
    mvc.month = mt.month AND
    mvc.voucher_id = mt.voucher_id
ORDER BY 
    mvc.voucher_name ASC,
    mvc.year DESC,
    mvc.month DESC;


-- name: GetYearAmountVouchersFailed :many
WITH
    active_vouchers AS (
        SELECT voucher_id, name AS voucher_name
        FROM vouchers
        WHERE deleted_at IS NULL
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_voucher_combos AS (
        SELECT
            ry.year,
            av.voucher_id,
            av.voucher_name
        FROM report_years ry
        CROSS JOIN active_vouchers av
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(v.price * v.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    yvc.voucher_id,
    yvc.voucher_name,
    yvc.year::text,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_voucher_combos yvc
LEFT JOIN yearly_transactions yt ON
    yvc.year = yt.year AND
    yvc.voucher_id = yt.voucher_id
ORDER BY 
    yvc.voucher_name ASC,
    yvc.year DESC;


-- name: GetMonthMethodVouchersSuccess :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    active_voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
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
            avm.voucher_id,
            avm.voucher_name,
            avm.payment_method
        FROM all_months am
        CROSS JOIN active_voucher_methods avm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(v.price * v.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.voucher_id = ms.voucher_id AND
    ac.payment_method = ms.payment_method
ORDER BY 
    ac.activity_month,
    ac.voucher_name,
    ac.payment_method;



-- name: GetYearMethodVouchersSuccess :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    active_voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
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
            avm.voucher_id,
            avm.voucher_name,
            avm.payment_method
        FROM all_years ay
        CROSS JOIN active_voucher_methods avm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(v.price * v.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.voucher_id = ys.voucher_id AND
    ac.payment_method = ys.payment_method
ORDER BY 
    ac.year DESC,
    ac.voucher_name,
    ac.payment_method;



-- name: GetMonthMethodVouchersFailed :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    active_voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
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
            avm.voucher_id,
            avm.voucher_name,
            avm.payment_method
        FROM all_months am
        CROSS JOIN active_voucher_methods avm
    ),
    monthly_stats AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(v.price * v.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(ms.total_transactions, 0) AS total_transactions,
    COALESCE(ms.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_stats ms ON 
    ac.activity_month = ms.activity_month AND
    ac.voucher_id = ms.voucher_id AND
    ac.payment_method = ms.payment_method
ORDER BY 
    ac.activity_month,
    ac.voucher_name,
    ac.payment_method;



-- name: GetYearMethodVouchersFailed :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    active_voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
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
            avm.voucher_id,
            avm.voucher_name,
            avm.payment_method
        FROM all_years ay
        CROSS JOIN active_voucher_methods avm
    ),
    yearly_stats AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(v.price * v.quantity), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(ys.total_transactions, 0) AS total_transactions,
    COALESCE(ys.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_stats ys ON 
    ac.year = ys.year AND
    ac.voucher_id = ys.voucher_id AND
    ac.payment_method = ys.payment_method
ORDER BY 
    ac.year DESC,
    ac.voucher_name,
    ac.payment_method;


-- name: GetMonthAmountVouchersSuccessById :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_voucher AS (
        SELECT 
            v.voucher_id,
            v.name AS voucher_name
        FROM vouchers v
        WHERE 
            v.deleted_at IS NULL
            AND v.voucher_id = $5  
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
    month_voucher_combos AS (
        SELECT
            rm.year,
            rm.month,
            tv.voucher_id,
            tv.voucher_name
        FROM report_months rm
        CROSS JOIN target_voucher tv
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(v.price * v.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND v.voucher_id = $5  
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    mvc.voucher_id,
    mvc.voucher_name,
    mvc.year::text,
    TO_CHAR(TO_DATE(mvc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_voucher_combos mvc
LEFT JOIN monthly_transactions mt ON
    mvc.year = mt.year AND
    mvc.month = mt.month AND
    mvc.voucher_id = mt.voucher_id
ORDER BY 
    mvc.year DESC,
    mvc.month DESC;


-- name: GetYearAmountVouchersSuccessById :many
WITH
    target_voucher AS (
        SELECT 
            v.voucher_id,
            v.name AS voucher_name
        FROM vouchers v
        WHERE 
            v.deleted_at IS NULL
            AND v.voucher_id = $2  
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_voucher_combos AS (
        SELECT
            ry.year,
            tv.voucher_id,
            tv.voucher_name
        FROM report_years ry
        CROSS JOIN target_voucher tv
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(v.price * v.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
            AND v.voucher_id = $2 
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    yvc.voucher_id,
    yvc.voucher_name,
    yvc.year::text,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_voucher_combos yvc
LEFT JOIN yearly_transactions yt ON
    yvc.year = yt.year AND
    yvc.voucher_id = yt.voucher_id
ORDER BY 
    yvc.year DESC;


-- name: GetMonthAmountVouchersFailedById :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_voucher AS (
        SELECT 
            v.voucher_id,
            v.name AS voucher_name
        FROM vouchers v
        WHERE 
            v.deleted_at IS NULL
            AND v.voucher_id = $5  
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
    month_voucher_combos AS (
        SELECT
            rm.year,
            rm.month,
            tv.voucher_id,
            tv.voucher_name
        FROM report_months rm
        CROSS JOIN target_voucher tv
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(v.price * v.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND v.voucher_id = $5  
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    mvc.voucher_id,
    mvc.voucher_name,
    mvc.year::text,
    TO_CHAR(TO_DATE(mvc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_voucher_combos mvc
LEFT JOIN monthly_transactions mt ON
    mvc.year = mt.year AND
    mvc.month = mt.month AND
    mvc.voucher_id = mt.voucher_id
ORDER BY 
    mvc.year DESC,
    mvc.month DESC;


-- name: GetYearAmountVouchersFailedById :many
WITH
    target_voucher AS (
        SELECT 
            v.voucher_id,
            v.name AS voucher_name
        FROM vouchers v
        WHERE 
            v.deleted_at IS NULL
            AND v.voucher_id = $2  
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    year_voucher_combos AS (
        SELECT
            ry.year,
            tv.voucher_id,
            tv.voucher_name
        FROM report_years ry
        CROSS JOIN target_voucher tv
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(v.price * v.quantity), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
            AND v.voucher_id = $2 
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    yvc.voucher_id,
    yvc.voucher_name,
    yvc.year::text,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_voucher_combos yvc
LEFT JOIN yearly_transactions yt ON
    yvc.year = yt.year AND
    yvc.voucher_id = yt.voucher_id
ORDER BY 
    yvc.year DESC;


-- name: GetMonthMethodVouchersSuccessById :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND v.voucher_id = $2 
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
            vm.voucher_id,
            vm.voucher_name,
            vm.payment_method
        FROM all_months am
        CROSS JOIN voucher_methods vm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND v.voucher_id = $2  
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month AND
    ac.voucher_id = mt.voucher_id AND
    ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month,
    ac.voucher_name,
    ac.payment_method;


-- name: GetYearMethodVouchersSuccessById :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND v.voucher_id = $2 
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
            vm.voucher_id,
            vm.voucher_name,
            vm.payment_method
        FROM all_years ay
        CROSS JOIN voucher_methods vm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND v.voucher_id = $2 
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year AND
    ac.voucher_id = yt.voucher_id AND
    ac.payment_method = yt.payment_method
ORDER BY 
    ac.year DESC,
    ac.voucher_name,
    ac.payment_method;





-- name: GetMonthMethodVouchersFailedById :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND v.voucher_id = $2 
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
            vm.voucher_id,
            vm.voucher_name,
            vm.payment_method
        FROM all_months am
        CROSS JOIN voucher_methods vm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND v.voucher_id = $2  
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month AND
    ac.voucher_id = mt.voucher_id AND
    ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month,
    ac.voucher_name,
    ac.payment_method;


-- name: GetYearMethodVouchersFailedById :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND v.voucher_id = $2 
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
            vm.voucher_id,
            vm.voucher_name,
            vm.payment_method
        FROM all_years ay
        CROSS JOIN voucher_methods vm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND v.voucher_id = $2 
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year AND
    ac.voucher_id = yt.voucher_id AND
    ac.payment_method = yt.payment_method
ORDER BY 
    ac.year DESC,
    ac.voucher_name,
    ac.payment_method;



-- name: GetMonthAmountVouchersSuccessByMerchant :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_vouchers AS (
        SELECT DISTINCT v.voucher_id, v.name AS voucher_name
        FROM vouchers v
        JOIN transactions t ON t.voucher_id = v.voucher_id
        WHERE 
            v.deleted_at IS NULL
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
    month_voucher_combos AS (
        SELECT
            rm.year,
            rm.month,
            tv.voucher_id,
            tv.voucher_name
        FROM report_months rm
        CROSS JOIN target_vouchers tv
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end)
            OR (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $5
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    mvc.voucher_id,
    mvc.voucher_name,
    mvc.year::text AS year,
    TO_CHAR(TO_DATE(mvc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_voucher_combos mvc
LEFT JOIN monthly_transactions mt 
    ON mvc.year = mt.year
    AND mvc.month = mt.month
    AND mvc.voucher_id = mt.voucher_id
ORDER BY 
    mvc.year DESC,
    mvc.month DESC;


-- name: GetYearAmountVouchersSuccessByMerchant :many
WITH
    target_vouchers AS (
        SELECT DISTINCT v.voucher_id, v.name AS voucher_name
        FROM vouchers v
        JOIN transactions t ON t.voucher_id = v.voucher_id
        WHERE 
            v.deleted_at IS NULL
            AND t.deleted_at IS NULL
            AND t.merchant_id = $2
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT ($1::integer - 1) AS year
    ),
    year_voucher_combos AS (
        SELECT
            ry.year,
            tv.voucher_id,
            tv.voucher_name
        FROM report_years ry
        CROSS JOIN target_vouchers tv
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND t.merchant_id = $2
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, ($1::integer - 1))
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    yvc.voucher_id,
    yvc.voucher_name,
    yvc.year::text AS year,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_voucher_combos yvc
LEFT JOIN yearly_transactions yt
    ON yvc.year = yt.year
    AND yvc.voucher_id = yt.voucher_id
ORDER BY 
    yvc.year DESC;


-- name: GetMonthAmountVouchersFailedByMerchant :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
    ),
    target_vouchers AS (
        SELECT DISTINCT v.voucher_id, v.name AS voucher_name
        FROM vouchers v
        JOIN transactions t ON t.voucher_id = v.voucher_id
        WHERE 
            v.deleted_at IS NULL
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
    month_voucher_combos AS (
        SELECT
            rm.year,
            rm.month,
            tv.voucher_id,
            tv.voucher_name
        FROM report_months rm
        CROSS JOIN target_vouchers tv
    ),
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end)
            OR (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $5
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    mvc.voucher_id,
    mvc.voucher_name,
    mvc.year::text AS year,
    TO_CHAR(TO_DATE(mvc.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_voucher_combos mvc
LEFT JOIN monthly_transactions mt 
    ON mvc.year = mt.year
    AND mvc.month = mt.month
    AND mvc.voucher_id = mt.voucher_id
ORDER BY 
    mvc.year DESC,
    mvc.month DESC;


-- name: GetYearAmountVouchersFailedByMerchant :many
WITH
    target_vouchers AS (
        SELECT DISTINCT v.voucher_id, v.name AS voucher_name
        FROM vouchers v
        JOIN transactions t ON t.voucher_id = v.voucher_id
        WHERE 
            v.deleted_at IS NULL
            AND t.deleted_at IS NULL
            AND t.merchant_id = $2
    ),
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT ($1::integer - 1) AS year
    ),
    year_voucher_combos AS (
        SELECT
            ry.year,
            tv.voucher_id,
            tv.voucher_name
        FROM report_years ry
        CROSS JOIN target_vouchers tv
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            v.voucher_id,
            v.name AS voucher_name,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND t.merchant_id = $2
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, ($1::integer - 1))
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name
    )
SELECT
    yvc.voucher_id,
    yvc.voucher_name,
    yvc.year::text AS year,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_voucher_combos yvc
LEFT JOIN yearly_transactions yt
    ON yvc.year = yt.year
    AND yvc.voucher_id = yt.voucher_id
ORDER BY 
    yvc.year DESC;

-- name: GetMonthMethodVouchersSuccessByMerchant :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND v.voucher_id = $2
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
            vm.voucher_id,
            vm.voucher_name,
            vm.payment_method
        FROM all_months am
        CROSS JOIN voucher_methods vm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND v.voucher_id = $2
            AND t.merchant_id = $3
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month AND
    ac.voucher_id = mt.voucher_id AND
    ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month,
    ac.voucher_name,
    ac.payment_method;


-- name: GetYearMethodVouchersSuccessByMerchant :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND v.voucher_id = $2
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
            vm.voucher_id,
            vm.voucher_name,
            vm.payment_method
        FROM all_years ay
        CROSS JOIN voucher_methods vm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'success'
            AND v.voucher_id = $2
            AND t.merchant_id = $3
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year AND
    ac.voucher_id = yt.voucher_id AND
    ac.payment_method = yt.payment_method
ORDER BY 
    ac.year DESC,
    ac.voucher_name,
    ac.payment_method;





-- name: GetMonthMethodVouchersFailedByMerchant :many
WITH 
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND v.voucher_id = $2
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
            vm.voucher_id,
            vm.voucher_name,
            vm.payment_method
        FROM all_months am
        CROSS JOIN voucher_methods vm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id 
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND v.voucher_id = $2
            AND t.merchant_id = $3
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) 
                                AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month AND
    ac.voucher_id = mt.voucher_id AND
    ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month,
    ac.voucher_name,
    ac.payment_method;


-- name: GetYearMethodVouchersFailedByMerchant :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    voucher_methods AS (
        SELECT DISTINCT
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND v.voucher_id = $2
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
            vm.voucher_id,
            vm.voucher_name,
            vm.payment_method
        FROM all_years ay
        CROSS JOIN voucher_methods vm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            v.voucher_id,
            v.name AS voucher_name,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        JOIN vouchers v ON t.voucher_id = v.voucher_id  
        WHERE
            t.deleted_at IS NULL
            AND v.deleted_at IS NULL
            AND t.status = 'failed'
            AND v.voucher_id = $2
            AND t.merchant_id = $3
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN 
                (SELECT start_year FROM year_range) AND 
                (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            v.voucher_id,
            v.name,
            t.payment_method
    )
SELECT
    ac.year,
    ac.voucher_id,
    ac.voucher_name,
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year AND
    ac.voucher_id = yt.voucher_id AND
    ac.payment_method = yt.payment_method
ORDER BY 
    ac.year DESC,
    ac.voucher_name,
    ac.payment_method;


-- Create Voucher
-- name: CreateVoucher :one
INSERT INTO
    vouchers (
        merchant_id,
        category_id,
        name,
        image_name
    )
VALUES ($1, $2, $3, $4)
RETURNING
    *;

-- Get Voucher by ID
-- name: GetVoucherByID :one
SELECT
    voucher_id,
    merchant_id,
    category_id,
    name,
    image_name,
    created_at,
    updated_at,
    deleted_at
FROM vouchers
WHERE
    voucher_id = $1
    AND deleted_at IS NULL;

-- Update Voucher
-- name: UpdateVoucher :one
UPDATE vouchers
SET
    name = $2,
    image_name = $3,
    category_id = $4,
    updated_at = CURRENT_TIMESTAMP
WHERE
    voucher_id = $1
    AND deleted_at IS NULL
RETURNING
    *;

-- Trash Voucher (Soft Delete)
-- name: TrashVoucher :one
UPDATE vouchers
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    voucher_id = $1
    AND deleted_at IS NULL
RETURNING
    *;

-- Restore Trashed Voucher
-- name: RestoreVoucher :one
UPDATE vouchers
SET
    deleted_at = NULL
WHERE
    voucher_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    *;

-- Delete Voucher Permanently
-- name: DeleteVoucherPermanently :exec
DELETE FROM vouchers
WHERE
    voucher_id = $1
    AND deleted_at IS NOT NULL;

-- Restore All Trashed Vouchers
-- name: RestoreAllVouchers :exec
UPDATE vouchers SET deleted_at = NULL WHERE deleted_at IS NOT NULL;

-- Delete All Trashed Vouchers Permanently
-- name: DeleteAllPermanentVouchers :exec
DELETE FROM vouchers WHERE deleted_at IS NOT NULL;