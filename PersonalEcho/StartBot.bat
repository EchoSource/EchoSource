:start
@echo off
SETLOCAL ENABLEDELAYEDEXPANSION
SET old=new
SET new=PersonalEcho
for /f "tokens=*" %%f in ('dir /b *.exe') do (
  SET newname=%%f
  SET newname=!newname:%old%=%new%!
  move "%%f" "!newname!"
)
TITLE Personal Echo Bot
TIMEOUT 3
COLOR 03
PersonalEcho.exe
goto start