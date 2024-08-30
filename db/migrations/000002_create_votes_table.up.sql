CREATE TABLE votes (
    voter_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    voted_user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    pass_vote INTEGER NOT NULL CHECK (pass_vote BETWEEN 1 AND 5),
    shot_vote INTEGER NOT NULL CHECK (shot_vote BETWEEN 1 AND 5),
    marking_vote INTEGER NOT NULL CHECK (marking_vote BETWEEN 1 AND 5),
    quality_vote INTEGER NOT NULL CHECK (quality_vote BETWEEN 1 AND 5),
    velocity_vote INTEGER NOT NULL CHECK (velocity_vote BETWEEN 1 AND 5),
    UNIQUE (voter_id, voted_user_id)
);
