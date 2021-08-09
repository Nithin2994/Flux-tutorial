flux bootstrap git --url=https://github.com/Nithin2994/Flux-tutorial.git --username=Nithin2994 --password=Nithin524 --token-auth=true --branch=master --path=clusters/my-cluster

mkdir app-podinfo

Flux create source git podinfo --url="https://github.com/stefanprodan/podinfo" --branch=master --interval=30s --export > ./clusters/my-cluster/app-podinfo/podinfo-source.yaml

Flux create kustomization podinfo --source=podinfo --path="./kustomize" --prune=true --validation=client interval=5m export > ./clusters/my-cluster/app-podinfo/podinfo-kustomization.yaml
