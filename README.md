# Consensus Algorithms

Consensus algorithms are used in distributed systems to ensure that all nodes in the system agree on a common state or value. There are different types of consensus algorithms with varying trade-offs between performance, fault tolerance, and complexity. In this document, we will explore some of the most popular consensus algorithms.

## Practical Byzantine Fault Tolerance (PBFT)

Practical Byzantine Fault Tolerance (PBFT) is a consensus algorithm used in distributed systems to tolerate up to a third of the nodes being Byzantine (malicious). It is designed for systems that require high throughput and low latency. In PBFT, each node has a copy of the state, and transactions are broadcasted to the network. The nodes then exchange messages to reach a consensus on the state of the system.

## Raft

Raft is a consensus algorithm used to ensure fault tolerance in distributed systems. It is designed to be easy to understand and implement and is used in systems that require moderate to high availability. In Raft, nodes elect a leader, and the leader is responsible for managing the state of the system. The leader replicates the state to the other nodes in the system, ensuring that all nodes have an up-to-date copy of the state.

## Paxos

Paxos is a consensus algorithm used in distributed systems to ensure that all nodes agree on a common value. It is designed to be highly fault-tolerant and is used in systems that require high availability. In Paxos, nodes propose values, and the nodes exchange messages to reach a consensus on the value to be agreed upon.

## Proof of Work (PoW)

Proof of Work (PoW) is a consensus algorithm used in blockchain systems to validate transactions and create new blocks. In PoW, nodes compete to solve a cryptographic puzzle, and the first node to solve the puzzle is rewarded with new coins. PoW is used in systems like Bitcoin, where security is paramount.

## Proof of Stake (PoS)

Proof of Stake (PoS) is a consensus algorithm used in blockchain systems that aims to address the energy consumption issue of PoW. In PoS, nodes are chosen to validate transactions based on the amount of cryptocurrency they hold, rather than their computational power. PoS is used in systems like Ethereum, where energy efficiency is a significant concern.

## Delegated Proof of Stake (DPoS)

Delegated Proof of Stake (DPoS) is a consensus algorithm used in blockchain systems that allows token holders to delegate their voting power to a trusted set of nodes. These nodes are responsible for validating transactions and creating new blocks. DPoS is used in systems like EOS and BitShares, where performance and scalability are critical.

## Proof of Elapsed Time (PoET)

Proof of Elapsed Time (PoET) is a consensus algorithm used in distributed systems that aims to be more energy-efficient than PoW. In PoET, nodes wait for a random amount of time before proposing a new block, and the node that waits the longest is selected to create the block. PoET is used in systems like Hyperledger Sawtooth.

## Federated Byzantine Agreement (FBA)

Federated Byzantine Agreement (FBA) is a consensus algorithm used in distributed systems that combines the benefits of PBFT and PoS. In FBA, nodes are organized into groups, and each group has a set of trusted nodes responsible for validating transactions and creating new blocks. FBA is used in systems like Stellar.

## Tendermint Core

Tendermint Core is a consensus algorithm used in blockchain systems that uses a modified version of the PBFT algorithm. It is designed to be fast and secure and is used in systems like Cosmos and Binance Chain.

These are just some of the most popular consensus algorithms used in Distributed Systems
