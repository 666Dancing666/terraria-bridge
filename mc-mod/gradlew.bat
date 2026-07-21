@echo off
set GRADLE_VERSION=8.5
set GRADLE_DIR=%USERPROFILE%\.gradle\wrapper\dists\gradle-%GRADLE_VERSION%-bin
set GRADLE_ZIP=%GRADLE_DIR%\gradle-%GRADLE_VERSION%-bin.zip
set GRADLE_EXTRACTED=%GRADLE_DIR%\gradle-%GRADLE_VERSION%

if not exist "%GRADLE_EXTRACTED%\bin\gradle.bat" (
    mkdir "%GRADLE_DIR%"
    curl -L -o "%GRADLE_ZIP%" "https://services.gradle.org/distributions/gradle-%GRADLE_VERSION%-bin.zip"
    tar -xf "%GRADLE_ZIP%" -C "%GRADLE_DIR%"
)

call "%GRADLE_EXTRACTED%\bin\gradle.bat" %*
