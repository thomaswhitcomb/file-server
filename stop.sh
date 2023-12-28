PID=`ps -eo pid,comm | grep file-server | cut -f3 -d ' '`
echo Killing $PID
kill -9 $PID
