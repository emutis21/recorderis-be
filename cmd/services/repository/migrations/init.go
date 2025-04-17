package migrations

import "gorm.io/gorm"

func InitDB(db *gorm.DB) error {
	if err := db.Exec(`DO $$ BEGIN
		CREATE TYPE user_role AS ENUM ('admin', 'user');
    EXCEPTION
        WHEN duplicate_object THEN null;
    END $$;`).Error; err != nil {
		return err
	}

	if err := db.Exec(`DO $$ BEGIN
        CREATE TYPE device_type AS ENUM ('web', 'mobile', 'tablet');
    EXCEPTION
        WHEN duplicate_object THEN null;
    END $$;`).Error; err != nil {
		return err
	}

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`).Error; err != nil {
		return err
	}

	return nil
}
