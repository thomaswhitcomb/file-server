PID=`ps | grep file-server | cut -f2 -d ' '`
echo Killing $PID
kill $PID
