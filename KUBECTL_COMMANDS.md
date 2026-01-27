# Kubectl - Guia de Comandos

## Verificar Conexão

#### Ver informações do cluster
```bash
docker compose exec app kubectl cluster-info
```

#### Ver nodes
```bash
docker compose exec app kubectl get nodes
```

#### Ver contexto atual
```bash
docker compose exec app kubectl config current-context
```

#### Listar contextos disponíveis
```bash
docker compose exec app kubectl config get-contexts
```

#### Trocar contexto
```bash
docker compose exec app kubectl config use-context <context-name>
```

## Listar Recursos

#### Listar pods em todos os namespaces
```bash
docker compose exec app kubectl get pods -A
```

#### Listar pods em namespace específico
```bash
docker compose exec app kubectl get pods -n <namespace>
```

#### Listar deployments
```bash
docker compose exec app kubectl get deployments -A
```

#### Listar services
```bash
docker compose exec app kubectl get services -A
```

#### Listar nodes
```bash
docker compose exec app kubectl get nodes
```

#### Listar namespaces
```bash
docker compose exec app kubectl get namespaces
```

#### Listar todos os recursos
```bash
docker compose exec app kubectl get all -A
```

#### Listar recursos com mais detalhes
```bash
docker compose exec app kubectl get pods -o wide
```

#### Listar com formato customizado
```bash
docker compose exec app kubectl get pods -o custom-columns=NAME:.metadata.name,STATUS:.status.phase
```

## Ver Detalhes

#### Descrever pod
```bash
docker compose exec app kubectl describe pod <pod-name> -n <namespace>
```

#### Descrever deployment
```bash
docker compose exec app kubectl describe deployment <deployment-name> -n <namespace>
```

#### Descrever service
```bash
docker compose exec app kubectl describe service <service-name> -n <namespace>
```

#### Descrever node
```bash
docker compose exec app kubectl describe node <node-name>
```

#### Ver YAML de recurso
```bash
docker compose exec app kubectl get pod <pod-name> -n <namespace> -o yaml
```

#### Ver JSON de recurso
```bash
docker compose exec app kubectl get pod <pod-name> -n <namespace> -o json
```

## Logs

#### Ver logs de pod
```bash
docker compose exec app kubectl logs <pod-name> -n <namespace>
```

#### Ver logs em tempo real
```bash
docker compose exec app kubectl logs -f <pod-name> -n <namespace>
```

#### Ver logs de container específico
```bash
docker compose exec app kubectl logs <pod-name> -c <container-name> -n <namespace>
```

#### Ver logs anteriores (após restart)
```bash
docker compose exec app kubectl logs <pod-name> -n <namespace> --previous
```

#### Ver logs com timestamp
```bash
docker compose exec app kubectl logs <pod-name> -n <namespace> --timestamps
```

#### Ver últimas N linhas
```bash
docker compose exec app kubectl logs <pod-name> -n <namespace> --tail=50
```

## Executar Comandos

#### Abrir shell no pod
```bash
docker compose exec app kubectl exec -it <pod-name> -n <namespace> -- /bin/bash
```

#### Abrir shell em container específico
```bash
docker compose exec app kubectl exec -it <pod-name> -c <container-name> -n <namespace> -- /bin/bash
```

#### Executar comando único
```bash
docker compose exec app kubectl exec <pod-name> -n <namespace> -- <comando>
```

#### Executar comando com argumentos
```bash
docker compose exec app kubectl exec <pod-name> -n <namespace> -- ls -la /app
```

## Port Forward

#### Fazer port forward de um pod
```bash
docker compose exec app kubectl port-forward <pod-name> <local-port>:<pod-port> -n <namespace>
```

#### Fazer port forward de um service
```bash
docker compose exec app kubectl port-forward svc/<service-name> <local-port>:<service-port> -n <namespace>
```

#### Port forward em background
```bash
docker compose exec app kubectl port-forward <pod-name> 8080:80 -n <namespace> &
```

## Editar Recursos

#### Editar deployment
```bash
docker compose exec app kubectl edit deployment <deployment-name> -n <namespace>
```

#### Editar service
```bash
docker compose exec app kubectl edit service <service-name> -n <namespace>
```

#### Editar configmap
```bash
docker compose exec app kubectl edit configmap <configmap-name> -n <namespace>
```

#### Aplicar mudanças de arquivo
```bash
docker compose exec app kubectl apply -f <file.yaml>
```

#### Aplicar mudanças de diretório
```bash
docker compose exec app kubectl apply -f <directory>/
```

## Escalar Recursos

#### Escalar deployment
```bash
docker compose exec app kubectl scale deployment <deployment-name> --replicas=3 -n <namespace>
```

#### Escalar replicaset
```bash
docker compose exec app kubectl scale replicaset <replicaset-name> --replicas=5 -n <namespace>
```

#### Auto-escalar deployment
```bash
docker compose exec app kubectl autoscale deployment <deployment-name> --min=2 --max=10 --cpu-percent=80 -n <namespace>
```

## Gerenciar Deployments

#### Reiniciar deployment
```bash
docker compose exec app kubectl rollout restart deployment <deployment-name> -n <namespace>
```

#### Ver status do rollout
```bash
docker compose exec app kubectl rollout status deployment <deployment-name> -n <namespace>
```

#### Ver histórico de rollout
```bash
docker compose exec app kubectl rollout history deployment <deployment-name> -n <namespace>
```

#### Fazer rollback
```bash
docker compose exec app kubectl rollout undo deployment <deployment-name> -n <namespace>
```

#### Rollback para revisão específica
```bash
docker compose exec app kubectl rollout undo deployment <deployment-name> --to-revision=2 -n <namespace>
```

## Criar Recursos

#### Criar namespace
```bash
docker compose exec app kubectl create namespace <namespace-name>
```

#### Criar deployment
```bash
docker compose exec app kubectl create deployment <deployment-name> --image=<image> -n <namespace>
```

#### Criar service
```bash
docker compose exec app kubectl expose deployment <deployment-name> --port=80 --target-port=8080 -n <namespace>
```

#### Criar configmap de arquivo
```bash
docker compose exec app kubectl create configmap <configmap-name> --from-file=<file-path> -n <namespace>
```

#### Criar secret
```bash
docker compose exec app kubectl create secret generic <secret-name> --from-literal=key1=value1 -n <namespace>
```

## Deletar Recursos

#### Deletar pod
```bash
docker compose exec app kubectl delete pod <pod-name> -n <namespace>
```

#### Deletar deployment
```bash
docker compose exec app kubectl delete deployment <deployment-name> -n <namespace>
```

#### Deletar service
```bash
docker compose exec app kubectl delete service <service-name> -n <namespace>
```

#### Deletar namespace (e todos os recursos dentro)
```bash
docker compose exec app kubectl delete namespace <namespace-name>
```

#### Deletar por arquivo
```bash
docker compose exec app kubectl delete -f <file.yaml>
```

#### Forçar deleção
```bash
docker compose exec app kubectl delete pod <pod-name> -n <namespace> --force --grace-period=0
```

## Copiar Arquivos

#### Copiar arquivo do pod para local
```bash
docker compose exec app kubectl cp <namespace>/<pod-name>:<path-in-pod> <local-path>
```

#### Copiar arquivo local para pod
```bash
docker compose exec app kubectl cp <local-path> <namespace>/<pod-name>:<path-in-pod>
```

#### Copiar de container específico
```bash
docker compose exec app kubectl cp <namespace>/<pod-name>:<path-in-pod> <local-path> -c <container-name>
```

## Monitoramento

#### Ver uso de recursos dos nodes
```bash
docker compose exec app kubectl top nodes
```

#### Ver uso de recursos dos pods
```bash
docker compose exec app kubectl top pods -A
```

#### Ver uso de recursos de pods em namespace específico
```bash
docker compose exec app kubectl top pods -n <namespace>
```

#### Ver eventos do cluster
```bash
docker compose exec app kubectl get events -A --sort-by='.lastTimestamp'
```

#### Ver eventos de namespace específico
```bash
docker compose exec app kubectl get events -n <namespace> --sort-by='.lastTimestamp'
```

## Configuração

#### Ver configuração atual
```bash
docker compose exec app kubectl config view
```

#### Ver configuração sem dados sensíveis
```bash
docker compose exec app kubectl config view --minify
```

#### Definir namespace padrão
```bash
docker compose exec app kubectl config set-context --current --namespace=<namespace>
```

#### Atualizar kubeconfig do EKS
```bash
docker compose exec app sh -c "AWS_PROFILE=<your-profile> aws eks update-kubeconfig --name <cluster-name> --region <region>"
```

## Troubleshooting

#### Verificar status dos componentes do cluster
```bash
docker compose exec app kubectl get componentstatuses
```

#### Ver pods com problemas
```bash
docker compose exec app kubectl get pods -A --field-selector=status.phase!=Running
```

#### Ver pods que foram reiniciados
```bash
docker compose exec app kubectl get pods -A --sort-by='.status.containerStatuses[0].restartCount'
```

#### Verificar recursos que consomem mais CPU
```bash
docker compose exec app kubectl top pods -A --sort-by=cpu
```

#### Verificar recursos que consomem mais Memória
```bash
docker compose exec app kubectl top pods -A --sort-by=memory
```

#### Debug de pod
```bash
docker compose exec app kubectl debug <pod-name> -n <namespace> -it --image=busybox
```

## Filtros e Seletores

#### Filtrar por labels
```bash
docker compose exec app kubectl get pods -l app=nginx
```

#### Filtrar por múltiplas labels
```bash
docker compose exec app kubectl get pods -l app=nginx,version=v1
```

#### Filtrar por field selector
```bash
docker compose exec app kubectl get pods --field-selector=status.phase=Running
```

#### Combinar label e field selector
```bash
docker compose exec app kubectl get pods -l app=nginx --field-selector=status.phase=Running
```

## Referências
- [Kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [Kubectl Reference](https://kubernetes.io/docs/reference/kubectl/)
- [AWS EKS Documentation](https://docs.aws.amazon.com/eks/)