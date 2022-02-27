#!/bin/bash

# 一键导出所有pod描述
for i in $(kubectl -n qfusion-admin get po|awk '{print $1}'|grep -v 'NAME'); do
    echo "$i"
   kubectl -n qfusion-admin get po "$i" -oyaml > /root/ppp/"$i".yml;
done


exit 0
