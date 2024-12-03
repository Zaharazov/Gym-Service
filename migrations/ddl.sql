CREATE TABLE "users"(
    "user_id" BIGINT NOT NULL,
    "login" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "active" BOOLEAN NOT NULL,
    "name" TEXT NOT NULL,
    "age" BIGINT NOT NULL,
    "sex" TEXT NOT NULL,
    "phone" TEXT NOT NULL,
    "pass_id" BIGINT NOT NULL,
    "gym_id" BIGINT NOT NULL
);
ALTER TABLE
    "users" ADD PRIMARY KEY("user_id");
CREATE TABLE "admins"(
    "admin_id" BIGINT NOT NULL,
    "login" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "name" TEXT NOT NULL,
    "access level" BIGINT NOT NULL,
    "phone" TEXT NOT NULL,
    "gym_id" BIGINT NOT NULL
);
ALTER TABLE
    "admins" ADD PRIMARY KEY("admin_id");
CREATE TABLE "gym_passes"(
    "gp_id" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "price" BIGINT NOT NULL,
    "description" TEXT NOT NULL
);
ALTER TABLE
    "gym_passes" ADD PRIMARY KEY("gp_id");
CREATE TABLE "gyms"(
    "gym_id" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "address" TEXT NOT NULL,
    "description" TEXT NOT NULL
);
ALTER TABLE
    "gyms" ADD PRIMARY KEY("gym_id");
CREATE TABLE "coaches"(
    "coach_id" BIGINT NOT NULL,
    "login" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "name" TEXT NOT NULL,
    "age" BIGINT NOT NULL,
    "sex" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "gym_id" BIGINT NOT NULL
);
ALTER TABLE
    "coaches" ADD PRIMARY KEY("coach_id");
CREATE TABLE "fitness_equipment"(
    "equipment_id" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL
);
ALTER TABLE
    "fitness_equipment" ADD PRIMARY KEY("equipment_id");
CREATE TABLE "events"(
    "event_id" BIGINT NOT NULL,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "coach_id" BIGINT NOT NULL
);
ALTER TABLE
    "events" ADD PRIMARY KEY("event_id");
CREATE TABLE "availability"(
    "id" BIGINT NOT NULL,
    "gym_id" BIGINT NOT NULL,
    "equipment_id" BIGINT NOT NULL,
    "amount" BIGINT NOT NULL
);
ALTER TABLE
    "availability" ADD PRIMARY KEY("id");
CREATE TABLE "current_events"(
    "id" BIGINT NOT NULL,
    "gym_id" BIGINT NOT NULL,
    "event_id" BIGINT NOT NULL,
    "time" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "data" DATE NOT NULL
);
ALTER TABLE
    "current_events" ADD PRIMARY KEY("id");
ALTER TABLE
    "events" ADD CONSTRAINT "events_coach_id_foreign" FOREIGN KEY("coach_id") REFERENCES "coaches"("coach_id");
ALTER TABLE
    "coaches" ADD CONSTRAINT "coaches_gym_id_foreign" FOREIGN KEY("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE
    "availability" ADD CONSTRAINT "availability_gym_id_foreign" FOREIGN KEY("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE
    "admins" ADD CONSTRAINT "admins_gym_id_foreign" FOREIGN KEY("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_pass_id_foreign" FOREIGN KEY("pass_id") REFERENCES "gym_passes"("gp_id");
ALTER TABLE
    "users" ADD CONSTRAINT "users_gym_id_foreign" FOREIGN KEY("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE
    "current_events" ADD CONSTRAINT "current_events_gym_id_foreign" FOREIGN KEY("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE
    "current_events" ADD CONSTRAINT "current_events_event_id_foreign" FOREIGN KEY("event_id") REFERENCES "events"("event_id");
ALTER TABLE
    "availability" ADD CONSTRAINT "availability_equipment_id_foreign" FOREIGN KEY("equipment_id") REFERENCES "fitness_equipment"("equipment_id");