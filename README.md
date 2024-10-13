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
sourceUrl := "https://storage.googleapis.com/studio-design-asset-files/projects/7kadpxLza3/s-1616x792_v-fms_webp_b5774a47-5fde-4867-987a-3a2bb4664066.webp"
path := fmt.Sprintf("/%s/plain/%s", processingOptions, sourceUrl)
```
### 署名付きURL作成
```
docker compose run --rm tool go run create_signed_path.go
```
### 出力されたパスにアクセスする
```
open http://localhost:10000/#{hogehoge}
```
