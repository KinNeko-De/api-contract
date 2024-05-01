#!/bin/bash

git fetch --prune --prune-tags

existingfeaturebranches=$(git branch -r | sed 's/^ *//;s/ *$//' | egrep -v "(^\*|main)")
echo "Existing feature branches: $existingfeaturebranches"

for tag in $(git tag --list '*v[0-9]*\.[0-9]*\.[0-9]*-*')
do
  echo "Tag: $tag"
  # regex to match the tag name: pattern is v1.2.3-featurename.buildnumber
  [[ $tag =~ ^golang/kinnekode/.+/v[0-9]*\.[0-9]*\.[0-9]*-(.*)\.([0-9]*)$ ]]
  featurebranchname=${BASH_REMATCH[1]}
  echo "Feature branch name: $featurebranchname"
  if [[ $existingfeaturebranches =~ "feature/$featurebranchname" ]]
  then
    echo "Branch $featurebranchname exists, so not deleting tag $tag"
  else
    echo "Branch $featurebranchname does not exist, so deleting tag $tag"
    git push origin --delete $tag
    git tag -d $tag
  fi
done

read -p "Press any key to end."