-- Get Transactions with Pagination and Total Count
-- name: GetTransactions :many
SELECT *, COUNT(*) OVER () AS total_count
FROM transactions
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR payment_method ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3;

-- Get Active Transactions with Pagination and Total Count
-- name: GetTransactionsActive :many
SELECT *, COUNT(*) OVER () AS total_count
FROM transactions
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR payment_method ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3;

-- Get Trashed Transactions with Pagination and Total Count
-- name: GetTransactionsTrashed :many
SELECT *, COUNT(*) OVER () AS total_count
FROM transactions
WHERE
    deleted_at IS NOT NULL
    AND (
        $1::TEXT IS NULL
        OR payment_method ILIKE '%' || $1 || '%'
    )
ORDER BY created_at DESC
LIMIT $2
OFFSET
    $3;


-- GetTransactionByMerchant: Retrieves merchant-specific transactions with pagination
-- Purpose: List transactions filtered by merchant ID
-- Parameters:
--   $1: search_term - Optional text to filter transactions
--   $2: merchant_id - Optional merchant ID to filter by (NULL for all merchants)
--   $3: limit - Pagination limit
--   $4: offset - Pagination offset
-- Returns:
--   Transaction records with total_count
-- Business Logic:
--   - Combines merchant filtering with search functionality
--   - Maintains same sorting and pagination as other transaction queries
--   - Useful for merchant-specific transaction reporting
--   - NULL merchant_id parameter returns all merchants' transactions
-- name: GetTransactionByMerchant :many
SELECT *, COUNT(*) OVER () AS total_count
FROM transactions
WHERE
    deleted_at IS NULL
    AND (
        $1::TEXT IS NULL
        OR payment_method ILIKE '%' || $1 || '%'
        OR payment_status ILIKE '%' || $1 || '%'
    )
    AND (
        $2::INT IS NULL
        OR merchant_id = $2
    )
ORDER BY created_at DESC
LIMIT $3
OFFSET
    $4;


-- GetMonthlyAmountTransactionSuccess: Retrieves monthly success transaction metrics
-- Purpose: Generate monthly reports of successful transactions for analysis
-- Parameters:
--   $1: Start date of first comparison period (timestamp)
--   $2: End date of first comparison period (timestamp)
--   $3: Start date of second comparison period (timestamp)
--   $4: End date of second comparison period (timestamp)
-- Returns:
--   year: Year as text
--   month: 3-letter month abbreviation (e.g. 'Jan')
--   total_success: Count of successful transactions
--   total_amount: Sum of successful transaction amounts
-- Business Logic:
--   - Only includes successful (payment_status = 'success') transactions
--   - Excludes deleted transactions
--   - Compares two customizable time periods
--   - Includes gap-filling for months with no transactions
--   - Returns 0 values for months with no successful transactions
--   - Orders by most recent year/month first
-- name: GetMonthlyAmountTransactionSuccess :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
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
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'success'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at)
    )
SELECT
    rm.year::text,
    TO_CHAR(TO_DATE(rm.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_success, 0) AS total_success,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM report_months rm
LEFT JOIN monthly_transactions mt ON
    rm.year = mt.year AND
    rm.month = mt.month
ORDER BY 
    rm.year DESC,
    rm.month DESC;

-- GetYearlyAmountTransactionSuccess: Retrieves yearly success transaction metrics
-- Purpose: Generate annual reports of successful transactions
-- Parameters:
--   $1: Reference year for comparison (current year as integer)
-- Returns:
--   year: Year as text
--   total_success: Count of successful transactions
--   total_amount: Sum of successful transaction amounts
-- Business Logic:
--   - Compares current year with previous year automatically
--   - Only includes successful (payment_status = 'success') transactions
--   - Excludes deleted transactions
--   - Includes gap-filling for years with no transactions
--   - Returns 0 values for years with no successful transactions
--   - Orders by most recent year first
-- name: GetYearlyAmountTransactionSuccess :many
WITH
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at)
    )
SELECT
    ry.year::text,
    COALESCE(yt.total_success, 0) AS total_success,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM report_years ry
LEFT JOIN yearly_transactions yt ON
    ry.year = yt.year
ORDER BY 
    ry.year DESC;

-- GetMonthlyAmountTransactionFailed: Retrieves monthly failed transaction metrics
-- Purpose: Generate monthly reports of failed transactions for analysis
-- Parameters:
--   $1: Start date of first comparison period (timestamp)
--   $2: End date of first comparison period (timestamp)
--   $3: Start date of second comparison period (timestamp)
--   $4: End date of second comparison period (timestamp)
-- Returns:
--   year: Year as text
--   month: 3-letter month abbreviation (e.g. 'Jan')
--   total_failed: Count of failed transactions
--   total_amount: Sum of failed transaction amounts
-- Business Logic:
--   - Only includes failed (payment_status = 'failed') transactions
--   - Excludes deleted transactions
--   - Compares two customizable time periods
--   - Includes gap-filling for months with no failed transactions
--   - Returns 0 values for months with no failed transactions
--   - Orders by most recent year/month first
-- name: GetMonthlyAmountTransactionFailed :many
WITH 
    date_ranges AS (
        SELECT
            $1::timestamp AS range1_start,
            $2::timestamp AS range1_end,
            $3::timestamp AS range2_start,
            $4::timestamp AS range2_end
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
    monthly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            EXTRACT(MONTH FROM t.created_at)::integer AS month,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        JOIN date_ranges dr ON (
            (t.created_at BETWEEN dr.range1_start AND dr.range1_end) OR
            (t.created_at BETWEEN dr.range2_start AND dr.range2_end)
        )
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'failed'
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            EXTRACT(MONTH FROM t.created_at)
    )
SELECT
    rm.year::text,
    TO_CHAR(TO_DATE(rm.month::text, 'MM'), 'Mon') AS month,
    COALESCE(mt.total_failed, 0) AS total_failed,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM report_months rm
LEFT JOIN monthly_transactions mt ON
    rm.year = mt.year AND
    rm.month = mt.month
ORDER BY 
    rm.year DESC,
    rm.month DESC;

-- GetYearlyAmountTransactionFailed: Retrieves yearly failed transaction metrics
-- Purpose: Generate annual reports of failed transactions
-- Parameters:
--   $1: Reference year for comparison (current year as integer)
-- Returns:
--   year: Year as text
--   total_failed: Count of failed transactions
--   total_amount: Sum of failed transaction amounts
-- Business Logic:
--   - Compares current year with previous year automatically
--   - Only includes failed (payment_status = 'failed') transactions
--   - Excludes deleted transactions
--   - Includes gap-filling for years with no transactions
--   - Returns 0 values for years with no failed transactions
--   - Orders by most recent year first
-- name: GetYearlyAmountTransactionFailed :many
WITH
    report_years AS (
        SELECT $1::integer AS year
        UNION
        SELECT $1::integer - 1 AS year
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) IN ($1::integer, $1::integer - 1)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at)
    )
SELECT
    ry.year::text,
    COALESCE(yt.total_failed, 0) AS total_failed,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM report_years ry
LEFT JOIN yearly_transactions yt ON
    ry.year = yt.year
ORDER BY 
    ry.year DESC;

-- GetMonthlyTransactionMethodsSuccess: Analyzes successful payment method usage by month
-- Purpose: Track monthly trends in successful payment method preferences
-- Parameters:
--   $1: Reference date (timestamp) - determines the 12-month analysis period
-- Returns:
--   month: 3-letter month abbreviation (e.g. 'Jan')
--   payment_method: The payment method used
--   total_transactions: Count of successful transactions
--   total_amount: Total amount successfully processed by this method
-- Business Logic:
--   - Analyzes a rolling 12-month period from reference date
--   - Only includes successful (payment_status = 'success') transactions
--   - Excludes deleted transactions
--   - Groups by month and payment method
--   - Returns formatted month names for reporting
--   - Orders chronologically by month then by payment method
-- name: GetMonthlyTransactionMethodsSuccess :many
WITH
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    payment_methods AS (
        SELECT DISTINCT
            payment_method
        FROM transactions
        WHERE deleted_at IS NULL
    ),
    all_months AS (
        SELECT generate_series(
            (SELECT start_date FROM date_range),
            (SELECT end_date FROM date_range),
            interval '1 month'
        )::date AS activity_month
    ),
    all_combinations AS (
        SELECT 
            am.activity_month,
            pm.payment_method
        FROM all_months am
        CROSS JOIN payment_methods pm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'success'
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            t.payment_method
    )
SELECT 
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month
    AND ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month, 
    ac.payment_method;


-- GetYearlyTransactionMethodsSuccess: Analyzes successful payment method usage by year
-- Purpose: Track annual trends in successful payment method preferences
-- Parameters:
--   $1: Reference date (timestamp) - determines the 5-year analysis window
-- Returns:
--   year: 4-digit year as text
--   payment_method: The payment method used
--   total_transactions: Count of successful transactions
--   total_amount: Total amount successfully processed by this method
-- Business Logic:
--   - Covers current year plus previous 4 years (5-year total window)
--   - Only includes successful (payment_status = 'success') transactions
--   - Excludes deleted transactions
--   - Groups by year and payment method
--   - Orders chronologically by year then by payment method
--   - Useful for identifying long-term successful payment trends
-- name: GetYearlyTransactionMethodsSuccess :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    payment_methods AS (
        SELECT DISTINCT
            payment_method
        FROM transactions
        WHERE deleted_at IS NULL
    ),
    all_years AS (
        SELECT generate_series(
            (SELECT start_year FROM year_range),
            (SELECT end_year FROM year_range)
        )::int AS year
    ),
    all_combinations AS (
        SELECT 
            ay.year::text AS year,  
            pm.payment_method
        FROM all_years ay
        CROSS JOIN payment_methods pm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'success'
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN (SELECT start_year FROM year_range) AND (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            t.payment_method
    )
SELECT 
    ac.year,  
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year
    AND ac.payment_method = yt.payment_method
ORDER BY 
    ac.year,
    ac.payment_method;





-- GetMonthlyTransactionMethodsFailed: Analyzes failed payment method usage by month
-- Purpose: Track monthly trends in failed payment method attempts
-- Parameters:
--   $1: Reference date (timestamp) - determines the 12-month analysis period
-- Returns:
--   month: 3-letter month abbreviation (e.g. 'Jan')
--   payment_method: The payment method attempted
--   total_transactions: Count of failed transactions
--   total_amount: Total amount attempted (not successfully processed)
-- Business Logic:
--   - Analyzes a rolling 12-month period from reference date
--   - Only includes failed (payment_status = 'failed') transactions
--   - Excludes deleted transactions
--   - Groups by month and payment method
--   - Returns formatted month names for reporting
--   - Orders chronologically by month then by payment method
-- name: GetMonthlyTransactionMethodsFailed :many
WITH
    date_range AS (
        SELECT 
            date_trunc('month', $1::timestamp) AS start_date, 
            date_trunc('month', $1::timestamp) + interval '1 year' - interval '1 day' AS end_date
    ),
    payment_methods AS (
        SELECT DISTINCT
            payment_method
        FROM transactions
        WHERE deleted_at IS NULL
    ),
    all_months AS (
        SELECT generate_series(
            (SELECT start_date FROM date_range),
            (SELECT end_date FROM date_range),
            interval '1 month'
        )::date AS activity_month
    ),
    all_combinations AS (
        SELECT 
            am.activity_month,
            pm.payment_method
        FROM all_months am
        CROSS JOIN payment_methods pm
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS activity_month,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'failed'
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) AND (SELECT end_date FROM date_range)
        GROUP BY
            date_trunc('month', t.created_at),
            t.payment_method
    )
SELECT 
    TO_CHAR(ac.activity_month, 'Mon') AS month,
    ac.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN monthly_transactions mt ON 
    ac.activity_month = mt.activity_month
    AND ac.payment_method = mt.payment_method
ORDER BY 
    ac.activity_month, 
    ac.payment_method;


-- GetYearlyTransactionMethodsFailed: Analyzes failed payment method usage by year
-- Purpose: Track annual trends in failed payment method attempts
-- Parameters:
--   $1: Reference date (timestamp) - determines the 5-year analysis window
-- Returns:
--   year: 4-digit year as text
--   payment_method: The payment method attempted
--   total_transactions: Count of failed transactions
--   total_amount: Total amount attempted (not successfully processed)
-- Business Logic:
--   - Covers current year plus previous 4 years (5-year total window)
--   - Only includes failed (payment_status = 'failed') transactions
--   - Excludes deleted transactions
--   - Groups by year and payment method
--   - Orders chronologically by year then by payment method
--   - Useful for identifying long-term failed payment trends
-- name: GetYearlyTransactionMethodsFailed :many
WITH
    year_range AS (
        SELECT 
            EXTRACT(YEAR FROM $1::timestamp)::int - 4 AS start_year,
            EXTRACT(YEAR FROM $1::timestamp)::int AS end_year
    ),
    payment_methods AS (
        SELECT DISTINCT
            payment_method
        FROM transactions
        WHERE deleted_at IS NULL
    ),
    all_years AS (
        SELECT generate_series(
            (SELECT start_year FROM year_range),
            (SELECT end_year FROM year_range)
        )::int AS year
    ),
    all_combinations AS (
        SELECT 
            ay.year::text AS year,  
            pm.payment_method
        FROM all_years ay
        CROSS JOIN payment_methods pm
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::text AS year,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            COALESCE(SUM(t.amount), 0)::NUMERIC AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'failed'
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN (SELECT start_year FROM year_range) AND (SELECT end_year FROM year_range)
        GROUP BY
            EXTRACT(YEAR FROM t.created_at),
            t.payment_method
    )
SELECT 
    ac.year, 
    ac.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM all_combinations ac
LEFT JOIN yearly_transactions yt ON 
    ac.year = yt.year
    AND ac.payment_method = yt.payment_method
ORDER BY 
    ac.year,  
    ac.payment_method;




-- GetMonthlyAmountTransactionSuccessByMerchant: Retrieves monthly success transaction metrics by merchant_id
-- Purpose: Generate monthly reports of successful transactions for analysis
-- Parameters:
--   $1: Start date of first comparison period (timestamp)
--   $2: End date of first comparison period (timestamp)
--   $3: Start date of second comparison period (timestamp)
--   $4: End date of second comparison period (timestamp)
--   $5: Merchant ID
-- Returns:
--   year: Year as text
--   month: 3-letter month abbreviation (e.g. 'Jan')
--   total_success: Count of successful transactions
--   total_amount: Sum of successful transaction amounts
-- Business Logic:
--   - Only includes successful (payment_status = 'success') transactions
--   - Excludes deleted transactions
--   - Compares two customizable time periods
--   - Includes gap-filling for months with no transactions
--   - Returns 0 values for months with no successful transactions
--   - Orders by most recent year/month first
-- name: GetMonthlyAmountTransactionSuccessByMerchant :many
WITH
    month_series AS (
        SELECT generate_series(
            date_trunc('month', LEAST($1::timestamp, $3::timestamp)),
            date_trunc('month', GREATEST($2::timestamp, $4::timestamp)),
            INTERVAL '1 month'
        ) AS month_start
    ),
    monthly_data AS (
        SELECT
            date_trunc('month', t.created_at) AS month_start,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'success'
            AND t.merchant_id = $5
            AND t.created_at BETWEEN LEAST($1::timestamp, $3::timestamp) AND GREATEST($2::timestamp, $4::timestamp)
        GROUP BY date_trunc('month', t.created_at)
    )
SELECT
    EXTRACT(YEAR FROM m.month_start)::text AS year,
    TO_CHAR(m.month_start, 'Mon') AS month,
    COALESCE(md.total_success, 0) AS total_success,
    COALESCE(md.total_amount, 0) AS total_amount
FROM month_series m
LEFT JOIN monthly_data md ON m.month_start = md.month_start
ORDER BY m.month_start DESC;

-- GetYearlyAmountTransactionSuccessByMerchant: Retrieves yearly success transaction metrics
-- Purpose: Generate annual reports of successful transactions by merchant_id
-- Parameters:
--   $1: Reference year for comparison (current year as integer)
--   $2: Merchant ID
-- Returns:
--   year: Year as text
--   total_success: Count of successful transactions
--   total_amount: Sum of successful transaction amounts
-- Business Logic:
--   - Compares current year with previous year automatically
--   - Only includes successful (payment_status = 'success') transactions
--   - Excludes deleted transactions
--   - Includes gap-filling for years with no transactions
--   - Returns 0 values for years with no successful transactions
--   - Orders by most recent year first
-- name: GetYearlyAmountTransactionSuccessByMerchant :many
WITH
    year_series AS (
        SELECT unnest(ARRAY[
            $1::integer,
            $1::integer - 1
        ]) AS year
    ),
    yearly_data AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            COUNT(*) AS total_success,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'success'
            AND t.merchant_id = $2
            AND EXTRACT(YEAR FROM t.created_at) IN ($1, $1 - 1)
        GROUP BY EXTRACT(YEAR FROM t.created_at)
    )
SELECT
    ys.year::text,
    COALESCE(yd.total_success, 0) AS total_success,
    COALESCE(yd.total_amount, 0) AS total_amount
FROM year_series ys
LEFT JOIN yearly_data yd ON ys.year = yd.year
ORDER BY ys.year DESC;

-- GetMonthlyAmountTransactionFailedByMerchant: Retrieves monthly failed transaction metrics
-- Purpose: Generate monthly reports of failed transactions for analysis by merchant_id
-- Parameters:
--   $1: Start date of first comparison period (timestamp)
--   $2: End date of first comparison period (timestamp)
--   $3: Start date of second comparison period (timestamp)
--   $4: End date of second comparison period (timestamp)
--   $5: Merchant ID
-- Returns:
--   year: Year as text
--   month: 3-letter month abbreviation (e.g. 'Jan')
--   total_failed: Count of failed transactions
--   total_amount: Sum of failed transaction amounts
-- Business Logic:
--   - Only includes failed (payment_status = 'failed') transactions
--   - Excludes deleted transactions
--   - Compares two customizable time periods
--   - Includes gap-filling for months with no failed transactions
--   - Returns 0 values for months with no failed transactions
--   - Orders by most recent year/month first
-- name: GetMonthlyAmountTransactionFailedByMerchant :many
WITH
    month_series AS (
        SELECT generate_series(
            date_trunc('month', LEAST($1::timestamp, $3::timestamp)),
            date_trunc('month', GREATEST($2::timestamp, $4::timestamp)),
            INTERVAL '1 month'
        ) AS month_start
    ),
    monthly_data AS (
        SELECT
            date_trunc('month', t.created_at) AS month_start,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'failed'
            AND t.merchant_id = $5
            AND t.created_at BETWEEN LEAST($1::timestamp, $3::timestamp) AND GREATEST($2::timestamp, $4::timestamp)
        GROUP BY date_trunc('month', t.created_at)
    )
SELECT
    EXTRACT(YEAR FROM m.month_start)::text AS year,
    TO_CHAR(m.month_start, 'Mon') AS month,
    COALESCE(md.total_failed, 0) AS total_failed,
    COALESCE(md.total_amount, 0) AS total_amount
FROM month_series m
LEFT JOIN monthly_data md ON m.month_start = md.month_start
ORDER BY m.month_start DESC;



-- GetYearlyAmountTransactionFailedByMerchant: Retrieves yearly failed transaction metrics
-- Purpose: Generate annual reports of failed transactions by merchant_id
-- Parameters:
--   $1: Reference year for comparison (current year as integer)
--   $2: Merchant ID
-- Returns:
--   year: Year as text
--   total_failed: Count of failed transactions
--   total_amount: Sum of failed transaction amounts
-- Business Logic:
--   - Compares current year with previous year automatically
--   - Only includes failed (payment_status = 'failed') transactions
--   - Excludes deleted transactions
--   - Includes gap-filling for years with no transactions
--   - Returns 0 values for years with no failed transactions
--   - Orders by most recent year first
-- name: GetYearlyAmountTransactionFailedByMerchant :many
WITH
    year_series AS (
        SELECT unnest(ARRAY[
            $1::integer,
            $1::integer - 1
        ]) AS year
    ),
    yearly_data AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            COUNT(*) AS total_failed,
            COALESCE(SUM(t.amount), 0)::integer AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'failed'
            AND t.merchant_id = $2
            AND EXTRACT(YEAR FROM t.created_at) IN ($1, $1 - 1)
        GROUP BY EXTRACT(YEAR FROM t.created_at)
    )
SELECT
    ys.year::text,
    COALESCE(yd.total_failed, 0) AS total_failed,
    COALESCE(yd.total_amount, 0) AS total_amount
FROM year_series ys
LEFT JOIN yearly_data yd ON ys.year = yd.year
ORDER BY ys.year DESC;



-- Purpose: Analyze the monthly payment method usage for a specific merchant with a focus on successful transactions.
-- Parameters:
--   $1: Reference date (timestamp) - determines the 12-month analysis period.
--   $2: Merchant ID - filters transactions for a specific merchant.
-- Returns:
--   month: 3-letter month abbreviation (e.g. 'Jan').
--   payment_method: The payment method used.
--   total_transactions: Count of successful transactions.
--   total_amount: Total amount processed by this payment method.
-- Business Logic:
--   - Analyzes a rolling 12-month period starting from the reference date ($1).
--   - Only includes transactions with a payment status of 'success' (payment_status = 'success').
--   - Excludes deleted transactions (where deleted_at IS NULL).
--   - Groups the results by month and payment method.
--   - For each month, reports the number of successful transactions and total amount processed.
--   - Orders the result chronologically by month, followed by payment method.
--   - Returns a report useful for tracking successful payment trends over the last 12 months.
-- name: GetMonthlyTransactionMethodsSuccessByMerchant :many
WITH
    date_range AS (
        SELECT
            date_trunc('month', $1::timestamp) AS start_date,
            date_trunc('month', $1::timestamp) + INTERVAL '1 year' - INTERVAL '1 day' AS end_date
    ),
    month_series AS (
        SELECT generate_series(
            (SELECT start_date FROM date_range),
            (SELECT end_date FROM date_range),
            INTERVAL '1 month'
        ) AS month_start
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS month_start,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            SUM(t.amount)::NUMERIC AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'success'
            AND t.merchant_id = $2
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) AND (SELECT end_date FROM date_range)
        GROUP BY
            month_start,
            t.payment_method
    ),
    payment_methods AS (
        SELECT DISTINCT payment_method
        FROM transactions
        WHERE deleted_at IS NULL
    )
SELECT
    TO_CHAR(m.month_start, 'Mon') AS month,
    pm.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_series m
CROSS JOIN payment_methods pm
LEFT JOIN monthly_transactions mt
    ON m.month_start = mt.month_start
    AND pm.payment_method = mt.payment_method
ORDER BY m.month_start, pm.payment_method;


-- Purpose: Analyze the yearly payment method usage for a specific merchant with a focus on successful transactions.
-- Parameters:
--   $1: Reference date (timestamp) - determines the 5-year analysis window.
--   $2: Merchant ID - filters transactions for a specific merchant.
-- Returns:
--   year: 4-digit year as text (e.g., '2023').
--   payment_method: The payment method used.
--   total_transactions: Count of successful transactions.
--   total_amount: Total amount processed by this payment method.
-- Business Logic:
--   - Analyzes a 5-year period, covering the current year and the previous 4 years.
--   - Only includes transactions with a payment status of 'success' (payment_status = 'success').
--   - Excludes deleted transactions (where deleted_at IS NULL).
--   - Groups the results by year and payment method.
--   - For each year, reports the number of successful transactions and total amount processed.
--   - Orders the result chronologically by year, followed by payment method.
--   - Useful for identifying long-term successful payment trends for the merchant.
-- name: GetYearlyTransactionMethodsSuccessByMerchant :many
WITH
    year_series AS (
        SELECT generate_series(
            EXTRACT(YEAR FROM $1::timestamp)::integer - 4,
            EXTRACT(YEAR FROM $1::timestamp)::integer,
            1
        ) AS year
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            SUM(t.amount)::NUMERIC AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'success'
            AND t.merchant_id = $2
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
        GROUP BY
            year,
            t.payment_method
    ),
    payment_methods AS (
        SELECT DISTINCT payment_method
        FROM transactions
        WHERE deleted_at IS NULL
    )
SELECT
    ys.year::text AS year,
    pm.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_series ys
CROSS JOIN payment_methods pm
LEFT JOIN yearly_transactions yt
    ON ys.year = yt.year
    AND pm.payment_method = yt.payment_method
ORDER BY ys.year, pm.payment_method;



-- Purpose: Analyze the monthly payment method usage for a specific merchant with a focus on failed transactions.
-- Parameters:
--   $1: Reference date (timestamp) - determines the 12-month analysis period.
--   $2: Merchant ID - filters transactions for a specific merchant.
-- Returns:
--   month: 3-letter month abbreviation (e.g. 'Jan').
--   payment_method: The payment method used.
--   total_transactions: Count of failed transactions.
--   total_amount: Total amount attempted for this payment method (including failed transactions).
-- Business Logic:
--   - Analyzes a rolling 12-month period starting from the reference date ($1).
--   - Only includes transactions with a payment status of 'failed' (payment_status = 'failed').
--   - Excludes deleted transactions (where deleted_at IS NULL).
--   - Groups the results by month and payment method.
--   - For each month, reports the number of failed transactions and the total amount attempted.
--   - Orders the result chronologically by month, followed by payment method.
--   - Returns a report useful for tracking failed payment trends over the last 12 months.
-- name: GetMonthlyTransactionMethodsFailedByMerchant :many
WITH
    date_range AS (
        SELECT
            date_trunc('month', $1::timestamp) AS start_date,
            date_trunc('month', $1::timestamp) + INTERVAL '1 year' - INTERVAL '1 day' AS end_date
    ),
    month_series AS (
        SELECT generate_series(
            (SELECT start_date FROM date_range),
            (SELECT end_date FROM date_range),
            INTERVAL '1 month'
        ) AS month_start
    ),
    monthly_transactions AS (
        SELECT
            date_trunc('month', t.created_at) AS month_start,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            SUM(t.amount)::NUMERIC AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'failed'
            AND t.merchant_id = $2
            AND t.created_at BETWEEN (SELECT start_date FROM date_range) AND (SELECT end_date FROM date_range)
        GROUP BY
            month_start,
            t.payment_method
    ),
    payment_methods AS (
        SELECT DISTINCT payment_method
        FROM transactions
        WHERE deleted_at IS NULL
    )
SELECT
    TO_CHAR(m.month_start, 'Mon') AS month,
    pm.payment_method,
    COALESCE(mt.total_transactions, 0) AS total_transactions,
    COALESCE(mt.total_amount, 0) AS total_amount
FROM month_series m
CROSS JOIN payment_methods pm
LEFT JOIN monthly_transactions mt
    ON m.month_start = mt.month_start
    AND pm.payment_method = mt.payment_method
ORDER BY m.month_start, pm.payment_method;



-- Purpose: Analyze the yearly payment method usage for a specific merchant with a focus on failed transactions.
-- Parameters:
--   $1: Reference date (timestamp) - determines the 5-year analysis window.
--   $2: Merchant ID - filters transactions for a specific merchant.
-- Returns:
--   year: 4-digit year as text (e.g., '2023').
--   payment_method: The payment method used.
--   total_transactions: Count of failed transactions.
--   total_amount: Total amount attempted for this payment method (including failed transactions).
-- Business Logic:
--   - Analyzes a 5-year period, covering the current year and the previous 4 years.
--   - Only includes transactions with a payment status of 'failed' (payment_status = 'failed').
--   - Excludes deleted transactions (where deleted_at IS NULL).
--   - Groups the results by year and payment method.
--   - For each year, reports the number of failed transactions and the total amount attempted.
--   - Orders the result chronologically by year, followed by payment method.
--   - Useful for identifying long-term failed payment trends for the merchant.
-- name: GetYearlyTransactionMethodsFailedByMerchant :many
WITH
    year_series AS (
        SELECT generate_series(
            EXTRACT(YEAR FROM $1::timestamp)::integer - 4,
            EXTRACT(YEAR FROM $1::timestamp)::integer,
            1
        ) AS year
    ),
    yearly_transactions AS (
        SELECT
            EXTRACT(YEAR FROM t.created_at)::integer AS year,
            t.payment_method,
            COUNT(t.transaction_id) AS total_transactions,
            SUM(t.amount)::NUMERIC AS total_amount
        FROM transactions t
        WHERE
            t.deleted_at IS NULL
            AND t.payment_status = 'failed'
            AND t.merchant_id = $2
            AND EXTRACT(YEAR FROM t.created_at) BETWEEN (EXTRACT(YEAR FROM $1::timestamp) - 4) AND EXTRACT(YEAR FROM $1::timestamp)
        GROUP BY
            year,
            t.payment_method
    ),
    payment_methods AS (
        SELECT DISTINCT payment_method
        FROM transactions
        WHERE deleted_at IS NULL
    )
SELECT
    ys.year::text AS year,
    pm.payment_method,
    COALESCE(yt.total_transactions, 0) AS total_transactions,
    COALESCE(yt.total_amount, 0) AS total_amount
FROM year_series ys
CROSS JOIN payment_methods pm
LEFT JOIN yearly_transactions yt
    ON ys.year = yt.year
    AND pm.payment_method = yt.payment_method
ORDER BY ys.year, pm.payment_method;


-- Create Transaction
-- name: CreateTransaction :one
INSERT INTO
    transactions (
        user_id,
        merchant_id,
        voucher_id,
        nominal_id,
        bank_id,
        payment_method,
        status,
        amount
    )
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING
    *;

-- Get Transaction by ID
-- name: GetTransactionByID :one
SELECT *
FROM transactions
WHERE
    transaction_id = $1
    AND deleted_at IS NULL;

-- Update Transaction
-- name: UpdateTransaction :one
UPDATE transactions
SET
    user_id = $2,
    merchant_id = $3,
    voucher_id = $4,
    nominal_id = $5,
    bank_id = $6,
    payment_method = $7,
    status = $8,
    amount = $9,
    updated_at = CURRENT_TIMESTAMP
WHERE
    transaction_id = $1
    AND deleted_at IS NULL
RETURNING
    *;

-- Update Transaction Status
-- name: UpdateTransactionStatus :exec
UPDATE transactions
SET
    status = $2,
    updated_at = CURRENT_TIMESTAMP
WHERE
    transaction_id = $1
    AND deleted_at IS NULL;

-- Trash Transaction (Soft Delete)
-- name: TrashTransaction :one
UPDATE transactions
SET
    deleted_at = CURRENT_TIMESTAMP
WHERE
    transaction_id = $1
    AND deleted_at IS NULL
RETURNING
    *;

-- Restore Trashed Transaction
-- name: RestoreTransaction :one
UPDATE transactions
SET
    deleted_at = NULL
WHERE
    transaction_id = $1
    AND deleted_at IS NOT NULL
RETURNING
    *;

-- Delete Transaction Permanently
-- name: DeleteTransactionPermanently :exec
DELETE FROM transactions
WHERE
    transaction_id = $1
    AND deleted_at IS NOT NULL;

-- Restore All Trashed Transactions
-- name: RestoreAllTransactions :exec
UPDATE transactions
SET
    deleted_at = NULL
WHERE
    deleted_at IS NOT NULL;

-- Delete All Trashed Transactions Permanently
-- name: DeleteAllPermanentTransactions :exec
DELETE FROM transactions WHERE deleted_at IS NOT NULL;