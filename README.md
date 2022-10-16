# K8S-Client-go-接口调用

## 基于kubernetes中的client-go接口调用各类api资源，并封装。

## 预计提供接口：

### Namespace
1.**CreateNamespace**: 创建Namespace

2.**DeleteNamespace**: 删除Namespace

3.**GetNamespaceList**: 获取Namespace list
### Deployment
1.**CreateDeployment**: 创建deployment

2.**DeleteDeployment**: 删除deployment

3.**UpdateDeploymentImage**: 更新deployment镜像

4.**UpdateDeploymentReplica**: 更新deployment副本

5.**GetDeploymentList**: 获取deployment list

6.**GetDeployment**: 获取特定deployment
### Pod
1.**CreatePod**: 创建Pod

2.**DeletePod**: 删除Pod

3.**GetPodList**: 获取Pod list

4.**GetPod**: 获取特定Pod
### Service
1.**CreateService**: 创建Service

2.**DeleteService**: 删除Service

3.**GetServiceList**: 获取Service list

4.**GetService**: 获取特定Service
### Job
