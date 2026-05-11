# cloudcanal-openapi-cli

CloudCanal OpenAPI 的命令行工具，支持：

- 单次命令执行
- `--output json` 机器可读输出
- 安装脚本默认配置 zsh / bash shell 补全

完整命令说明见 [docs/cloudcanal-cli-usage.md](docs/cloudcanal-cli-usage.md)。

## 快速开始

1. 安装

```bash
curl -fsSL https://raw.githubusercontent.com/ClouGence/cloudcanal-openapi-cli/main/scripts/install.sh | bash
```

2. 非交互初始化

```bash
cloudcanal config init --api-host https://cc.example.com --ak your-ak --sk your-sk
```

## 常用用法

命令执行：

```bash
cloudcanal version
cloudcanal --version
cloudcanal --help
cloudcanal jobs --help
cloudcanal config init --api-host https://cc.example.com --ak your-ak --sk your-sk
cloudcanal config profiles list
cloudcanal config profiles use dev
cloudcanal config profiles add test --api-host https://cc.example.com --ak test-ak --sk test-sk
cloudcanal config lang set zh
cloudcanal jobs list
cloudcanal jobs show 123
cloudcanal jobs create --body-file create-job.json
cloudcanal datasources list --type MYSQL
cloudcanal datasources add --body-file add-datasource.json
cloudcanal workers list --cluster-id 2
cloudcanal schemas list-trans-objs-by-meta --src-db demo --src-trans-obj orders
```

JSON 输出：

```bash
cloudcanal jobs list --type SYNC --output json
```

## 配置

配置文件默认保存在：

```text
~/.cloudcanal-cli/config.json
```

最小配置示例：

```json
{
  "language": "en",
  "currentProfile": "dev",
  "profiles": {
    "dev": {
      "apiBaseUrl": "https://cc.example.com",
      "accessKey": "your-ak",
      "secretKey": "your-sk"
    }
  }
}
```

如果你需要调整网络行为，也可以在具体 profile 下追加这些可选项：

```json
{
  "profiles": {
    "dev": {
      "httpTimeoutSeconds": 15,
      "httpReadMaxRetries": 2,
      "httpReadRetryBackoffMillis": 300
    }
  }
}
```

## 文档入口

- 安装、初始化、命令参数、示例：[docs/cloudcanal-cli-usage.md](docs/cloudcanal-cli-usage.md)
- SDK API 对照 CLI 命令：[docs/openapi-sdk-api-reference.md](docs/openapi-sdk-api-reference.md)
- 版本变更记录：[CHANGELOG.md](CHANGELOG.md)
- 机器可读输出：在查询命令后追加 `--output json`
- shell 补全由安装脚本自动配置

## 卸载

```bash
curl -fsSL https://raw.githubusercontent.com/ClouGence/cloudcanal-openapi-cli/main/scripts/uninstall.sh | bash
```
