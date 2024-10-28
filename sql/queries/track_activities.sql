-- name: CreateTrackActivities :exec
INSERT INTO track_activities(user_id, track_id, is_liked)
VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateTrackActivities :exec
UPDATE track_activities
SET is_liked = $3, updated_at = NOW()
WHERE user_id = $1 AND track_id = $2
RETURNING *;

-- name: GetTrackActivity :one
SELECT *
FROM track_activities
WHERE user_id = $1 AND track_id = $2;

-- name: GetTrackActivitiesForTracks :many
SELECT * 
FROM track_activities
WHERE user_id = $1 AND track_id = ANY(sqlc.arg(track_i_ds)::text[]);