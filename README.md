# Lazyaws-environment  <img src="https://upload.wikimedia.org/wikipedia/commons/9/93/Amazon_Web_Services_Logo.svg" alt="AWS" width="60" height="36"> <img src="https://raw.githubusercontent.com/derailed/k9s/master/assets/k9s.png" alt="K9s" width="60" height="36"> <img src="https://upload.wikimedia.org/wikipedia/commons/3/39/Kubernetes_logo_without_workmark.svg" alt="Kubernetes" width="36" height="36">

## The Idea :bulb:
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