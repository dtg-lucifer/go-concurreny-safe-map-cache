# Consurrency safe map cache in GO

This demo code demonstrates that how much mutext locks can be usable to the concurrency and combining them with vertical sharding of the basic go map, we can achieve a highly reliable local cache using a local map

Apperantly the map uses the hash of the keys to determine that where to put the values in the shards


### This is an example how vertical sharding works
![image](https://github.com/user-attachments/assets/038c82a8-e583-4a97-bc39-729b7a783b19)
