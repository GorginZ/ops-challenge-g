#!/usr/bin/env bash
# WIP this is in spin-up stack but fiddling with stuff here bc it's a lot to have in the stack-up script hope to have something more eligant soon still better than hard coding aws service Ips tho...

die() { echo “${1:-something went wrong}“; exit “${2:-1}“; }

hash wget || die "wget not found"

die() { echo “${1:-something went wrong}“; exit “${2:-1}“; }

hash jq || die "jq not found"



wget https://ip-ranges.amazonaws.com/ip-ranges.json


res=$(jq -r '[.prefixes[] | select(.region=="ap-southeast-2" and .service=="S3").ip_prefix] | .[]' < ip-ranges.json)

# don't think i need list count
no=$(echo "$res"| wc -l)
# echo $no
# are there always 5

set -- $res
s3cidrone=$1
s3cidrtwo=$2
s3cidrthree=$3
s3cidrfour=$4
s3cidfive=$5
# for value in $res
# do
#     echo $value

# done