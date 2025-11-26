@echo off
echo "mesh 启动开始"
cd mesh
start "mesh" main.exe
echo "mesh 启动完成"

timeout /t 2 /nobreak


echo "task 启动开始"
cd ..\task
start "task" main.exe
echo "task 启动完成"

echo "web 启动开始"
cd ..\web
start "web" main.exe
echo "web 启动完成"

echo "game\crash 启动开始"
cd ..\game\crash
start "game\crash" main.exe
echo "game\crash 启动完成"

timeout /t 2 /nobreak

echo "gate 启动开始"
cd ../../gate
start "gate" main.exe
echo "gate 启动完成"

pause