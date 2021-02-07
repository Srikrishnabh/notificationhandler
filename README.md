# notification handler
## Assignment problem statement 
    Most applications have the need to implement notifications for a variety of use cases and scenarios. Create a centralized generic service for notification that can be used by a variety consuming application for their notification needs e.g. an incident workflow system may use this system when each incident ticket moves from one state to another, similarly a order management system may use this service to notify the customer of the status of the order whenever it changes
    The system should allow for the following capabilities:
       1. Accept messages including from, to and subject
       2. Ability to notify on multiple channels (e.g email, slack, you can stub out/mock if required)
       3. Deliver messages in correct order for each consumer of this

## Approach
    1. notification handler service uses kafka to consume the messages async and send notification via channels such as email, slack etc
    2. pkg/pb/notifier/notifier.proto contains message aggreement to be passed to consumer.
    3. notification.to will be the partioin key for sending message via producer. for ordering of the messages.
    4. we can persist the consumed message to databas, and can schedule the retry for the messages failed to send.
    
## Steps to deploy
    1. Deploy strimzi kafka operator using https://strimzi.io/quickstarts/ 
    2. Name the kafka cluster as data-pipeline-cluster
        2.1 > wget https://strimzi.io/examples/latest/kafka/kafka-persistent-single.yaml
        2.2 > change name to data-pipeline-cluster
        2.3 > specify the storage as per need
        2.4 > kubectl apply -f kafka-persistent-single.yaml -n kafka
    3. Deploy producer (send dummy messages) and consumer to consume messages.
        3.1 > kubectl apply -f producer.yaml
        3.2 > kubectl apply -f consumer.yaml
    4. Note: notification_handler.properties is been set to above configurations to work.

## Custom deploy
    1. Bring up kafka cluster.
    2. edit notification_handler.properties 
        2.1 specify kafka broker address list in KafkaConsumerGroup and KafkaProducers
    2. run > make build.
    3. By default Docker image also be created.
    4. run producer and consumer in different terminals or as docker containers.
    5. consumer should print logs on consuming messages and sending notifications(dummy message)
