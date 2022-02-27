@echo off
@rem #"Makefile" for Windows projects.
@rem #Copyright (c) shiguanghuxian, 2022. All rights reserved.
@rem
SETLOCAL

@rem ###############
@rem # PRINT USAGE #
@rem ###############

if [%1]==[] goto usage

@rem ################
@rem # SWITCH BLOCK #
@rem ################

@rem # make build
if /I [%1]==[build] call :build

@rem # make clean
if /I [%1]==[clean] call :clean

goto :eof

@rem #############
@rem # FUNCTIONS #
@rem #############
:build
for /f "delims=" %%i in ('go version') do (set go_version=%%i)
for /f "delims=" %%i in ('git rev-parse HEAD') do (set git_hash=%%i)
@go build -ldflags "-X gohttp/gosrc/gohttp.VERSION=1.0.0 -X 'gohttp/gosrc/gohttp.BUILD_TIME=%DATE% %TIME%' -X 'gohttp/gosrc/gohttp.GO_VERSION=%go_version%' -X gohttp/gosrc/gohttp.GIT_HASH=%git_hash%" -o bin\windows\libgohttp.dll ./
exit /B %ERRORLEVEL%

:clean
@del /S /F /Q "bin\windows\libgohttp.dll"
exit /B 0

:usage
@echo Usage: %0 ^[ build ^| clean ^]
exit /B 1