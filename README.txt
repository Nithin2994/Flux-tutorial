Flux create source git podinfo \
  --url="https://github.com/stefanprodan/podinfo" \
  --branch=master \
  --interval=30s \
  --export > ./clusters/my-cluster/app-podinfo/podinfo-source.yaml