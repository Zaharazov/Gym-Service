CREATE TABLE "users" (
    "user_id" SERIAL PRIMARY KEY, -- Автоинкрементируемый ID
    "login" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "active" BOOLEAN,
    "name" TEXT NOT NULL,
    "age" BIGINT,
    "sex" TEXT,
    "phone" TEXT,
    "pass_id" BIGINT,
    "gym_id" BIGINT
);

CREATE TABLE "admins" (
    "admin_id" SERIAL PRIMARY KEY, -- Автоинкрементируемый ID
    "login" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "name" TEXT NOT NULL,
    "access_level" BIGINT,
    "phone" TEXT,
    "gym_id" BIGINT
);

CREATE TABLE "coaches" (
    "coach_id" SERIAL PRIMARY KEY, -- Автоинкрементируемый ID
    "login" VARCHAR(255) NOT NULL,
    "password" VARCHAR(255) NOT NULL,
    "name" TEXT NOT NULL,
    "age" BIGINT,
    "sex" TEXT,
    "description" TEXT,
    "gym_id" BIGINT
);

CREATE TABLE "gym_passes" (
    "gp_id" SERIAL PRIMARY KEY, -- Автоинкрементируемый ID
    "name" TEXT NOT NULL,
    "price" BIGINT NOT NULL,
    "description" TEXT NOT NULL
);

CREATE TABLE "gyms" (
    "gym_id" SERIAL PRIMARY KEY, -- Автоинкрементируемый ID
    "name" TEXT NOT NULL,
    "address" TEXT NOT NULL,
    "description" TEXT NOT NULL
);

CREATE TABLE "fitness_equipment" (
    "equipment_id" SERIAL PRIMARY KEY, -- Автоинкрементируемый ID
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL
);

CREATE TABLE "events" (
    "event_id" SERIAL PRIMARY KEY, -- Автоинкрементируемый ID
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "coach_id" BIGINT NOT NULL
);

CREATE TABLE "availability" (
    "id" SERIAL PRIMARY KEY, -- Автоинкрементируемый ID
    "gym_id" BIGINT NOT NULL,
    "equipment_id" BIGINT NOT NULL,
    "amount" BIGINT NOT NULL
);

CREATE TABLE "current_events" (
    "id" SERIAL PRIMARY KEY, -- Автоинкрементируемый ID
    "gym_id" BIGINT NOT NULL,
    "event_id" BIGINT NOT NULL,
    "time" TIME(0) WITHOUT TIME ZONE NOT NULL,
    "data" DATE NOT NULL
);

-- Внешние ключи
ALTER TABLE "events" ADD CONSTRAINT "events_coach_id_foreign" FOREIGN KEY ("coach_id") REFERENCES "coaches"("coach_id");
ALTER TABLE "coaches" ADD CONSTRAINT "coaches_gym_id_foreign" FOREIGN KEY ("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE "availability" ADD CONSTRAINT "availability_gym_id_foreign" FOREIGN KEY ("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE "admins" ADD CONSTRAINT "admins_gym_id_foreign" FOREIGN KEY ("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE "users" ADD CONSTRAINT "users_pass_id_foreign" FOREIGN KEY ("pass_id") REFERENCES "gym_passes"("gp_id");
ALTER TABLE "users" ADD CONSTRAINT "users_gym_id_foreign" FOREIGN KEY ("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE "current_events" ADD CONSTRAINT "current_events_gym_id_foreign" FOREIGN KEY ("gym_id") REFERENCES "gyms"("gym_id");
ALTER TABLE "current_events" ADD CONSTRAINT "current_events_event_id_foreign" FOREIGN KEY ("event_id") REFERENCES "events"("event_id");
ALTER TABLE "availability" ADD CONSTRAINT "availability_equipment_id_foreign" FOREIGN KEY ("equipment_id") REFERENCES "fitness_equipment"("equipment_id");
