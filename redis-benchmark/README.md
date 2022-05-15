## 第五次作业
1. 使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

利用redis-benchmark 命令来测试
```
# d 代表value 的大小，单位byte
redis-benchmark -h 127.0.0.1 -p 6383 -d 10 -t get,set
redis-benchmark -h 127.0.0.1 -p 6383 -d 20 -t get,set

```
测试结果

value size (byte) | set  性能 | get  性能
---|---|---
10 | 100000 requests completed in 0.54 seconds| 100000 requests completed in 0.50 seconds
20 | 100000 requests completed in 0.54 seconds| 100000 requests completed in 0.50 seconds
50 | 100000 requests completed in 0.55 seconds| 100000 requests completed in 0.50 seconds


2. 写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。