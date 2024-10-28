-- +goose Up
-- +goose StatementBegin
CREATE TABLE track_activities (
  id SERIAL PRIMARY KEY,
  user_id UUID NOT NULL,
  track_id TEXT NOT NULL,
  is_liked BOOLEAN NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

  CONSTRAINT fk_useractivities FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE track_activities;
-- +goose StatementEnd
