# CI/CD 使用教程

[Github CI](https://docs.github.com/en/actions/configuring-and-managing-workflows/configuring-and-managing-workflow-files-and-runs)
- 创建workflow, 在github项目根目录下创建workflow, [参考](https://docs.github.com/en/actions/configuring-and-managing-workflows/configuring-a-workflow#creating-a-workflow-file). 
![](workflow%20build.png)
    - 在仓库创建`.github/workflows`目录
    - 在`.github/workflows`目录下创建`.yaml`文件的workflow file(可以创建多个)
- 配置`workflow`, [参考](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions). 也可使用
[样板](https://docs.github.com/en/actions/getting-started-with-github-actions/starting-with-preconfigured-workflow-templates)

## 示例
- [github官方示例](../../.github/workflows/workflow.yaml)
- [本项目CI文件](../../.github/workflows/bookNote.yml)

## 参考
[repository checkout工具](https://github.com/actions/checkout)