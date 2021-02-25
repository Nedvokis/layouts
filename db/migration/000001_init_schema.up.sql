CREATE TABLE "sta_room" (
	"id" bigserial PRIMARY KEY,
	"bitrix_id" bigint UNIQUE NOT NULL,
	"type_name" varchar
);
CREATE TABLE "sta_statuses" (
	"id" bigserial PRIMARY KEY,
	"bitrix_id" bigint UNIQUE NOT NULL,
	"type_name" varchar
);
CREATE TABLE "sta_types" (
	"id" bigserial PRIMARY KEY,
	"bitrix_id" bigint UNIQUE NOT NULL,
	"type_name" varchar
);
CREATE TABLE "complexes" (
	"id" bigserial PRIMARY KEY,
	"bitrix_id" bigint UNIQUE NOT NULL,
	"name" varchar
);
CREATE TABLE "litters" (
	"id" bigserial PRIMARY KEY,
	"parent" bigint NOT NULL,
	"bitrix_id" bigint UNIQUE NOT NULL,
	"name" varchar
);
CREATE TABLE "layouts" (
	"id" bigserial PRIMARY KEY,
	"parent" bigint NOT NULL,
	"area" float,
	"citchen_area" float,
	"door" int,
	"floor" int,
	"bitrix_id" int,
	"layout_id" int,
	"living_area" float,
	"num" varchar,
	"price" int,
	"status" int,
	"type" int,
	"room" int,
	"layouts_url" varchar,
	"svg_path" text
);
ALTER TABLE "litters"
ADD FOREIGN KEY ("parent") REFERENCES "complexes" ("bitrix_id");
ALTER TABLE "layouts"
ADD FOREIGN KEY ("parent") REFERENCES "litters" ("bitrix_id");
ALTER TABLE "layouts"
ADD FOREIGN KEY ("room") REFERENCES "sta_room" ("bitrix_id");
ALTER TABLE "layouts"
ADD FOREIGN KEY ("status") REFERENCES "sta_statuses" ("bitrix_id");
ALTER TABLE "layouts"
ADD FOREIGN KEY ("type") REFERENCES "sta_types" ("bitrix_id");