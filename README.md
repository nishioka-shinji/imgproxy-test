# imgproxy-test
## imgproxyドキュメント
https://docs.imgproxy.net/
## 下準備
### .envファイル用意
```
cp .env.sample .env
```
### tool用にGo環境用意（本体とは関係ない）
```
docker build --no-cache -f tool/Dockerfile -t my-golang-app .
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
