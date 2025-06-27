-- name: GetReport :one
SELECT reports.*, report_skills.*
FROM reports
JOIN report_skills ON reports.id = report_skills.report_id
WHERE reports.id = $1
LIMIT 1;

-- name: GetReports :many
SELECT reports.*, report_skills.*
FROM reports
JOIN report_skills ON reports.id = skills.report_id
WHERE reports.id = $1;

-- name: CreateReport :one
INSERT INTO reports (id)
VALUES (gen_random_uuid())
RETURNING id;

-- name: CreateReportPerson :exec
INSERT INTO report_persons (
  report_id,
  name,
  surname,
  patronymic,
  telegram,
  email
) VALUES ($1, $2, $3, $4, $5, $6);

-- name: CreateReportWork :exec
INSERT INTO report_works (
  report_id,
  position,
  grade,
  growth_message,
  tasks_message
) VALUES (
  $1,
  $2,
  $3,
  $4,
  $5
);

-- name: CreateReportSkills :exec
INSERT INTO report_skills (
  report_id, skills
) VALUES (
  $1,
  $2
);
