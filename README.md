# Lazyaws-environment

The idea behind this project is to use a tool that can connect to AWS to manage Kubernetes more easily and with better visualization than Cloudshell. If you don't have a tool like Rancher or Portainer available to manage Kubernetes, this tool has k9s built in.

I saw this project https://github.com/fuziontech/lazyaws and this other one https://github.com/victorhugorch/go-enviromment, and thought... why not combine the two projects into one? So I merged the two projects to have a complete environment with Docker, without having to install another tool to use.

That's how lazy-enviroment was created.

## How to use

1. `git clone https://github.com/walterfjustino/lazyaws-enviroment.git` 
2. `docker compose run --rm app lazyaws`
3. Access the `K9S_GUIDE.md` to see how to use, if you have any questions the k9s link: `https://github.com/derailed/k9s`.
4. enjoy ;)

## License

[MIT](https://github.com/walterfjustino/lazyaws-enviroment/blob/master/LICENSE)
