-- 001_initial.up.sql
CREATE FUNCTION update_modified() RETURNS TRIGGER AS $$ BEGIN NEW.modified = now(); RETURN NEW; END; $$ language plpgsql;

CREATE TABLE users (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  discord_id text NOT NULL,
  discord_username text NOT NULL,
  discord_tag text NOT NULL,
  created TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TRIGGER update_users_modified BEFORE UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_modified();

-- 001_initial.down.sql
-- DROP TABLE IF EXISTS users;
-- DROP FUNCTION IF EXISTS update_modified;
-- DROP TABLE IF EXISTS schema_migrations;
