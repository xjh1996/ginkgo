# 预置资源
对于测试用例依赖的资源，比如集成 Cargo，Elastic，并不适合每个用例执行前都去对接一次，然在执行后取消对接一次。对于这种公共的资源，由框架层面做处理，在 testsuit 执行前对资源进行对接处理，在整个用例库执行结束后再取消对接。

现有如下[预置资源](https://github.com/caicloud/zeus/tree/master/framework/config/preset) `Auth`, `cargo`, `cicd`, `elastic`, `storage`. 预置资源的执行规则是:

testsuit 执行前
***
a. 检查 testsuit 需要哪些预置资源

b. 检查配置文件/CLI 是否传入配置资源信息 

c.  检查传入资源是否可用  
* 可用  -->   资源对接至系统供，所有用例使用
* 不可用 -->  不使用使用传入资源，框架自行创建资源并做对接

testsuit 执行后
***
a. 判读资源是否由框架预置
* 不是 --> 结束运行，不解绑资源
* 是 --> 对框架预置的资源进行删除，并结束运行


## 典型资源
### 必须资源
某些资源是所有用例的公共基础资源，如预置不当几乎所有用例均会失败。例如 auth 在 config 中 的配置是
```
# If admin-tenant, admin-user and admin-password are not right, testcases will be fail in prepare phase.
# If tenant, password and password are not right, test framework will create right ones.
tenant: test
user: admin
password: Pwd123456
admin-tenant: sys-tenant
admin-user: admin
admin-password: Pwd123456
```
如果框架检测填入信息不全或者租户用户信息鉴权有误，则框架自建租户及用户并分配资源，并将该资源作用于所有用例。
> 如果预置过程中出现异常，框架结束预置资源动作并终止运行。

### 非必须资源
如果预置过程中出现异常，框架跳过该资源预置过程，记录异常且继续运行。影响范围为使用该预置资源部分的用例执行将失败。

例如被集成的 Cargo，此类资源并不是强需，默认框架不会预置，如需跑集成 cargo 的用例需要在配置中设置 `cargo-enable: true`。系统检查 `cargo-host` 资源是否存在，如存在则进行集成对接，反之框架自建一个 cargo 资源并做系统对接。
```
# If set true, must input a available cargo-host service or cargo-info.
cargo-enable: false
cargo-host: https://cargo-31.test.caicloud.xyz
```
> 