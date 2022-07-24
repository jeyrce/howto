

### docker swarm 不支持特权模式?

```text
root@uos-PC:/tmp# docker stack deploy collect -c udcp-collect.yml
Ignoring unsupported options: network_mode, privileged, restart

Ignoring deprecated options:

container_name: Setting the container name is not supported.

Creating network collect_default
Creating service collect_udcp_collect
root@uos-PC:/tmp# docker stack ls
NAME                SERVICES            ORCHESTRATOR
base                6                   Swarm
collect             1                   Swarm
udcp                22                  Swarm
```
