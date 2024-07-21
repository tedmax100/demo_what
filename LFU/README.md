# LFU Cache
----

## Data structure
- Hash table: The key is the keyword key, and the value is *DLinkNode, which maps to the node address in the linked list. This is used to find the corresponding cache in O(1) time complexity.
- Double linked list: Elements are organized by their usage frequency. Elements with the same frequency are ordered by their access time. The list maintains elements with the same frequency, with elements near the head being the most recently used within that frequency, and elements near the tail being the least recently used within that frequency.

## Action
Implementing LFU requires two actions:
- Get data (Get(key int)): If the key exists in the cache, retrieve the value of the key and update its frequency; otherwise, return -1.
- Write data (Put(key, value int)): If the key already exists, update its value and increase its frequency. If it does not exist, insert the key-value pair into the cache with a frequency of 1. If the insertion operation causes the number of keys to exceed the capacity, the least frequently used key should be evicted. If there is a tie in frequency, the least recently used key within that frequency should be evicted.