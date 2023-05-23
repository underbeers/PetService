ALTER TABLE pet_card RENAME COLUMN photo to origin_photo;
ALTER TABLE pet_card ALTER COLUMN origin_photo TYPE text[];
ALTER TABLE pet_card ALTER COLUMN thumbnail_photo TYPE text[];