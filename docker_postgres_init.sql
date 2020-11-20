-- public.users definition

CREATE TABLE public.users (
	id bigserial NOT NULL,
	username text NOT NULL,
	"name" text NOT NULL,
	birthday timestamptz NOT NULL,
	CONSTRAINT users_pkey PRIMARY KEY (id)
);


-- public.questions definition

CREATE TABLE public.questions (
	id bigserial NOT NULL,
	"statement" text NOT NULL,
	answer text NULL,
	user_id int8 NOT NULL,
	CONSTRAINT questions_pkey PRIMARY KEY (id)
);


-- public.questions foreign keys

ALTER TABLE public.questions ADD CONSTRAINT fk_questions_user FOREIGN KEY (user_id) REFERENCES users(id) ON UPDATE RESTRICT ON DELETE RESTRICT;