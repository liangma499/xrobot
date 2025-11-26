#!/bin/bash

# 远程主机
#readonly remote_host='root@192.168.0.108'
readonly remote_host='root@24.144.94.62'
# 二进制执行目录
readonly bin_dir='E:/game-run/xrobot'
readonly bin_dir_bin='E:/game-run'
readonly bin_dir_bin_pkg='./xrobot'
readonly pack_name_prefix='tron-robot'

# 启动环境
readonly env='prd'
# 临时编译目录
readonly temp_make_dir='./temp'
# 当前目录
readonly directory=$(cd "$(dirname "$0")" && pwd)
# 服务信息
readonly services="${directory}/services.info"
# 部署脚本
readonly script="${directory}/deploy.sh"
# 配置文件
readonly etc_file='etc.toml'
# 配置文件
readonly sshkey='./crroot'
readonly not_copy_script='true'
# 罗列所有服务
list() {
  < "${services}" tr ' ' ' ' | nl
}

# 上传部署脚本
copy_script_to_remote() {

  if [ "${not_copy_script}" == 'true' ] 
  then
    if [ "${sshkey}" ]
    then
      echo "开始上传<部署脚本> ssh key no_script" 
      # scp -i "${sshkey}" "${services}" ${remote_host}:${bin_dir}
      echo "结束上传<部署脚本> > ssh key no_script"
    else
      echo "开始上传<部署脚本> passwd no_script"
      # scp "${services}" ${remote_host}:${bin_dir}
      echo "结束上传<部署脚本> passwd no_script"
    fi
  else
    if [ "${sshkey}" ]
    then
      echo "开始上传<部署脚本> ssh key"
      scp -i "${sshkey}" "${services}" ${remote_host}:${bin_dir}
      scp -i "${sshkey}" "${script}" ${remote_host}:${bin_dir}
      echo "结束上传<部署脚本> > ssh key"
    else
      echo "开始上传<部署脚本> passwd"
      scp "${services}" ${remote_host}:${bin_dir}
      scp "${script}" ${remote_host}:${bin_dir}
      echo "结束上传<部署脚本> passwd"
    fi
  fi
}

# 复制部署脚本
copy_script_to_local() {
  if [ "${not_copy_script}" == 'true' ] 
  then
    echo "开始拷贝<部署脚本 no_script>"
   # cp "${services}" "${bin_dir}"
    echo "结束拷贝<部署脚本>"
  else
    echo "开始拷贝<部署脚本 no_script>"
    cp "${services}" "${bin_dir}"
    cp "${script}" "${bin_dir}"
    echo "结束拷贝<部署脚本>"
  fi

}

# 编译服务
make() {
  case $1 in
  'local')
    make_to_local "$2"
    ;;
  'remote')
    make_to_remote "$2"
    ;;
  esac
}

# 编译到本地服务器
make_to_local() {
  case $1 in
  'all')
    make_all_to_local
    copy_script_to_local
    cd ${bin_dir_bin}
    packName=${pack_name_prefix}-bin-$(date "+%Y-%m-%d-%H-%M-%S").tar.gz
    echo $packName
    tar czvf $packName ${bin_dir_bin_pkg}/*
    ;;
  *)
    make_one_to_local "$1"
    copy_script_to_local
    tar_one_to_local "$1"
    ;;
  esac
}

# 编译所有项目到本地服务器

tar_one_to_local() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0 ]
    then
      IFS=" " read -r -a arr <<< "${line}"
      if [ "${arr[0]}" = "$1" ]
      then
        
        cd ${bin_dir_bin}
        name=$(echo "$1" | tr 'A-Z' 'a-z')

        #path = "${arr[1]/./|/}"

        runPath="${arr[1]}"
        runPath=${runPath:2}

        packName=${pack_name_prefix}-$name-bin-$(date "+%Y-%m-%d-%H-%M-%S").tar.gz
        echo $packName
        tar czvf $packName ${bin_dir_bin_pkg}/$runPath/*
        break
      fi
    fi
    index=$index+1
  done < "${services}"
}

# 编译到远程服务器
make_to_remote() {
  case $1 in
  'all')
    make_all_to_remote
    cd ${bin_dir}
    tar czvf ${pack_name_prefix}-bin-$(date "+%Y-%m-%d-%H-%M-%S").tar.gz  ./*
    copy_script_to_remote
    ;;
  *)
    make_one_to_remote "$1"
    copy_script_to_remote
    ;;
  esac
}

# 编译所有项目到本地服务器
make_all_to_local() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0  ]
    then
      make_line "${line}"
      copy_line_to_local "${line}"
    fi
    index=$index+1
  done < "${services}"

}

# 编译所有服务到远程服务器
make_all_to_remote() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0  ]
    then
      make_line "${line}"
      copy_line_to_remote "${line}"
    fi
    index=$index+1
  done < "${services}"
}

# 编译单个服务
make_one_to_remote() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0 ]
    then
      IFS=" " read -r -a arr <<< "${line}"
      if [ "${arr[0]}" = "$1" ]
      then
        make_line "${line}"
        copy_line_to_remote "${line}"
      fi
    fi
    index=$index+1
  done < "${services}"
}

# 编译单个服务到本地
make_one_to_local() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0 ]
    then
      IFS=" " read -r -a arr <<< "${line}"
      if [ "${arr[0]}" = "$1" ]
      then
        make_line "${line}"
        copy_line_to_local "${line}"
      fi
    fi
    index=$index+1
  done < "${services}"
}

# 执行编译
make_line() {
  # 移除换行符
  line=$(echo "$1" | tr -d '\r')
  # 转换成数组
  IFS=" " read -r -a arr <<< "${line}"
  # 工作目录
  work_dir="${arr[1]#*/}"
  # 项目准备目录
  ready_dir="${temp_make_dir}/${work_dir}"
  # 项目上传目录
  upload_dir="${temp_make_dir}/${work_dir%/*}"

  # 创建临时编译目录
  cd "${arr[1]}" || return
  mkdir -p "${ready_dir}"

  # 开始编译
  echo "开始编译<${arr[0]}>"

  # 交叉编译
  CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

  # 复制二进制文件
  mv main "${ready_dir}"

  # 复制etc配置目录
  if [ "${arr[3]}" != '' ]
  then
    cp -f -r "${arr[3]}" "${ready_dir}"
  fi

  # 复制config配置目录
  if [ "${arr[4]}" != '' ]
  then
    cp -f -r "${arr[4]}" "${ready_dir}"
  fi

  echo "结束编译<${arr[0]}>"

  # 退出到根目录
  cd "${directory}" || return
}

# 拷贝配置服务
copy() {
  case $1 in
  'local')
    copy_to_local "$2" "$3"
    ;;
  'remote')
    copy_to_remote "$2" "$3"
    ;;
  esac
}

# 复制配置到本地服务器
copy_to_local() {
  case $1 in
  'all')
    copy_all_to_local
    ;;
  *)
    copy_one_to_local "$1"
    ;;
  esac
}

# 复制配置到远程服务器
make_conf_to_remote() {
  case $1 in
  'all')
    copy_script_to_remote
    ;;
  *)
    copy_one_to_remote "$1"
    ;;
  esac
}

# 编译所有项目配置到本地服务器
copy_all_to_local() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0  ]
    then
      copy_line_to_local "${line}"
    fi
    index=$index+1
  done < "${services}"
}

# 复制到远程服务器
copy_line_to_remote() {
  # 移除换行符
  line=$(echo "$1" | tr -d '\r')
  # 转换成数组
  IFS=" " read -r -a arr <<< "${line}"
  # 工作目录
  work_dir="${arr[1]#*/}"
  # 项目准备目录
  ready_dir="${temp_make_dir}/${work_dir}"
  # 项目上传目录
  upload_dir="${temp_make_dir}/${work_dir%/*}"

  # 进入项目目录
  cd "${arr[1]}" || return

  # scp上传到服务器
  if [ "${sshkey}" ]
  then
      echo "开始上传<${arr[0]}> sskey"
      scp -i "${sshkey}" -r "${upload_dir}" ${remote_host}:${bin_dir}
      echo "结束上传<${arr[0]}> sskey"
  else
      echo "开始上传<${arr[0]}>"
      scp  -r "${upload_dir}" ${remote_host}:${bin_dir}
      echo "结束上传<${arr[0]}>"
     
  fi
  # 删除临时编译目录
  rm -rf "${temp_make_dir}"

  # 退出到根目录
  cd "${directory}" || return
}

# 复制到本地服务器
copy_line_to_local() {
  # 没有目录创建目录
  mkdir -p "${bin_dir}"
  # 移除换行符
  line=$(echo "$1" | tr -d '\r')
  # 转换成数组
  IFS=" " read -r -a arr <<< "${line}"
  # 工作目录
  work_dir="${arr[1]#*/}"
  # 项目准备目录
  ready_dir="${temp_make_dir}/${work_dir}"
  # 项目上传目录
  upload_dir="${temp_make_dir}/${work_dir%/*}"

  # 进入项目目录
  cd "${arr[1]}" || return

  # scp上传到服务器
  echo "开始拷贝<${arr[0]}>"
  cp -r "${upload_dir}" "${bin_dir}"
  echo "结束拷贝<${arr[0]}>"

  # 删除临时编译目录
  rm -rf "${temp_make_dir}"

  # 退出到根目录
  cd "${directory}" || return
}

# 启动服务
start() {
  case $1 in
  'all')
    start_all
    ;;
  *)
    start_one "$1"
    ;;
  esac
}

# 启动所有服务
start_all() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0  ]
    then
      start_line "${line}"
    fi
    index=$index+1
  done < "${services}"
}

# 启动单个服务
start_one() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0 ]
    then
      IFS=" " read -r -a arr <<< "${line}"
      if [ "${arr[0]}" = "$1" ]
      then
        start_line "${line}"
      fi
    fi
    index=$index+1
  done < "${services}"
}

start_line() {
  # 移除换行符
  line=$(echo "$1" | tr -d '\r')
  # 转换成数组
  IFS=" " read -r -a arr <<< "${line}"

  # 进入项目目录
  cd "${arr[1]}" || return

  # 停止之前的进程
  if [ -e "${arr[2]}" ]
  then
    pid=$(cat "${arr[2]}")
    if ps -ef | grep "${pid}" | grep -v grep > /dev/null
    then
      kill -9 "${pid}"
      sleep 1s
    fi
  fi

  # 执行权限
  chmod +x main

  # 切换配置文件
  etc_file_ext=${etc_file##*.}
  etc_file_name=${etc_file%.*}
  if [ -e "${arr[3]}/${etc_file_name}-${env}.${etc_file_ext}" ]
  then
    mv "${arr[3]}/${etc_file_name}-${env}.${etc_file_ext}" "${arr[3]}/${etc_file}"
  fi

  # 启动项目
  echo "开始启动<${arr[0]}>"
  ./main &
  sleep 1s
  echo "结束启动<${arr[0]}>"

  # 退出到根目录
  cd "${directory}" || return
}

# 停止服务
stop() {
  case $1 in
  'all')
    stop_all
    ;;
  *)
    stop_one "$1"
    ;;
  esac
}

# 停止所有服务
stop_all() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0  ]
    then
      stop_line "${line}"
    fi
    index=$index+1
  done < "${services}"
}

# 停止单个服务
stop_one() {
  index=0
  while IFS= read -r line; do
    if [ $index != 0 ]
    then
      IFS=" " read -r -a arr <<< "${line}"
      if [ "${arr[0]}" = "$1" ]
      then
        stop_line "${line}"
      fi
    fi
    index=$index+1
  done < "${services}"
}

stop_line() {
  # 移除换行符
  line=$(echo "$1" | tr -d '\r')
  # 转换成数组
  IFS=" " read -r -a arr <<< "${line}"

  # 进入项目目录
  cd "${arr[1]}" || return

  # 停止项目
  echo "开始停止<${arr[0]}>"
  if [ -e "${arr[2]}" ]
  then
    pid=$(cat "${arr[2]}")
    if ps -ef | grep "${pid}" | grep -v grep > /dev/null
    then
      kill -9 "${pid}"
      sleep 1s
    fi
  fi
  echo "结束停止<${arr[0]}>"

  # 退出到根目录
  cd "${directory}" || return
}

case $1 in
'list')
    list
    ;;
'make')
    make "$2" "$3"
    ;;
'copy')
    copy "$2" "$3"
    ;;
'start')
    start "$2"
    ;;
'stop')
    stop "$2"
    ;;
'restart')
    start "$2"
    ;;
*)
    echo "list make copy start stop restart"
    ;;
esac
