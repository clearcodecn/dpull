# dpull


### workflow 

1. user init repo
2. user add image and push to remote repo
3. the aliyun ci will build it and push to aliyun mirror
4. client download it. 

by top 4 points, we can do:
1. docker query image from aliyun, if exist download
2. use normal ways

git + ci workflow
for example: we want to pull image: `k8s.gcr.io/pause:3.2`
in bash view:

1. clone our repo
git clone git@xxx.com/clearcodecn/repo.git

2. modify docker file
`echo 'k8s.gcr.io/pause:3.2' > Dockerfile`

3. commit and tag and push
```bash
tag=$(echo -n $image | base64)
git add . 
git commit -m 'add image'
git tag -a -m 'add image'
git push -u origin release-v$tag
```
4. ci will build an image like `registry.aliyun.com/xxx/xxx/mirror:release-v${tag}`
5. pull image and rename
```shell script
docker pull registry.aliyun.com/xxx/xxx/mirror:release-v${tag}
image=$(base64 -d ${tag})
docker tag registry.aliyun.com/xxx/xxx/mirror:release-v${tag} $image 
docker rmi registry.aliyun.com/xxx/xxx/mirror:release-v${tag}
```
 
more info please [中文文档](docs/zh_cn.md)