-- Create tables.
DROP TABLE IF EXISTS "disk";
DROP TABLE IF EXISTS "server";
CREATE TABLE "server"
(
    "server_id"   SERIAL PRIMARY KEY,
    "name" VARCHAR(50) NOT NULL UNIQUE,
    "cpu_count" int
);
CREATE TABLE "disk"
(
    "disk_id"   SERIAL PRIMARY KEY,
    "space" numeric,
    "server_id" integer REFERENCES server(server_id)
);

-- Insert demo data.
INSERT INTO "server" (name, cpu_count) VALUES ('server-1', 4);
INSERT INTO "server" (name, cpu_count) VALUES ('server-2', 2);
INSERT INTO "server" (name, cpu_count) VALUES ('server-3', 8);
INSERT INTO "disk" (space, server_id) VALUES (100000, (select server_id from server where name='server-1'));
INSERT INTO "disk" (space, server_id) VALUES (300000, (select server_id from server where name='server-2'));
INSERT INTO "disk" (space, server_id) VALUES (230000, (select server_id from server where name='server-3'));
INSERT INTO "disk" (space) VALUES (120000);
INSERT INTO "disk" (space) VALUES (560000);
INSERT INTO "disk" (space) VALUES (220000);
INSERT INTO "disk" (space) VALUES (10000);
