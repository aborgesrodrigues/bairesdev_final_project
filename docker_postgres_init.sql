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


-- populate users table

INSERT INTO public.users(username, "name", birthday) VALUES('username1', 'User Name 1', '2000-01-01');
INSERT INTO public.users(username, "name", birthday) VALUES('username2', 'User Name 2', '1995-05-02');
INSERT INTO public.users(username, "name", birthday) VALUES('username3', 'User Name 3', '2001-07-03');
INSERT INTO public.users(username, "name", birthday) VALUES('username4', 'User Name 4', '1980-01-10');
INSERT INTO public.users(username, "name", birthday) VALUES('username5', 'User Name 5', '2005-04-12');