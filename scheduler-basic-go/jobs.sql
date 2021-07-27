CREATE TABLE IF NOT EXISTS "public"."jobs" (   
   "id"      SERIAL PRIMARY KEY,   
   "name"    varchar(50) NOT NULL,   
   "payload" text,   
   "runAt"   TIMESTAMP NOT NULL  
)