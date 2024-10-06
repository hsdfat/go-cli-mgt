CREATE TABLE "user" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar,
  "password" varchar(255),
  "full_name" varchar,
  "email" varchar,
  "phone_number" varchar,
  "created_date_unix" integer,
  "disable_date_unix" integer,
  "active" bool,
  "create_by" varchar,
  "deactivate_by" varchar,

    UNIQUE(id),
    UNIQUE(username),
    UNIQUE(email),
    UNIQUE(phone_number)
);

CREATE TABLE "network_element" (
  "id" SERIAL PRIMARY KEY,
  "name" varchar,
  "type" varchar,
  "namespace" varchar,
  "master_ip_config" varchar,
  "master_port_config" integer,
  "slave_ip_config" varchar,
  "slave_port_config" integer,
  "base_url" varchar,
  "ip_command" varchar,
  "port_command" integer
);

CREATE TABLE "user_ne" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer,
  "ne_id" integer,

  UNIQUE(user_id, ne_id)
);

CREATE TABLE "role" (
  "id" SERIAL PRIMARY KEY,
  "role_name" varchar,
  "description" varchar,
  "priority" varchar
);

CREATE TABLE "user_role" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer,
  "role_id" integer,
  
  UNIQUE(user_id, role_id)
);

CREATE TABLE "login_history" (
  "id" SERIAL PRIMARY KEY,
  "user_id" integer,
  "time_login" timestamp,
  "result" bool,
  "cause" varchar
);

CREATE TABLE "operation_history" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar,
  "command" varchar,
  "executed_time" timestamp,
  "user_ip" varchar,
  "result" bool,
  "ne_name" varchar,
  "mode" varchar
);

CREATE TABLE "server_info" (
  "id" SERIAL PRIMARY KEY,
  "ip_ssh" varchar,
  "ip_internal" varchar,
  "name" varchar,
  "pass" varchar,
  "role" varchar,
  "type" varchar,
  "user" varchar
);

CREATE TABLE "mme_subscribers" (
  "id" SERIAL PRIMARY KEY,
  "file_index" integer,
  "imsi" varchar,
  "msisdn" varchar,
  "imei" varchar,
  "mcc" varchar,
  "mnc" varchar,
  "tac" varchar,
  "eci" varchar,
  "apn" varchar,
  "msc_name" varchar,
  "ecm_state" varchar
);

CREATE TABLE "mme_files" (
  "id" SERIAL PRIMARY KEY,
  "file_name" varchar,
  "executed_time" varchar
);

ALTER TABLE "user_ne" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_ne" ADD FOREIGN KEY ("ne_id") REFERENCES "network_element" ("id");

ALTER TABLE "user_role" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_role" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "login_history" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "mme_subscribers" ADD FOREIGN KEY ("file_index") REFERENCES "mme_files" ("id");

-- ALTER TABLE "operation_history" ADD FOREIGN KEY ("username") REFERENCES "user" ("username");

INSERT INTO "role" (role_name, description, priority)
VALUES
    ('super admin', 'admin role', 'highest'),
    ('admin', 'admin role', 'highest'),
    ('editor', 'editor role', 'medium'),
    ('viewer', 'viewer role', 'lowest');

INSERT INTO "user" (username, password, full_name, email, phone_number, created_date_unix, disable_date_unix, active, create_by, deactivate_by)
VALUES ('userTest1', '$2a$14$75uobJvaxno9tJBZeg0nLeg26uajEhVLOqTMsO30QgQGvx4XUKbKS', 'Le Chi Phat', 'userTest1', '0854267485', '1728147459', '123', true, 'phatlc', 'phatlc');


INSERT INTO "user_role" (user_id, role_id)
VALUES
    (1, 1),
    (1, 2),
    (1, 3),
    (1, 4);

INSERT INTO network_element (name, type, namespace, master_ip_config, master_port_config, slave_ip_config, slave_port_config, base_url, ip_command, port_command)
VALUES ('phatlc-computer', 'dunno', 'Co Nhue, Ha Noi', 'dunno', 1234, 'dunno', 1234, 'dunno', 'dunno', 1234);

INSERT INTO user_ne (user_id, ne_id)
VALUES (1, 1);