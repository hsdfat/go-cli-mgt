CREATE TABLE "user" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "password" varchar,
  "full_name" varchar,
  "email" varchar,
  "phone_number" varchar,
  "created_date" timestamp,
  "disable_date" timestamp,
  "active" bool,

    UNIQUE(id),
    UNIQUE(username),
    UNIQUE(email),
    UNIQUE(phone_number)
);

CREATE TABLE "network_element" (
  "id" integer PRIMARY KEY,
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
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "ne_id" integer,

  UNIQUE(user_id, ne_id)
);

CREATE TABLE "role" (
  "id" integer PRIMARY KEY,
  "role_name" varchar,
  "description" varchar,
  "priority" varchar
);

CREATE TABLE "user_role" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "role_id" integer,
  
  UNIQUE(user_id, role_id)
);

CREATE TABLE "login_history" (
  "id" integer PRIMARY KEY,
  "user_id" integer,
  "time_login" timestamp,
  "result" bool,
  "cause" varchar
);

CREATE TABLE "operation_history" (
  "id" integer PRIMARY KEY,
  "username" varchar,
  "command" varchar,
  "executed_time" timestamp,
  "user_ip" varchar,
  "result" bool,
  "ne_name" varchar
);

CREATE TABLE "server_info" (
  "id" integer PRIMARY KEY,
  "ip_ssh" varchar,
  "ip_internal" varchar,
  "name" varchar,
  "pass" varchar,
  "role" varchar,
  "type" varchar,
  "user" varchar
);

CREATE TABLE "mme_subscribers" (
  "id" integer PRIMARY KEY,
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
  "id" integer PRIMARY KEY,
  "file_name" varchar,
  "executed_time" varchar
);

ALTER TABLE "user_ne" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_ne" ADD FOREIGN KEY ("ne_id") REFERENCES "network_element" ("id");

ALTER TABLE "user_role" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "user_role" ADD FOREIGN KEY ("role_id") REFERENCES "role" ("id");

ALTER TABLE "login_history" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "mme_subscribers" ADD FOREIGN KEY ("file_index") REFERENCES "mme_files" ("id");

ALTER TABLE "operation_history" ADD FOREIGN KEY ("username") REFERENCES "user" ("username");