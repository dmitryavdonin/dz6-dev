helm repo add datawire https://www.getambassador.io
helm repo update
kubectl apply -f https://app.getambassador.io/yaml/edge-stack/3.5.1/aes-crds.yaml
kubectl wait --timeout=90s --for=condition=available deployment emissary-apiext -n emissary-system

helm install edge-stack --namespace ambassador datawire/edge-stack --set emissary-ingress.createDefaultListeners=true --set emissary-ingress.agent.cloudConnectToken=eyJhbGciOiJQUzUxMiIsInR5cCI6IkpXVCJ9.eyJsaWNlbnNlX2tleV92ZXJzaW9uIjoidjIiLCJjdXN0b21lcl9pZCI6ImRtaXRyeS5hdmRvbmluQGdtYWlsLmNvbS0xNjgyNDUyMjQ3IiwiY3VzdG9tZXJfZW1haWwiOiJkbWl0cnkuYXZkb25pbkBnbWFpbC5jb20iLCJlbmFibGVkX2ZlYXR1cmVzIjpbIiIsImZpbHRlciIsInJhdGVsaW1pdCIsInRyYWZmaWMiLCJkZXZwb3J0YWwiXSwiZW5mb3JjZWRfbGltaXRzIjpbeyJsIjoiZGV2cG9ydGFsLXNlcnZpY2VzIiwidiI6NX0seyJsIjoicmF0ZWxpbWl0LXNlcnZpY2UiLCJ2Ijo1fSx7ImwiOiJhdXRoZmlsdGVyLXNlcnZpY2UiLCJ2Ijo1fSx7ImwiOiJ0cmFmZmljLXVzZXJzIiwidiI6NX1dLCJtZXRhZGF0YSI6e30sImV4cCI6MTcxMzk4ODI0NywiaWF0IjoxNjgyNDUyMjQ3LCJuYmYiOjE2ODI0NTIyNDd9.FVmpt-QNcsiJwNbx6EKHkhr9CmQbQ_TytKLUpkt7zvpipg9e93geK-OiTtV0v1qgkmtPAltk9_9E3aauHi2aLUty-TVrsztZmtMc_Sf6eKE0bLEq3IIB8ZnaO3KISLVE7bmK3UAJJ3QR57bGWLzwwIzNsZD_jE4gyqlQnZ4qSHvxHXysHgpc-xXTE5ftDn6iCym5f-MsTFLDVRwJ-acMu5KtWlbHwZXjJCKx8s_-7rb63bwdBM-IMFrpfpEO0Eh3v9pW0oJqHE5C6hToW5VPtvn9lwF1BIKpPfDqqqZ24qVJr0EjWnd0ucBdRjzWzuvyroJIOxoqyIH0OMMnSytvXA