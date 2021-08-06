#!/bin/bash

CURR_DIR=$PWD
LOGO=$(cat "$CURR_DIR"/common/shell/logo.txt)
GITHUB_PACKAGE_PATH="$CURR_DIR"/github/hooks
GITHUB_PATH="$CURR_DIR"/.git

source "$CURR_DIR"/common/shell/colors.sh
source "$CURR_DIR"/common/shell/ascii.sh

echo "${bold}${blue}$LOGO${normal}"

makeCmd() {
  cmd=$1
  for DIR in "$CURR_DIR"/*; do
    FILE="$DIR"/Makefile
    if [ -f "$FILE" ]; then
      if grep -q "$cmd:" "$FILE"; then
        (
          cd "$DIR" || exit
          echo ""
          border "########### Executing ${magenta}make $1${normal} command in package ${bold}${blue}$DIR${normal} ###########"
          make $cmd
        )
        err=$?
        if [ $err -ne 0 ]; then
          echo -e "\n${bold}${red}An error has occurred during test process ${bold}[FAIL]${norma}\n"
          exit 1
        fi
      fi
    fi
  done
}

checkHooks() {
    err=0
    echo "Checking github hooks..."
    for FILE in "$GITHUB_PACKAGE_PATH"/*; do
      f="$(basename -- $FILE)"
      FILE2="$GITHUB_PATH"/hooks/$f
      if [ -f "$FILE2" ]; then
        if cmp -s "$FILE" "$FILE2"; then
          lineOk "Hook file ${underline}$f${normal} installed and updated"
        else
          lineError "Hook file ${underline}$f${normal} ${red}installed but out-of-date [OUT-OF-DATE]"
          err=1
        fi
      else
        lineError "Hook file ${underline}$f${normal} ${red}not installed [NOT INSTALLED]"
        err=1
      fi
      if [ $err -ne 0 ]; then
        echo -e "\nRun ${bold}make setup-env${normal} to setup your development environment, then try again.\n"
        exit 1
      fi
    done
}

lint() {
  title1 "STARTING LINT"
  out=$(golint $CURR_DIR/... | tee /dev/tty)
  out_err=$?
  err=0
  if [ $out_err -ne 0 ]; then
    echo -e "\n${bold}${red}An error has occurred during lint process\n"
    err=1
    exit 1
  fi
  if [ -n "$out" ]; then
    echo -e "\n${red}Some lint rules are broken ${bold}[WARNING]${normal}\n"
    err=1
    exit 1
  fi
  if [ ! $err -ne 0 ]; then
    lineOk "\nAll lint rules passed"
  fi
}

checkEnvs() {
  title1 "STARTING SECRETS CHECK"
  err="0"
  for DIR in "$CURR_DIR"/*; do
    FILE="$DIR"/.env
    if [ -f "$FILE" ]; then
      data=`cat $FILE`
      if [[ $data =~ SECRET=[0-9A-Za-z]{1,} ]]; then
        lineError "Secret exposed in file file:////$FILE"
        err="1"
      fi
      if [[ $data =~ MONGO_CONNECTION_STRING= && ! $data =~ localhost ]]; then
        lineError "Data base connection string exposed in file file://$FILE"
        err="1"
      fi
    fi
  done
  if [ $err -eq 1 ];then
    exit 1
  else
    lineOk "All env file passed"
  fi
}

echo -e "\n\n"
title1 "STARTING PRE-COMMIT SCRIPT"

checkHooks

if [ "$1" == "lint" ]; then
  lint
elif [ "$1" == "checkEnvs" ]; then
  checkEnvs
else
  echo "Executing with parameter $1"
  makeCmd "$1"
fi

if [ "$1" != "clean" -a "$1" != "lint" -a "$1" != "checkEnvs" -a "$1" != "doc" ]; then
  checkEnvs
  lint
fi
