# readme
10.21.2.219
默认连接http://10.21.2.219:9200，可在交互模式和命令行模式中修改
```shell
/opt/script/elktools es -a http://127.0.0.1:9200 -u user -p password index
```

## 使用方法
```shell
/opt/script/elktools # 进入交互模式
/opt/script/elktools es index # 命令行模式执行
/opt/script/elktools -h # 帮助文档
```

## 支持的功能
* 交互模式下tab键自动补全
* 查看health状态
* 查看nodes列表
* 查看index列表，
* 查看pendingtasks列表
* 查看shards列表
* 重试分配失败的分片

## 例子
```shell
/opt/script/elktools es index
/opt/script/elktools es index -1
/opt/script/elktools es index 2020.01.01
/opt/script/elktools es index --health yellow
/opt/script/elktools es index all --health yellow
/opt/script/elktools es index --less
/opt/script/elktools es health --interval 3

# 自定义参数参考elastic search api
/opt/script/elktools es index h=health,status,index,pri.store.size

# 重试分配失败的分片
/opt/script/elktools es route retry
```

## 过滤
```shell
--day      # 过滤指定时间的索引列表
--desc     # 默认采用升序输出数据，可指定降序输出
--grep     # 过滤字符串
--health   # 过滤指定状态的索引，[green yellow red] (default: green)
--interval # 持续输出数据时间间隔，单位秒
--less     # 支持linux的less浏览数据
--number   # 由0开始对所有输出的行数编号
--wc       # 统计输出的行数
```