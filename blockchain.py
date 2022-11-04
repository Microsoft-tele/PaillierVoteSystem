import hashlib
import json
from datetime import datetime
from uuid import uuid4
from flask import Flask, jsonify, request
from urllib.parse import urlparse
import requests


class Blockchain(object):
    def __init__(self):
        self.current_data = []
        self.chain = []
        # 创建创世区块
        self.new_block(previous_hash="", proof=100)
        self.nodes = set()

    def new_block(self, proof, previous_hash):
        """
        生成新块
        """
        block = {
            'index': len(self.chain) + 1,
            'timestamp': datetime.now().strftime("%Y-%m-%d %H:%M:%S"),
            'data': self.current_data,
            'proof': proof,
            'previous_hash': previous_hash,
        }
        # 清空当前数据列表
        self.current_data = []
        self.chain.append(block)
        return block

    def new_data(self, data):
        self.current_data.append(data)
        return self.last_block['index'] + 1

    @property
    def last_block(self):
        return self.chain[-1]

    @staticmethod
    def hash(block):
        """
        生成块的 SHA-256 hash值
        """
        # We must make sure that the Dictionary is Ordered, or we'll have inconsistent hashes
        block_string = json.dumps(block, sort_keys=True).encode()
        return hashlib.sha256(block_string).hexdigest()

    def proof_of_work(self, last_proof):
        """
        工作量证明:
        查找一个 p' 使得 hash(pp') 以4个0开头
        p 是上一个块的证明,  p' 是当前的证明
        """
        proof = 0
        while self.valid_proof(last_proof, proof) is False:
            proof += 1
        return proof

    @staticmethod
    def valid_proof(last_proof, proof):
        """
        验证证明: 是否hash(last_proof, proof)以4个0开头
        """
        guess = f'{last_proof}{proof}'.encode()
        guess_hash = hashlib.sha256(guess).hexdigest()
        return guess_hash[:4] == "0000"

    def register_node(self, address):
        """
        向节点列表中添加节点地址
        """
        parsed_url = urlparse(address)
        self.nodes.add(parsed_url.netloc)

    def valid_chain(self, chain):
        """
        验证链的有效性
        """
        last_block = chain[0]
        current_index = 1
        while current_index < len(chain):
            block = chain[current_index]
            print(f'{last_block}')
            print(f'{block}')
            print("-----------")
            # 验证hash值
            if block['previous_hash'] != self.hash(last_block):
                return False
            # 验证proof
            if not self.valid_proof(last_block['proof'], block['proof']):
                return False
            last_block = block
            current_index += 1
        return True

    def resolve_conflicts(self):
        """
        共识算法解决冲突
        使用网络中最长的链
        """
        neighbours = self.nodes
        new_chain = None
        # 寻找最长链
        max_length = len(self.chain)
        # 遍历列表中所有节点
        for node in neighbours:
            response = requests.get(f'http://{node}/chain')
            if response.status_code == 200:
                length = response.json()['length']
                chain = response.json()['chain']
                # 判断是否为更长链以及链的有效性
                if length > max_length and self.valid_chain(chain):
                    max_length = length
                    new_chain = chain
        # 若有更长有效链，则替换并返回True
        if new_chain:
            self.chain = new_chain
            return True
        # 若未替换链，返回False
        return False


# 创建节点
app = Flask(__name__)
# 基于伪随机数随机命名当前节点
node_identifier = str(uuid4()).replace('-', '')
# 实例化Blockchain类
blockchain = Blockchain()


@app.route('/mine', methods=['GET'])
def mine():
    # 挖矿
    last_block = blockchain.last_block
    last_proof = last_block['proof']
    proof = blockchain.proof_of_work(last_proof)
    # 添加区块
    block = blockchain.new_block(proof, blockchain.hash(blockchain.chain[-1]))
    response = {
        'message': "New Block Forged",
        'index': block['index'],
        'data': block['data'],
        'proof': block['proof'],
        'previous_hash': block['previous_hash'],
    }
    return jsonify(response), 200


@app.route('/data/new', methods=['POST'])
def new_data():
    values = request.get_data()
    index = blockchain.new_data(values.decode())
    response = {'message': f'data will be added to Block {index}'}
    mine()
    return jsonify(response), 201


@app.route('/chain', methods=['GET'])
def full_chain():
    response = {
        'chain': blockchain.chain,
        'length': len(blockchain.chain),
    }
    return jsonify(response), 200


@app.route('/nodes/register', methods=['POST'])
def register_nodes():
    """
    注册节点
    POST传参格式:
    {
        "nodes":["http://ip_1:port_1","http://ip_2:port_2", ......]
    }
    :return:
    """
    values = request.get_json()
    nodes = values.get('nodes')
    if nodes is None:
        return "Error: Please supply a valid list of nodes", 400
    for node in nodes:
        blockchain.register_node(node)
    response = {
        'message': 'New nodes have been added',
        'total_nodes': list(blockchain.nodes),
    }
    return jsonify(response), 201


@app.route('/nodes/resolve', methods=['GET'])
def consensus():
    replaced = blockchain.resolve_conflicts()
    if replaced:
        response = {
            'message': 'Our chain was replaced',
            'new_chain': blockchain.chain
        }
    else:
        response = {
            'message': 'Our chain is authoritative',
            'chain': blockchain.chain
        }
    return jsonify(response), 200


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000)
