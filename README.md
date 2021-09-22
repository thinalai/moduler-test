# moduler-test

## Task 1. Scripting language

Build site via Hugo.

- Create a new post with random content by run `cd task1 && ./generate.sh <Title>`.
- Deploy website by `[GitHub Actions](https://gohugo.io/hosting-and-deployment/hosting-on-github/)` when `main-pages` branch updates.

### Dependencies tools

Website and random post generator needs `hugo` and `fortune` CLI, if you are using macOS, you can install them by package manager `brew`.

```bash
brew install hugo
brew install fortune
```

## Task 2. Configuration management

Configure the applications with `ansible` playbook.

See more details in [task2](./task2) directory.

## Task 6. CI/CD

Do CI with Github Actions and CD with Azure DevOps pipelines.

See more details in [task6](./task6) directory.
