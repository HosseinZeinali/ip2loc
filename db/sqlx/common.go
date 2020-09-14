package sqlx

import (
	"fmt"
	"strings"
)

func (db *Database) Truncate(tableName string) {
	db.Exec(fmt.Sprintf(`TRUNCATE TABLE %s`, tableName))
}

func (db *Database) RenameTable(from string, to string) {
	db.Exec(`ALTER TABLE %s RENAME TO %s`, from, to)
}

func (db *Database) DoesTableExist(table string) bool {
	_, err := db.Query("select * from " + table + ";")

	return err == nil
}

func (db *Database) CreateIpv4Table(tableName string) {
	query := replace(`
-- auto-generated definition
create table {table}
(
    id       serial not null
        constraint {table}_pk
            primary key,
    ip_type  varchar,
    ip_start bigint not null,
    ip_end   bigint not null,
    ip_count integer,
    status   varchar,
    cc       varchar,
    date     date
);

alter table {table}
    owner to project;

create unique index {table}_id_uindex
    on {table} (id);

create index {table}_ip_start_ip_end_index
    on {table} (ip_start, ip_end);

create index {table}_ip_start_ip_end_index_2
    on {table} (ip_start desc, ip_end desc);

create index {table}_ip_end_index
    on {table} (ip_end desc);

create index {table}_ip_start_index
    on {table} (ip_start desc);

create index {table}_ip_end_index_2
    on {table} (ip_end);

create index {table}_ip_start_index_2
    on {table} (ip_start);

`, "table", tableName)
	db.Exec(query)
}

func (db *Database) CreateIpv6Table(tableName string) {
	query := replace(`-- auto-generated definition
create table {table}
(
    id       serial      not null
        constraint {table}_pk
            primary key,
    ip_type  varchar,
    ip_start numeric(39) not null,
    ip_end   numeric(39) not null,
    ip_count integer,
    status   varchar,
    cc       varchar,
    date     date
);

alter table {table}
    owner to project;

create unique index {table}_id_uindex
    on {table} (id);

create index {table}_ip_start_ip_end_index
    on {table} (ip_start, ip_end);

create index {table}_ip_start_ip_end_index_2
    on {table} (ip_start desc, ip_end desc);

create index {table}_ip_end_index
    on {table} (ip_end desc);

create index {table}_ip_start_index
    on {table} (ip_start desc);

create index {table}_ip_end_index_2
    on {table} (ip_end);

create index {table}_ip_start_index_2
    on {table} (ip_start);

`, "table", tableName)

	db.Exec(query)
}

func (db *Database) CreateTableTable() {
	query := `
create table tables
(
    id    serial not null
        constraint tables_pk
            primary key,
    name  varchar,
    extra varchar,
    date  date,
    type  varchar
);

alter table tables
    owner to project;

create unique index tables_id_uindex
    on tables (id);

`
	db.Exec(query)
}

func replace(format string, args ...string) string {
	for i, v := range args {
		if i%2 == 0 {
			args[i] = "{" + v + "}"
		}
	}
	r := strings.NewReplacer(args...)

	return r.Replace(format)
}
