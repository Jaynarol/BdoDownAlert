rmdir bin /s /q

rsrc -ico res/icon.ico -o src/icon.syso
cd src
go build -o ../bin/BdoDownAlert.exe
cd ..

xcopy src\assets\* bin\assets\ /d /y /e
copy src\setting.ini bin\setting.ini

del /f src\icon.syso
pause