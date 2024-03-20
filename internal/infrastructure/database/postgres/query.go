package postgres

const CreateTableQuery = `create table unit_id_%d
(
    user_id bigint    not null
        constraint unit_id_%d_user_id_check
            check ((user_id >= %d) AND (user_id <= %d)),
    time    timestamp not null
);

create index unit_id_%d_user_id_index
    on unit_id_%d (user_id);`

const CheckHaveTableQuery = `select user_id from unit_id_%d limit 1`

const DeleteLastMarksQuery = `delete from unit_id_%d where time < now() - interval '%d seconds' and user_id = %d;`

const CheckQuery = `select count(time) from unit_id_%d where user_id = %d and time > now() - interval '%d seconds' group by user_id;`

const InsertMarkCheckQuery = `insert into unit_id_%d values (%d, now());`
