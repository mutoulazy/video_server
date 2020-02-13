# video_server
go video server

## database
### users
id int(10) PK auto_increment
login_name varchar(64) UNK
pwd text

### video_info 
id varchar(64) PK
author_id int(10) 
name text
display_ctime text
create_time datetime CURRENT_TIMESTAMP

### comments
id varchar(64) PK
video_id varchar(64)
author_id int(10)
content text
time datetime CURRENT_TIMESTAMP

### sessions
session_id tinytext PK
TTL tinytext
login_name text