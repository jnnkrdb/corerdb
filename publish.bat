@echo off
git add .

set /p commitgmsg=CommitMessage:

git commit -m "%commitgmsg%"

git push origin master

git tag

set /p tag=Tag:

git tag %tag%

git push origin master