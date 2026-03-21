## AWS(EC2/Ubuntu)環境でのOpenVPN疎通トラブルと解決策

### 事象
AWS上にTerraformで構築した攻撃拠点（Ubuntu）から、TryHackMeのVPNに接続しようとした際、以下のエラーが発生。
Options error: option 'auth-user-pass' is not expected to be inline

### 原因分析
THMのAccessページで「macOS用」の.ovpnをダウンロードし、Linux CLIのopenvpnコマンドで読み込もうとしたため。

### 解決策

THM側で「Linux」用プロファイルを再取得する。

Terraformのuser_dataにおいて、Linux用設定を直接流し込み、インスタンス起動時にVPN接続を自動完結させる（IaC化）。

### 考察
ローカル環境（Mac）とクラウド環境（EC2）の「OS作法の違い」を意識することの重要性を再認識。