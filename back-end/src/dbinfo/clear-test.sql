use dynauth_backend;
drop table auth2;
drop table auth3;
delete from testPass;
delete from testLocks;
delete from testKeysDisplay;
delete from testPassDisplay;
update testUsers set init = false;