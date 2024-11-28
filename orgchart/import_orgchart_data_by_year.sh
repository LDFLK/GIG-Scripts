#!/bin/bash

# Check if a year is passed as an argument
if [ -z "$1" ]; then
  echo "Usage: $0 <year>"
  exit 1
fi

YEAR=$1

# List of all files to process
FILES=(
  "extracted/2010-2020/gazette-2006-1-2.csv"
  "extracted/2010-2020/gazette-2006-1-3.csv"
  "extracted/2010-2020/gazette-2006-1-4.csv"
  "extracted/2010-2020/gazette-2006-1-5.csv"
  "extracted/2010-2020/gazette-2010-4-30.csv"
  "extracted/2010-2020/gazette-2010-11-22.csv"
  "extracted/2010-2020/gazette-2011-3-17.csv"
  "extracted/2010-2020/gazette-2013-2-6.csv"
  "extracted/2010-2020/gazette-2013-2-18.csv"
  "extracted/2010-2020/gazette-2013-6-25.csv"
  "extracted/2010-2020/gazette-2013-8-16.csv"
  "extracted/2010-2020/gazette-2013-10-17.csv"
  "extracted/2010-2020/gazette-2014-8-14.csv"
  "extracted/2010-2020/gazette-2014-9-22.csv"
  "extracted/2010-2020/gazette-2014-11-11.csv"
  "extracted/2010-2020/gazette-2015-1-10.csv"
  "extracted/2010-2020/gazette-2015-1-18.csv"
  "extracted/2010-2020/gazette-2015-3-9.csv"
  "extracted/2010-2020/gazette-2015-4-6.csv"
  "extracted/2010-2020/gazette-2015-5-5.csv"
  "extracted/2010-2020/gazette-2015-9-21.csv"
  "extracted/2010-2020/gazette-2015-10-15.csv"
  "extracted/2010-2020/gazette-2015-11-24.csv"
  "extracted/2010-2020/gazette-2015-12-18.csv"
  "extracted/2010-2020/gazette-2016-3-18.csv"
  "extracted/2010-2020/gazette-2016-8-8.csv"
  "extracted/2010-2020/gazette-2017-6-9.csv"
  "extracted/2010-2020/gazette-2017-7-6.csv"
  "extracted/2010-2020/gazette-2017-8-17.csv"
  "extracted/2010-2020/gazette-2017-9-8.csv"
  "extracted/2010-2020/gazette-2017-12-20.csv"
  "extracted/2010-2020/gazette-2018-2-20.csv"
  "extracted/2010-2020/gazette-2018-2-20-2.csv"
  "extracted/2010-2020/gazette-2018-3-28.csv"
  "extracted/2010-2020/gazette-2018-5-12.csv"
  "extracted/2010-2020/gazette-2018-5-22.csv"
  "extracted/2010-2020/gazette-2018-5-31.csv"
  "extracted/2010-2020/gazette-2018-12-28.csv"
  "extracted/2010-2020/gazette-2018-12-28-2.csv"
  "extracted/2010-2020/gazette-2019-1-10.csv"
  "extracted/2010-2020/gazette-2019-1-16.csv"
  "extracted/2010-2020/gazette-2019-3-18.csv"
  "extracted/2010-2020/gazette-2019-6-4.csv"
  "extracted/2010-2020/gazette-2019-12-10.csv"
  "extracted/2010-2020/gazette-2019-12-31.csv"
  "extracted/2010-2020/gazette-2020-1-9.csv"
  "extracted/2010-2020/gazette-2020-1-13.csv"
  "extracted/2010-2020/gazette-2020-1-22.csv"
  "extracted/2010-2020/gazette-2020-1-24.csv"
  "extracted/2020-2030/gazette-2020-2-1.csv"
  "extracted/2020-2030/gazette-2020-2-7.csv"
  "extracted/2020-2030/gazette-2020-3-17.csv"
  "extracted/2020-2030/gazette-2020-4-8.csv"
  "extracted/2020-2030/gazette-2020-5-22.csv"
  "extracted/2020-2030/gazette-2020-8-9.csv"
  "extracted/2020-2030/gazette-2020-9-25.csv"
  "extracted/2020-2030/gazette-2020-10-6.csv"
  "extracted/2020-2030/gazette-2020-11-20.csv"
  "extracted/2020-2030/gazette-2020-12-11.csv"
  "extracted/2020-2030/gazette-2021-5-3.csv"
  "extracted/2020-2030/gazette-2021-5-17.csv"
  "extracted/2020-2030/gazette-2021-6-3.csv"
  "extracted/2020-2030/gazette-2021-6-18.csv"
  "extracted/2020-2030/gazette-2021-7-7.csv"
  "extracted/2020-2030/gazette-2021-7-8.csv"
  "extracted/2020-2030/gazette-2021-7-29.csv"
  "extracted/2020-2030/gazette-2021-8-16.csv"
  "extracted/2020-2030/gazette-2021-10-6.csv"
  "extracted/2020-2030/gazette-2021-11-17.csv"
  "extracted/2020-2030/gazette-2021-12-2.csv"
)

# Iterate through the list and execute `go run` for files matching the given year
for FILE in "${FILES[@]}"; do
  if [[ $FILE == *"$YEAR"* ]]; then
    echo "Processing: $FILE"
    go run import_csv.go "$FILE"
  fi
done
