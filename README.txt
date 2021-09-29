# bootstrap

flux bootstrap git --url=https://github.com/Nithin2994/Flux-tutorial.git --components-extra=image-reflector-controller,image-automation-controller --username=Nithin2994 --password=Nithin524 --token-auth=true --branch=main --path=clusters/my-cluster

flux bootstrap github   --owner=Nithin2994 --repository=Flux-tutorial --branch=master --path=clusters/my-cluster --personal

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

flux create image update game-server --git-repo-ref=game-server --git-repo-path="./kustomize" --checkout-branch=master --push-branch=master --author-name=Nithin2994 --author-email=nithinreddy2994@gamil.com --commit-template="{{range .Updated.Images}}{{println .}}{{end}}" --export > ./clusters/my-cluster/game-server/game-server-automation.yaml


flux bootstrap github --hostname=github.com --ssh-hostname=github.com --owner=Nithin2994 --repository=GameServer-golang --path=clusters/my-cluster --personal

flux create secret git game-server-secret --url="https://github.com/Nithin2994/GameServer-golang" --username=Nithin2994 --password=Nithin524

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

public key: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDhe9cR38X8hT1dBxsLKnHdhL442JtP4A943RUf5WyUnOVKsVyo7Fwb4xGVGnLCbyneCjXIH8UomWRgWUpVLBpFfWEqklXw1h6AD8RUwYiNST54/ePWYRlQ48bIobXYNHtS1jRSUw6JGEvt9dzoUbmAdxxIoT+Cj+TJ9PYDvGFeRrrcmEBV8/qNXHo/b/uSodU57OkBWdhkGZorRZTodeFS5ZlfVyZ6yGMZfLff90mq8kbVHHxOnlMl3HsbcaAzjMBaaiIR0YxuL+ZE6lfdtHpWSAbQFX4Qs+G77Ei1q911gUoa7SymXUGIuycuRysbLwQMdj8OSakYLRWwR9p92I3

flux bootstrap github \
  --owner=$GITHUB_USER \
  --repository=Flux-tutorial \
  --branch=main \
  --components-extra=image-reflector-controller,image-automation-controller \
  --path=./clusters/new-cluster \
  --personal

  flux create source git gameserver-golang \
  --url=https://github.com/Nithin2994/GameServer-golang \
  --branch=master \
  --interval=30s \
  --export > ./clusters/new-cluster/GameServer-golang-source.yaml

  flux create kustomization gameserver-golang \
  --source=gameserver-golang \
  --path="./kustomize" \
  --prune=true \
  --validation=client \
  --interval=5m \
  --export > ./clusters/new-cluster/gameserver-golang-kustomization.yaml