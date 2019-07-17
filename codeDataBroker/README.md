
Resources
===============================================================================================================================
For Docker based Kafka cluster setup and kafkacat: https://medium.com/@yusufs/getting-started-with-kafka-in-golang-14ccab5fa26

For Kafka architecture: http://cloudurable.com/blog/kafka-architecture/index.html
===============================================================================================================================


DockerCompose
===============================================================================================================================
Edit .env file to change variables used in docker-compose.yml.Ex for ${MY_IP}
===============================================================================================================================


Use kafkacat cli for standalone test of kafka cluster.
===============================================================================================================================
First, open new terminal and type:

Consumer: `kafkacat -C -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 0`

It will listen to topic “foo” in partition 0 (Kafka start the partition index from 0).

Then, from the other terminal you can publish a message specific into partition 0 using this command:

Producer: `echo 'publish to partition 0' | kafkacat -P -b localhost:19092,localhost:29092,localhost:39092 -t foo -p 0`

If success, the first command will retrieve “publish to partition 0” message which sent by second command. You can do it respectively for partition 1, 2, and 3. You must ensure that all partition can receive the message as well as the consumer can subscribe it.
===============================================================================================================================


MongoDB CourseWare:
===============================================================================================================================
  http://portquiz.net:27017/      -> To test outgoing port
===============================================================================================================================
  
