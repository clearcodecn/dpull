# dpull 镜像拉取工具

原理：`dpull` 主要用于解决拉取国外镜像困难的痛点，事先本项目在国内gitlab的代码托管方建了一个仓库用于保存dockerfile
在阿里云镜像服务绑定了该账号，该项目一旦发生更改，会立即执行镜像构建

1. 项目中只有一个Dockerfile文件
2. 每次需要拉取镜像，会先对镜像base64，然后从阿里云仓库拉取tag为base64的镜像，如果没有的话则会推送代码进行构建
3. 构建成功，拉取镜像，并将镜像名更改成原名字

工具无法保证阿里云构建的速度，但是一旦构建完毕，拉取是非常快的