CREATE TABLE decision_logs (
    decision_id TEXT NOT NULL,
    "path" TEXT NOT NULL,
    input TEXT NOT NULL,
    result TEXT NOT NULL,
    "timestamp" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);