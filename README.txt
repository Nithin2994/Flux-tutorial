---------NEW CLUSTER------------------
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

public key: ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDDhe9cR38X8hT1dBxsLKnHdhL442JtP4A943RUf5WyUnOVKsVyo7Fwb4xGVGnLCbyneCjXIH8UomWRgWUpVLBpFfWEqklXw1h6AD8RUwYiNST54/ePWYRlQ48bIobXYNHtS1jRSUw6JGEvt9dzoUbmAdxxIoT+Cj+TJ9PYDvGFeRrrcmEBV8/qNXHo/b/uSodU57OkBWdhkGZorRZTodeFS5ZlfVyZ6yGMZfLff90mq8kbVHHxOnlMl3HsbcaAzjMBaaiIR0YxuL+ZE6lfdtHpWSAbQFX4Qs+G77Ei1q911gUoa7SymXUGIuycuRysbLwQMdj8OSakYLRWwR9p92I3

flux bootstrap github \
  --owner=$GITHUB_USER \
  --repository=Flux-tutorial \
  --branch=main \
  --components-extra=image-reflector-controller,image-automation-controller \
  --path=./clusters/new-cluster \
  --log-level=debug \
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
  --interval=5m \
  --export > ./clusters/new-cluster/gameserver-golang-kustomization.yaml

flux create secret git gameserver-golang-credentials \
    --url=https://github.com/Nithin2994/GameServer-golang \
    --username=Nithin2994 \
    --password=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

flux create image repository gameserver-golang-image-repo --image=nithin524/gameserver-golang --interval=1m --export > ./clusters/my-cluster/game-server/game-server-image-registry.yaml

flux create image policy gameserver-golang-image-policy \
    --image-ref=game-server \
    --select-semver=1.0.x

flux create image update gameserver-golang-image-update --git-repo-ref=gameserver-golang --git-repo-path="./kustomize" --checkout-branch=master --push-branch=master --author-name=Nithin2994 --author-email=nithinreddy2994@gamil.com --commit-template="{{range .Updated.Images}}{{println .}}{{end}}" --export > ./clusters/new-cluster/game-server-automation.yaml


gpg --export-secret-keys \
  --armor 366435D2CEA9151EC95B1183FAD044D593279BD4 |
kubectl create secret generic git-gpg \
  --namespace=flux-system \
  --from-file=git.asc=/dev/stdin


  pub   rsa4096 2021-09-29 [SC]
      366435D2CEA9151EC95B1183FAD044D593279BD4
uid                      nithinreddy <nithinreddy2994@gmail.com>
sub   rsa4096 2021-09-29 [E]

kubectl create secret generic github-game-server --namespace flux-system --from-literal=token=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

-------HELM-----------
pwd : /Users/nithin.reddy/Documents/GitHub/Flux-tutorial/clusters/new-cluster/helm-charts

helm create game-server
helm install game-server-helm game-server/
helm uninstall game-server-helm