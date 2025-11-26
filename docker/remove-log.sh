#!/bin/bash

curdate=$(date -d '-1 day' +'%Y%m%d')
runnow=$(date  +'%Y-%m-%d %H:%M:%S')
removelogInfo="/data/cr_vn_bin/*/log/due.${curdate}.log*"

rm -rf ${removelogInfo}
echo "${runnow} remove ${removelogInfo}" >> /data/cr_vn_bin/remove.log

removelogInfo2="/data/cr_vn_bin/*/*/log/due.${curdate}.log*"

rm -rf ${removelogInfo2}
echo "${runnow} remove ${removelogInfo2}" >> /data/cr_vn_bin/remove.log