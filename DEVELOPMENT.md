## Contributing Guide

* Create pull latest code of `main` branch to your local `main` branch.
```sh
User:(any-branch)$ git checkout main
User:(main)$ git pull origin main
```

* Create a new branch from `main` branch, use proper naming convension, i.e.- `feature-xyz`, `ISSUE-ID-xyz` etc.
```sh
User:(main)$ git checkout -b feature-xyz
```

* User proper commit message, i.e- `#ISSUE-ID TYPE: your commit message`. TYPE are `chore`, `fix`, `docs`, `test`, `feature`, `refactor`.
```sh
User:(feature-xyz)$ git add -p
User:(feature-xyz)$ git commit -m "commit message"
```

* Rebase your code with letest code of `main` branch`.
```sh
User:(feature-xyz)$ git pull --all
User:(feature-xyz)$ git rebase main
```

* Push your code, and give a Pull Request.
```sh
User:(feature-xyz)$ git push origin feature-xyz
```

# Developing Guideline

* Steps to generate source code from proto file.
```sh
$ cd grpc
$ protoc --go_out=plugins=grpc:<package-name> <proto file name>
# For example
$ protoc --go_out=plugins=grpc:todo todo.proto
```