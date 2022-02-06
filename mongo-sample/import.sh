#! /bin/bash
cd /docker-entrypoint-initdb.d
for coll in *; do
    if [ -d "$coll" ]; then
        for file in $coll/*; do
            mongoimport --db "$coll" --collection "$(basename $file .json)" --file $file
        done
    fi
done