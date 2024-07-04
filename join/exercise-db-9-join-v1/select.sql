-- TODO: answer here
SELECT reports.id as id, students.fullname AS fullname, students.class AS class, students.status AS status, reports.study AS study, reports.score AS score
FROM students
INNER JOIN reports
ON reports.student_id = students.id
WHERE score < 70
ORDER BY score ASC