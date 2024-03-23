resource "aws_vpc" "eq_api_vpc" {
  cidr_block = "10.0.0.0/16"
  enable_dns_support = true
  enable_dns_hostnames = true

  tags = {
    Name = "eq-api-vpc"
  }
}

resource "aws_subnet" "eq_api_subnet_1" {
  vpc_id            = aws_vpc.eq_api_vpc.id
  cidr_block        = "10.0.1.0/24"
  availability_zone = "us-west-2a"

  tags = {
    Name = "eq-api-subnet-1"
  }
}

resource "aws_subnet" "eq_api_subnet_2" {
  vpc_id            = aws_vpc.eq_api_vpc.id
  cidr_block        = "10.0.2.0/24"
  availability_zone = "us-west-2b"

  tags = {
    Name = "eq-api-subnet-2"
  }
}

resource "aws_security_group" "eq_api_sg" {
  name        = "eq-api-security-group"
  description = "Allow inbound traffic"
  vpc_id      = aws_vpc.eq_api_vpc.id

  ingress {
    from_port   = 443
    to_port     = 443
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }

  tags = {
    Name = "eq-api-security-group"
  }
}

resource "aws_route_table" "eq_api_rt" {
  vpc_id = aws_vpc.eq_api_vpc.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.eq_api_igw.id
  }

  tags = {
    Name = "eq-api-rt"
  }
}

resource "aws_route_table_association" "eq_api_rta_subnet_1" {
  subnet_id      = aws_subnet.eq_api_subnet_1.id
  route_table_id = aws_route_table.eq_api_rt.id
}

resource "aws_route_table_association" "eq_api_rta_subnet_2" {
  subnet_id      = aws_subnet.eq_api_subnet_2.id
  route_table_id = aws_route_table.eq_api_rt.id
}
