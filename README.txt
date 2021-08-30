# bootstrap

flux bootstrap git --url=https://github.com/Nithin2994/Flux-tutorial.git --components-extra=image-reflector-controller,image-automation-controller --username=Nithin2994 --password=Nithin524 --token-auth=true --branch=master --path=clusters/my-cluster

flux bootstrap github --components-extra=image-reflector-controller,image-automation-controller --owner=Nithin2994 --repository=Flux-tutorial --branch=master --path=clusters/my-cluster --read-write-key --personal

mkdir app-podinfo

# podinfo app (Thirdparty app)

Flux create source git podinfo --url="https://github.com/stefanprodan/podinfo" --branch=master --interval=30s --export > ./clusters/my-cluster/app-podinfo/podinfo-source.yaml

Flux create kustomization podinfo --source=podinfo --path="./kustomize" --prune=true --validation=client interval=5m --export > ./clusters/my-cluster/app-podinfo/podinfo-kustomization.yaml

# game-server-golang app (My App)

Flux create source git game-server --url="https://github.com/Nithin2994/GameServer-golang" --branch=master --interval=30s --export > ./clusters/my-cluster/game-server/game-server-source.yaml

Flux create source git game-server --url="git@github.com:Nithin2994/Flux-tutorial.git" --branch=master --interval=30s --export > ./clusters/my-cluster/game-server/game-server-source.yaml

Flux create kustomization game-server --source=GameServer-golang --path="./kustomize" --export > ./clusters/my-cluster/game-server/game-server-kustomization.yaml

Github token -> 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

flux create image repository game-server --image=nithin524/gameserver-golang --interval=1m --export > ./clusters/my-cluster/game-server/game-server-registry.yaml

flux create image policy game-server --image-ref=game-server --select-semver=1.x --export > ./clusters/my-cluster/game-server/game-server-policy.yaml

flux create image update game-server --git-repo-ref=game-server --git-repo-path="./clusters/my-cluster" --checkout-branch=master --push-branch=master --author-name=Nithin2994 --author-email=nithinreddy2994@gamil.com --commit-template="{{range .Updated.Images}}{{println .}}{{end}}" --export > ./clusters/my-cluster/game-server/game-server-automation.yaml


flux bootstrap github --hostname=github.com --ssh-hostname=github.com --owner=Nithin2994 --repository=GameServer-golang --path=clusters/my-cluster --personal