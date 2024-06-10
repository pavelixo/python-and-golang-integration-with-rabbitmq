from json import dumps as json
from time import sleep
from pika import BlockingConnection, ConnectionParameters
from pika.exceptions import AMQPConnectionError

def test_rabbitmq_connection():
  is_connected = False
  while not is_connected:
    try:
      connection = BlockingConnection(ConnectionParameters('rabbitmq')); connection.close()
      print("Successfully connected to RabbitMQ"); is_connected = True
    except AMQPConnectionError:
      print("Failed to connect to RabbitMQ. Retrying in 5 seconds...")
      sleep(5)

def main():
  connection = BlockingConnection(ConnectionParameters('rabbitmq'))
  channel = connection.channel()

  message = json({"message": "Hello, Golang"}); print(message)

  channel.queue_declare(queue='golang')
  channel.basic_publish(exchange='', routing_key='golang', body=message)

  connection.close()

if __name__ == '__main__':
  test_rabbitmq_connection(); main()