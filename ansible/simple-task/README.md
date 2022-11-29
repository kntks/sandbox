# simple-task
```
$ pwd
xxxx/ansible/simple-task
$ ansible-playbook playbook.yml -i localhost, -e ansible_connection=local
```

リリースバージョンの値は以下のコマンドで取得する
```
$ aws ssm get-parameter --name /aws/service/eks/optimized-ami/1.24/amazon-linux-2/recommended --query "Parameter.Value" --output text  | jq .release_version
```