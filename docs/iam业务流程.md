# iam业务流程

## init
创建root用户，写入“system”级别policy：`{id: 1, subject: root, obj: system::*, effect: allow}`,则此时`root.res_domain=[system::*]`

## 正式开始
1.user1 自行注册

2. root创建tenant1, 同时默认创建角色tenant1::admin, 并将`tenant1.admin`与user1绑定,同时写入“tenant”级别policy：`{id: 2, subject: system::tenant1:admin, obj: system::tenant1::*, effect: allow}`,则此时`user1.res_domain=[system::tenant1::*]`
```
sub: root
resource: system::tenant1,
action: write
use_policy: 1
allow: true
```

3. user2 自行注册

4. user1创建角色：tenant1::role1
```
sub: system::tenant::admin
resource: system::tenant1::role,
action: write
use_policy: 2
allow: true
```

5. user1将user2与tenant1::role1进行绑定
```
sub: system::tenant::admin
resource: system::tenant1::role,
action: write
use_policy: 2
allow: true
```

6. user1写入“tenant”级别policy： `{id: 3, subject: system::tenant1::role1, obj: system::tenant1::article, action: read, effect: allow}`, 则此时user2可对system::tenant1::res1进行读, user2.res_domain=[system::tenant1::article(read)]
```
sub: system::tenant::admin
resource: system::tenant1::policy,
action: write
use_policy: 2
allow: true
```

7. user2读取tenant1中的某篇article
```
sub: system::tenant1::role1
resource: system::tenant1::article,
action: read
use_policy: 3
allow: true
```

