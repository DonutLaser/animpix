@echo off
go build -ldflags -H=windowsgui
ResourceHacker -open animpix.exe -save animpix.exe -action addskip -res assets/images/icon.ico -mask ICONGROUP,MAIN,
xcopy /s /y assets\* D:\Programo\custom\animpix\assets\
xcopy /y animpix.exe D:\Programos\custom\animpix\