create table if not exists urls (
	id bigserial primary key,
	original_url text not null,
	short_url text not null,
	created_at timestampz not null default now(),

	constraint unique_short_url unique (short_url)
);

create table if not exists statistics (
	id bigserial primary key,
	url_id bigint not null,
	clicks bigint not null default 0,
	last_accessed timestampz default null,

	constraint fk_statistics_urls foreign key (url_id) references urls (id) on delete cascade
);


create table if not exists fingerprint (
	id bigserial primary key,
	url_id bigint not null,
	ip text not null,
	user_agent text not null,
	created_at timestampz not null default now(),

	constraint fk_fingerprint_urls foreign key (url_id) references urls (id) on delete cascade
);

-- FK индексы
create index idx_statistics_urls on statistics (url_id);
create index idx_fingerprint_urls on fingerprint (url_id);

-- Индексы оптимизации
create index idx_urls_short_url on urls (short_url);
create index idx_statistics_clicks on statistics (clicks);
create index idx_statistics_last_accessed on statistics (last_accessed);
create index idx_fingerprint_ip on fingerprint (ip);
create index idx_fingerprint_user_agent on fingerprint (user_agent);
