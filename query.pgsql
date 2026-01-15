EXPLAIN SELECT id , url from links
LIMIT 5 
OFFSET 250;


EXPLAIN SELECT id , url from links
where id > 250  
LIMIT 5; 

SELECT to_char(date, 'YYYY-MM') as period, SUM(clicks)
FROM stats 
WHERE date BETWEEN '2025-07-20' and '2025-07-21'
GROUP By period;    