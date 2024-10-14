# imgproxy-test
## imgproxyドキュメント
https://docs.imgproxy.net/
## 下準備
### .envファイル用意
```
cp .env.sample .env
```
### 署名付きURL用のKEY/SALT作成
```
echo $(xxd -g 2 -l 64 -p /dev/random | tr -d '\n')
```
出力されたものを.envに貼る
## imgproxy+envoy起動方法
```
docker compose run -d
```
## 利用方法
### tool/create_signed_path.goで条件と画像ファイルを設定
```
processingOptions := "rs:fill:400:400"
sourceUrl := {画像ファイルのパス}
encodedUrl := base64.URLEncoding.EncodeToString([]byte(sourceUrl))
path := fmt.Sprintf("/%s/%s", processingOptions, encodedUrl)
```
### 署名付きURL作成
```
docker compose run --rm tool go run create_signed_path.go
```
### 出力されたパスにアクセスする
```
open http://localhost:10000/{hogehoge}
```
