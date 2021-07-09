package migrations

import (
	"gorm.io/gorm"
)

const up46 = `
ALTER TABLE job_proposals
ADD COLUMN uuid UUID NOT NULL;

CREATE UNIQUE INDEX idx_job_proposals_uuid ON job_proposals(uuid);
`

const down46 = `
ALTER TABLE job_proposals
DROP COLUMN uuid;
`

func init() {
	Migrations = append(Migrations, &Migration{
		ID: "0046_add_uuid_to_job_proposals",
		Migrate: func(db *gorm.DB) error {
			return db.Exec(up46).Error
		},
		Rollback: func(db *gorm.DB) error {
			return db.Exec(down46).Error
		},
	})
}
