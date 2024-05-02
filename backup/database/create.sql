CREATE TABLE public."user" (
	id serial8 ,
	user_id varchar ,
	password varchar ,
	email varchar ,
	status varchar NULL,
	actice_email varchar NULL,
	connection_id varchar null,
	created_time TIMESTAMP  NULL,
	updated_time TIMESTAMP  NULL
);
