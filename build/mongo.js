use admin

db.createUser({ user: "root",
  pwd: "root",
  roles: [ "userAdminAnyDatabase",
    "dbAdminAnyDatabase",
    "readWriteAnyDatabase"
  ]})

db.auth('root', 'root')

use geodata

db.createUser({ user: "guest",
  pwd: "pass",
  roles: [{role: "dbOwner", db: "geodata"}]}
)