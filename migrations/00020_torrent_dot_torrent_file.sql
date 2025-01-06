-- +goose Up
-- +goose StatementBegin

create table torrent_torrent_dot_file
(
  info_hash    bytea                    not null
    primary key
    references torrents
      on delete cascade,
  binary_file       bytea                    not null,
  created_at   timestamp with time zone not null
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP table IF EXISTS torrent_torrent_dot_file;



-- +goose StatementEnd
