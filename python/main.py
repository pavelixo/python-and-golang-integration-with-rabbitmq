from json import dumps as json
from pika import BlockingConnection, ConnectionParameters

connection = BlockingConnection(ConnectionParameters('rabbitmq'))
channel = connection.channel()

message = json({"message": "Hello, Golang"})

channel.queue_declare(queue='golang')
channel.basic_publish(exchange='', routing_key='golang', body=message)

connection.close()