```shell
# 消费者组的消费情况
./bin/kafka-consumer-groups.sh --describe --group {group_name} --bootstrap-server localhost:9092

# 查看版本号
./bin/kafka-broker-api-versions.sh --version --bootstrap-server localhost:9092 

# 查看topics
./bin/kafka-topics.sh --list --bootstrap-server localhost:9092
```