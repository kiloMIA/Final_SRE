provider "aws" {
    region = "us-east-1" 
}

resource "aws_instance" "web_server" {
    ami           = "ami-0fc5d935ebf8bc3bc"  
    instance_type = "t2.micro"
    key_name      = "sre_key_pair"  
    security_groups = ["launch-wizard-1"]  


    user_data = <<-EOF
        #!/bin/bash
    EOF

    tags = {
        Name = "WebServer"
    }
}

output "web_server_public_ip" {
    value = aws_instance.web_server.public_ip
}
