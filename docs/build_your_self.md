# 自行构建教程

### 依赖
* gitlab(共有私有都可以，这里使用的是：`https://codechina.csdn.net/`)
* 阿里云容器镜像服务 (需要注册阿里云，开通容器镜像服务, 地址是：`https://cr.console.aliyun.com/cn-hangzhou/instances/repositories`)

### 创建gitlab账号和仓库
1. 注册gitlab
2. 创建仓库，可以直接fork我的项目
![](./images/5.png)
3. 请注意仓库必须先初始化
4. 记住仓库的ssh clone地址
5. 进入仓库的设置页面
![6](./images/6.png)
7. 使用ssh生成一对秘钥
![7](./images/7.png)
8.把公钥填写到部署秘钥的值里面，标题随意，需要允许写入权限
![8](./images/8.png)


### gitlab和仓库绑定
1. 进入gitlab, 地址是 [https://codechina.csdn.net/profile/personal_access_tokens](https://codechina.csdn.net/profile/personal_access_tokens)
2. 如图创建一个访问令牌 
![2](./images/2.png)
3. 保存好刚刚的令牌
4. 进入阿里云控制面板 [https://cr.console.aliyun.com/cn-beijing/instances/source](https://cr.console.aliyun.com/cn-beijing/instances/source)
5. 选择 gitlab，绑定你的gitlab的账号 如图所示，
url填写gitlab的首页 这里是`https://codechina.csdn.net/`
用户名填写你的用户名：
![4](./images/4.png)
token就是第3步获取的令牌. 
![3](./images/3.png)
6. 确定绑定完成

### 在阿里云创建镜像仓库
地址：[https://cr.console.aliyun.com/cn-beijing/instances/repositories](https://cr.console.aliyun.com/cn-beijing/instances/repositories)
1. 按照如图完成
![11](./images/11.png)
![12](./images/12.png)

### 初始化 `dpull`

1. 执行 `init`
```shell script
dpull init 
```

2. 修改配置
在`~/.dpull/config.yaml`输出默认配置文件，如下所示：
```yaml
repo_option:
  # 修改成 你自己仓库的ssh克隆地址
  ssh_url: git@codechina.csdn.net:clearcodecn/dpull.git
  # 不能修改
  ssh_user: git
  # 修改成你创建的ssh秘钥的私钥
  private_key: |-
    -----BEGIN OPENSSH PRIVATE KEY-----
    b3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAABFwAAAAdzc2gtcn
    NhAAAAAwEAAQAAAQEA1fZpC6/CjrdDl7DbQmKohLboTYWc9PjclN4FYhPsizRbxBR098RU
    7gTlbJ3vKMIlqHlMCUY2OwEb1rpoAuTu5JEoQl1HKX8eKq41CsfLrkc+VZuTqdteMPqnhK
    JQqrRXnx3CMsWO6KED6v+3UIkRTENxrVyeVXKsKzUsnlVU47KzVEl+RTofHMo4+CaCBFFs
    8ZyebGOHBTKi0JUeLD6PcXrzXCwqPQbCj+8JEda1zx68MU5GArqFuxsgEXI0jsJ3hJxM6x
    DsfdZl+PTdpkkII2kl5gN4HfInrjHc+rp+yvQ0F4CNm2Lx5SvMjiG46eDZDxOCdlqgYKki
    sPco4zQTOwAAA7irpi1bq6YtWwAAAAdzc2gtcnNhAAABAQDV9mkLr8KOt0OXsNtCYqiEtu
    hNhZz0+NyU3gViE+yLNFvEFHT3xFTuBOVsne8owiWoeUwJRjY7ARvWumgC5O7kkShCXUcp
    fx4qrjUKx8uuRz5Vm5Op214w+qeEolCqtFefHcIyxY7ooQPq/7dQiRFMQ3GtXJ5VcqwrNS
    yeVVTjsrNUSX5FOh8cyjj4JoIEUWzxnJ5sY4cFMqLQlR4sPo9xevNcLCo9BsKP7wkR1rXP
    HrwxTkYCuoW7GyARcjSOwneEnEzrEOx91mX49N2mSQgjaSXmA3gd8ieuMdz6un7K9DQXgI
    2bYvHlK8yOIbjp4NkPE4J2WqBgqSKw9yjjNBM7AAAAAwEAAQAAAQBIP8bE7XqzGms2o7/G
    MO5asjDLTJztk8NYeYgz0CqF7w41rfq5V5CeNwUJomMJzlVNCHiGgTD6x6sQ3S0WHRwWDn
    Ybwseu2X/kRaMfmsvKc8A2xCwepTavL1S10uGOYwtbbX8QCenx370k82iBR2eR6wxN0AKf
    M/OzO2dvp7zcjfLjMRKFSbroDOm33mMEqOSfpGJthCWMpHSZzTNlygFrYFfj6ouCzcNJLY
    ibUsoDCelDkhBCDEkIsuGfx3Kcja/AkESrwBvdqDs1VXiwgCnsVue8Un4liNl/UPG5ITkl
    O+0E61LE3XVw2oCm9cawfRNTDHLuJJUi7+/aYsSuxQlBAAAAgQCv+Ms/eFMD8JV1NtH0G9
    QHJXpiyS84G+hFrWSN7QZejrFOEFH0MvLVJypF47k0Ca0P3qh5nLLfHrtoDcM8uOAVtbE6
    LRWjknLJKwprQn5X4wY2iZKLU5mEj3riIeP0Z5RDsePaSKlAmy7yrfbMq4mlN9xzY672am
    UIV0f57FEROAAAAIEA+WdxDkSbnYmcHcu5bvrU242WlqIisj0Z38wLHRVsO9Ci8Sx6LoiY
    4wE8XOf6BLToAvltmnmZsVsoOAymR3M4Mx+EpwgZ3082RfiOA8QJ0iII7EsmYBVwzs2XxD
    H7/MofseeL5hSFEcRFYs7hvKlQ4QHQ53MDD7pO7AwKMqAmgpEAAACBANufBDDoBc5jl2ac
    bSIt1ZSZtejbvMgYDgMl1YeLUX3KDUDEJ22iUnTALSg1e7yImLli30LAWP3G4cgtCnVVGV
    whKjBVEBbPNiEAXupJm9e4ZTLJ627K/4rNjk79sQobbJQhj6VrVQuquNFVTiAiEmOKUB4A
    mDE4JKT55aN8kocLAAAAAAEC
    -----END OPENSSH PRIVATE KEY-----
  # 密码填写创建的密码，没有就填空
  private_key_password: ""
  # 以下几项随意填写
  git_username: dpull
  git_email: dpull@clearcode.cn
  # 在电脑上删除以下仓库的目录
  store_path: /Users/zhangmengjin/.dpull/repo
mirror_option:
  # 不修改，未实现
  state_url: https://dpull.clearcode.cn/api/state
  # 这是阿里云镜像仓库的基础地址 在创建的镜像仓库查看. 
  image_base_path: registry.cn-hangzhou.aliyuncs.com/clearcodecn/mirror
```
3. 配置完成
重新执行 `dpull init`, 或者指定配置文件 `dpull init -c ./config.yaml`

4. 拉取镜像 
`dpull -c ./config.yaml gcr.io/google_containers/kube-scheduler-amd64:v1.9.0`
`dpull gcr.io/google_containers/kube-scheduler-amd64:v1.9.0` (默认配置文件在 ~/.dpull/config.yaml)