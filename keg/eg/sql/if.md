# IF Clause
```sql
SELECT
    COUNT(
        DISTINCT IF(
            item_id is null,
            "value_is_none",
            item_id
        ),
        IF(
            date is null, 
            "manta_repeat_value_none", 
            date
        )
    ) distinct_count,
    COUNT(*) total_count
from
    my_db.my_table
where
    (`p_date` = "20210101")
limit
    1
```