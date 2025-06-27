-- name: GetReport :one
SELECT
  reports.id,
  report_persons.name,
  report_persons.surname,
  report_persons.patronymic,
  report_persons.email,
  report_persons.telegram,
  report_works.position,
  report_works.grade,
  report_works.growth_message,
  report_works.tasks_message,
  report_skills.skills
FROM reports
LEFT JOIN report_skills ON reports.id = report_skills.report_id
LEFT JOIN report_works ON reports.id = report_works.report_id
LEFT JOIN report_persons ON reports.id = report_persons.report_id
WHERE reports.id = $1
LIMIT 1;

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
