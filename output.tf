# Terratestでのサンプルテスト用の出力設定
# ※Terratestのサンプルテスト向けに構築した値を「terraform.tfstate」ファイルにoutputとして出力します
#   サンプルでは、2台のEC2の設定値をテストしています
#
#
#
## EC2サーバaの設定値確認用の出力設定
output "ec2-a_ami" {
  value = aws_instance.test-a.ami
}
output "ec2-a_instance_type" {
  value = aws_instance.test-a.instance_type
}
output "ec2-a_private_ip" {
  value = aws_instance.test-a.private_ip
}
output "ec2-a_volume_type" {
  value = aws_instance.test-a.root_block_device[0].volume_type
}
output "ec2-a_volume_size" {
  value = aws_instance.test-a.root_block_device[0].volume_size
}

## EC2サーバcの設定値確認用の出力設定
output "ec2-c_ami" {
  value = aws_instance.test-c.ami
}
output "ec2-c_instance_type" {
  value = aws_instance.test-c.instance_type
}
output "ec2-c_private_ip" {
  value = aws_instance.test-c.private_ip
}
output "ec2-c_volume_type" {
  value = aws_instance.test-c.root_block_device[0].volume_type
}
output "ec2-c_volume_size" {
  value = aws_instance.test-c.root_block_device[0].volume_size
}