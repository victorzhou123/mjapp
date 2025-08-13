#!/bin/bash
set -e

# 从.env文件加载环境变量
if [ -f ".env" ]; then
  echo "Loading environment variables from .env file"
  
  # 读取.env文件并设置环境变量
  while IFS= read -r line || [ -n "$line" ]; do
    # 跳过注释和空行
    case $line in
      \#*|'')
        continue
        ;;
    esac
    
    # 提取变量名和值
    if echo $line | grep -q '='; then
      var_name=$(echo $line | cut -d '=' -f 1)
      var_value=$(echo $line | cut -d '=' -f 2-)
      
      # 处理特殊字符
      var_value=$(echo "$var_value" | sed 's/\r//g')
      
      # 导出环境变量
      export $var_name="$var_value"
      echo "Set environment variable: $var_name"
    fi
  done < .env
  
  echo "Environment variables loaded successfully"
else
  echo "Warning: .env file not found"
fi

# 执行传入的命令或默认命令
exec "$@"