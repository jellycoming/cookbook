```shell
# 消费者组的消费情况
./bin/kafka-consumer-groups.sh --describe --group {group_name} --bootstrap-server localhost:9092

# 查看版本号
./bin/kafka-broker-api-versions.sh --version --bootstrap-server localhost:9092 

#### Topics
# 查看所有topics
./bin/kafka-topics.sh --list --bootstrap-server localhost:9092
# 查看指定topic的详细信息
./bin/kafka-topics.sh --describe --topic {topicName} --bootstrap-server localhost:9092
# 创建topic
./bin/kafka-topics.sh --create --topic {topicName} --bootstrap-server localhost:9092 --partitions 2 --replication-factor 1
```