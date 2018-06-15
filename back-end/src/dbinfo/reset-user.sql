use dynauth_backend;


drop table auth25;
delete from testLocks where userid = 25;
delete from testKeysDisplay where userid = 25;
delete from testConfigLog where userid = 25;
update testUsers set init = false where id = 25;

delete from testPass where userid = 9;
delete from testPassDisplay where userid = 9;
delete from testConfigLog where userid = 9;
update testUsers set init = false where id = 9;