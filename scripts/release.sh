#!/bin/bash

latest_tag=$(git tag --sort=-v:refname | head -n1)

if [ -z "$latest_tag" ]; then
    new_tag="v0.0.1"
else
    IFS='.' read -r major minor patch <<<"${latest_tag#v}"
    patch=$((patch + 1))
    new_tag="v${major}.${minor}.${patch}"
fi

while git rev-parse "$new_tag" >/dev/null 2>&1; do
    patch=$((patch + 1))
    new_tag="v${major}.${minor}.${patch}"
done

echo "Last tag: $latest_tag"
echo "New tag: $new_tag"

git tag "$new_tag"
git push origin "$new_tag"
