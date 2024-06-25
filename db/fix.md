## Quick Fix for Dirty Database Issue

```
select * from schema_migrations;

update schema_migrations set dirty =false where version=XXXX;
```