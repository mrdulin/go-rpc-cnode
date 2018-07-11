# schedule-commit

部署到`heroku`上： https://dashboard.heroku.com/apps/schedule-commit

查看应用日志:

```bash
☁  schedule-commit [master] heroku logs --app=schedule-commit -t
```

查看`heroku`当前登录的用户:

```bash
☁  schedule-commit [master] ⚡  heroku auth:whoami
 ›   Warning: heroku update available from 7.0.47 to 7.5.11
novaline@aliyun.com
```

在应用部署的虚拟机上运行终端：

```bash
☁  schedule-commit [master] ⚡  heroku run bash --app=schedule-commit
 ›   Warning: heroku update available from 7.0.47 to 7.5.11
Running bash on ⬢ schedule-commit... !
 ▸    ECONNREFUSED: connect ECONNREFUSED 50.19.103.36:5000
```
