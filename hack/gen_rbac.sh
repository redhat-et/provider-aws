#!/bin/bash

cat ./hack/boilerplate.go.txt
echo ""
echo "package main"

echo ""
echo "// +kubebuilder:rbac:groups=\"\",resources=secrets,verbs=get;list;watch"

files=$(find ./config/crd/bases/*.yaml -printf "%f\n" | sort)
for value in $files
do
    if [[ -f "config/crd/bases/$value" ]]; then
        name=$(cat ./config/crd/bases/$value | yq .spec.names.plural | tr -d '"')
        group=$(cat ./config/crd/bases/$value | yq .spec.group | tr -d '"')
        echo ""
        echo "// +kubebuilder:rbac:groups=$group,resources=$name,verbs=get;list;watch;create;update;patch;delete"
        echo "// +kubebuilder:rbac:groups=$group,resources=$name/status,verbs=get;update;patch"
        echo "// +kubebuilder:rbac:groups=$group,resources=$name/finalizers,verbs=update"
    fi
done
