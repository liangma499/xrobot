#!/bin/bash
# 当前目录
readonly directory=$(cd "$(dirname "$0")" && pwd)
# 服务信息
readonly services="${directory}/services.info"
readonly restart_log="restart_log"

# 
restart_services() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0  ]
    then
      #echo "${line}"
      line2=$(echo "${line}" | tr -d '\r')
      IFS2=" " read -r -a arr <<< "${line2}"
      #echo "${arr[2]}"
        # 进入项目目录
      if cd "${arr[1]}" 
      then
       # echo "${arr[1]}"
        # 停止之前的进程
        if [ -e "${arr[2]}" ]
        then
          pid=$(cat "${arr[2]}")
      
          if [ "$pid" ]; then
            #echo "${pid}"
            runnow=$(date  +'%Y-%m-%d %H:%M:%S')
            if ps -ef | grep "${pid}" | grep -v grep > /dev/null
            then
              runnowDay=$(date  +'%Y-%m-%d')
              echo "${runnow} <${arr[0]}> is live" >> "${directory}/${restart_log}/live_${runnowDay}.log"
            else
              chmod +x main

              # 切换配置文件
              etc_file_ext=${etc_file##*.}
              etc_file_name=${etc_file%.*}
              if [ -e "${arr[3]}/${etc_file_name}-${env}.${etc_file_ext}" ]
              then
                mv "${arr[3]}/${etc_file_name}-${env}.${etc_file_ext}" "${arr[3]}/${etc_file}"
              fi
              # 启动项目
              ./main &
              sleep 1s
              runnowDay=$(date  +'%Y-%m-%d')
              echo "${runnow} <${arr[0]}> restart success" >> "${directory}/${restart_log}/start_${runnowDay}.log"
            fi
          fi
        fi
      fi 
    fi
    cd "${directory}"
    index=$index+1
  done < "$1"

}
for i in {1..60}; do
  restart_services "${services}"
  sleep 1s
done
