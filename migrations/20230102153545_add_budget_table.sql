-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS budget (
    id serial NOT NULL,
    user_id VARCHAR NOT NULL,
    amount JSONB NOT NULL,
    current_amount JSONB NOT NULL,
    name VARCHAR NOT NULL,
    description VARCHAR,
    type VARCHAR,
    start_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    end_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT now(),
    deleted_at TIMESTAMP(0) WITH TIME ZONE DEFAULT null,
    CONSTRAINT "PK_budget" PRIMARY KEY ("id")
);

CREATE INDEX IF NOT EXISTS "idx_budget_user_id"
    ON "budget"("user_id") WHERE "deleted_at" IS NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS budget;
DROP INDEX IF EXISTS "idx_budget_user_id";
-- +goose StatementEnd
