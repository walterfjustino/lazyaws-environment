# K9s - Guia de Comandos

## Acessando o K9s

Após atualizar o kubeconfig no LazyAWS (pressione `K` no cluster) pressione `9` e irá abrir o k9s.

## Atalhos Principais do K9s

### Navegação Básica
- `↑/↓` ou `j/k` - Navegar entre itens
- `Enter` - Ver detalhes do recurso selecionado
- `Esc` - Voltar para a tela anterior
- `:q` ou `Ctrl+C` - Sair do k9s

### Visualização de Recursos
- `:pods` ou `:po` - Ver pods
- `:deployments` ou `:deploy` - Ver deployments
- `:services` ou `:svc` - Ver services
- `:nodes` - Ver nodes do cluster
- `:namespaces` ou `:ns` - Ver namespaces
- `:configmaps` ou `:cm` - Ver configmaps
- `:secrets` - Ver secrets
- `:ingress` ou `:ing` - Ver ingress
- `:pvc` - Ver persistent volume claims
- `:jobs` - Ver jobs
- `:cronjobs` - Ver cronjobs

### Filtros e Busca
- `0` - Mostrar recursos de TODOS os namespaces
- `1-9` - Filtrar por namespace específico
- `/` - Buscar/filtrar recursos
- `Esc` - Limpar filtro

### Ações em Recursos
- `d` - Descrever recurso (describe)
- `y` - Ver YAML do recurso
- `e` - Editar recurso
- `l` - Ver logs do pod
- `s` - Abrir shell no container
- `Ctrl+D` - Deletar recurso
- `Ctrl+K` - Matar/reiniciar pod

### Logs
- `l` - Ver logs do pod selecionado
- `p` - Ver logs do pod anterior (útil após crash)
- `f` - Seguir logs em tempo real (tail -f)
- `w` - Wrap/unwrap linhas longas
- `s` - Salvar logs em arquivo

### Pulse - Vizualize um painel de controle sobre o estado atual do seu cluster
- `Shift+:` e digite `pulse`- Vizualize um painel de controle sobre o estado atual do seu cluster

### XRay - Explore os recursos do seu cluster e visualize suas dependências
- `Shift+:` e digite `xray`- Visualize os recursos do seu cluster e visualize suas dependências

### RBAC - Explore os recursos do seu cluster e visualize suas dependências
- `Shift+:` e digite `rbac`- Veja quem/o quê/como das autorizações no seu cluster

### Outras Funcionalidades
- `?` - Mostrar todos os atalhos disponíveis
- `Ctrl+A` - Mostrar todos os recursos
- `r` - Atualizar/refresh
- `z` - Alternar entre modos de visualização
- `x` - Executar comando no pod

## Troubleshooting

### K9s não mostra pods
1. Verifique se o kubeconfig foi atualizado corretamente:
   ```bash
   docker compose exec app cat ~/.kube/config
   ```
2. Teste a conexão com kubectl:
   ```bash
   docker compose exec app kubectl get pods -A
   ```
3. Verifique se está usando o profile AWS correto:
   ```bash
   docker compose exec app sh -c "AWS_PROFILE=<your-profile> kubectl get pods -A"
   ```
4. Verifique se o cluster está na região correta:
   ```bash
   docker compose exec app sh -c "AWS_PROFILE=<your-profile> aws eks describe-cluster --name <cluster-name> --region <region>"
   ```

### Atualizar kubeconfig manualmente
```bash
# Dentro do container
docker compose exec app sh -c "AWS_PROFILE=<your-profile> aws eks update-kubeconfig --name <cluster-name> --region <region>"
```

## Dicas

1. **Use o namespace correto**: Muitos pods estão em namespaces específicos, não no `default`. Pressione `0` no k9s para ver todos.

2. **Verifique permissões**: Certifique-se de que suas credenciais AWS têm permissão para acessar o cluster EKS.

3. **Use filtros**: No k9s, use `/` para filtrar pods por nome, facilitando a busca.

4. **Logs em tempo real**: Pressione `l` no pod e depois `f` para seguir os logs em tempo real.

5. **Shell rápido**: Pressione `s` no pod para abrir um shell interativo rapidamente.

## Referências

- [K9s Documentation](https://k9scli.io/)
- [Kubectl Cheat Sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)
- [AWS EKS Documentation](https://docs.aws.amazon.com/eks/)
