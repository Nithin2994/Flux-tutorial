# bootstrap

flux bootstrap git --url=https://github.com/Nithin2994/Flux-tutorial.git --components-extra=image-reflector-controller,image-automation-controller --username=Nithin2994 --password=Nithin524 --token-auth=true --branch=master --path=clusters/my-cluster


flux bootstrap github --components-extra=image-reflector-controller,image-automation-controller --owner=Nithin2994 --repository=Flux-tutorial --branch=main --path=clusters/my-cluster --read-write-key --personal

mkdir app-podinfo

# podinfo app (Thirdparty app)

Flux create source git podinfo --url="https://github.com/stefanprodan/podinfo" --branch=master --interval=30s --export > ./clusters/my-cluster/app-podinfo/podinfo-source.yaml

Flux create kustomization podinfo --source=podinfo --path="./kustomize" --prune=true --validation=client interval=5m --export > ./clusters/my-cluster/app-podinfo/podinfo-kustomization.yaml

# game-server-golang app (My App)

Flux create source git podinfo --url="https://github.com/Nithin2994/GameServer-golang" --branch=master --interval=30s --export > ./clusters/my-cluster/game-server/game-server-source.yaml

Flux create kustomization game-server --source=GameServer-golang --path="./kustomize" --export > ./clusters/my-cluster/game-server/game-server-kustomization.yaml

Github token -> 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

flux create image repository game-server --image=nithin524/gameserver-golang --interval=1m --export > ./clusters/my-cluster/game-server/game-server-registry.yaml

flux create image policy podinfo --image-ref=podinfo --select-semver=1.x --export > ./clusters/my-cluster/game-server/game-server-policy.yaml