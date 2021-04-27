# Cast String as Date
```sql
SELECT cast('2020-01-02' as DATE);
```

# Format Date/String to other String Format
```sql
SELECT date_format('2020-01-02', '%Y%m%d');
```

# Convert Unix Time to Date
```sql
SELECT from_unixtime(1447430881, '%Y%m%d');
```