@echo off
git add .

set /p commitgmsg=CommitMessage:

git commit -m "%commitgmsg%"

git push origin master

git tag

set /p tag=Tag:

if "%tag%" == "" goto END

echo Go to END
goto END

git tag %tag%

git push origin master

:END