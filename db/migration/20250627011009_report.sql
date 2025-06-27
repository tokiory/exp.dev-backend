-- +goose Up
-- +goose StatementBegin
CREATE TABLE reports (id UUID PRIMARY KEY DEFAULT gen_random_uuid());

CREATE TABLE report_persons (
  report_id UUID REFERENCES reports (id),
  name TEXT NOT NULL,
  surname TEXT NOT NULL,
  patronymic TEXT DEFAULT '',
  telegram TEXT NOT NULL,
  email TEXT NOT NULL
);

CREATE TABLE report_works (
  report_id UUID REFERENCES reports (id),
  position VARCHAR(255) NOT NULL,
  grade VARCHAR(255) NOT NULL, -- TODO: <tokiory> Maybe we should use ENUM here?
  growth_message TEXT NOT NULL,
  tasks_message TEXT NOT NULL
);

CREATE TABLE report_skills (
  report_id UUID REFERENCES reports (id),
  skills JSONB
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS report_skills;
DROP TABLE IF EXISTS report_works;
DROP TABLE IF EXISTS report_persons;
DROP TABLE IF EXISTS reports;
-- +goose StatementEnd
