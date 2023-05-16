import random

class Node:
    def __init__(self, id):
        self.id = id

    def send_message(self, to, data):
        print(f"Node {self.id} sends message to Node {to}: {data}")

    def receive_message(self, message):
        print(f"Node {self.id} receives message from Node {message['from']}: {message['data']}")

def main():
    # Create a set of nodes
    nodes = [Node(1), Node(2), Node(3)]

    # Simulate sending and receiving messages
    while True:
        from_node = random.choice(nodes)
        to_node = random.choice(nodes)
        if from_node != to_node:
            data = f"Message from Node {from_node.id} to Node {to_node.id}"
            message = {'from': from_node.id, 'to': to_node.id, 'data': data}
            to_node.receive_message(message)

if __name__ == '__main__':
    main()
