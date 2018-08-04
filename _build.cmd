rmdir bin /s /q

rsrc -ico resources/icon.ico -o icon.syso -arch="amd64"
go build -o bin/BdoDownAlert.exe

xcopy assets\* bin\assets\ /d /y /e
copy setting.ini bin\setting.ini

del /f icon.syso
pause