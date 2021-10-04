# terratest_test
Terratestのテストコード
Terraformでapply（構築）時にリソースに設定した属性値をファイル（terraform.tfstate）に出力し、その出力値とTerratestテストコードに記述した設定値とを比較します。　

#環境条件
実行環境は Go 1.13 以上のバージョンであることが必要になります。
